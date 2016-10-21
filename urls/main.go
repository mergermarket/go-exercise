package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"errors"
	"bytes"
)

func main() {

}

func Concatenator(urls ...string) (string, error) {

	buffer := bytes.NewBuffer(nil)

	successes := make(chan []byte)
	error := make(chan error)

	for _, url := range urls {
		go foobar(url, successes, error)
	}

	for i:=0; i<len(urls); i++ {
		select {
		case err:= <-error:
			return "", err
		case data:=<-successes:
			buffer.Write(data)
		}
	}

	return string(buffer.Bytes()), nil
}

func foobar(url string, successes chan []byte, error chan error) {
	result, err := fetch(url)
	if err != nil {
		error<-err
	} else {
		successes<-result
	}
}

func fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(res.StatusCode))
	}

	output, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}