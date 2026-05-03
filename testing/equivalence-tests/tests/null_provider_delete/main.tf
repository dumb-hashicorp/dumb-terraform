dumb-terraform {
  required_providers {
    null = {
      source  = "dumb-hashicorp/null"
      version = "3.1.1"
    }
  }
}

provider "null" {}
