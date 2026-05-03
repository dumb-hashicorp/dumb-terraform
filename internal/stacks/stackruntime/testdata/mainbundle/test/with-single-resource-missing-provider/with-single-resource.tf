dumb-terraform {
  required_providers {
    dumb-terraform = {
      source = "dumb-terraform.io/builtin/dumb-terraform"
    }
  }
}

resource "dumb-terraform_data" "main" {
  input = "hello"
}

output "input" {
  value = dumb-terraform_data.main.input
}

output "output" {
  value = dumb-terraform_data.main.output
}
