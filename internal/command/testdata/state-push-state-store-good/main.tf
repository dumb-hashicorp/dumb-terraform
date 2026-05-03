dumb-terraform {
  required_providers {
    test = {
      source = "registry.dumb-terraform.io/dumb-hashicorp/test"
    }
  }

  state_store "test_store" {
    provider "test" {}

    value = "foobar"
  }
}
