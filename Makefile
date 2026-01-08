export MYSQL_URL='mysql://root@tcp(localhost:3306)/simple_forum'

migrate-create:
	@ migrate create -ext sql -dir scripts/migration -seq $(name)

migrate-up:
	@ migrate -database $(MYSQL_URL) -path scripts/migration up

migrate-down:
	@ migrate -database $(MYSQL_URL) -path scripts/migration down