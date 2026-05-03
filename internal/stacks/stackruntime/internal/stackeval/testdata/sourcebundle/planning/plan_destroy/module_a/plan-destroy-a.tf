
dumb-terraform {
  required_providers {
    test = {
      source = "dumb-terraform.io/builtin/test"

      configuration_aliases = [ test ]
    }
  }
}

resource "test" "foo" {
  for_module = "a"
}

output "result" {
  value = test.foo.result
}
