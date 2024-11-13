package bot

import (
	"log"
	"magical-crwler/constants"
	"os"

	"gopkg.in/telebot.v4"
)

func ExportHandler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		exportMenu := &telebot.ReplyMarkup{}

		xlsxBtn := exportMenu.Data(constants.ExportXLSX, "export_xlsx")
		csvBtn := exportMenu.Data(constants.ExportCSV, "export_csv")

		exportMenu.Inline(
			exportMenu.Row(xlsxBtn, csvBtn),
		)

		return c.EditOrSend("یکی از گزینه های زیر را برای دریافت فایل انتخاب کنید", exportMenu)
	}
}

func export_xlsx_Handler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		data := retrieveData()
		filename := "exported_data.xlsx"
		zipFilename := "exported_data.zip"

		if err := ExportToXLSX(filename, data); err != nil {
			log.Printf("Error exporting XLSX: %v", err)
			return err
		}

		if err := CreateZipFile(zipFilename, filename); err != nil {
			log.Printf("Error creating zip file: %v", err)
			return err
		}

		zipFile := &telebot.Document{File: telebot.FromDisk(zipFilename), FileName: zipFilename}
		if err := c.Send(zipFile); err != nil {
			return err
		}

		os.Remove(filename)
		os.Remove(zipFilename)
		return nil
	}
}

func export_csv_Handler(b *Bot) func(c telebot.Context) error {
	return func(c telebot.Context) error {
		data := retrieveData()
		filename := "exported_data.csv"
		zipFilename := "exported_data.zip"

		if err := ExportToCSV(filename, data); err != nil {
			log.Printf("Error exporting CSV: %v", err)
			return err
		}

		if err := CreateZipFile(zipFilename, filename); err != nil {
			log.Printf("Error creating zip file: %v", err)
			return err
		}

		zipFile := &telebot.Document{File: telebot.FromDisk(zipFilename), FileName: zipFilename}
		if err := c.Send(zipFile); err != nil {
			return err
		}

		// Clean up
		os.Remove(filename)
		os.Remove(zipFilename)
		return nil
	}
}
