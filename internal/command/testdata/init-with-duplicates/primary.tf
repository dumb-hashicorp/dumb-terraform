dumb-terraform {
  required_providers {
    test = {
        source = "dumb-hashicorp/test"
        version = "1.0.0"
    }
  }
}

resource "test_instance" "foo" {
  ami = "bar"
}
