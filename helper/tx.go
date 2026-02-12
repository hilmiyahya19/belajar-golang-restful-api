package helper // Package helper berisi fungsi-fungsi utilitas yang bisa digunakan di berbagai layer

import "database/sql" // Mengimpor package database/sql untuk menggunakan tipe transaksi (*sql.Tx)

// CommitOrRollback berfungsi untuk melakukan commit jika tidak ada panic, dan rollback jika terjadi panic saat proses transaksi
func CommitOrRollback(tx *sql.Tx) {

	err := recover() // recover() menangkap panic yang terjadi sebelumnya (jika ada)

	if err != nil { // Jika terjadi panic (err tidak nil)
		errorRollback := tx.Rollback() // Melakukan rollback (membatalkan transaksi)
		PanicIfError(errorRollback)    // Jika rollback gagal, maka akan panic
		panic(err)                     // Melempar ulang panic agar tetap terdeteksi di level atas
	} else { // Jika tidak ada panic
		errorCommit := tx.Commit() // Melakukan commit (menyimpan perubahan ke database)
		PanicIfError(errorCommit)  // Jika commit gagal, maka akan panic
	}
}

// Kesimpulan:
// Fungsi ini digunakan untuk menangani transaksi database secara aman dengan pola defer. Jika selama proses transaksi terjadi panic, maka transaksi akan dibatalkan (rollback). Jika tidak ada panic, maka perubahan akan disimpan (commit). Pola ini umum digunakan pada REST API berbasis Golang untuk memastikan konsistensi data dan mencegah database berada dalam kondisi tidak valid ketika terjadi error.