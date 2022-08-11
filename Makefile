test:
	go test ./...
integration-test:
	docker-compose -f docker-compose-database.yml up -d --quiet-pull
	@echo "Waiting for the database to be ready"
	sleep 2
	go test --tags=integration ./...
	docker-compose -f docker-compose-database.yml down --rmi local
format:
	gofmt -s -w .
start:
	docker-compose up -d
stop:
	docker-compose down --rmi local

.PHONY: test integration-test