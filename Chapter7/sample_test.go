package sample_test

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://www.ardanlabs.com/blog/index.xml"
	statusCode := 200

	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("unable to issue GET on URL: %s, %s", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != statusCode {
		t.Log("exp:", statusCode)
		t.Log("got:", resp.StatusCode)
		t.Fatal("status codes don't match")
	}
}
func TestUpload(t *testing.T) {}
