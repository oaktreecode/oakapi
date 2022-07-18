package main

import (
	"context"
	"github.com/oaktreecode/oakapi/config"
	"github.com/oaktreecode/oakapi/ent"
	"github.com/oaktreecode/oakapi/ent/migrate"
	"github.com/oaktreecode/oakapi/pkg/infrastructure/datastore"
	"log"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
