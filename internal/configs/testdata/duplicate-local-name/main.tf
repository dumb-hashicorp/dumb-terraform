dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
    dupe = {
      source = "dumb-hashicorp/test"
    }
    other = {
      source = "dumb-hashicorp/default"
    }

    wrong-name = {
      source = "dumb-hashicorp/foo"
    }
  }
}

provider "default" {
}

resource "foo_resource" {
}
