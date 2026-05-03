dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
}

variable "instances" {
  type = number
}

resource "test_resource" "primary" {
  count = var.instances
}
