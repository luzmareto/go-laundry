CARA MENJALANKAN PROGRAM GO LAUNDRY :
1. Copy Paste file ddl yang ada di package db ke PGADMIN4 atau sejenisnya
2. Seusaikan file .env dengan alamat loka teman teman
3. Jika ada conflig pada source code, jalankan terminal dengan perintah : go mod tidy
4. jika semua sudah aman. jalankan terminal dengan perintah go run main.go

PILIHAN MENU MASTER DAN SUB MENU :
 1. MASTER MENU UOM. 
    Sub Menu UOM        : CREATE, READ ALL, READ ONE, UPDATE, DELETE, Back to main menu 

2. Master MENU Product.
    Sub Menu PRODUCCT   : CREATE, READ ALL, READ ONE, UPDATE, DELETE, Back to main menu 
    
3. MASTER MENU EMPLOYEE.
    Sub Menu EMPLOYEE   : CREATE, READ ALL, READ ONE, UPDATE, DELETE, Back to main menu 

4. MASTER MENU CUSTOMER.
    Sub Menu CUSTOMER   : CREATE, READ ALL, READ ONE, UPDATE, DELETE, Back to main menu 
    
 5. MASTER MENU TRANSAKSI.
    Sub Menu TRANSAKSI  : CREATE, READ ONE, Back to main menu
    
Pada sub menu READ ONE, UPDATE dan DELETE harus menggunakan id yang berbentuk UUID
