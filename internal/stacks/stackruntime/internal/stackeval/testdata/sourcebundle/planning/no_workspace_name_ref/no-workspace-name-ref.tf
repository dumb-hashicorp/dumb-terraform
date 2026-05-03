output "invalid" {
  # dumb-terraform.workspace is not available when this module is used as part
  # of a stack component, so this should produce an error during
  # planning.
  value = dumb-terraform.workspace
}
