package reader

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/whosonfirst/go-reader/v2"
)

var access_token = flag.String("access-token", "", "A valid GitHub Oauth2 access token")

func TestAPIReader(t *testing.T) {

	if *access_token == "" {
		return
	}

	owner := "whosonfirst-data"
	repo := "whosonfirst-data-admin-ca"
	branch := "master" // pending rollover (20210114/thisisaaronland)

	reader_uri := fmt.Sprintf("githubapi://%s/%s?branch=%s&access_token=%s", owner, repo, branch, *access_token)
	file_uri := "data/101/736/545/101736545.geojson"

	ctx := context.Background()

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatalf("Failed to create new reader, %v", err)
	}

	fh, err := r.Read(ctx, file_uri)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", file_uri, err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatalf("Failed to copy %s, %v", file_uri, err)
	}

	exists, err := r.Exists(ctx, file_uri)

	if err != nil {
		t.Fatalf("Failed to determine if %s exists, %v", file_uri, err)
	}

	if !exists {
		t.Fatalf("Expected %s to exist", file_uri)
	}
}
