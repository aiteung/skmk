package skmk

import (
	"database/sql"
	"encoding/base64"
	"log"
	"regexp"
	"strings"

	"fmt"

	"github.com/aiteung/module/model"
)

func Handler(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
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

	if strings.Contains(strings.ToLower(Pesan.Message), "minta") {
		// Mengekstrak kata setelah "minta" menggunakan ekspresi reguler
		re := regexp.MustCompile(`minta\s+(\w+)`)
		matches := re.FindStringSubmatch(Pesan.Message)
		if len(matches) > 1 {
			// matches[1] berisi kata setelah "minta"
			if strings.Contains(strings.ToLower(matches[1]), "skmk") {
				// Handle logika untuk "minta skmk"
				SkmkFile := ReplaceSkmk(dataMhs, *tahunakademik, db)
				baseString := base64.StdEncoding.EncodeToString(SkmkFile)
				filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)

				attachEmail := Files{
					Mimetype: "application/pdf",
					Name:     filename,
					Base64:   baseString,
				}

				log.Printf("Sending email to: %s, Subject: %s", dataMhs.Email, "Surat Keterangan Masih Kuliah ULBI")

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
				// Handle logika untuk "minta" tanpa "skmk"
				return "Minta apa broww???"
			}
		} else {
			// Handle logika jika tidak ada kata setelah "minta"
			return "Yaiyaa minta, minta nya minta apaa???"
		}
	} else {
		// Handle logika jika tidak ada kata "minta" dalam pesan
		return "skmk apaan???"
	}
}
