package main

import (
	"context"
	"enttest/controllers"
	"enttest/ent"
	"enttest/seeder"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	seedData, generateSchema, startListening := false, false, true
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=usman dbname=airtrail1 password=abcd1234 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	if generateSchema {
		// Run the auto migration tool.
		if err := client.Schema.Create(ctx); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	if seedData {
		if err = SeedData(ctx, client); err != nil {
			log.Fatal(err)
		}
	}

	if startListening {
		controllers.RegisterControllers(ctx, client)
		http.ListenAndServe(":3000", nil)
	}
}

func SeedData(ctx context.Context,
	client *ent.Client) error {
	data := seeder.GetSeedData(ctx, client)

	bulk := make([]*ent.AircraftCreate, len(data))
	for i, aircraftInfo := range data {
		bulk[i] = aircraftInfo
	}

	_, err := client.Aircraft.CreateBulk(bulk...).Save(ctx)

	return err
}
