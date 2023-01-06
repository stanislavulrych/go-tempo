//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package tempo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type TempoApi struct {
	AuthToken string
	client    *http.Client
}

const baseUrl = "https://api.tempo.io/4"
const limit = 5000

func NewTempoApi(authToken string) *TempoApi {
	return &TempoApi{
		AuthToken: authToken,
		client:    &http.Client{},
	}
}
func (r TempoApi) request(suffix string, params map[string]string) []byte {

	req, _ := http.NewRequest("GET", baseUrl+suffix, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.AuthToken)

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}

	var res []byte
	q.Set("offset", strconv.Itoa(0))
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()
	resp, err := r.client.Do(req)

	if err != nil {
		fmt.Println("Error on response: ", err)
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on reading: ", err)
	}

	var tr TempoResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println("Error on unmarshal META: ", err)
	}

	if tr.Metadata.Count >= limit {
		fmt.Println("ERROR: not all data loaded")
		// TODO handle pagination
	}
	res = body

	return res
}

func (r TempoApi) GetWorklogs(params map[string]string) *TempoWorklogs {

	url := "/worklogs"
	bytes := r.request(url, params)

	var res TempoResponseWorklogs
	err := json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println("Error on unmarshal: ", err)
		fmt.Println(string(bytes))
	}

	return &res.Results
}

func (r TempoApi) GetAccounts() *TempoAccounts {
	params := map[string]string{}

	url := "/accounts"
	bytes := r.request(url, params)

	var res TempoResponseAccounts
	err := json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println("Error on unmarshal: ", err)
		fmt.Println(string(bytes))
	}

	return &res.Results
}

func (r TempoApi) GetHolidaySchemes() *TempoHolidaySchemes {
	params := map[string]string{}

	url := "/holiday-schemes"
	bytes := r.request(url, params)

	var res TempoResponseHolidaySchemes
	err := json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println("Error on unmarshal: ", err)
		fmt.Println(string(bytes))
	}

	return &res.Results
}
