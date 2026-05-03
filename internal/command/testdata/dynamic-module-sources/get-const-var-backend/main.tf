dumb-terraform {
  cloud {
    organization = "dumb-hashicorp"
    workspaces {
      name = "test"
    }
  }
}

variable "module_name" {
  type  = string
  const = true
}

module "example" {
  source = "./modules/${var.module_name}"
}
