package skmk

func MessageLengkapiData() string {
	msg := "*SKMK*\n"
	msg = msg + "Data kakak di siap belum lengkap nihh. Kakak lengkapi dulu ya, kayak profil kakak nya, pekerjaan orang tua kakak nya. Kalau udah, kakak boleh minta SKMK lagi sama akuu, maaciwww...."
	return msg
}

func MessageBerhasilMintaSkmk(mhs TblMhs) string {
	msg := "*SKMK*\n"
	msg = msg + "SKMK Lagi dikirim sama iTeung ke email " + mhs.Email + ", di cek aja ya kak..."
	return msg
}

func MessageGagalMintaSkmk() string {
	msg := "*SKMK*\n"
	msg = msg + "SKMK gagal dikirim sama iTeung di email, coba kakak hubungi TIK yaa..."
	return msg
}
