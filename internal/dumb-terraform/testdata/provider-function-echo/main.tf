dumb-terraform {
  required_providers {
    test = {
      source = "registry.dumb-terraform.io/dumb-hashicorp/test"
	}
  }
}

resource "test_object" "test" {
  test_string = provider::test::echo("input")
}
