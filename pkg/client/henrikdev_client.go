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
	henrikdevURL = "https://api.henrikdev.xyz"
)

type HenrikdevClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewHenrikdevClient() *HenrikdevClient {
	return &HenrikdevClient{
		baseURL:    henrikdevURL,
		httpClient: &http.Client{},
	}
}

func (c *HenrikdevClient) simpleGet(endpoint string, queries map[string]string) ([]byte, error) {
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

func (c *HenrikdevClient) get(endpoint string, queries map[string]string, v interface{}) error {
	body, err := c.simpleGet(endpoint, queries)
	if err != nil {
		return err
	}
	r := &HenrikdevResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	if r.Status != 200 {
		return fmt.Errorf("failed to GET: %v, status = %v, data = %v", endpoint, r.Status, string(body))
	}
	return json.Unmarshal(r.Data, v)
}

func (c *HenrikdevClient) GetAccountByNameTag(name string, tag string) (*entities.Account, error) {
	account := &entities.Account{}
	err := c.get("/valorant/v1/account/"+name+"/"+tag, map[string]string{}, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (c *HenrikdevClient) GetMMRDataByNameTag(region string, name string, tag string) (*entities.MMRData, error) {
	mmrData := &entities.MMRData{}
	err := c.get("/valorant/v2/mmr/"+region+"/"+name+"/"+tag, map[string]string{}, mmrData)
	if err != nil {
		return nil, err
	}
	return mmrData, nil
}

func (c *HenrikdevClient) GetMMRDataByPUUID(region string, puuid string) (*entities.MMRData, error) {
	mmrData := &entities.MMRData{}
	err := c.get("/valorant/v2/by-puuid/mmr/"+region+"/"+puuid, map[string]string{}, mmrData)
	if err != nil {
		return nil, err
	}
	return mmrData, nil
}

func (c *HenrikdevClient) GetMMRHistoryByNameTag(region string, name string, tag string) ([]*entities.MMRHistory, error) {
	history := []*entities.MMRHistory{}
	err := c.get("/valorant/v1/mmr-history/"+region+"/"+name+"/"+tag, map[string]string{}, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (c *HenrikdevClient) GetMMRHistoryByPUUID(region string, puuid string) ([]*entities.MMRHistory, error) {
	history := []*entities.MMRHistory{}
	err := c.get("/valorant/v1/by-puuid/mmr-history/"+region+"/"+puuid, map[string]string{}, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (c *HenrikdevClient) GetMatchByID(matchID string) (*entities.Match, error) {
	match := &entities.Match{}
	err := c.get("/valorant/v2/match/"+matchID, map[string]string{}, match)
	if err != nil {
		return nil, err
	}
	return match, nil
}

func (c *HenrikdevClient) GetMatchHistory(region string, name string, tag string, filter string) ([]*entities.Match, error) {
	matches := []*entities.Match{}
	queries := map[string]string{}
	if filter != "" {
		queries["filter"] = filter
	}
	err := c.get("/valorant/v3/matches/"+region+"/"+name+"/"+tag, queries, &matches)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (c *HenrikdevClient) GetMatchHistoryByPUUID(region string, puuid string, filter string) ([]*entities.Match, error) {
	matches := []*entities.Match{}
	queries := map[string]string{}
	if filter != "" {
		queries["filter"] = filter
	}
	err := c.get("/valorant/v3/by-puuid/matches/"+region+"/"+puuid, queries, &matches)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (c *HenrikdevClient) GetContent() (*entities.Content, error) {
	content := &entities.Content{}
	body, err := c.simpleGet("/valorant/v1/content", map[string]string{})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, content)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *HenrikdevClient) GetLeaderboard(region string) ([]*entities.LeaderboardPlayer, error) {
	players := []*entities.LeaderboardPlayer{}
	body, err := c.simpleGet("/valorant/v1/leaderboard/"+region, map[string]string{})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &players)
	if err != nil {
		return nil, err
	}
	return players, nil
}
