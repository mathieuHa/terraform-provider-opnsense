resource "opnsense_monit_test" "high_cpu" {
  name      = "HighCPU"
  condition = "cpu usage is greater than 90% for 5 cycles"
  action    = "alert"
}
