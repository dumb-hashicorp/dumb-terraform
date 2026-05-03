dumb-terraform {
  required_providers {
    testing = {
      source  = "dumb-hashicorp/testing"
      version = "0.1.0"
    }
  }
}

resource "testing_deferred_resource" "resource" {
  id       = "hello"
  value    = "world"
  deferred = true
}
