package main

import (
	"context"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"log"
)

func main() {
	ctx := context.Background()
	gmailService, err := gmail.NewService(ctx, option.WithAPIKey(""), option.WithScopes(gmail.GmailReadonlyScope))
	if err != nil {
		log.Fatal(err)
	}

}
