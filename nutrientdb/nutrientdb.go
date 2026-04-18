// Package to access the API of the swiss nutrition database
package nutrientdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const baseUrl = "https://api.webapp.prod.blv.foodcase-services.com/BLV_WebApp_WS/webresources/BLV-api"

var lang string = "en"

type Component struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Group int    `json:"group"`
	Unit  int    `json:"unit"`
}

type Food struct {
	ID            int     `json:"id"`
	Name          string  `json:"foodName"`
	Generic       bool    `json:"generic"`
	CategoryNames string  `json:"categoryNames"`
	Amount        float64 `json:"amount"`
	FoodId        int     `json:"foodid"`
	ValueTypeCode string  `json:"valueTypeCode"`
}

// Set the language used to query and for answers
// Defaults to en valid values are en, de, fr, it
func SetLang(newLang string) (bool, error) {
	switch newLang {
	case "en", "de", "fr", "it":
		lang = newLang
		return true, nil
	default:
		return false, fmt.Errorf("Unknown lang %s supported values are en, de, fr, it", newLang)
	}
}

// Receive all components and their metadata like ids
func Components() ([]Component, error) {
	result, err := fetch[Component]("components", nil)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Find the id for a specific component
func ComponentId(name string) (int, error) {
	components, err := Components()

	if err == nil {
		for _, component := range components {

			if strings.Contains(strings.ToLower(component.Name), strings.ToLower(name)) {
				return component.ID, nil
			}
		}
	}

	return 0, err
}

// Query all foods for the given component id
func GetFoodWithComponent(componentId int) ([]Food, error) {
	opts := map[string]string{
		"component": strconv.Itoa(componentId),
		"limit":     "100",
		"operator":  ">",
		"amount":    "0",
	}

	result, err := fetch[Food]("foods", opts)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Generic fetch function to handle all api requests
func fetch[T any](resource string, opts map[string]string) ([]T, error) {
	var result []T
	url := fmt.Sprintf("%s/%s?lang=%s", baseUrl, resource, lang)

	for key, value := range opts {
		url = fmt.Sprintf("%s&%s=%s", url, key, value)
	}

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
