package booksearch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type Query struct {
	Query Match `json:"query"`
}

type Match struct {
	Target map[string]string `json:"match"`
}

func doByJson(data interface{}, mode string) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return do(bytes.NewReader(json))
}

func do(reader io.Reader) error {
	// URL はとりあえず固定
	req, err := http.NewRequest("GET", "http://localhost:9200/my_index/_search", reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		return err
	}
	fmt.Println(string(dump))
	return nil
}

// Register register book
func Register(filepath string) error {
	// TODO Do not load full contents of the file
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	s := base64.StdEncoding.EncodeToString(f)

	return doByJson(map[string]string{"data": s})
}

// Search search for the specified word in all books
func Search(word string) error {
	return doByJson(Query{Match{map[string]string{"attachment.content": word}}})
}
