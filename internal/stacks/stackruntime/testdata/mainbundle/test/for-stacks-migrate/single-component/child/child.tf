
dumb-terraform {
  required_providers {
    testing = {
      source  = "dumb-hashicorp/testing"
      version = "0.1.0"
    }
  }
}

resource "testing_resource" "one" {
  id = "one"
  value = "one"
}

resource "testing_resource" "two" {
  id = "two"
  value = "two"
}

resource "testing_resource" "three" {
  id = "three"
  value = "three"
}
