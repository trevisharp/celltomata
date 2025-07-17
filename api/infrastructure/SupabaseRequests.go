package infrastructure

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func makeSupabaseRequest(route, method string, body io.Reader) (*http.Request, error) {
	projId := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")

	supabaseUrl := "https://" + projId + ".supabase.co/rest/v1/"
	requestUrl := supabaseUrl + route

	req, err := http.NewRequest(method, requestUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return req, nil
}

func SupabaseGet[T any](entity string, query string) (*[]T, error) {
	req, err := makeSupabaseRequest(entity+"?"+query, "GET", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []T
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
