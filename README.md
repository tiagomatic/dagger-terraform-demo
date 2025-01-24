# Terraform Plan Dagger Demo

In this folder is a minimal, fully working demo showing how to run a Terraform plan locally using Dagger. This example uses Go for the pipeline code, but the workflow is the same conceptually for other supported languages. The steps are:

1. Create a small Terraform configuration.
2. Write a Go program that uses Dagger to invoke Terraform in a container.
3. Run the pipeline locally using the Dagger CLI.

## Prerequisites
* Go 1.18+ (so we can write the pipeline in Go)
* [Docker Desktop](https://www.docker.com/) or [Podman](https://podman.io/) installed and running
* Dagger CLI installed locally
    * Installation instructions: [Dagger Quickstart](https://docs.dagger.io/quickstart/)
* Terraform is not required locally. We'll run it with a Docker image through Dagger.

## Project Structure
```
├── dagger
│   └── main.go         # Go code for our Dagger pipeline
├── terraform
│   ├── main.tf         # Our sample Terraform config
│   └── terraform.tf    # Terraform providers
└── go.mod              # Go module definition
```

## Running Terraform Pipeline
To run this Dagger pipeline, go to the root of this module and run the following command:
```
dagger run go run dagger/main.go -tfImage=hashicorp/terraform:1.4.0
```

The initial run will take some time at first, but subsequent runs will have cached images and run much faster.
