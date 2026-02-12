package web // Package web berisi struktur umum untuk request dan response pada REST API

// WebResponse adalah struktur response standar yang digunakan untuk membungkus semua response API agar konsisten
type WebResponse struct {
	Code   int    `json:"code"`   
	// Code berisi kode status (biasanya HTTP status code seperti 200, 400, 404, 500)
	// Tag `json:"code"` memastikan field ini dikirim dalam JSON dengan key "code"

	Status string `json:"status"` 
	// Status berisi keterangan singkat (misalnya: "OK", "BAD REQUEST", "NOT FOUND")
	// Akan dikirim dalam JSON dengan key "status"

	Data   any    `json:"data"`   
	// Data berisi payload utama response (bisa berupa object, array, atau null)
	// Tipe `any` (alias dari interface{}) memungkinkan menampung berbagai jenis data
	// Akan dikirim dalam JSON dengan key "data"
}

// Kesimpulan:
// Struct WebResponse digunakan sebagai format response standar agar seluruh endpoint memiliki struktur JSON yang konsisten (Code, Status, dan Data). Penggunaan tipe `any` membuat field Data fleksibel untuk menampung berbagai jenis response, sehingga memudahkan pengelolaan API dan menjaga konsistensi komunikasi antara server dan client.