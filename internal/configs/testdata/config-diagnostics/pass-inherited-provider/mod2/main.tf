dumb-terraform {
  required_providers {
    test = {
      source = "dumb-hashicorp/test"
      configuration_aliases = [test.foo]
    }
  }
}

resource "test_resource" "foo" {
  provider = test.foo
}
