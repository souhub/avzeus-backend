package dmm

import (
	"errors"
	"fmt"
	"os"

	"github.com/dmmlabo/dmm-go-sdk/api"
	"github.com/souhub/avzeus-backend/pkg/model"
)

var ApiID = os.Getenv("DMM_API_ID")
var AffiliateID = os.Getenv("DMM_AFFILIATE_ID")

// レコメンドされた女優データに、DMMから取得したデータ追加
func AddDataToActresses(recommendedActresses model.Actresses) (updatedActresses model.Actresses, err error) {
	for _, recommendedActress := range recommendedActresses {
		updatedActress, err := addDataToActress(recommendedActress)
		if err != nil {
			msg := fmt.Sprintf("Failed to fetch Actress from DMM. %v", err)
			err = errors.New(msg)
		}
		updatedActresses = append(updatedActresses, updatedActress)
	}
	return
}

func addDataToActress(recommendedActress model.Actress) (updatedActress model.Actress, err error) {
	var actressAPI = api.NewActressService(AffiliateID, ApiID)
	actressAPI.SetKeyword(recommendedActress.Name)
	actressResp, err := actressAPI.Execute()
	if err != nil {
		msg := fmt.Sprintf("Failed to get an actress from DMM API. %v", err)
		err = errors.New(msg)
		return
	}
	dmmActress := actressResp.Actresses[0]
	updatedActress = model.Actress{
		ID:        recommendedActress.ID,
		Name:      recommendedActress.Name,
		ImagePath: dmmActress.ImageURL.Large,
		ListURL:   dmmActress.ListURL,
	}
	return
}
