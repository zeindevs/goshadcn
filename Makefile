run: build
	./bin/goshadcnui.exe

build:
	go build -ldflags="-s -w" -o bin/goshadcnui.exe ./cmd/web/main.go

templ:
	templ generate --watch

tailwind:
	tailwindcss -i public/style.scss -o cmd/web/static/style.css -c public/tailwind.config.js --minify --watch

winres-make:
	go-winres make

migrate:
	go run cmd/migrate/main.go

seed:
	go run cmd/seed/main.go

drop:
	go run cmd/drop/main.go
