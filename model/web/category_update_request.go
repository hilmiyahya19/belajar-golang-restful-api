package web // Package web berisi struktur request dan response yang digunakan pada layer HTTP/API

// CategoryUpdateRequest merepresentasikan struktur data yang diterima
// dari client saat melakukan update data category
type CategoryUpdateRequest struct {
	Id   int    `validate:"required"` 
	// Id adalah identitas kategori yang akan diupdate
	// Tag `validate:"required"` memastikan Id wajib diisi saat proses validasi

	Name string `validate:"required,min=1,max=200" json:"name"` 
	// Name adalah nama kategori terbaru yang akan menggantikan data lama
	// Tag `validate`:
	// - required  -> wajib diisi
	// - min=1     -> minimal 1 karakter
	// - max=200   -> maksimal 200 karakter
	// Tag `json:"name"` untuk mapping field ini dengan key "name" pada JSON request body
}

// Kesimpulan:
// Struct CategoryUpdateRequest digunakan sebagai DTO pada layer web untuk menerima data update kategori dari client. Field Id digunakan untuk menentukan data mana yang akan diperbarui, sedangkan Name adalah nilai barunya. Validasi diterapkan melalui tag struct untuk memastikan data yang masuk sesuai aturan sebelum diproses lebih lanjut ke layer service atau repository.