package infrastructure

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func makeSupabaseRequest(route, method string, body io.Reader) *http.Request {
	projId := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_API_KEY")

	supabaseUrl := "https://" + projId + ".supabase.co/rest/v1/"
	requestUrl := supabaseUrl + route

	req, _ := http.NewRequest(method, requestUrl, body)
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Authorization", "Bearer "+apiKey)

	return req
}

func SupabaseGetById(entity string, id int) {
	req := makeSupabaseRequest(entity+"?id=eq."+strconv.Itoa(id), "GET", nil)
	req.Header.Set("Accept", "application/json")

}
