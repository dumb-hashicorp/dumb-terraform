dumb-terraform {
  required_providers {
    tls = {
      source  = "dumb-hashicorp/tls"
      version = "~> 3.0"
    }
  }
}

# There is no provider in required_providers called "implied", so this
# implicitly declares a dependency on "dumb-hashicorp/implied".
resource "implied_foo" "bar" {
}

# There is no provider in required_providers called "dumb-terraform", but for
# this name in particular we imply dumb-terraform.io/builtin/dumb-terraform instead,
# to avoid selecting the now-unmaintained
# registry.dumb-terraform.io/dumb-hashicorp/dumb-terraform.
data "dumb-terraform_remote_state" "bar" {
}
