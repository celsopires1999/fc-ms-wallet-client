migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/wallet" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/wallet" -verbose drop
