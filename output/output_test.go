package output

import (
	"testing"

	"github.com/balle/goSwissNutrientDb/nutrientdb"
)

var foodList = []nutrientdb.Food{
	{ID: 348976, Name: "Cocoa powder, deoiled, without sugar", Generic: true, CategoryNames: "Chocolate and cocoa products", Amount: 51, FoodId: 978, ValueTypeCode: "BE"},
	{ID: 349191, Name: "Cinnamon", Generic: true, CategoryNames: "Salt, spices and flavours", Amount: 38, FoodId: 911, ValueTypeCode: "X"},
	{ID: 350054, Name: "Paprika (spice)", Generic: true, CategoryNames: "Salt, spices and flavours", Amount: 24, FoodId: 663, ValueTypeCode: "X"},
	{ID: 350723, Name: "Salt, himalaya salt", Generic: true, CategoryNames: "Salt, spices and flavours", Amount: 24, FoodId: 14088, ValueTypeCode: "W"}}

func TestGeneratePdf(t *testing.T) {
	err := GeneratePdf("test.pdf", "Test", foodList)

	if err != nil {
		t.Errorf("Failed generating pdf: %s", err)
	}
}

func TestGenerateJson(t *testing.T) {
	err := GenerateJson("test.json", foodList)

	if err != nil {
		t.Errorf("Failed generating json: %s", err)
	}
}
