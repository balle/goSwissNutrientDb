package nutrientdb

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"
)

func TestFoodListSort(t *testing.T) {
	foods := FoodList{
		{Name: "Salt, himalaya salt", Amount: 24.2},
		{Name: "Cocoa powder, deoiled, without sugar", Amount: 51},
		{Name: "Paprika (spice)", Amount: 24.4},
		{Name: "Cinnamon", Amount: 38},
	}

	sort.Sort(foods)

	got := []string{foods[0].Name, foods[1].Name, foods[2].Name, foods[3].Name}
	want := []string{
		"Cocoa powder, deoiled, without sugar",
		"Cinnamon",
		"Paprika (spice)",
		"Salt, himalaya salt",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected sort order: got %v want %v", got, want)
	}
}

func TestComponentsAndComponentId(t *testing.T) {
	originalLang := lang
	originalBaseURL := baseURL
	originalHTTPGet := httpGet

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		components := []Component{
			{ID: 10, Name: "Natrium"},
			{ID: 20, Name: "Eisen"},
		}

		if err := json.NewEncoder(w).Encode(components); err != nil {
			t.Fatalf("failed to encode response: %v", err)
		}
	}))

	lang = "de"
	httpGet = http.Get
	baseURL = server.URL

	t.Cleanup(func() {
		server.Close()
		baseURL = originalBaseURL
		lang = originalLang
		httpGet = originalHTTPGet
	})

	components, err := Components()

	if err != nil {
		t.Fatalf("Components returned error: %v", err)
	}

	if len(components) != 2 {
		t.Fatalf("Components returned %d entries, want 2", len(components))
	}

	if components[0].ID != 10 || components[1].Name != "Eisen" {
		t.Fatalf("unexpected components result: %+v", components)
	}

	componentID, err := ComponentId("Eisen")

	if err != nil {
		t.Fatalf("ComponentId returned error: %v", err)
	}

	if componentID != 20 {
		t.Fatalf("ComponentId returned %d, want 20", componentID)
	}
}

func TestGetFoodWithComponent(t *testing.T) {
	originalLang := lang
	originalBaseURL := baseURL
	originalHTTPGet := httpGet

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		foods := []Food{
			{Name: "Salt, himalaya salt", Amount: 24},
			{Name: "Cocoa powder, deoiled, without sugar", Amount: 51},
			{Name: "Paprika (spice)", Amount: 24},
			{Name: "Cinnamon", Amount: 38},
		}

		if err := json.NewEncoder(w).Encode(foods); err != nil {
			t.Fatalf("failed to encode response: %v", err)
		}
	}))

	lang = "it"
	httpGet = http.Get
	baseURL = server.URL

	t.Cleanup(func() {
		server.Close()
		baseURL = originalBaseURL
		lang = originalLang
		httpGet = originalHTTPGet
	})

	foods, err := GetFoodWithComponent(123)

	if err != nil {
		t.Fatalf("GetFoodWithComponent returned error: %v", err)
	}

	want := []string{
		"Cocoa powder, deoiled, without sugar",
		"Cinnamon",
		"Paprika (spice)",
		"Salt, himalaya salt",
	}
	got := []string{foods[0].Name, foods[1].Name, foods[2].Name, foods[3].Name}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected sorted foods: got %v want %v", got, want)
	}
}
