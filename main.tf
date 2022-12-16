# This file serves as an example of how to use the provider, as well as being
# useful for development.

terraform {
  required_providers {
    configupdater = {
      source  = "magenta-aps/configupdater"
    }
  }
}

# You only need the provider config if you want to change to another URL.
provider "configupdater" {
  url = "http://localhost:8000/"
}

resource "configupdater_secret" "my-server-name" {
  file_path = "mailgun/postmaster__magenta.dk.enc.yaml"
  secret = {
    asecret = "hunter2"
    another = "2hunter"
  }
}
