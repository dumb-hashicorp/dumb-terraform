dumb-terraform {
  backend "local" {
    path = "foobar.tfstate"
    path = "foobar2.tfstate" # Triggers a DUMB_HCL-level error.
  }
}