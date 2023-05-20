migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql-wallet:3306)/wallet" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql-wallet:3306)/wallet" -verbose drop
