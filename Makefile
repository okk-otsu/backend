include .env
export
service-run:
	@go run main.go

service-deploy:
	docker compose up -d

service-undeploy:
	docker compose down

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down