package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "http://127.0.0.1:7891"
	method := "POST"

	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8"?>` +
		"" +
		`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="urn:TC" xmlns:xsd="http://www.w3.org/1999/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" SOAP-ENV:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">` +
		"" +
		`    <SOAP-ENV:Body>` +
		"" + `
        <ns1:executeCommand>` +
		"" + `<command>server info</command>` +
		"" + ` </ns1:executeCommand>` +
		"" + `
    </SOAP-ENV:Body>` + "" + `
</SOAP-ENV:Envelope>`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	username := "1#1"
	password := "07049809"

	stringToEncode := strings.Join([]string{username, password}, ":")
	fmt.Printf("\n %s \n", stringToEncode)

	origBytes := []byte(stringToEncode)
	encodedBytes := make([]byte, base64.StdEncoding.EncodedLen(len(origBytes)))
	base64.StdEncoding.Encode(encodedBytes, origBytes)

	fmt.Printf("\n encoded byte string %v \n", encodedBytes)

	encodedText := base64.StdEncoding.EncodeToString(origBytes)

	fmt.Printf("\n Encoded Text: %s \n", encodedText)

	auth := strings.Join([]string{"Basic", encodedText}, " ")

	req.Header.Add("Content-Type", "application/xml")
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
