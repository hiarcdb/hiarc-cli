package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	HiarcDirName          = ".hiarc"
	HiarcConfigFileName   = "config"
	HiarcConfigFileFormat = ".json"
)

var (
	url      string
	adminKey string
	profile  string
)

type HiarcConfigValues struct {
	Url         string `json:"url"`
	AdminKey    string `json:"adminKey"`
	ProfileName string `json:"profile"`
}

type HiarcConfig struct {
	Configs map[string]*HiarcConfigValues `json:"configs"`
	cfg     *ConfigPath                   `json:"-"`
}

type ConfigPath struct {
	cfgPath     string
	cfgFilePath string
	fileName    string
	fileType    string
}

func NewDefaultConfigPath() *ConfigPath {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	cfgPath := filepath.Join(user.HomeDir, HiarcDirName)
	cfgFilePath := filepath.Join(cfgPath, HiarcConfigFileName+HiarcConfigFileFormat)
	return &ConfigPath{
		cfgPath:     cfgPath,
		fileType:    HiarcConfigFileFormat,
		cfgFilePath: cfgFilePath,
		fileName:    HiarcConfigFileName,
	}
}

func NewDefaultHiarcConfig() *HiarcConfig {
	return &HiarcConfig{
		Configs: make(map[string]*HiarcConfigValues),
		cfg:     NewDefaultConfigPath(),
	}
}

func NewHiarcConfig(cp ConfigPath) *HiarcConfig {
	return &HiarcConfig{
		Configs: make(map[string]*HiarcConfigValues),
		cfg:     &cp,
	}
}

func (hc *HiarcConfig) AddNewConfig(adminKey string, url string, profile string) {
	if profile == "" {
		profile = "default"
	}
	if v, ok := hc.Configs[profile]; ok {
		v.Url = url
		v.AdminKey = adminKey
	} else {
		hc.Configs[profile] = &HiarcConfigValues{
			Url:         url,
			AdminKey:    adminKey,
			ProfileName: profile,
		}
	}
}
func (hc *HiarcConfig) AddEditUrl(url string, profile string) {
	if profile == "" {
		profile = "default"
	}
	if v, ok := hc.Configs[profile]; ok {
		v.Url = url
	}
}

func (hc *HiarcConfig) AddEditAdminKey(key string, profile string) {
	if profile == "" {
		profile = "default"
	}
	if v, ok := hc.Configs[profile]; ok {
		v.AdminKey = key
	}
}

func (hc *HiarcConfig) GetConfigPath() string {
	return hc.cfg.cfgPath
}
func (hc *HiarcConfig) GetConfigFilePath() string {
	return hc.cfg.cfgFilePath
}

func MakeCredentialsFolderIfNotExists(path string) error {
	if _, fileErr := os.Stat(path); fileErr != nil {
		fmt.Println(fileErr)
		dirErr := os.MkdirAll(path, os.ModePerm)
		if dirErr != nil {
			fmt.Println(dirErr)
			return dirErr
		}
	}
	return nil
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "do something against configuration",
	Run:   nil,
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "create your config file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := NewDefaultHiarcConfig()
		fmt.Println(url)
		fmt.Println(adminKey)
		cfg.AddNewConfig(adminKey, url, profile)
		for key, value := range cfg.Configs {
			viper.Set(key, value)
		}
		if err := MakeCredentialsFolderIfNotExists(cfg.GetConfigPath()); err != nil {
			fmt.Println("Something went wrong creating the credentials folder.")
		}
		if err := viper.SafeWriteConfigAs(cfg.GetConfigFilePath()); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Config created")
		}
	},
}
var addConfigCmd = &cobra.Command{
	Use:   "add [profile name]",
	Short: "add a new profile to your config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := NewDefaultHiarcConfig()
		c := viper.AllSettings()
		fmt.Println(c)
		for p := range c {
			fmt.Println(c[p])
			jsonbody, err := json.Marshal(c[p])
			if err != nil {
				fmt.Println(err)
				return
			}
			config := HiarcConfigValues{}
			if err := json.Unmarshal(jsonbody, &config); err != nil {
				fmt.Println(err)
				return
			}
			cfg.AddNewConfig(config.AdminKey, config.Url, config.ProfileName)
		}

		cfg.AddNewConfig(adminKey, url, args[0])
		fmt.Println(cfg)
		for key, value := range cfg.Configs {
			viper.Set(key, value)
		}
		if err := MakeCredentialsFolderIfNotExists(cfg.GetConfigPath()); err != nil {
			fmt.Println("Something went wrong creating the credentials folder.")
		}
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Config profile added.")
		}
	},
}

var deleteConfigCmd = &cobra.Command{
	Use:   "delete [profile name]",
	Short: "delete a profile from your config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dp := viper.Get(args[0])
		if dp == nil {
			fmt.Println("Couldn't find this profile")
		} else {
			configMap := viper.AllSettings()
			delete(configMap, args[0])
			encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
			err := viper.ReadConfig(bytes.NewReader(encodedConfig))
			if err != nil {
				log.Fatal(err)
			}
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("Config profile deleted.")
			}
		}
	},
}

var viewConfigCmd = &cobra.Command{
	Use:   "view [profile name]",
	Short: "view a profile in your config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.Get(args[0])
		if p == nil {
			fmt.Println(fmt.Sprintf("Couldn't find a profile named %s", args[0]))
		} else {
			encodedConfig, _ := json.MarshalIndent(p, "", " ")
			log.Println(string(encodedConfig))
		}
	},
}

var viewAllConfigCmd = &cobra.Command{
	Use:   "all",
	Short: "view all of your configs",
	Run: func(cmd *cobra.Command, args []string) {
		configMap := viper.AllSettings()
		encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
		log.Println(string(encodedConfig))
	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "set a value in your config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Provide item from config to set")
	},
}

var setUrlConfigCmd = &cobra.Command{
	Use:   "url",
	Short: "set a URL in your config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Needs URL as argument")
		}
		url := args[0]
		viper.Set("url", url)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("URL updated.")
		}
		return nil
	},
}
var setAdminKeyConfigCmd = &cobra.Command{
	Use:   "adminKey",
	Short: "set an admin key in your config file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Needs your admin key as argument")
		}
		adminKey := args[0]
		viper.Set("adminKey", adminKey)
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Admin key updated.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(addConfigCmd)
	configCmd.AddCommand(deleteConfigCmd)

	setConfigCmd.AddCommand(setUrlConfigCmd)
	setConfigCmd.AddCommand(setAdminKeyConfigCmd)

	viewConfigCmd.AddCommand(viewAllConfigCmd)

	initConfigCmd.Flags().StringVar(&url, "url", "", "Hiarc API URL (required)")
	initConfigCmd.MarkFlagRequired("url")
	initConfigCmd.Flags().StringVar(&adminKey, "adminKey", "", "Hiarc Admin Key (required)")
	initConfigCmd.MarkFlagRequired("adminKey")
	initConfigCmd.Flags().StringVar(&profile, "profile", "default", "Hiarc Profile")

	addConfigCmd.Flags().StringVar(&url, "url", "", "Hiarc API URL (required)")
	addConfigCmd.MarkFlagRequired("url")
	addConfigCmd.Flags().StringVar(&adminKey, "adminKey", "", "Hiarc Admin Key (required)")
	addConfigCmd.MarkFlagRequired("adminKey")
	// addConfigCmd.Flags().StringVar(&profile, "profile", "", "Hiarc Profile name")
	// addConfigCmd.MarkFlagRequired("profile")

	setUrlConfigCmd.Flags().StringVar(&profile, "profile", "default", "Hiarc Profile")
	setAdminKeyConfigCmd.Flags().StringVar(&profile, "profile", "default", "Hiarc Profile")
}
