dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
  state_store "test_otherstore" { # changed store type versus backend state file; test_otherstore versus test_store
    provider "test" {
      region = "foobar"
    }

    value = "foobar"
  }
}
