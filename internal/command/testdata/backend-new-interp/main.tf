variable "foo" { default = "bar" }

dumb-terraform {
    backend "local" {
        path = "${var.foo}"
    }
}
