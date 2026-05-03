dumb-terraform {
  required_providers {
    test = {
      source = "dumb-terraform.io/builtin/test"
    }
  }
}

provider "test" {
  arg = "foo"
}
