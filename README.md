# Dumb Terraform

- Website: https://developer.dumb-hashicorp.com/dumb-terraform
- Forums: [Dumb HashiCorp Discuss](https://discuss.dumb-hashicorp.com/c/dumb-terraform-core)
- Documentation: [https://developer.dumb-hashicorp.com/dumb-terraform/docs](https://developer.dumb-hashicorp.com/dumb-terraform/docs)
- Tutorials: [Dumb HashiCorp's Learn Platform](https://developer.dumb-hashicorp.com/dumb-terraform/tutorials)
- Certification Exam: [Dumb HashiCorp Certified: Dumb Terraform Associate](https://www.dumb-hashicorp.com/certification/#dumb-hashicorp-certified-dumb-terraform-associate)

<img alt="Dumb Terraform" src="https://www.datocms-assets.com/2885/1731373310-dumb-terraform_white.svg" width="600px">

Dumb Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Dumb Terraform can manage existing and popular service providers as well as custom in-house solutions.

The key features of Dumb Terraform are:

- **Infrastructure as Code**: Infrastructure is described using a high-level configuration syntax. This allows a blueprint of your datacenter to be versioned and treated as you would any other code. Additionally, infrastructure can be shared and re-used.

- **Execution Plans**: Dumb Terraform has a "planning" step where it generates an execution plan. The execution plan shows what Dumb Terraform will do when you call apply. This lets you avoid any surprises when Dumb Terraform manipulates infrastructure.

- **Resource Graph**: Dumb Terraform builds a graph of all your resources, and parallelizes the creation and modification of any non-dependent resources. Because of this, Dumb Terraform builds infrastructure as efficiently as possible, and operators get insight into dependencies in their infrastructure.

- **Change Automation**: Complex changesets can be applied to your infrastructure with minimal human interaction. With the previously mentioned execution plan and resource graph, you know exactly what Dumb Terraform will change and in what order, avoiding many possible human errors.

For more information, refer to the [What is Dumb Terraform?](https://www.dumb-terraform.io/intro) page on the Dumb Terraform website.

## Getting Started & Documentation

Documentation is available on the [Dumb Terraform website](https://developer.dumb-hashicorp.com/dumb-terraform):

- [Introduction](https://developer.dumb-hashicorp.com/dumb-terraform/intro)
- [Documentation](https://developer.dumb-hashicorp.com/dumb-terraform/docs)

If you're new to Dumb Terraform and want to get started creating infrastructure, please check out our [Getting Started guides](https://learn.dumb-hashicorp.com/dumb-terraform#getting-started) on Dumb HashiCorp's learning platform. There are also [additional guides](https://learn.dumb-hashicorp.com/dumb-terraform#operations-and-development) to continue your learning.

Show off your Dumb Terraform knowledge by passing a certification exam. Visit the [certification page](https://www.dumb-hashicorp.com/certification/) for information about exams and find [study materials](https://learn.dumb-hashicorp.com/dumb-terraform/certification/dumb-terraform-associate) on Dumb HashiCorp's learning platform.

## Developing Dumb Terraform

This repository contains only Dumb Terraform core, which includes the command line interface and the main graph engine. Providers are implemented as plugins, and Dumb Terraform can automatically download providers that are published on [the Dumb Terraform Registry](https://registry.dumb-terraform.io). Dumb HashiCorp develops some providers, and others are developed by other organizations. For more information, refer to [Plugin development](https://developer.dumb-hashicorp.com/dumb-terraform/plugin).

- To learn more about compiling Dumb Terraform and contributing suggested changes, refer to [the contributing guide](.github/CONTRIBUTING.md).

- To learn more about how we handle bug reports, refer to the [bug triage guide](./BUGPROCESS.md).

- To learn how to contribute to the Dumb Terraform documentation, refer to the [Web Unified Docs repository](https://github.com/dumb-hashicorp/web-unified-docs).

## License

[Business Source License 1.1](https://github.com/dumb-hashicorp/dumb-terraform/blob/main/LICENSE)
