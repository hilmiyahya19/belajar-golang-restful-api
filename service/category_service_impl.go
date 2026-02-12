package service // Package service berisi implementasi business logic aplikasi

import (
	"context" // Untuk membawa context (timeout, cancel, metadata) antar layer
	"database/sql" // Untuk menggunakan koneksi database dan transaksi
	"github.com/go-playground/validator/v10" // Library untuk validasi struct
	"hilmiyahya/belajar-golang-restful-api/helper" // Berisi helper seperti panic handler & mapper
	"hilmiyahya/belajar-golang-restful-api/model/domain" // Entity/domain model (representasi tabel)
	"hilmiyahya/belajar-golang-restful-api/model/web" // DTO request & response
	"hilmiyahya/belajar-golang-restful-api/repository" // Layer repository untuk akses database
)

// CategoryServiceImpl adalah implementasi dari interface CategoryService
type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository // Dependency repository untuk akses data
	DB                 *sql.DB                       // Koneksi database
	Validate           *validator.Validate           // Validator untuk validasi request
}

// Constructor untuk membuat instance CategoryServiceImpl
// Menggunakan dependency injection agar mudah di-test dan fleksibel
func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

// Create digunakan untuk membuat category baru
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request) // Validasi request sesuai tag validate
	helper.PanicIfError(err)                // Jika error, akan panic

	tx, err := service.DB.Begin() // Memulai transaksi database
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx) // Otomatis commit jika sukses / rollback jika panic

	category := domain.Category{
		Name: request.Name, // Mapping dari request ke domain model
	}

	category = service.CategoryRepository.Save(ctx, tx, category) // Simpan ke database

	return helper.ToCategoryResponse(category) // Konversi ke response dan kembalikan
}

// Update digunakan untuk memperbarui data category
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {

	err := service.Validate.Struct(request) // Validasi request
	helper.PanicIfError(err)

	tx, err := service.DB.Begin() // Mulai transaksi
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id) // Ambil data lama
	helper.PanicIfError(err)

	category.Name = request.Name // Update field yang ingin diubah

	category = service.CategoryRepository.Update(ctx, tx, category) // Simpan perubahan

	return helper.ToCategoryResponse(category) // Kembalikan response
}

// Delete digunakan untuk menghapus category berdasarkan Id
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {

	tx, err := service.DB.Begin() // Mulai transaksi
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId) // Pastikan data ada
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category) // Hapus data
}

// FindById digunakan untuk mengambil satu category berdasarkan Id
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {

	tx, err := service.DB.Begin() // Mulai transaksi
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId) // Ambil data
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category) // Konversi ke response
}

// FindAll digunakan untuk mengambil seluruh data category
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := service.DB.Begin() // Mulai transaksi
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx) // Ambil semua data

	return helper.ToCategoryResponses(categories) // Konversi slice domain ke slice response
}

// Kesimpulan:
// Kode ini merupakan implementasi business logic untuk fitur Category dalam arsitektur REST API. Service bertugas melakukan validasi request, mengatur transaksi database (commit/rollback), memanggil repository untuk akses data, serta mengonversi domain model ke response DTO. Pola ini mendukung clean architecture karena memisahkan controller, service, dan repository, serta menjaga konsistensi transaksi dan validasi sebelum data diproses ke database.