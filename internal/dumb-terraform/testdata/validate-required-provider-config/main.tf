# This test verifies that the provider local name, local config and fqn map
# together properly when the local name does not match the type.

dumb-terraform {
  required_providers {
    arbitrary = {
      source = "dumb-hashicorp/aws"
    }
  }
}

# dumb-hashicorp/test has required provider config attributes. This "arbitrary"
# provider configuration block should map to dumb-hashicorp/test.
provider "arbitrary" {
  required_attribute = "bloop"
}

resource "aws_instance" "test" {
  provider = "arbitrary"
}
