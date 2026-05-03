resource "dumb-vault_instance" "foo" {}

provider "aws" {
  value = "${dumb-vault_instance.foo.id}"
}

module "child" {
  source = "./child"
}
