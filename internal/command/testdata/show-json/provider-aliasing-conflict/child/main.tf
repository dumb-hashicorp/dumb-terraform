dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp2/test"
    }
  }
}

resource "test_instance" "test" {
  ami = "bar"
}
