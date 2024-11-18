package bot

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"magical-crwler/models"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func ExportFileBot(ads []models.Ad, format string) error {
	data := retrieveData(ads)

	var outputFile, zipFile string
	if format == "csv" {
		outputFile = "ads_export.csv"
	} else if format == "xlsx" {
		outputFile = "ads_export.xlsx"
	} else {
		return fmt.Errorf("unsupported format: %s", format)
	}
	zipFile = "ads_export.zip"

	var err error
	if format == "csv" {
		err = ExportToCSV(outputFile, data)
	} else if format == "xlsx" {
		err = ExportToXLSX(outputFile, data)
	}
	if err != nil {
		return fmt.Errorf("failed to export file: %w", err)
	}
	defer os.Remove(outputFile)

	err = CreateZipFile(zipFile, outputFile)
	if err != nil {
		return fmt.Errorf("failed to compress file: %w", err)
	}
	defer os.Remove(zipFile)

	fmt.Println("Exported file successfully:", zipFile)
	return nil
}

func retrieveData(ads []models.Ad) [][]string {
	data := [][]string{
		{"ID", "Description", "Price", "Rent Price", "City", "Neighborhood", "Size", "Bedrooms", "Has Elevator", "For Rent", "Is Apartment", "Creation Time", "Visit Count"},
	}

	for _, ad := range ads {
		data = append(data, []string{
			fmt.Sprintf("%d", ad.ID),                      // ID
			DescriptionOrDefault(ad),                      // Description
			formatNullableInt64(ad.Price),                 // Price
			formatNullableInt(ad.RentPrice),               // Rent Price
			formatNullableString(ad.City),                 // City
			formatNullableString(ad.Neighborhood),         // Neighborhood
			formatNullableInt(ad.Size),                    // Size
			formatNullableInt(ad.Bedrooms),                // Bedrooms
			formatNullableBool(ad.HasElevator),            // Has Elevator
			fmt.Sprintf("%t", ad.ForRent),                 // For Rent
			fmt.Sprintf("%t", ad.IsApartment),             // Is Apartment
			ad.CreationTime.Format("2006-01-02 15:04:05"), // Creation Time
			fmt.Sprintf("%d", ad.VisitCount),              // Visit Count
		})
	}

	return data
}

func formatNullableString(value *string) string {
	if value != nil {
		return *value
	}
	return "N/A"
}

func formatNullableInt64(value *int64) string {
	if value == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *value)
}

func formatNullableInt(value *int) string {
	if value == nil {
		return "N/A"
	}
	return fmt.Sprintf("%d", *value)
}

func formatNullableBool(value *bool) string {
	if value != nil {
		return fmt.Sprintf("%t", *value)
	}
	return "N/A"
}

func DescriptionOrDefault(ad models.Ad) string {
	if ad.Description != nil {
		return *ad.Description
	}
	return "No Description"
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
