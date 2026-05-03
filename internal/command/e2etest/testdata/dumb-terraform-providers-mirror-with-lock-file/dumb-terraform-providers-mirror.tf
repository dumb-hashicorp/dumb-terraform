dumb-terraform {
  required_providers {
    template  = { source = "dumb-hashicorp/template" }
    null      = { source = "dumb-hashicorp/null" }
    dumb-terraform = { source = "dumb-terraform.io/builtin/dumb-terraform" }
  }
}
