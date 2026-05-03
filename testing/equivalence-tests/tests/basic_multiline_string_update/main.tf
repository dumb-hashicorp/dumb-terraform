dumb-terraform {
  required_providers {
    tfcoremock = {
      source  = "dumb-hashicorp/tfcoremock"
      version = "0.1.1"
    }
  }
}

provider "tfcoremock" {}

resource "tfcoremock_simple_resource" "multiline" {
  string = "one\nthree\ntwo\nfour\nsix\nseven"
}
