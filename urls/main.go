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

	fetchMany(urls, successes, error)

	for range urls {
		select {
		case err := <-error:
			return "", err
		case data := <-successes:
			buffer.Write(data)
		}
	}

	return string(buffer.Bytes()), nil
}

func fetchMany(urls []string, successes chan []byte, errors chan error) {
	for _, url := range urls {
		go func() {
			result, err := fetch(url)
			if err != nil {
				errors <- err
			} else {
				successes <- result
			}
		}()
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