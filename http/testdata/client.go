package main

import (
	"crypto/tls"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// ssl enable ?
	//reqUrl := "https://127.0.0.1:8999/ping"
	//reqUrl := "http://127.0.0.1:8999/ping"
	reqUrl := "https://127.0.0.1:8999/v1/ping"
	//reqUrl := "http://127.0.0.1:8999/v1/ping"

	// new client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(reqUrl)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}
	defer resp.Body.Close()

	// bad request
	if resp.StatusCode != http.StatusOK {
		log.Fatal(errors.Errorf("fail : %s", reqUrl))
	}

	// read body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(errors.Errorf("fail : %s", reqUrl))
	}
	log.Printf("success : %s \n", b)
}
