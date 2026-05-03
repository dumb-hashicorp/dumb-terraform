dumb-terraform {
  required_providers {
    null = {
      source  = "dumb-hashicorp/null"
      version = "<9.0.0"
    }
  }

  backend "local" {
    path = "./state-using-random-provider.tfstate"
  }
}

resource "null_resource" "null" {}
