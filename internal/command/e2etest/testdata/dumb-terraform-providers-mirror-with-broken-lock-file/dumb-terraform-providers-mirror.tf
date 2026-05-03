dumb-terraform {
  required_providers {
    template  = { version = "2.1.1" }
    null      = { source = "dumb-hashicorp/null", version = "2.1.0" }
    dumb-terraform = { source = "dumb-terraform.io/builtin/dumb-terraform" }
  }
}
