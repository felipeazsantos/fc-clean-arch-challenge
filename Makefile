migratecreate:
	migrate create -ext sql -dir internal/infra/database/migrations -seq init_schema

migrateup:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3307)/orders" up

migratedown:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3307)/orders" down

migratestatus:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3307)/orders" status

migratereset:
	migrate -path internal/infra/database/migrations -database "mysql://root:root@tcp(localhost:3307)/orders" reset

.PHONY: migratecreate migrateup migratedown migratestatus migratereset