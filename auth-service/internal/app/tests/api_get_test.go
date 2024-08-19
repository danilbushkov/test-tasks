package tests

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func TestApiNotFound(t *testing.T) {
	app, err := getTestApp()
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest("GET", "/test", nil)
	resp, _ := app.Api().Test(req)
	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("Status code is not 404. Code is %d", resp.StatusCode)
	}
}

func TestApiGetTokensWithEmptyJsonObject(t *testing.T) {
	app, err := getTestApp()
	defer app.Close()
	if err != nil {
		t.Fatal(err)
	}
	body := `{}`
	req := httptest.NewRequest("POST", "/api/auth/get", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := app.Api().Test(req)
	if resp.StatusCode != fiber.StatusUnprocessableEntity {
		t.Fatalf("Status code is not 422. Code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var e structs.Error
	err = json.Unmarshal(respBody, &e)
	if err != nil {
		t.Fatal(err)
	}
	if e.Message != "uuid field is not set" {
		t.Fatal("Body contains invalid value")
	}
}
