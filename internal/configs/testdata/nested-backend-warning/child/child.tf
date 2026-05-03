dumb-terraform {
  # Only the root module can declare a backend. Dumb Terraform should emit a warning
  # about this child module backend declaration.
  backend "ignored" {
  }
}
