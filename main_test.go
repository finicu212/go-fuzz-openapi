package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

// TODO: Native Go Fuzz integration with complex data (struct) fuzzing? https://github.com/AdaLogics/go-fuzz-headers

type Tag struct {
	Id   int32  // integer, int64
	Name string // string
}

type Category struct {
	Id   int32  // integer, example: 1
	Name string // string, example: Dogs
}

// Pet:
//
//	required:
//	- name
//	- photoUrls
type Pet struct {
	Category  Category
	Id        int32    // integer, int64, example: 10
	Name      string   // string, example: doggie
	Photourls []string // array, items: string
	Status    string   // enum: [available, pending, sold]
	Tags      []Tag    // schemas/Tag
}

func encodeAndSend(v any, method, endpoint string) (*http.Response, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("https://petstore3.swagger.io/api/v3/%s", endpoint), &buf)
	if err != nil {
		log.Fatalf("Failed creating request with data: %+v", buf)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	return client.Do(req)
}

func Fuzz_PetPost_OnlyRequired(f *testing.F) {
	f.Add("doggie", 2, "https://my-doggie-picture.com/pic123.jpg")
	test := func(t *testing.T, name string, photoUrlsLen int, Photourl string) {
		if photoUrlsLen < 1 {
			t.Skip()
		}
		Photourls := make([]string, photoUrlsLen)
		for i := range Photourls {
			Photourls[i] = Photourl
		}
		p := Pet{Name: name, Photourls: Photourls}

		// Encode the Pet object to a JSON object, which is held in a bytes.Buffer (which extends io.Reader)
		resp, err := encodeAndSend(p, "POST", "pet")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode/100 == 5 {
			t.Fatalf("Found a failed state: %s, using %+v", resp.Status, p)
		}

		if resp.StatusCode/100 == 4 {
			t.Logf("Unexpected forbidden state: %s, using %+v", resp.Status, p)
		}

		fmt.Println(string(body))
	}
	f.Fuzz(test)
}
