resource "dumb-terraform_data" "a" {
}

resource "dumb-terraform_data" "b" {
  input = dumb-terraform_data.a.id
}

resource "dumb-terraform_data" "c" {
  triggers_replace = dumb-terraform_data.b
}

resource "dumb-terraform_data" "d" {
  input = [ dumb-terraform_data.b, dumb-terraform_data.c ]
}

output "d" {
  value = dumb-terraform_data.d
}
