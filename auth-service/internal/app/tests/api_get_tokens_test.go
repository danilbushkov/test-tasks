package tests

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	auth_service "github.com/danilbushkov/test-tasks/internal/app/services/auth"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

func TestApiGetTokensWithInvalidUUIDLength(t *testing.T) {
	app, err := getTestApp()
	defer app.Close()
	if err != nil {
		t.Fatal(err)
	}
	body := `{ "uuid": "test" }`
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
	if e.Message != "invalid UUID length: 4" {
		t.Fatal("Body contains invalid value")
	}
}

func TestApiGetTokensWithInvalidUUID(t *testing.T) {
	app, err := getTestApp()
	defer app.Close()
	if err != nil {
		t.Fatal(err)
	}
	body := `{ "uuid": "674859e1-7772-4a6a-9dfz-11fccfa4e144" }`
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
	if e.Message != "invalid UUID format" {
		t.Fatal("Body contains invalid value")
	}

}

func TestApiGetTokensWithValidUUID(t *testing.T) {
	setEnv(t)
	app, err := getTestApp()
	defer app.Close()
	if err != nil {
		t.Fatal(err)
	}
	body := `{ "uuid": "674859e1-7772-4a6a-9df1-11fccfa4e144" }`
	req := httptest.NewRequest("POST", "/api/auth/get", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := app.Api().Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Status code is not 200. Code is %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var tokens auth_service.Tokens

	err = json.Unmarshal(respBody, &tokens)
	if err != nil {
		t.Fatal(err)
	}
	accessClaims := new(auth_service.TokenClaims)
	_, err = jwt.ParseWithClaims(tokens.AccessToken, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("a_key"), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if accessClaims.UUID != "674859e1-7772-4a6a-9df1-11fccfa4e144" {
		t.Fatal("UUID is invalid in token in response")
	}
	refreshClaims := new(auth_service.TokenClaims)
	_, err = jwt.ParseWithClaims(tokens.RefreshToken, refreshClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("r_key"), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if refreshClaims.UUID != "674859e1-7772-4a6a-9df1-11fccfa4e144" {
		t.Fatal("UUID is invalid in token in response ")
	}

}
