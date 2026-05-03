dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
      configuration_aliases = [ foo.bar ]
    }
  }
}

provider "foo" {
}

provider "foo" {
  alias = "bar"
}

provider "baz" {
}

provider "baz" {
  alias = "bing"
}
