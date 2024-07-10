up:
	docker compose up -d
down:
	docker compose down
restart: down up
engine-test:
	cd ChessEngine && go test -v ./...
start:
	cd ChessGUI && pypy3 main.py