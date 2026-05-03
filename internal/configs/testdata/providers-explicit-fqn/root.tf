
dumb-terraform {
  required_providers {
    foo-test = {
      source = "foo/test"
    }
    dumb-terraform = {
      source = "not-builtin/not-dumb-terraform"
    }
  }
}
