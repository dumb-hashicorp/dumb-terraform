dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_store" {
    provider "test" {
      region = "new-value" # changed versus backend state file
    }

    value = "foobar"
  }
}
