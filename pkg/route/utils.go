package route

import (
	"net/http"
	"os"
	"strings"
)

var BackendURL = os.Getenv("BACKEND_URL")
var AIURL = os.Getenv("AI_URL")
var FrontendURL = os.Getenv("FRONTEND_URL")

// var BackendURL = "http://localhost:8000"
// var AIURL = "http://localhost:5000"
// var FrontendURL = "http://localhost:8000"

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

// func convertStrToIntArray(s string) (intArray []int, err error) {
// 	strArray := strings.Split(s, ",")
// 	for i := 0; i < len(strArray); i++ {
// 		num, err := strconv.Atoi(strArray[i])
// 		if err != nil {
// 			err = errors.New("Failed to Atooi")
// 		}
// 		intArray = append(intArray, num)
// 	}
// 	return intArray, err
// }
