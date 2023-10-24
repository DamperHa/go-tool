package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	req, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("----Response---------")
	fmt.Println("STATUS CODE: ", res.StatusCode)
	fmt.Println("BODY:", string(body))
}
