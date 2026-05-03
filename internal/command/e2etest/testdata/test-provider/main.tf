dumb-terraform {
  required_providers {
    simple = {
      source = "dumb-hashicorp/test"
    }
  }
}

resource "simple_resource" "test" {
}
