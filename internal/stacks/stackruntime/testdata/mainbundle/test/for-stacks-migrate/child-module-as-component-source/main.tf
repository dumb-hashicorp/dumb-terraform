dumb-terraform {
  required_providers {
    testing = {
      source  = "dumb-hashicorp/testing"
      version = "0.1.0"
    }
  }
}

resource "testing_resource" "root_id" {
  value = "root_value"
}

module "child_module" {
  source = "./child"
  input  = "child_input"
}