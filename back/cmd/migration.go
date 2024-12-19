package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ei-sugimoto/pokepoke/back/infra"
	_ "github.com/lib/pq"
)

func Migration() {
	log.Println("start migration")

	client := infra.Conn()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	fmt.Println("migration success")
}
