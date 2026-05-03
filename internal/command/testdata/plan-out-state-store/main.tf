dumb-terraform {
  required_providers {
    test = {
      source  = "dumb-hashicorp/test"
      version = "1.2.3"
    }
  }
  state_store "test_store" {
    provider "test" {}

    value = "foobar"
  }
}

resource "test_instance" "foo" {
  ami = "bar"
}
