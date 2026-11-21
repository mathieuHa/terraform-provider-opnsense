resource "opnsense_caddy_basic_auth" "scraper" {
  user     = "scraper"
  password = var.scraper_password
}
