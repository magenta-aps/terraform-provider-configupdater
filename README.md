# terraform-provider-configupdater

Terraform provider for
[config-updater](https://labs.docs.magenta.dk/config-updater.html).


## Usage

```
terraform {
  required_providers {
    configupdater = {
      version = "change-this"
      source  = "magenta-aps/configupdater"
    }
  }
}

provider "configupdater" {
  # basic auth
  username = "admin"
  password = "password1"
}

resource "configupdater_secret" "my-server-name" {
  file_path = "mailgun/postmaster__magenta.dk.enc.yaml"
  secret = {
    asecret = "hunter2"
    another = "2hunter"
  }
}
```


## Development

You will need `go`, `terraform` and `make`.

1. Start `config-updater` in `salt-automation/config-updater`
2. `make` -- this will build the provider and init terraform (see `main.tf`)
3. `terraform apply` (`terraform destroy` is also useful)


## License

Mozilla Public License Version 2.0, see LICENSE.
