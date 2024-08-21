package tests

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danilbushkov/test-tasks/internal/app/services/auth/tokens"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/pashagolub/pgxmock/v4"
)

func TestApiNotFound(t *testing.T) {
	app, _, err := getTestApp()
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
	app, _, err := getTestApp()
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
	app, _, err := getTestApp()
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
	app, _, err := getTestApp()
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

type AnyHash struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyHash) Match(v interface{}) bool {
	return true
}

func TestApiGetTokensWithValidUUID(t *testing.T) {
	u := "674859e1-7772-4a6a-9df1-11fccfa4e144"
	ip := "0.0.0.0"
	setEnv(t)
	app, mock, err := getTestApp()
	defer app.Close()

	mock.ExpectExec("INSERT INTO auth").WithArgs(AnyHash{}).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	body := `{ "uuid": "` + u + `" }`
	req := httptest.NewRequest("POST", "/api/auth/get", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	req.RemoteAddr = ip

	resp, _ := app.Api().Test(req)
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Status code is not 200. Code is %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var ts structs.Tokens

	err = json.Unmarshal(respBody, &ts)
	if err != nil {
		t.Fatal(err)
	}
	access, err := tokens.NewAccessFromString(ts.AccessToken, "a_key")
	if err != nil {
		t.Fatal(err)
	}
	if access.UUID() != u {
		t.Fatal("UUID is invalid in token in response")
	}

}
