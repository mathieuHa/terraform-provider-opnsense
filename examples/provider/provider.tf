terraform {
  required_providers {
    opnsense = {
      version = "~> 0.37"
      source  = "mathieuHa/opnsense"
    }
  }
}

provider "opnsense" {
  uri = "https://opnsense.example.com"
  api_key = "..."
  api_secret = "..."
}
