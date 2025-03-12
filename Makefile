.PHONY: run server migrate rollback build clean

# Menjalankan service utama
run-user:
	go run cmd/app/user/*.go

# Menjalankan migrasi database (up)
migrate:
	go run cmd/migration/*.go -type=run

# Menjalankan rollback migrasi (down)
rollback:
	go run cmd/migration/*.go -type=rollback

# Build binary untuk service utama dan migrasi
build-user:
	go build -o bin/serve cmd/app/user/*.go

# Membersihkan binary build
clean:
	rm -f bin/serve

compose-up:
	docker-compose up -d