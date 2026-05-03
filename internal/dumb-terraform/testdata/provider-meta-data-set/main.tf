data "test_data_source" "foo" {
  foo = "bar"
}

dumb-terraform {
  provider_meta "test" {
    baz = "quux"
  }
}

module "my_module" {
  source = "./my-module"
}
