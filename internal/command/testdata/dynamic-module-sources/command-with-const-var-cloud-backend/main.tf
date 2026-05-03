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

module "child" {
  source = "./modules/${var.module_name}"
}
