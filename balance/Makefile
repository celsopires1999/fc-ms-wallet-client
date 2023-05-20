migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql-balance:3306)/balance" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql-balance:3306)/balance" -verbose drop
