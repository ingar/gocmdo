package firebase

import (
	"encoding/json"
	"fmt"
	"github.com/ingar/barglebot/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func makeUrl(path string) (s string) {
	s = fmt.Sprintf("https://gocmdo-dev.firebaseio.com/%s.json?auth=%s", path, os.Getenv("FIREBASE_SECRET"))
	fmt.Println("Constructed firebase URL:", s)
	return
}

func makeRequest(method string, path string, data io.Reader) (buf []byte, err error) {
	req, err := http.NewRequest(method, makeUrl(path), data)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// debug
	var o interface{}
	err = json.Unmarshal(buf, &o)
	fmt.Println("Deserialized firebase response:", util.FmtJSON(o))

	return
}

func Post(path string, data io.Reader) (buf []byte, err error) {
	return makeRequest("POST", path, data)
}

func Patch(path string, data io.Reader) (buf []byte, err error) {
	return makeRequest("PATCH", path, data)
}

func Put(path string, data io.Reader) (buf []byte, err error) {
	return makeRequest("PUT", path, data)
}

func Get(path string) (buf []byte, err error) {
	return makeRequest("GET", path, nil)
}
