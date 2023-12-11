package skmk

type TblMhs struct {
	NamaMhs         string `json:"Nama_Mhs"`
	TempatTglLahir  string `json:"ttl"`
	IDAgama         string `json:"id_agama"`
	NamaAgama       string `json:"nama_agama"`
	AlamatMhs       string `json:"alamat_mhs"`
	KodeJp          string `json:"Kode_Jp"`
	Prodi           string `json:"prodi"`
	Nim             string `json:"Nim"`
	NamaAyah        string `json:"Nama_Ayah"`
	IDPekerjaanAyah string `json:"id_pekerjaan_ayah"`
	NamaPekerjaan   string `json:"nama_pekerjaan"`
	AlamatOrangTua  string `json:"AlamatOrangTua"`
	KotaKodePos     string `json:"kota_kodepos"`
}
