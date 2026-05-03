dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {}

    custom_attr = "override"
  }
}
