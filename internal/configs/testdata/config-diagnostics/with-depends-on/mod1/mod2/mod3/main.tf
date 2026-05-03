dumb-terraform {
  required_providers {
    foo = {
      source = "dumb-hashicorp/foo"
      configuration_aliases = [ foo.bar ]
    }
  }
}

resource "foo_resource" "a" {
  providers = foo.bar
}
