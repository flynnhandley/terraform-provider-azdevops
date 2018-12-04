package azdevops

import (
	"net/http"
	"testing"
)

func TestApiClient(t *testing.T) {

	client := &AzDevOpsClient{
		Account:    "",
		Token:      "",
		HTTPClient: &http.Client{},
	}

	for _, err := client.Get("_apis/projects/?api-version=4.1-preview"); err != nil; {
		t.Fatalf("err: %s", err)
	}

}
