package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	hiarc "github.com/allenmichael/hiarcgo"
	"github.com/spf13/viper"
)

func ConfigureHiarcClientWithValues(url string, adminKey string) *hiarc.APIClient {
	cfg := hiarc.NewConfiguration()
	cfg.BasePath = url
	cfg.AddDefaultHeader("X-Hiarc-Api-Key", adminKey)
	return hiarc.NewAPIClient(cfg)
}
func ConfigureHiarcClientWithToken(url string, token string) *hiarc.APIClient {
	cfg := hiarc.NewConfiguration()
	cfg.BasePath = url
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
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
	if profile == "default" {
		k := os.Getenv(HiarcProfileEnvVar)
		if k != "" {
			profile = k
		}
	}
	url, admin := GetConfigValuesByProfile(profile)
	token, _ := rootCmd.Flags().GetString("token")
	if token != "" {
		return ConfigureHiarcClientWithToken(url, token)
	}
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
func IsValidAccessLevel(a string) bool {
	levels := []string{string(hiarc.CO_OWNER), string(hiarc.READ_WRITE), string(hiarc.READ_ONLY), string(hiarc.UPLOAD_ONLY)}
	for _, item := range levels {
		if item == a {
			return true
		}
	}
	return false
}

func GetAccessLevelFromString(a string) (hiarc.AccessLevel, error) {
	if a == string(hiarc.CO_OWNER) {
		return hiarc.CO_OWNER, nil
	}
	if a == string(hiarc.READ_WRITE) {
		return hiarc.READ_WRITE, nil
	}
	if a == string(hiarc.READ_ONLY) {
		return hiarc.READ_ONLY, nil
	}
	if a == string(hiarc.UPLOAD_ONLY) {
		return hiarc.UPLOAD_ONLY, nil
	}
	var emp hiarc.AccessLevel
	return emp, errors.New("Couldn't convert this Access Level")
}
