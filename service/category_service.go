package service // Package service berisi business logic / aturan bisnis aplikasi

import (
	"context" // Digunakan untuk membawa context (timeout, cancel, metadata) antar layer
	"hilmiyahya/belajar-golang-restful-api/model/web" // Mengimpor DTO request & response dari layer web
)

// CategoryService adalah interface yang mendefinisikan kontrak untuk semua operasi bisnis terkait Category
type CategoryService interface {

	// Create digunakan untuk membuat category baru
	// Menerima context dan request dari client, lalu mengembalikan response
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse

	// Update digunakan untuk memperbarui data category
	// Menerima context dan data update, lalu mengembalikan response terbaru
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse

	// Delete digunakan untuk menghapus category berdasarkan Id
	// Tidak mengembalikan data, hanya melakukan proses penghapusan
	Delete(ctx context.Context, categoryId int)

	// FindById digunakan untuk mengambil satu category berdasarkan Id
	// Mengembalikan data category dalam bentuk response
	FindById(ctx context.Context, categoryId int) web.CategoryResponse

	// FindAll digunakan untuk mengambil seluruh data category
	// Mengembalikan slice (array) CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}

// Kesimpulan:
// Interface CategoryService mendefinisikan kontrak business logic untuk fitur Category, seperti Create, Update, Delete, FindById, dan FindAll. Interface ini memisahkan layer controller dari implementasi detail service, sehingga mendukung prinsip clean architecture, dependency inversion, dan memudahkan testing (misalnya dengan mock).