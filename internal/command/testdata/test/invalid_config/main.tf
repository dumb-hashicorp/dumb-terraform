dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
    }
  }
}

resource "test_resource" "foo" {
  nein = "foo"
}
