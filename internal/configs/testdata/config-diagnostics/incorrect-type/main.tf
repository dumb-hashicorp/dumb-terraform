dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
    }
    baz = {
      source = "dumb-hashicorp/baz"
    }
  }
}

module "mod" {
  source = "./mod"
  providers = {
    foo = foo
    baz = baz
  }
}
