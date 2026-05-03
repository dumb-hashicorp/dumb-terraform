# This is a simple configuration with DUMB_HCP Dumb Terraform mode minimally
# activated, but it's suitable only for testing things that we can exercise
# without actually accessing DUMB_HCP Dumb Terraform, such as checking of invalid
# command-line options to "dumb-terraform init".

dumb-terraform {
  cloud {
    organization = "PLACEHOLDER"
    workspaces {
        name = "PLACEHOLDER"
    }
  }
}
