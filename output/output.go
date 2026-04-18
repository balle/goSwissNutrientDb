package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/balle/goSwissNutrientDb/nutrientdb"
	"github.com/flopp/go-findfont"
	"github.com/signintech/gopdf"
)

func GeneratePdf(filename string, title string, foodList []nutrientdb.Food) error {
	var x, y float64 = 15, 30
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	fontPath, err := findfont.Find("arial.ttf")

	if err != nil {
		fontPath, _ = findfont.Find("liberation-sans.ttf")
	}

	pdf.AddTTFFont("my-font", fontPath)
	pdf.AddTTFFont("my-big-font", strings.Replace(fontPath, ".ttf", "_Bold.ttf", 1))
	err = pdf.SetFont("my-big-font", "", 24)

	if err != nil {
		return fmt.Errorf("Cannot set big font: %s", err)
	}

	pdf.Text(title)
	y += 30
	pdf.SetXY(x, y)

	err = pdf.SetFont("my-font", "", 12)

	if err != nil {
		return fmt.Errorf("Cannot set font: %s", err)
	}

	for _, food := range foodList {
		pdf.Text(food.Name)
		y += 15
		pdf.SetXY(x, y)
	}

	if err := pdf.WritePdf(filename); err != nil {
		return fmt.Errorf("Cannot write pdf: %s", err)
	}

	return nil
}

func GenerateJson(filename string, foodlist []nutrientdb.Food) error {
	fh, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return fmt.Errorf("Cannot write file %s: %s", filename, err)
	}

	defer fh.Close()
	data, err := json.MarshalIndent(foodlist, "", "    ")

	if err != nil {
		return fmt.Errorf("Cannot encode data to json: %s", err)
	}

	fh.Write(data)

	return nil
}
