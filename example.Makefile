.PHONY: gen_listings clear_listing gen_users gen_application clear_application gen_file_service clear_file_service gen_frontend clear_frontend clear_users migration migrate rollback drop

gen_listings:
	protoc -I=./proto/listings --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/listings ./proto/listings/*.proto

clear_listing:
	rm gen/listings/*.go

gen_users:
	protoc -I=./proto/users --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/users ./proto/users/*.proto

clear_users:
	rm gen/users/*.go

gen_frontend:
	protoc -I=./proto/listings ./proto/listings/*.proto --js_out=import_style=commonjs:./frontend/src/proto/listings --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto/listings
	protoc -I=./proto/users ./proto/users/*.proto --js_out=import_style=commonjs:./frontend/src/proto/users --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/proto/users

clear_frontend:
	rm frontend/src/proto/listings/*.js
	rm frontend/src/proto/users/*.js

gen_application:
	protoc -I=./proto/application --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/application ./proto/application/*.proto

clear_application:
	rm gen/application/*.go

gen_file_service:
	protoc -I=./proto/file_service --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/file_service ./proto/file_service/*.proto

clear_file_service:
	rm gen/file_service/*.go

migration_listings:
	@read -p "Enter migration name: " name; migrate create -ext sql -dir backend/services/listings/database/migrations $$name

migrate_listings:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable

rollback_listings:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable down

drop_listings:
	migrate -source=file://backend/services/listings/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable drop

migration_application:
	@read -p "Enter migration name: " name; migrate create -ext sql -dir backend/services/application/database/migrations $$name

migrate_application:
	migrate -source=file://backend/services/application/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable up

rollback_application:
	migrate -source=file://backend/services/application/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable down

drop_application:
	migrate -source=file://backend/services/application/database/migrations -database postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName?sslmode=disable drop