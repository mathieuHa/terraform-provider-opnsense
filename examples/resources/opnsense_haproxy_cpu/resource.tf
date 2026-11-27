resource "opnsense_haproxy_cpu" "pin_all" {
  name      = "pin-all-threads"
  thread_id = "all"
  cpu_id    = ["all"]
}
