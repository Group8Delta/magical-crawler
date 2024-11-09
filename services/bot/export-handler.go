package bot

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
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

func ExportToCSV(data []map[string]string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Price", "Occupancy"} //bayad hame data haro namayesh bede
	writer.Write(headers)

	for _, item := range data {
		row := []string{item["price"], item["occupancy"]}
		writer.Write(row)
	}

	return nil

}

func ExportToXLSX(data []map[string]string, fileName string) error {
	file := excelize.NewFile()
	sheet := "Sheet1"

	file.SetCellValue(sheet, "A1", "Price")
	file.SetCellValue(sheet, "B1", "Occupancy")

	for i, item := range data {
		file.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), item["price"])
		file.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), item["occupancy"])
	}

	return file.SaveAs(fileName)
}
