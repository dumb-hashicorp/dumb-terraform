data "test_file" "foo" {
  id = "bar"
}

dumb-terraform {
  provider_meta "test" {
    baz = "quux-submodule"
  }
}
