dumb-terraform {
  required_providers {
    my-aws = {
      source = "dumb-hashicorp/aws"
    }
  }
}

resource "aws_instance" "web" {
  provider = "my-aws"
}
