package cmd

import (
	"encoding/json"
	"fmt"

	hiarc "github.com/allenmichael/hiarcgo"
	"github.com/spf13/viper"
)

func ConfigureHiarcClientWithValues(url string, adminKey string) *hiarc.APIClient {
	cfg := hiarc.NewConfiguration()
	cfg.BasePath = url
	cfg.AddDefaultHeader("X-Hiarc-Api-Key", adminKey)
	return hiarc.NewAPIClient(cfg)
}

func GetConfigUrlByProfile(profile string) string {
	return viper.GetString(fmt.Sprintf("%s.url", profile))
}

func GetConfigAdminKeyByProfile(profile string) string {
	return viper.GetString(fmt.Sprintf("%s.adminKey", profile))
}

func GetConfigValuesByProfile(profile string) (string, string) {
	return viper.GetString(fmt.Sprintf("%s.url", profile)), viper.GetString(fmt.Sprintf("%s.adminKey", profile))
}

func ConfigureHiarcClient() *hiarc.APIClient {
	profile, _ := rootCmd.Flags().GetString("profile")
	url, admin := GetConfigValuesByProfile(profile)
	return ConfigureHiarcClientWithValues(url, admin)
}

func ConvertMetadataStringToObject(md string) (map[string]interface{}, error) {
	var mdo map[string]interface{}
	err := json.Unmarshal([]byte(md), &mdo)
	if err != nil {
		return nil, err
	}
	return mdo, nil
}
func ConvertQueryToObject(q string) (map[string]interface{}, error) {
	var qo map[string]interface{}
	err := json.Unmarshal([]byte(q), &qo)
	if err != nil {
		return nil, err
	}
	return qo, nil
}
