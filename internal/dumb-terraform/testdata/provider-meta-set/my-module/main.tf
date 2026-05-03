resource "test_resource" "bar" {
  value = "bar"
}

dumb-terraform {
  provider_meta "test" {
    baz = "quux-submodule"
  }
}
