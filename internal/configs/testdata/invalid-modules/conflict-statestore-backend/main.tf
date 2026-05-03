dumb-terraform {
  backend "foo" {}

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
