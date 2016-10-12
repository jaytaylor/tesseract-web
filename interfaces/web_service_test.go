package interfaces

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gigawattio/go-commons/pkg/web"
)

const (
	TestDataPath      = "../test-data/"
	TestImageFilename = "hn_comment_12694592.png"
	ExpectedJson      = `"Although Google's API is certainly better, Tesseract.js should work similarly if you increase the font size.\nScreenshots taken on 'retina' devices are around the smallest text it can handle well.\n\nEdit:\nA screenshot of the same text at a higher resolution: https:[ZimguncomlalWHGu\nTesseract.js output: https:[[imgur.com[a[nilfM"`
)

func TestWebServicePostBytes(t *testing.T) {
	withRunningWebService(t, func(webService *WebService) {
		imageFh, err := os.OpenFile(TestDataPath+TestImageFilename, os.O_RDONLY, 0x777)
		if err != nil {
			t.Fatal(err)
		}
		response, err := http.Post(fmt.Sprintf("http://%v/v1/tesseract", webService.Addr()), "image/png", imageFh)
		if err != nil {
			t.Fatal(err)
		}
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}
		body := string(bodyBytes)
		expected := ExpectedJson
		if body != expected {
			t.Errorf("\nExpected response body (%v characters):\n---\n%v\n---\n\nBut got actual body (%v characters):\n---\n%v\n---\n", len(expected), expected, len(body), body)
		}
	})
}

func TestWebServiceRemoteUrl(t *testing.T) {
	withRunningWebService(t, func(webService *WebService) {
		fsServer := web.NewFsServer("127.0.0.1:0", TestDataPath)
		if err := fsServer.Start(); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := fsServer.Stop(); err != nil {
				t.Error(err)
			}
		}()
		http.Post(fmt.Sprintf("http://%v/v1/tesseract/http://%v/%v", webService.Addr(), fsServer.Addr(), TestImageFilename), "", nil)
	})
}

func withRunningWebService(t *testing.T, fn func(webService *WebService)) {
	webService := NewWebService("127.0.0.1:0")
	if err := webService.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := webService.Stop(); err != nil {
			t.Error(err)
		}
	}()
	fn(webService)
}
