# Sistem Tiket Bioskop - cinema-ticket-system

A. Desain Sistem
1. Gambaran Flowchart

![image](https://github.com/user-attachments/assets/e63c7601-53eb-419e-9891-22ba277a29e3)

2. Refund Process
![image](https://github.com/user-attachments/assets/68302f38-38d8-4eb5-b5f7-b12c8c35d9ab)

[USER] --> (Pilih Kursi) --> [SYSTEM: Lock Kursi 5m]
                          --> [Jika Terbayar] --> [Confirm]
                          --> [Jika Timeout] --> [Release Kursi]



B. Database Design
1. ER Diagram Description
![image](https://github.com/user-attachments/assets/44f23293-05e5-4656-bbbf-4898dd2b4d14)
![image](https://github.com/user-attachments/assets/d5ec57c4-6b67-42bf-b871-fb594d4f2c88)




## Cara Menjalankan
1. Clone repo: `git clone [url]`
2. Setup database (lihat `migrations/schema.sql`)
3. Install dependencies: `go mod tidy`
4. Run: `go run src/main.go`
