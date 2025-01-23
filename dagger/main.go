package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	tfImageFlag := flag.String("tfImage", "", "REQUIRED: Terraform Docker image to use")

	flag.Parse()

	if *tfImageFlag == "" {
		log.Fatal("Error: -tfImage flag is required. Example: -tfImage=hashicorp/terraform:1.4.0")
	}

	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		log.Fatalf("Failed to connect to Dagger: %v", err)
	}
	defer client.Close()

	// 1. Pull Terraform image + clear entrypoint
	tfContainer := client.Container().
		From(*tfImageFlag).
		WithEntrypoint([]string{})

	// 2. Mount your local terraform/ folder into /src
	tfContainer = tfContainer.
		WithDirectory("/src", client.Host().Directory("./terraform")).
		WithWorkdir("/src")

	// -- Step A: terraform providers lock -platform=linux_amd64 --
	lockContainer := tfContainer.WithExec([]string{
		"terraform", "providers", "lock",
		"-platform=linux_amd64",
	})

	lockOutput, err := lockContainer.Stdout(ctx)
	if err != nil {
		log.Fatalf("terraform providers lock failed: %v", err)
	}
	fmt.Println("==== providers lock output ====")
	fmt.Println(lockOutput)
	fmt.Println("==== end of providers lock output ====")

	// -- Step B: terraform init --
	initContainer := lockContainer.WithExec([]string{
		"terraform", "init", "-upgrade",
	})

	initOutput, err := initContainer.Stdout(ctx)
	if err != nil {
		log.Fatalf("terraform init failed: %v", err)
	}
	fmt.Println("==== terraform init output ====")
	fmt.Println(initOutput)
	fmt.Println("==== end of init output ====")

	// -- Step C: terraform plan --
	planContainer := initContainer.WithExec([]string{
		"terraform", "plan",
	})

	planOutput, err := planContainer.Stdout(ctx)
	if err != nil {
		log.Fatalf("terraform plan failed: %v", err)
	}
	fmt.Println("==== terraform plan output ====")
	fmt.Println(planOutput)
	fmt.Println("==== end of plan output ====")
}
