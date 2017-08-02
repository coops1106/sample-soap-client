package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getXMLHeader() string {
	return "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
}

func getBody(countryName string) []byte {
	type CitiesByCountryRequest struct {
		XMLName     xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountry"`
		CountryName string   `xml:"CountryName"`
	}

	c := &CitiesByCountryRequest{CountryName: countryName}

	output, err := xml.MarshalIndent(c, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return output
}

func getSoapRequest(countryName string) string {
	body := getXMLHeader()
	body += "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">"
	body += "<soap:Body>"
	body += string(getBody(countryName))
	body += "</soap:Body>"
	body += "</soap:Envelope>"

	return body
}

func main() {
	url := "http://www.webservicex.net/globalweather.asmx?wsdl"

	reqBody := getSoapRequest("Thailand")

	//fmt.Println(reqBody)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/xml")
	resp, err := client.Do(req)
	if err != nil {
		//return err
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(errors.New(fmt.Sprint(resp)))
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println("\n***********************************************************")
	fmt.Println("post:\n", string(res))
	fmt.Println("\n***********************************************************")
}
