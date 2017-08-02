package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://www.webservicex.net/globalweather.asmx?wsdl"

	body := "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	body += "<soap:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap=\"http://schemas.xmlsoap.org/soap/envelope/\">"
	body += "<soap:Body>"
	body += "<GetCitiesByCountry xmlns=\"http://www.webserviceX.NET\">"
	body += "<CountryName>Thailand</CountryName>"
	body += "</GetCitiesByCountry>"
	body += "</soap:Body>"
	body += "</soap:Envelope>"

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
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
	fmt.Println("post:\n", string(res))
}
