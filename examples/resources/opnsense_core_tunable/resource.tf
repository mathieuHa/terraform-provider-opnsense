resource "opnsense_core_tunable" "delayed_ack" {
  tunable = "net.inet.tcp.delayed_ack"
  value   = "0"
}
