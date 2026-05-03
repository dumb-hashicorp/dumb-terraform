dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {
      unknown = "this isn't in the test provider's schema" # Should trigger an error
    }
  }
}
