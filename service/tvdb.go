package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const TVDB_API_KEY = "9ad18ccf-ef99-409d-9abb-079357109c9c"
const TVDB_API_URI = "https://api4.thetvdb.com/v4/"

type ITVDB interface {
	Search(query string) *SearchResponse
}

type TVDB struct {
	Token string
}

func NewTVDB() ITVDB {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOiIiLCJhcGlrZXkiOiI5YWQxOGNjZi1lZjk5LTQwOWQtOWFiYi0wNzkzNTcxMDljOWMiLCJjb21tdW5pdHlfc3VwcG9ydGVkIjpmYWxzZSwiZXhwIjoxNzE2MjUyNzAzLCJnZW5kZXIiOiIiLCJoaXRzX3Blcl9kYXkiOjEwMDAwMDAwMCwiaGl0c19wZXJfbW9udGgiOjEwMDAwMDAwMCwiaWQiOiIyNDYyODI5IiwiaXNfbW9kIjpmYWxzZSwiaXNfc3lzdGVtX2tleSI6ZmFsc2UsImlzX3RydXN0ZWQiOmZhbHNlLCJwaW4iOiIiLCJyb2xlcyI6W10sInRlbmFudCI6InR2ZGIiLCJ1dWlkIjoiIn0.H0aLw5-mu3ZBS8iGf9C9AampsUw9BxhwKFUXau0xz5Qytd7ZtCEjrujbPbDCjDR5xitXuEktxM6ZNscE5vncgO57CEO1r6gYzzuRNSg4Ori9mIsgx3Q31WhTZqUI3xzqDFVVy5LYBgg481ZzvGsKsjpcc95cziYG0iRPEw_U92QZncekz3hXFRtfdo9yxC65FayIScZr7UBwJzxEZEtiDwFUrc8IVQxe5xABx8udkLaHeOcGA7R-0ecBwcPwoQSyXjtXSwNC5xKeS1xlgLwcZILy32x_xzw3TndQ_MSpWFx0enB_-sv-fQ9kMEdOZeQD97Lq4ntjFutRp2DBrlszaY8I7xL4PUxF4m72IyRuoYNYUi5VGqZswnTu1EzVSwwuRhRIy4gI87R0nFyosZEJ34LGi3l9B1w9IshAQ0tqrcU9z-WZtbvtSgm-bAvcFHjVMX0i0KSn5sKQsZy6iU7rpl9ObD8RWq20RL30nPQoq3gNwrigSPqi2a57-ozraPDzmp7PnWya6JT3f5iopznjJDqx-rfRND51KQ6ueCvj59BNuZ3HkLVjYmJwnOeTRjWaQyZoSKPhyF3p3StiE5lKdAKijgcM6K7cHLm9bokf3UaqfhBBvbrz0IJwp6dSNQ3ACwpyaIxGqHO3UWsXfJ9TfLyLzHN8tU9mCB5xi6gw8XY"

	if len(token) == 0 {
		res := &LoginResponse{}
		req := &LoginRequest{
			APIKey: TVDB_API_KEY,
		}
		jsonBody, _ := json.Marshal(req)
		bodyReader := bytes.NewReader(jsonBody)

		if resp, err := http.Post(TVDB_API_URI+"login", "application/json", bodyReader); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()
			if body, err := io.ReadAll(io.Reader(resp.Body)); err != nil {
				panic(err)
			} else {
				if err := json.Unmarshal(body, res); err != nil {
					panic(err)
				}

				token = res.Data.Token
			}
		}
	}

	return &TVDB{
		Token: token,
	}
}

func (t *TVDB) Search(query string) *SearchResponse {
	var bearer = "Bearer " + t.Token

	res := &SearchResponse{}

	if req, err := http.NewRequest("GET", TVDB_API_URI+"search", nil); err != nil {
		panic(err)
	} else {
		req.Header.Add("Authorization", bearer)

		q := req.URL.Query()
		q.Add("query", query)
		req.URL.RawQuery = q.Encode()

		client := &http.Client{}
		if resp, err := client.Do(req); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()

			if body, err := io.ReadAll(io.Reader(resp.Body)); err != nil {
				panic(err)
			} else {
				if err := json.Unmarshal(body, res); err != nil {
					panic(err)
				}
			}
		}
	}

	return res
}

type LoginRequest struct {
	APIKey string `json:"apikey"`
	PIN    string `json:"pin"`
}
type LoginResponse struct {
	Status string `json:"status"`

	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type SearchResponse struct {
	Status string `json:"status"`

	Data []struct {
		ID string `json:"id"`

		Name string `json:"name"`
		Year string `json:"year"`
	} `json:"data"`

	Links struct {
		Prev string `json:"prev"`
		Self string `json:"self"`
		Next string `json:"next"`

		TotalItems int `json:"total_items"`
		PageSize   int `json:"page_size"`
	} `json:"links"`
}
