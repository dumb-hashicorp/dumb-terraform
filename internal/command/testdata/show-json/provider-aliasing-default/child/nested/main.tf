dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
}

resource "test_instance" "test" {
  ami = "baz"
}
