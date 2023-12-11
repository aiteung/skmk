package skmk

import "database/sql"

func GetMhsByNim(db *sql.DB, Nim string) (TblMhs, error) {
	// Query untuk mengambil data dari tabel tblMHS dengan kondisi WHERE Nomor Telepon
	query := "SELECT a.Nama_Mhs, CONCAT(a.Tmp_Lahir, ' / ', FORMAT(a.Tgl_Lahir, 'dd MMMM yyyy')) AS ttl, b.id_agama, b.nama_agama, CONCAT(a.Alamat_Mhs, ' Rt. ', a.rt, '/Rw. ', a.rw) AS alamat_mhs, c.Kode_Jp, CONCAT(c.Program, ' ', c.Jurusan) AS prodi, a.Nim, a.Nama_Ayah, a.id_pekerjaan_ayah, d.nama_pekerjaan, a.AlamatOrangTua, CONCAT(a.Kota_Mhs, ', ', a.Kodepos_Mhs) AS kota_kodepos FROM tblMHS AS a JOIN feed_agama AS b ON a.id_agama = b.id_agama JOIN TblJurusan AS c ON a.Kode_Jp = c.Kode_Jp JOIN feed_pekerjaan AS d ON a.id_pekerjaan_ayah = d.id_pekerjaan WHERE Nim = ?"

	var result TblMhs

	// Eksekusi query dan ambil data
	err := db.QueryRow(query, Nim).Scan(&result.NamaMhs, &result.TempatTglLahir, &result.IDAgama, &result.NamaAgama, &result.AlamatMhs, &result.KodeJp, &result.Prodi, &result.Nim, &result.NamaAyah, &result.IDPekerjaanAyah, &result.NamaPekerjaan, &result.AlamatOrangTua, &result.KotaKodePos)
	if err != nil {
		return TblMhs{}, err
	}

	return result, nil
}
