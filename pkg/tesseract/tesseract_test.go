package tesseract

import (
	"io/ioutil"
	"testing"
)

func TestFromReader(t *testing.T) {
	// reader := ioutil.R
}

func TestFromUrl(t *testing.T) {

}

func TestFromBytes(t *testing.T) {
	bs, err := ioutil.ReadFile("../../test-data/hn_comment_12694592.png")
	if err != nil {
		t.Fatal(err)
	}
	ocr, err := FromBytes(bs)
	if err != nil {
		t.Fatal(err)
	}
	ioutil.WriteFile("foo.txt", []byte(ocr), 0x7777)
	expected := `Although Google's API is certainly better, Tesseract.js should work similarly if you increase the font size.
Screenshots taken on 'retina' devices are around the smallest text it can handle well.

Edit:
A screenshot of the same text at a higher resolution: https:[ZimguncomlalWHGu
Tesseract.js output: https:[[imgur.com[a[nilfM`
	if ocr != expected {
		t.Errorf("\nExpected output (%v characters):\n---\n%v\n---\n\nActual output (%v characters):\n---\n%v\n---", len(expected), expected, len(ocr), ocr)
	}
}
