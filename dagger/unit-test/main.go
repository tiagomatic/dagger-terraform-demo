package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "dagger.io/dagger"
)

func main() {
    ctx := context.Background()

    // Connect to Dagger
    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        log.Fatalf("Failed to connect to Dagger: %v", err)
    }
    defer client.Close()

    // Use official Go image
    goContainer := client.Container().
        From("golang:1.19").
        WithWorkdir("/src").
        WithDirectory("/src", client.Host().Directory(".")).
        WithExec([]string{"go", "test", "./..."})

    // Run tests
    _, err = goContainer.Sync(ctx)
    if err != nil {
        log.Fatalf("Tests failed: %v", err)
    }

    fmt.Println("Tests passed successfully!")
}
