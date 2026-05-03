dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_nonexistent" { # nonexistent is not a valid state store type in the mocked provider
    provider "test" {}
  }
}
