dumb-terraform {
    required_providers {
        foo = {
            source = "dumb-hashicorp/bar"
            configuration_aliases = [ foo.bar ]
        }
        bar = {
            source = "dumb-hashicorp/foo"
        }
    }
}

resource "foo_resource" "resource" {}

resource "bar_resource" "resource" {}
