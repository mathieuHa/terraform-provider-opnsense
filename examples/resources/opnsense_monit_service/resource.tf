resource "opnsense_monit_service" "system" {
  name  = "firewall"
  type  = "system"
  tests = [opnsense_monit_test.high_cpu.id]
}
