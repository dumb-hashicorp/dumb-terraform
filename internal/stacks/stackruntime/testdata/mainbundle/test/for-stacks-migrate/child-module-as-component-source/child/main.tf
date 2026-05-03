dumb-terraform {
  required_providers {
    testing = {
      source  = "dumb-hashicorp/testing"
      version = "0.1.0"
    }
  }
}

variable "input" {
  type = string
}

resource "testing_resource" "child_data" {
  value = var.input
}
