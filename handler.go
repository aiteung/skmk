package skmk

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"github.com/aiteung/module/model"
)

// func Handler(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
// 	// Cek apakah pesan mengandung kata "minta"
// 	if strings.Contains(strings.ToLower(Pesan.Message), "minta") {
// 		// Mengekstrak kata setelah "minta" menggunakan ekspresi reguler
// 		pesanSplit := strings.Fields(Pesan.Message)
// 		if len(pesanSplit) > 2 {
// 			// matches[1] berisi kata setelah "minta"
// 			if strings.Contains(strings.ToLower(pesanSplit[2]), "skmk") {
// 				// Handle logika untuk "minta skmk"
// 				dataMhs, err := GetMhsByPhoneNumber(db, Pesan.Phone_number)
// 				if err != nil {
// 					reply = MessageLengkapiData()
// 					return reply
// 				}

// 				tahunakademik, err := GetCurrentAcademicYear(db)
// 				if err != nil {
// 					reply = "Tidak ada tahun akademik aktif"
// 					return reply
// 				}

// 				// Handle logika untuk "minta skmk"
// 				SkmkFile := ReplaceSkmk(dataMhs, *tahunakademik, db)
// 				baseString := base64.StdEncoding.EncodeToString(SkmkFile)
// 				filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)

// 				attachEmail := Files{
// 					Mimetype: "application/pdf",
// 					Name:     filename,
// 					Base64:   baseString,
// 				}

// 				// log.Printf("Sending email to: %s, Subject: %s", dataMhs.Email, "Surat Keterangan Masih Kuliah ULBI")

// 				// Send email with attachment
// 				StatusSend := SendEmailTo(
// 					urlEmail,
// 					dataMhs.Email,
// 					"Surat Keterangan Masih Kuliah ULBI",
// 					fmt.Sprintf("Hai hai haiii Kakak %s, inii iteung bawain SKMK nih buat kakak. Gunakan seperlunya ya kakk.....", dataMhs.NamaMhs),
// 					attachEmail)

// 				if !StatusSend {
// 					log.Printf("Error sending email: %t", StatusSend)
// 					reply = "Gagal mengirim SKMK melalui email"
// 					return reply
// 				}

// 				return MessageBerhasilMintaSkmk(dataMhs)
// 			} else {
// 				// Handle logika untuk "minta" tanpa "skmk"
// 				return "Kirim apa broww???"
// 			}
// 		} else {
// 			// Handle logika jika tidak ada kata setelah "minta"
// 			return "Yaiyaa minta, minta nya minta apaa???"
// 		}
// 	} else {
// 		// Handle logika jika tidak ada kata "minta" dalam pesan
// 		return "Apa yang Kakak minta?"
// 	}
// }

func Handler(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
	if strings.Contains(Pesan.Message, "skmk") {
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
		fileskmk := ReplaceSkmk(dataMhs, *tahunakademik, db)
		baseString := base64.StdEncoding.EncodeToString(fileskmk)
		filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)
		attachEmail := Files{
			Mimetype: "application/pdf",
			Name:     filename,
			Base64:   baseString,
		}
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
		return MessageGagalMintaSkmk()
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
