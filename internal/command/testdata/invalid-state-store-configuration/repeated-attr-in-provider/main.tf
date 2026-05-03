dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {
      region = "region1"
      region = "region2" # Should trigger an error
    }
  }
}
