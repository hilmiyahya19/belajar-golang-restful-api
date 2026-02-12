package helper // Mendefinisikan package helper sebagai tempat fungsi-fungsi utilitas / konversi data

import (
	"hilmiyahya/belajar-golang-restful-api/model/domain" // Mengimpor package domain (biasanya representasi entity/database)
	"hilmiyahya/belajar-golang-restful-api/model/web"    // Mengimpor package web (biasanya untuk response/request API)
)

// ToCategoryResponse berfungsi untuk mengubah data dari domain.Category menjadi web.CategoryResponse (format yang dikirim ke client)
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,   // Mengisi field Id pada response dari Id milik domain
		Name: category.Name, // Mengisi field Name pada response dari Name milik domain
	}
}

// ToCategoryResponses berfungsi untuk mengubah slice (array) domain.Category menjadi slice web.CategoryResponse
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse // Membuat slice kosong untuk menampung hasil konversi

	// Melakukan perulangan pada setiap data category di dalam slice categories
	for _, category := range categories {
		// Menambahkan hasil konversi satu per satu ke dalam slice categoryResponses
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses // Mengembalikan slice hasil konversi
}

// Kesimpulan:
// Kode ini berfungsi sebagai mapper/helper untuk mengonversi data dari layer domain (entity/database) ke layer web (response API). Fungsi pertama mengubah satu objek Category, sedangkan fungsi kedua mengubah banyak objek (slice) Category. Pola ini umum digunakan dalam arsitektur REST API untuk memisahkan struktur data internal dengan struktur data yang dikirim ke client (separation of concerns).