dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
      version = "1.0.0"
    }
  }
}

module "mod" {
  source = "./mod"
  providers = {
    foo = foo
  }
}
