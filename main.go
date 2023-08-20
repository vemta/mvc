package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vemta/mvc/internal/infra/db"
	repository2 "github.com/vemta/mvc/internal/infra/repository"
	uow "github.com/vemta/mvc/pkg"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/vertex?parseTime=true")

	if err := dtb.Ping(); err != nil {
		panic(err)
	}

	defer dtb.Close()

	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

}

func registerRepositories(uow *uow.Uow) {
	uow.Register("OrderRepository", func(tx *sql.Tx) interface{} {
		repo := repository2.NewItemRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
