package route

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

var BackendURL = os.Getenv("BACKEND_URL")
var AIURL = os.Getenv("AI_URL")
var FrontendURL = os.Getenv("FRONTEND_URL")
var S3URL = os.Getenv("S3_BUCKET_URL")

func checkSelectionForm(data string, w http.ResponseWriter, r *http.Request) {
	// string型を[]stringに変換し、選択された人数を確認する準備
	formData := strings.Split(data, ",")

	// 5人選択されてなければリダイレクトさせる
	if len(formData) != 5 {
		endpoint := FrontendURL + "/selection"
		http.Redirect(w, r, endpoint, 301)
		return
	}
}

func createURL(baseURL string, queries map[string]string, pathes ...string) (createdURL string, err error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	for _, p := range pathes {
		parsedBaseURL.Path = path.Join(parsedBaseURL.Path, p)
	}
	q := parsedBaseURL.Query()
	for k, v := range queries {
		q.Add(k, v)
	}
	parsedBaseURL.RawQuery = q.Encode()
	createdURL = parsedBaseURL.String()
	return
}
