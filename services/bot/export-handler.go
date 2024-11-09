package bot

import (
	"encoding/csv"
	"os"

	"gopkg.in/telebot.v4"
)

func ExportHandler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		exportMenu := &telebot.ReplyMarkup{ResizeKeyboard: true}

		xslxBtn := exportMenu.Text("xslx(اکسل)")
		csvBtn := exportMenu.Text("csv")

		exportMenu.Reply(
			exportMenu.Row(xslxBtn, csvBtn),
		)

		return c.Send(" یکی از گزینه های زیر را برای دریافت فایل انتخاب کنید", exportMenu)

	}
}

func exportToCsv(data map[string]string, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"شهر",
		"قیمت",
		"محله",
		"متراژ",
		"اتاق",
		"سن بنا",
		"نوع خونه",
		"طبقه",
		"انباری",
		"آسانسور",
		"تاریخ",
	}
	writer.Write(headers)
	return nil

}
