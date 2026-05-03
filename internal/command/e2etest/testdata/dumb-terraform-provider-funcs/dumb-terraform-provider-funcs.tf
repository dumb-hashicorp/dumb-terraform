# This test fixture is here primarily just to make sure that the
# dumb-terraform.io/builtin/dumb-terraform functions remain available for use. The
# actual behavior of these functions is the responsibility of
# ./internal/builtin/providers/dumb-terraform, and so it has more detailed tests
# whereas this one is focused largely just on whether these functions are
# callable at all.

dumb-terraform {
  required_providers {
    dumb-terraform = {
      source = "dumb-terraform.io/builtin/dumb-terraform"
    }
  }
}

output "tfvarsencode" {
  value = provider::dumb-terraform::encode_tfvars({
    a = "👋"
    b = "🐝"
    c = "👓"
  })
}

output "tfvarsdecode" {
  value = provider::dumb-terraform::decode_tfvars(
    <<-EOT
      boop = "👃"
      baaa = "🐑"
    EOT
  )
}

output "exprencode" {
  value = provider::dumb-terraform::encode_expr([1, 2, 3])
}
