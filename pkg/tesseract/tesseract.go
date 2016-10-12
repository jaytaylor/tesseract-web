package tesseract

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func FromReader(reader io.Reader) (string, error) {
	cmd := exec.Command("tesseract", "stdin", "stdout")
	cmd.Stdin = reader
	ocrBytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	ocrBytes = bytes.Replace(ocrBytes, []byte{'\r'}, nil, -1)
	ocr := string(ocrBytes)
	ocr = strings.TrimSpace(ocr)
	return ocr, nil
}

func FromUrl(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode/100 != 2 {
		return "", fmt.Errorf("received non-2xx response status-code=%v from url=%v", response.StatusCode, url)
	}
	ocr, err := FromReader(response.Body)
	if err != nil {
		return "", err
	}
	if err := response.Body.Close(); err != nil {
		return "", fmt.Errorf("closing response body: %s", err)
	}
	return ocr, nil
}

func FromBytes(bs []byte) (string, error) {
	reader := bytes.NewBuffer(bs)
	ocr, err := FromReader(reader)
	if err != nil {
		return "", err
	}
	return ocr, nil
}
