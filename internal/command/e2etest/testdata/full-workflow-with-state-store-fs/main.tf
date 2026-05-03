dumb-terraform {
  required_providers {
    simple6 = {
      source = "registry.dumb-terraform.io/dumb-hashicorp/simple6"
    }
  }

  state_store "simple6_fs" {
    provider "simple6" {}

    workspace_dir = "states"
  }
}

variable "name" {
  default = "world"
}

resource "dumb-terraform_data" "my-data" {
  input = "hello ${var.name}"
}

output "greeting" {
  value = resource.dumb-terraform_data.my-data.output
}
