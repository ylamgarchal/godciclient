/*
Copyright (C) Red Hat, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dci

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Topic represents the topic resource from the server response
type Topic struct {
	Name string
	ID   string
}

// Client is the dci client api
type Client struct {
	baseURL    string
	httpClient *http.Client
	username   string
	password   string
}

// GetClient returns an initiliazed Client
func GetClient(baseURL, username, password string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		username:   username,
		password:   password,
	}
}

// GetTopicByName returns a Topic struct with its name
func (c *Client) GetTopicByName(name string) (*Topic, error) {
	req, _ := http.NewRequest("GET", c.baseURL+"/topics?where=name:"+name, nil)
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(req)
	if resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()

	var topics map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&topics)
	if err != nil {
		return nil, err
	}

	if _, in := topics["topics"]; !in {
		return nil, errors.New("topics key not in result")
	}

	if _, isSlice := topics["topics"].([]interface{}); !isSlice {
		return nil, errors.New("topics value is not a slice")
	}
	topicsSlice := topics["topics"].([]interface{})

	if len(topicsSlice) == 0 {
		return nil, errors.New("topic " + name + " not found")
	}

	if _, isMap := topicsSlice[0].(map[string]interface{}); !isMap {
		return nil, errors.New("topic structure is not a map")
	}
	var theTopic = topicsSlice[0].(map[string]interface{})

	return &Topic{
		Name: theTopic["name"].(string),
		ID:   theTopic["id"].(string),
	}, nil
}
