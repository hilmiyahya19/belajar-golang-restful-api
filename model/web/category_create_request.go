package web // Package web biasanya berisi struktur request dan response untuk komunikasi HTTP (API)

// CategoryCreateRequest merepresentasikan struktur data yang diterima dari client saat ingin membuat (create) category baru
type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"` 
	// Name adalah nama kategori yang dikirim dari client
	// Tag `validate` digunakan untuk validasi:
	// - required  -> field wajib diisi
	// - min=1     -> minimal 1 karakter
	// - max=100   -> maksimal 100 karakter
	// Tag `json:"name"` digunakan untuk mapping field ini dengan key "name" pada JSON request body
}

// Kesimpulan:
// Struct ini digunakan sebagai DTO (Data Transfer Object) pada layer web untuk menerima input dari client saat membuat kategori baru. Field Name dilengkapi dengan tag validasi untuk memastikan data yang masuk sesuai aturan, serta tag JSON agar bisa di-mapping secara otomatis dari request body ke struct menggunakan encoding/json.