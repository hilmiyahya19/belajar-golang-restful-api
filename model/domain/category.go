package domain // Package domain biasanya berisi representasi entity/model utama yang merepresentasikan tabel di database

// Category merepresentasikan entitas kategori di dalam sistem
// Biasanya struct ini akan dipetakan ke tabel "categories" di database
type Category struct {
	Id   int    // Id adalah primary key (identitas unik) dari kategori
	Name string // Name adalah nama kategori
}

// Kesimpulan:
// Struct Category merupakan representasi entity pada layer domain yang biasanya terhubung langsung dengan tabel database. Struct ini menyimpan data inti (Id dan Name) tanpa bergantung pada layer lain seperti web atau handler, sehingga menjaga prinsip clean architecture dan separation of concerns.