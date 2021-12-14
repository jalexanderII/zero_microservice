.PHONY: gen_listings clear_listing sqlc migration migrate rollback drop

gen_listings:
	protoc -I=./proto/listings --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/listings ./proto/listings/*.proto

clear_listing:
	rm gen/listings/*.go

sqlc:
	sqlc generate

migration:
	@read -p "Enter migration name: " name; migrate create -ext sql -dir backend/services/listings/database/migrations $$name

migrate:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://postgres:Qnacmg797y@127.0.0.1:5432/postgres?sslmode=disable up

rollback:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://postgres:Qnacmg797y@127.0.0.1:5432/postgres?sslmode=disable down

drop:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://postgres:Qnacmg797y@127.0.0.1:5432/postgres?sslmode=disable drop