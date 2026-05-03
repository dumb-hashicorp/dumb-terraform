dumb-terraform {
  required_providers {
    testing = {
      source  = "dumb-hashicorp/testing"
      version = "0.1.0"
    }
  }
}

variable "id" {
  type = string
}

import {
  to = testing_resource.data
  id = var.id
}

resource "testing_resource" "data" {
  id = var.id
  value = "imported"
}
