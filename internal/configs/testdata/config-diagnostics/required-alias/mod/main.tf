dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
      version = "1.0.0"
      configuration_aliases = [ foo.bar ]
    }
  }
}

resource "foo_resource" "a" {
  provider = foo.bar
}
