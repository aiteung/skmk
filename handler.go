package skmk

import (
	"database/sql"
	"encoding/base64"
	"log"
	"strings"

	"fmt"

	"github.com/aiteung/module/model"
)

func Handler(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
	// Cek apakah pesan mengandung kata "kirim"
	if strings.Contains(strings.ToLower(Pesan.Message), "kirim") {
		// Mengekstrak kata setelah "kirim" menggunakan ekspresi reguler
		log.Printf("Pesan.Message sebelum ekstraksi: %v", Pesan.Message)
		pesanSplit := strings.Fields(Pesan.Message)
		log.Printf("Pesan.Message setelah ekstraksi: %v", pesanSplit)

		kirimIndex := -1
		for i, word := range pesanSplit {
			log.Printf("Iterasi ke-%d, word: %s", i, word)
			if strings.EqualFold(word, "kirim") {
				kirimIndex = i
				break
			}
		}

		log.Printf("kirimIndex: %v", kirimIndex)

		if kirimIndex != -1 && kirimIndex+1 < len(pesanSplit) {
			log.Printf("Nilai pesanSplit[1]: %v", pesanSplit[1])
			log.Printf("Panjang pesanSplit[1]: %v", len(pesanSplit[1]))
			log.Printf("Karakter pesanSplit[1]: %v", pesanSplit[1])
			log.Printf("Masuk ke dalam kondisi 'kirim skmk'")

			// matches[1] berisi kata setelah "kirim"
			if strings.EqualFold(strings.ToLower(pesanSplit[kirimIndex+1]), "skmk") {
				// Handle logika untuk "kirim skmk"
				dataMhs, err := GetMhsByPhoneNumber(db, Pesan.Phone_number)
				if err != nil {
					reply = MessageLengkapiData()
					return reply
				}

				tahunakademik, err := GetCurrentAcademicYear(db)
				if err != nil {
					reply = "Tidak ada tahun akademik aktif"
					return reply
				}

				// Handle logika untuk "kirim skmk"
				SkmkFile := ReplaceSkmk(dataMhs, *tahunakademik, db)
				baseString := base64.StdEncoding.EncodeToString(SkmkFile)
				filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)

				attachEmail := Files{
					Mimetype: "application/pdf",
					Name:     filename,
					Base64:   baseString,
				}

				// log.Printf("Sending email to: %s, Subject: %s", dataMhs.Email, "Surat Keterangan Masih Kuliah ULBI")

				// Send email with attachment
				StatusSend := SendEmailTo(
					urlEmail,
					dataMhs.Email,
					"Surat Keterangan Masih Kuliah ULBI",
					fmt.Sprintf("Hai hai haiii Kakak %s, inii iteung bawain SKMK nih buat kakak. Gunakan seperlunya ya kakk.....", dataMhs.NamaMhs),
					attachEmail)

				if !StatusSend {
					log.Printf("Error sending email: %t", StatusSend)
					reply = "Gagal mengirim SKMK melalui email"
					return reply
				}

				return MessageBerhasilMintaSkmk(dataMhs)
			} else {
				// Handle logika untuk "kirim" tanpa "skmk"
				log.Printf("Masuk ke dalam kondisi 'kirim' tanpa 'skmk'")
				return "Kirim apa broww???"
			}
		} else {
			// Handle logika jika tidak ada kata setelah "kirim"
			return "Yaiyaa kirim, kirim nya kirim apaa???"
		}
	} else {
		// Handle logika jika tidak ada kata "kirim" dalam pesan
		return "Apa yang Kakak kirim?"
	}
}

// func Handler(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
// 	dataMhs, err := GetMhsByPhoneNumber(db, Pesan.Phone_number)
// 	if err != nil {
// 		reply = MessageLengkapiData()
// 		return reply
// 	}

// 	tahunakademik, err := GetCurrentAcademicYear(db)
// 	if err != nil {
// 		reply = "Tidak ada tahun akademik aktif"
// 		return reply
// 	}

// 	if strings.Contains(strings.ToLower(Pesan.Message), "skmk") {
// 		SkmkFile := ReplaceSkmk(dataMhs, *tahunakademik, db)
// 		baseString := base64.StdEncoding.EncodeToString(SkmkFile)
// 		filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)

// 		attachEmail := Files{
// 			Mimetype: "application/pdf",
// 			Name:     filename,
// 			Base64:   baseString,
// 		}

// 		log.Printf("Sending email to: %s, Subject: %s", dataMhs.Email, "Surat Keterangan Masih Kuliah ULBI")

// 		// Send email with attachment
// 		StatusSend := SendEmailTo(
// 			urlEmail,
// 			dataMhs.Email,
// 			"Surat Keterangan Masih Kuliah ULBI",
// 			fmt.Sprintf("Hai hai haiii Kakak %s, inii iteung bawain SKMK nih buat kakak. Gunakan seperlunya ya kakk.....", dataMhs.NamaMhs),
// 			attachEmail)

// 		if !StatusSend {
// 			log.Printf("Error sending email: %t", StatusSend)
// 			reply = "Gagal mengirim SKMK melalui email"
// 			return reply
// 		}

// 		return MessageBerhasilMintaSkmk(dataMhs)
// 	} else {
// 		// Jika tidak ada kata kunci "skmk" dalam pesan, berikan respons yang sesuai
// 		return "Keyword Anda Salah"
// 	}
// }
