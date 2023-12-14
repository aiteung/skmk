package skmk

import (
	"database/sql"
	"encoding/base64"
	"log"
	"strings"

	"fmt"

	"github.com/aiteung/module/model"
)

func SKMKCreator(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
	if strings.Contains(strings.ToLower(Pesan.Message), "skmk") {
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

		return MessageBerhasilMintaSkmk()
	} else {
		// Jika tidak ada kata kunci "skmk" dalam pesan, berikan respons yang sesuai
		return "Tidak ada permintaan SKMK dalam pesan"
	}
}
