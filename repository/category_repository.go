package repository // package repository berisi kontrak / interface untuk akses data ke database

import (
	"context"      // digunakan untuk mengatur lifecycle request (timeout, cancel, dll)
	"database/sql" // menyediakan interaksi dengan database SQL termasuk transaksi
	"hilmiyahya/belajar-golang-restful-api/model/domain" // import struct domain Category
)

// CategoryRepository adalah interface yang mendefinisikan operasi CRUD untuk Category
type CategoryRepository interface {
	// Save menyimpan data category baru ke database dan mengembalikan data yang sudah tersimpan
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category

	// Update memperbarui data category yang sudah ada di database
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category

	// Delete menghapus data category dari database berdasarkan data yang diberikan
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)

	// FindById mencari category berdasarkan id, mengembalikan error jika tidak ditemukan
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)

	// FindAll mengambil seluruh data category yang ada di database
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

// Secara keseluruhan, kode ini mendefinisikan kontrak (interface) untuk operasi database pada entitas Category. Interface ini memaksa setiap implementasi repository memiliki fungsi standar seperti simpan, update, hapus, ambil berdasarkan id, dan ambil semua data. Dengan adanya interface, logic bisnis dapat bergantung pada abstraksi tanpa perlu tahu detail implementasi database yang digunakan.
