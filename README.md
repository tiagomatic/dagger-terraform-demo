# Terraform Plan Dagger Demo

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
.
├── .github
│   └── workflows
├── dagger
│   ├── json-lint      # Concourse example
│   │   └── main.go    
│   ├── tf-plan        # Terraform plan example
│   │   └── main.go   
│   └── unit-test      # Unit tests example
│       └── main.go
Actions workflow
├── terraform
│   ├── main.tf        # A simple Terraform config
│   └── terraform.tf   # Terraform providers
└── go.mod             # Go module definition
```


## Running Locally

Below are instructions for running each example pipeline in this repo

### Running the Go Test Pipeline
The Go test pipeline runs go test ./... inside a container.

```bash
dagger run go run dagger/unit-test/main.go
```

If your tests pass, the pipeline will print a success message.

### Running the Terraform Pipeline
The Terraform pipeline initializes and plans a simple Terraform project.

```bash
dagger run go run dagger/tf-plan/main.go
```

You’ll see Terraform output (e.g. “Terraform has been successfully initialized!”). If the plan is successful, you’ll see no errors.


### Running the JSON Lint Pipeline


The JSON lint pipeline scans for all *.json files (excluding .git) and checks whether they’re valid JSON.

```bash
dagger run go run dagger/lint-json/main.go
```

For each valid JSON file, you’ll see a checkmark. If any file is invalid, the pipeline exits with a non-zero status, and you’ll see the parse error.
