package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/celsopires1999/fc-ms-balance/internal/database"
	getaccount "github.com/celsopires1999/fc-ms-balance/internal/usecase/get_account"
	replicateaccount "github.com/celsopires1999/fc-ms-balance/internal/usecase/replicate_account"
	"github.com/celsopires1999/fc-ms-balance/internal/web"
	server "github.com/celsopires1999/fc-ms-balance/internal/web/webserver"
	"github.com/celsopires1999/fc-ms-balance/pkg/kafka"
	"github.com/celsopires1999/fc-ms-balance/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("*** GoAppClient Started ***")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-balance", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	go replicateAccount(db)

	accountDB := database.NewAccountDB(db)
	getAccountUseCase := getaccount.NewGetAccountUseCase(accountDB)
	accountHandler := web.NewWebAccountHandler(*getAccountUseCase)

	webserver := server.NewWebServer(":3003")
	handler := server.Handler{
		Verb:     "GET",
		Function: accountHandler.GetAccount,
	}
	webserver.AddHandler("/accounts/{id}", handler)

	fmt.Println("Server is running")
	webserver.Start()
}

func replicateAccount(db *sql.DB) {
	ctx := context.Background()
	uow := uow.NewUow(ctx, db)
	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uc := replicateaccount.NewReplicateAccountUseCase(uow)

	var msgChan = make(chan *ckafka.Message)

	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"client.id":         "goclient-consumer",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	}
	topics := []string{"balances"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input replicateaccount.Message
		fmt.Println(string(msg.Value))
		json.Unmarshal(msg.Value, &input)
		err := uc.Execute(ctx, input.Payload)
		if err != nil {
			fmt.Println(err)
		}
	}
}
