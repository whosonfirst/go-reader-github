package reader

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/whosonfirst/go-reader"
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
	file_uri := "101/736/545/101736545.geojson"

	ctx := context.Background()

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatal(err)
	}

	fh, err := r.Read(ctx, file_uri)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}
