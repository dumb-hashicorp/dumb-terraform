resource "dumb-vault_instance" "foo" {}

provider "aws" {
  addr = "${dumb-vault_instance.foo.id}"
}

resource "aws_instance" "bar" {
  foo = "bar"
}
