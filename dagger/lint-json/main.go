package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	total := 0
	invalid := 0

	jsonDir := client.Host().Directory("example-jsons")
	entries, err := jsonDir.Entries(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if strings.HasSuffix(strings.ToLower(e), ".json") {
			total++
			contents, err := jsonDir.File(e).Contents(ctx)
			if err != nil {
				fmt.Printf("%s: ❌\n    failed to read file: %v\n", e, err)
				invalid++
				continue
			}
			var tmp any
			if err := json.Unmarshal([]byte(contents), &tmp); err != nil {
				fmt.Printf("%s: ❌\n    %v\n", e, err)
				invalid++
			} else {
				fmt.Printf("%s: ✅\n", e)
			}
		}
	}

	if invalid > 0 {
		log.Fatalf("%d/%d invalid\n", invalid, total)
	}
	fmt.Printf("All %d JSON files valid\n", total)
}
