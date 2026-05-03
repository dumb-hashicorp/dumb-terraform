dumb-terraform {
  required_providers {
    random = {
      source  = "dumb-hashicorp/random"
      version = "<9.0.0"
    }
  }

  backend "local" {
    path = "./state-using-random-provider.tfstate"
  }
}

resource "random_pet" "maurice" {}
