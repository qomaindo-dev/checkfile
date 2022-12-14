package checkfile

import (
	"errors"
	"net/http"
	"os"
	"strconv"
)

// Search file from url website
func ExternDir(url string) (bool, error) {
	// url file in web, ex: http://google.com/file/example-file.pdf
	if url == "" {
		return false, errorFieldResponse(url)
	}

	response, err := http.Head(url)

	return convRespExt(response, err)
}

// Search file on internal project
func InternDir(filePath string) (bool, error) {
	// Path local directory, ex: assets/file/example-file.pdf
	if filePath == "" {
		return false, errorFieldResponse(filePath)
	}

	_, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

func convRespExt(response *http.Response, err error) (bool, error) {
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return false, errors.New("Response data URL is " + strconv.Itoa(response.StatusCode))
	}

	return true, nil
}

func errorFieldResponse(param string) error {
	return errors.New("Missing data on field " + param)
}
