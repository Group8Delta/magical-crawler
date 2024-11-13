package bot

import (
	"archive/zip"
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func retrieveData() [][]string {
	return [][]string{
		{"ID", "Name", "Price", "Occupancy"},
		{"1", "Apartment A", "1000", "3"},
		{"2", "Apartment B", "1500", "2"},
		{"3", "Apartment C", "1200", "4"},
	}
}

func ExportToCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func ExportToXLSX(filename string, data [][]string) error {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName(f.GetSheetName(1), sheet)

	for i, row := range data {
		for j, cell := range row {
			cellRef, _ := excelize.CoordinatesToCellName(j+1, i+1)
			f.SetCellValue(sheet, cellRef, cell)
		}
	}

	return f.SaveAs(filename)
}

func CreateZipFile(zipFileName, fileToZip string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToAdd, err := os.Open(fileToZip)
	if err != nil {
		return err
	}
	defer fileToAdd.Close()

	w, err := zipWriter.Create(filepath.Base(fileToZip))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, fileToAdd)
	return err
}
