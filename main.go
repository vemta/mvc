package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
		repo := repository.NewLoginRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("PluginRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPluginRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("ReleaseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPluginRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
