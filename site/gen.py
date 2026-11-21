import json
import os
import re
import shutil
import sys

import yaml

ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
SITE = os.path.join(ROOT, "site")
BUILD = os.path.join(SITE, "build")
DOCS = os.path.join(BUILD, "docs")

SERVICES = {
    "auth": ("Users & groups", "core", None),
    "captiveportal": ("Captive portal", "core", None),
    "core": ("System tunables", "core", None),
    "cron": ("Cron", "core", None),
    "dhcrelay": ("DHCP relay", "core", None),
    "dnsmasq": ("Dnsmasq DNS & DHCP", "core", None),
    "firewall": ("Firewall", "core", None),
    "firmware": ("Firmware & plugins", "core", None),
    "hasync": ("High availability", "core", None),
    "hostdiscovery": ("Host discovery", "core", None),
    "ids": ("Intrusion detection (Suricata)", "core", None),
    "interface": ("Interfaces", "core", None),
    "interfaces": ("Interfaces", "core", None),
    "ipsec": ("IPsec", "core", None),
    "kea": ("Kea DHCP", "core", None),
    "monit": ("Monit", "core", None),
    "openvpn": ("OpenVPN", "core", None),
    "radvd": ("Router advertisements", "core", None),
    "route": ("Routing", "core", None),
    "routing": ("Routing", "core", None),
    "syslog": ("Syslog", "core", None),
    "trafficshaper": ("Traffic shaper", "core", None),
    "trust": ("Trust & certificates", "core", None),
    "unbound": ("Unbound DNS", "core", None),
    "wireguard": ("WireGuard", "core", None),
    "acmeclient": ("ACME client", "plugin", "os-acme-client"),
    "caddy": ("Caddy", "plugin", "os-caddy"),
    "chrony": ("Chrony", "plugin", "os-chrony"),
    "crowdsec": ("CrowdSec", "plugin", "os-crowdsec"),
    "haproxy": ("HAProxy", "plugin", "os-haproxy"),
    "nodeexporter": ("Prometheus node exporter", "plugin", "os-node_exporter"),
    "nut": ("Network UPS Tools", "plugin", "os-nut"),
    "proxy": ("Squid web proxy", "plugin", "os-squid"),
    "quagga": ("FRR routing (BGP)", "plugin", "os-frr"),
    "tailscale": ("Tailscale", "plugin", "os-tailscale"),
    "telegraf": ("Telegraf", "plugin", "os-telegraf"),
    "zerotier": ("ZeroTier", "plugin", "os-zerotier"),
}

KINDS = (("resources", "Resource"), ("data-sources", "Data source"), ("ephemeral-resources", "Ephemeral resource"))
CALLOUTS = {"~>": "warning", "->": "note", "!>": "danger"}


def slug(title):
    return re.sub(r"[^a-z0-9]+", "-", title.lower()).strip("-")


def service_of(name):
    for prefix in sorted(SERVICES, key=len, reverse=True):
        if name == prefix or name.startswith(prefix + "_"):
            return prefix
    raise SystemExit(f"no service for {name}; add its prefix to SERVICES in site/gen.py")


def split_frontmatter(text):
    if text.startswith("---\n"):
        end = text.index("\n---\n", 4)
        return yaml.safe_load(text[4:end]) or {}, text[end + 5:]
    return {}, text


def convert(body):
    out, lines, i = [], body.split("\n"), 0
    while i < len(lines):
        line = lines[i]
        m = re.match(r"^(~>|->|!>) (.*)$", line)
        if m:
            out.append(f"!!! {CALLOUTS[m.group(1)]}")
            out.append("")
            out.append("    " + m.group(2))
            i += 1
            while i < len(lines) and lines[i].strip():
                out.append("    " + lines[i])
                i += 1
            continue
        out.append(line)
        i += 1
    text = "\n".join(out)
    text = re.sub(r"\]\((\.\./(?:resources|data-sources|ephemeral-resources)/[a-z0-9_]+)\)", r"](\1.md)", text)
    text = re.sub(r"\]\(([a-z0-9_]+)\)", r"](\1.md)", text)
    return text


def first_sentence(desc):
    desc = re.sub(r"\s+", " ", (desc or "").strip())
    m = re.match(r"(.+?\.)(\s|$)", desc)
    return (m.group(1) if m else desc).replace("|", "\\|")


def collect():
    items = []
    for kind, label in KINDS:
        src = os.path.join(ROOT, "docs", kind)
        if not os.path.isdir(src):
            continue
        for f in sorted(os.listdir(src)):
            if not f.endswith(".md"):
                continue
            name = f[:-3]
            meta, body = split_frontmatter(open(os.path.join(src, f)).read())
            items.append({"name": name, "kind": kind, "label": label, "service": service_of(name),
                          "desc": meta.get("description", ""), "body": body,
                          "deprecated": bool(re.search(r"^~> \*\*Deprecated\*\*", body, re.M))})
    return items


def singleton_id(it):
    text = (it["desc"] or "") + "\n" + it["body"]
    if it["kind"] != "resources" or "singleton" not in text.lower():
        return None
    for pat in (r"import it with the ID `([a-z0-9_]+)`", r"fixed string `([a-z0-9_]+)`",
                r"import opnsense_[a-z0-9_]+\.[a-z0-9_]+ ([a-z0-9_]+)\s"):
        m = re.search(pat, text)
        if m:
            return m.group(1)
    raise SystemExit(f"singleton {it['name']} has no recognisable import ID")


def singletons_table(items):
    rows = ["| Resource | Import ID |", "|---|---|"]
    for it in sorted(items, key=lambda x: x["name"]):
        sid = singleton_id(it)
        if sid:
            rows.append(f"| [`opnsense_{it['name']}`](../reference/resources/{it['name']}.md) | `{sid}` |")
    return "\n".join(rows)


def plugins_table(items):
    rows = ["| Plugin | Package | Resources |", "|---|---|---|"]
    seen = {}
    for it in items:
        title, cat, pkg = SERVICES[it["service"]]
        if cat == "plugin" and it["kind"] == "resources":
            seen.setdefault((title, pkg), []).append(it["name"])
    for (title, pkg), names in sorted(seen.items()):
        anchor = slug(title)
        rows.append(f"| [{title}](../resources.md#{anchor}) | `{pkg}` | {len(names)} |")
    return "\n".join(rows)


def write(path, text):
    os.makedirs(os.path.dirname(path), exist_ok=True)
    open(path, "w").write(text)


def catalogue(items, stats):
    by = {}
    for it in items:
        by.setdefault(SERVICES[it["service"]][0], []).append(it)
    names = {(it["kind"], it["name"]) for it in items}
    out = ["# Supported resources", "",
           f"The provider manages **{stats['resources']} resources** and reads **{stats['data_sources']} data sources** "
           f"across **{stats['services']} services**: {stats['core']} built into OPNsense and {stats['plugins']} plugins. "
           "Plugin resources need the plugin installed, for example with "
           "[`opnsense_firmware_plugin`](reference/resources/firmware_plugin.md).", ""]
    for cat, heading in (("core", "OPNsense core"), ("plugin", "Plugins")):
        out += [f"## {heading}", ""]
        for title in sorted(t for t in by if SERVICES[service_of(by[t][0]["name"])][1] == cat):
            plugin = SERVICES[service_of(by[title][0]["name"])][2]
            out += [f"### {title}", ""]
            if plugin:
                out += [f"Plugin `{plugin}`.", ""]
            out += ["| Resource | Data source | Description |", "|---|---|---|"]
            for it in sorted(by[title], key=lambda x: x["name"]):
                if it["kind"] == "data-sources" and ("resources", it["name"]) in names:
                    continue
                rname = f"opnsense_{it['name']}"
                if it["kind"] == "resources":
                    res = f"[`{rname}`](reference/resources/{it['name']}.md)"
                    ds = f"[yes](reference/data-sources/{it['name']}.md)" if ("data-sources", it["name"]) in names else ""
                else:
                    res = f"{it['label']} only"
                    ds = f"[`{rname}`](reference/{it['kind']}/{it['name']}.md)"
                desc = first_sentence(it["desc"])
                if it["deprecated"]:
                    desc = "*Deprecated.* " + desc
                out.append(f"| {res} | {ds} | {desc} |")
            out.append("")
    return "\n".join(out)


def nav_reference(items):
    groups = {}
    for it in items:
        groups.setdefault(SERVICES[it["service"]][0], []).append(it)
    nav = []
    for title in sorted(groups):
        entries = []
        for it in sorted(groups[title], key=lambda x: ([k for k, _ in KINDS].index(x["kind"]), x["name"])):
            suffix = {"resources": "", "data-sources": " (data source)", "ephemeral-resources": " (ephemeral)"}[it["kind"]]
            entries.append({f"opnsense_{it['name']}{suffix}": f"reference/{it['kind']}/{it['name']}.md"})
        nav.append({title: entries})
    return nav


def main():
    if os.path.isdir(BUILD):
        shutil.rmtree(BUILD)
    items = collect()
    services = {SERVICES[it["service"]][0]: SERVICES[it["service"]][1] for it in items}
    stats = {
        "resources": sum(1 for it in items if it["kind"] == "resources"),
        "data_sources": sum(1 for it in items if it["kind"] == "data-sources"),
        "services": len(services),
        "core": sum(1 for v in services.values() if v == "core"),
        "plugins": sum(1 for v in services.values() if v == "plugin"),
    }
    for it in items:
        write(os.path.join(DOCS, "reference", it["kind"], it["name"] + ".md"), convert(it["body"]).lstrip("\n"))
    provider_meta, provider_body = split_frontmatter(open(os.path.join(ROOT, "docs", "index.md")).read())
    write(os.path.join(DOCS, "reference", "provider.md"), convert(provider_body).lstrip("\n"))
    write(os.path.join(DOCS, "resources.md"), catalogue(items, stats))
    content = os.path.join(SITE, "content")
    guides = []
    for dirpath, _, files in os.walk(content):
        for f in files:
            src = os.path.join(dirpath, f)
            rel = os.path.relpath(src, content)
            text = open(src).read()
            if f.endswith(".md"):
                for k, v in stats.items():
                    text = text.replace("{{ " + k + " }}", str(v))
                text = text.replace("{{ singletons_table }}", singletons_table(items))
                text = text.replace("{{ plugins_table }}", plugins_table(items))
                if rel.startswith("guides" + os.sep):
                    title = re.search(r"^# (.+)$", text, re.M).group(1)
                    guides.append((title, rel))
            write(os.path.join(DOCS, rel), text)
    order = ["guides/singletons.md", "guides/plugins.md"]
    guides.sort(key=lambda g: (order.index(g[1]) if g[1] in order else len(order), g[0]))
    cfg = yaml.safe_load(open(os.path.join(SITE, "mkdocs.base.yml")))
    cfg["docs_dir"] = "docs"
    cfg["site_dir"] = os.path.join(BUILD, "public")
    for ext in cfg["markdown_extensions"]:
        if isinstance(ext, dict) and "pymdownx.snippets" in ext:
            ext["pymdownx.snippets"]["base_path"] = [ROOT]
    cfg["nav"] = [
        {"Home": "index.md"},
        {"Quickstart": "quickstart.md"},
        {"Guides": [{t: r} for t, r in guides]},
        {"Supported resources": "resources.md"},
        {"Reference": [{"Provider configuration": "reference/provider.md"}] + nav_reference(items)},
    ]
    write(os.path.join(BUILD, "mkdocs.yml"), yaml.safe_dump(cfg, sort_keys=False, allow_unicode=True))
    json.dump(stats, sys.stdout)
    print()


if __name__ == "__main__":
    main()
