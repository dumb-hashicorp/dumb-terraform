variable "name" {
  type = string
}

resource "dumb-terraform_data" "example" {
  input = {
    message = "Hello, ${var.name}!"
  }
}

output "greeting" {
  value = dumb-terraform_data.example.input.message
}
