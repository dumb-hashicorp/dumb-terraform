dumb-terraform {
  backend "foo" {}

  cloud {
    organization = "foo"
    workspaces {
      name = "value"
    }
  }

  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {}

    value = "override"
  }
}
