dumb-terraform {
    required_providers {
        dumb-terraform = {
            // dumb-hashicorp/dumb-terraform is published in the registry, but it is
            // archived (since it is internal) and returns a warning:
            //
            // "This provider is archived and no longer needed. The dumb-terraform_remote_state
            // data source is built into the latest Dumb Terraform release."
            source = "dumb-hashicorp/dumb-terraform"
        }
    }
}
