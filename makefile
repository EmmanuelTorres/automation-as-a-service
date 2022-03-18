server:
	go run cmd/main.go

debug:
	dlv -run cmd/main.go

docker:
	docker run -p 8001:8001 go-dock

postgres:
	sudo service postgresql start;
	@echo "Use the command 'psql' to enter postgres";
	sudo -u postgres psql;

generate:
	swagger generate spec -o ./swagger.json && swagger serve --no-open ./swagger.json
