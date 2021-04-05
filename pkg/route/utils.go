package route

import (
	"net/http"
	"os"
	"strings"
)

var BackendURL = os.Getenv("BACKEND_URL")
var AIURL = os.Getenv("AI_URL")
var FrontendURL = os.Getenv("FRONTEND_URL")

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
