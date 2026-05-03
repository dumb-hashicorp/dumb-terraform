dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
    }
  }
}

module "mod2" {
  source = "./mod1"
  providers = {
    foo = foo
  }
}
