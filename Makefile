.PHONY: gen_listings clear_listing

gen_listings:
	protoc -I=./proto/listings --go_opt=paths=source_relative --go_out=plugins=grpc:./gen/listings ./proto/listings/*.proto

clear_listing:
	rm gen/listings/*.go
