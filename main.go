package main

import (
	"net/http"
	"log"
	"bytes"
)

func main(){
	//API Get Post API
	url := "https://api.apitel.co/sms"
	var jsonStr []byte

	//append(jsonStr, "":"")


	jsonStr = []byte(`{"table_code":"QT","bill_type":0}`)


	reqs, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	reqs.Header.Set("X-Custom-Header", "myvalue")
	reqs.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqs)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&new_doc_no); err != nil {
		log.Println(err)
	}

}