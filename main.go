package main

import (
	"context"
	"database/sql"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/vemta/mvc/internal/infra/kafka/consumer"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vemta/mvc/internal/infra/db"
	repository2 "github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/vemta?parseTime=true")

	if err := dtb.Ping(); err != nil {
		panic(err)
	}

	defer dtb.Close()

	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

	var topics = []string{"create_order", "create_item", "create_discount_rule"}
	msgChan := make(chan *kafka.Message)
	go consumer.Consume(topics, "host.docker.internal:9094", msgChan)
	consumer.ProcessEvents(ctx, msgChan, uow)

}

func registerRepositories(uow *uow.Uow) {
	uow.Register("OrderRepository", func(tx *sql.Tx) interface{} {
		repo := repository2.NewItemRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
