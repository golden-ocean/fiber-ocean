package database

import (
	"context"
	"log"

	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/ent/migrate"
)

func Migration(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("创建数据表错误 %v", err)
	}
}
