package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/delihiros/uno/pkg/entities"
)

const (
	unofficialVarorantAPIURL = "https://valorant-api.com"
)

type UnofficialValorantAPIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewUnofficialValorantAPIClient() *UnofficialValorantAPIClient {
	return &UnofficialValorantAPIClient{
		baseURL:    unofficialVarorantAPIURL,
		httpClient: &http.Client{},
	}
}

func (c *UnofficialValorantAPIClient) simpleGet(endpoint string, queries map[string]string) ([]byte, error) {
	url := &url.URL{}
	q := url.Query()
	for k, v := range queries {
		q.Set(k, v)
	}
	url.RawQuery = q.Encode()
	requestURL := c.baseURL + endpoint + url.String()
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func (c *UnofficialValorantAPIClient) get(endpoint string, queries map[string]string, v interface{}) error {
	body, err := c.simpleGet(endpoint, queries)
	if err != nil {
		return err
	}
	r := &UnofficialValorantResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	if r.Status != 200 {
		return fmt.Errorf("failed to GET: %v, status = %v, data = %v", endpoint, r.Status, string(body))
	}
	return json.Unmarshal(r.Data, v)
}

func (c *UnofficialValorantAPIClient) GetWeapons() ([]*entities.Weapon, error) {
	weapons := []*entities.Weapon{}
	err := c.get("/v1/weapons", map[string]string{}, &weapons)
	if err != nil {
		return nil, err
	}
	return weapons, nil
}
