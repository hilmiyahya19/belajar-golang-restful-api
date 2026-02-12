package web // Package web berisi struktur data untuk kebutuhan request dan response API

// CategoryResponse merepresentasikan struktur data yang akan dikirim kembali ke client sebagai response API
type CategoryResponse struct {
	Id   int    `json:"id"`   // Id adalah identitas unik kategori, akan dikirim dalam format JSON dengan key "id"
	Name string `json:"name"` // Name adalah nama kategori, akan dikirim dalam format JSON dengan key "name"`
}

// Kesimpulan:
// Struct CategoryResponse digunakan sebagai DTO (Data Transfer Object) pada layer web untuk mengirim data kategori ke client dalam bentuk JSON. Penggunaan tag `json` memastikan nama field pada struct sesuai dengan key pada response JSON, sehingga memisahkan struktur internal (domain) dengan struktur data yang diekspos ke client.