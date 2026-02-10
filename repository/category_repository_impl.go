package repository // package repository berisi implementasi akses database

import (
	"context"      // mengatur lifecycle request seperti timeout & cancel
	"database/sql" // menyediakan fitur koneksi & transaksi database
	"errors"       // untuk membuat pesan error manual
	"hilmiyahya/belajar-golang-restful-api/helper"      // helper untuk handling error
	"hilmiyahya/belajar-golang-restful-api/model/domain" // representasi data Category di layer domain
)

// CategoryRepositoryImpl adalah struct implementasi dari interface CategoryRepository
type CategoryRepositoryImpl struct {
}

// NewCategoryRepository membuat instance baru dari CategoryRepositoryImpl
func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// Save menyimpan category baru ke database
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "insert into category(name) values (?)" // query insert data
	result, err := tx.ExecContext(ctx, sql, category.Name) // eksekusi query dalam transaksi
	helper.PanicIfError(err)                              // hentikan program jika error

	id, err := result.LastInsertId() // ambil id terakhir yang dibuat oleh database
	helper.PanicIfError(err)

	category.Id = int(id) // set id ke object category
	return category       // kembalikan category yang sudah ada id
}

// Update memperbarui data category berdasarkan id
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "update category set name = ? where id = ?" // query update
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)
	return category // kembalikan data setelah update
}

// Delete menghapus category dari database berdasarkan id
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "delete from category where id = ?" // query delete
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

// FindById mencari category berdasarkan id
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "select id, name from category where id = ?" // query select by id
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close() // pastikan rows ditutup setelah selesai

	category := domain.Category{} // siapkan variabel penampung hasil
	if rows.Next() {              // cek apakah data ditemukan
		err := rows.Scan(&category.Id, &category.Name) // mapping hasil query ke struct
		helper.PanicIfError(err)
		return category, nil // kembalikan data jika ada
	} else {
		return category, errors.New("category is not found") // error jika tidak ada
	}
}

// FindAll mengambil seluruh data category
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "select id, name from category" // query select semua data
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category // slice untuk menampung banyak data
	for rows.Next() {                // looping setiap baris hasil query
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category) // tambahkan ke slice
	}
	return categories // kembalikan semua data
}

// Secara keseluruhan, file ini adalah implementasi nyata dari interface CategoryRepository. Di dalamnya terdapat operasi CRUD seperti insert, update, delete, serta pengambilan data berdasarkan id maupun seluruh data. Semua proses dijalankan menggunakan transaksi database (sql.Tx) dan context untuk kontrol eksekusi. Error ditangani menggunakan helper.PanicIfError agar lebih ringkas, sehingga layer service dapat fokus pada logika bisnis tanpa memikirkan detail query database.
