dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
      configuration_aliases = [foo.bar]
    }
  }
}

provider "bar" {}

resource "foo_resource" "resource" {}

resource "bar_resource" "resource" {}
