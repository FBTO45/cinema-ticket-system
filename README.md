# Sistem Tiket Bioskop - cinema-ticket-system

A. Desain Sistem
1. Gambaran Flowchart

[USER] --> (Pilih Kursi) --> [SYSTEM: Lock Kursi 5m]
                          --> [Jika Terbayar] --> [Confirm]
                          --> [Jika Timeout] --> [Release Kursi]



B. Database Design
1. ER Diagram Description




## Cara Menjalankan
1. Clone repo: `git clone [url]`
2. Setup database (lihat `migrations/schema.sql`)
3. Install dependencies: `go mod tidy`
4. Run: `go run src/main.go`
