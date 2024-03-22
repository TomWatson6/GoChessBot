up:
	docker compose up -d
down:
	docker compose down
restart: down up
engine-test:
	cd ChessEngine && go test -v ./...
start: up
	cd chess_ai && cargo run --release

