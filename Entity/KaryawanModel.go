package Entity

type Karyawan struct {
	Id      uint  `json:"id"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	Alamat   string `json:"alamat"`
}

func (b *Karyawan) TableName() string {
	return "karyawan"
}
