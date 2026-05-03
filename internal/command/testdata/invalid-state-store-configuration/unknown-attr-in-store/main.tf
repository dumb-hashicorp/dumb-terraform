dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {}
    unknown = "this isn't in test_store's schema" # Should trigger an error
  }
}
