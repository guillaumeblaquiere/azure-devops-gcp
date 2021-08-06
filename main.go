package main

import (
	functions "cloud.google.com/go/functions/apiv1"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"
)

func main() {
	ctx := context.Background()
	f, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	it := f.ListFunctions(ctx, &functionspb.ListFunctionsRequest{
		Parent: "projects/gdglyon-cloudrun/locations/us-central1",
	})
	for {
		i, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i.Name)
	}
}
