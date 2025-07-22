package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("supabase error (%d): %s", resp.StatusCode, string(body))
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

func SupabasePost[T any](entity string, data T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := makeSupabaseRequest(entity, "POST", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("supabase error (%d)", resp.StatusCode)
	}

	return nil
}

func SupabaseDelete(entity, query string) error {
	req, err := makeSupabaseRequest(entity+"?"+query, "DELETE", nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("supabase error (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}

func SupabasePatch[T any](entity string, id int, data T) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := makeSupabaseRequest(fmt.Sprintf(entity+"?ID.eq=%d", id), "PATCH", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("supabase error (%d)", resp.StatusCode)
	}

	return nil
}
