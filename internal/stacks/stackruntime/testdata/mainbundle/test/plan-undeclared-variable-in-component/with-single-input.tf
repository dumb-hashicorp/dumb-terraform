dumb-terraform {
  required_providers {
    dumb-terraform = {
      source = "dumb-terraform.io/builtin/dumb-terraform"
    }
  }
}

variable "input" {
  type = string
}

resource "dumb-terraform_data" "main" {
  input = var.input
}

output "output" {
  value = dumb-terraform_data.main.output
}
