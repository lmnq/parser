package estp

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var auctionParamKey string = "search"
var estpURL string = "http://estp.ru/_next/data/rSEH0esgmicC0bXaMGzcW/purchases.json"

func GetAuctions(keyWord string) (*Estp, error) {
	// reqURL := fmt.Sprintf("%s?%s=%s", estpURL, auctionParamKey, keyWord)
	reqURL := estpURL + "?" + auctionParamKey + "=" + url.QueryEscape(keyWord)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		log.Printf("server: could not create request: GetAuctions\nurl: %s", reqURL)
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: GetAuctions\nERROR: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("client: success!\nrequest URL: %s\nstatus code: %d", reqURL, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("client: could not read response body\nERROR: %s", err)
		return nil, err
	}

	estp := New()
	if err := json.Unmarshal(respBody, estp); err != nil {
		log.Printf("server: could not parse response body (json)\nERROR: %s", err)
		return nil, err
	}

	return estp, nil
}
