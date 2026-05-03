dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
}

provider "test" {}

resource "test_instance" "example" {
  ami = "ami-12345"
}
