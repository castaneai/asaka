package asaka

import (
	"context"
	"net/http"
	"testing"
)

func TestGetDoc(t *testing.T) {
	hc := &http.Client{}
	c, err := NewClient(hc, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	doc, err := c.GetDoc(ctx, "http://example.com")
	if err != nil {
		t.Fatal(err)
	}
	// extract text
	titleText := doc.Find("h1").Text()
	if titleText != "Example Domain" {
		t.Fatalf("expected: %s, actual: %s", "Example Domain", titleText)
	}
}
