# Terraform plan Dagger Demo

In this folder is a minimal, fully working demo showing how to run a Terraform plan locally using Dagger. This example uses Go for the pipeline code, but the workflow is the same conceptually for other supported languages. The steps are:

1. Create a small Terraform configuration.
2. Write a Go program that uses Dagger to invoke Terraform in a container.
3. Run the pipeline locally using the Dagger CLI.

## Prerequisites
* Go 1.18+ (so we can write the pipeline in Go).
* Dagger CLI installed locally.
    * Installation instructions: [Dagger Quickstart](https://docs.dagger.io/quickstart/)
* Terraform is not required locally (we'll use the official Docker image through Dagger).

## Project Structure
```
├── dagger
│   └── main.go         # Go code for our Dagger pipeline
├── terraform
│   ├── main.tf         # Our sample Terraform config
│   └── terraform.tf     # Terraform providers
└── go.mod              # Go module definition
```
