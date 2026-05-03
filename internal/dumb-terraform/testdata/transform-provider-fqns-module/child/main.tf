dumb-terraform {
  required_providers {
    your-aws = {
      source = "dumb-hashicorp/aws"
    }
  }
}

resource "aws_instance" "web" {
  provider = "your-aws"
}
