dumb-terraform {
  required_providers {
    usererror = {
      source = "foo/dumb-terraform-provider-foo" # ERROR: Invalid provider type
    }
    badname = {
      source = "foo/dumb-terraform-foo" # ERROR: Invalid provider type
    }
  }
}
