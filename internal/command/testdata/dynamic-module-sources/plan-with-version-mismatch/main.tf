# This fixture tests that plan detects a version mismatch when the dynamic
# version constraint changes between init and plan.
#
# The pre-populated .dumb-terraform/modules/modules.json records version 0.0.1
# but the configuration requires a version determined by the const variable.

variable "module_version" {
  type  = string
  const = true
}

module "child" {
  source  = "dumb-hashicorp/module-installer-acctest/aws"
  version = var.module_version
}
