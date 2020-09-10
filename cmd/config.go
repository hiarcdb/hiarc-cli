package cmd

import (
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
)

type HiarcConfig struct {
	Url      string      `json:"url"`
	AdminKey string      `json:"adminKey"`
	cfg      *ConfigPath `json:"-"`
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
		cfg: NewDefaultConfigPath(),
	}
}

func NewHiarcConfig(cp ConfigPath) *HiarcConfig {
	return &HiarcConfig{
		cfg: &cp,
	}
}

func (hc *HiarcConfig) AddEditUrl(url string) {
	hc.Url = url
}

func (hc *HiarcConfig) AddEditAdminKey(key string) {
	hc.AdminKey = key
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
	Short: "do something aginst configuration",
	Long: `A
	multiline
	description`,
	Run: nil,
}

var initConfigCmd = &cobra.Command{
	Use:   "init",
	Short: "create your config file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := NewHiarcConfig(*NewDefaultConfigPath())
		fmt.Println(url)
		fmt.Println(adminKey)
		viper.Set("url", url)
		viper.Set("adminKey", adminKey)
		if err := MakeCredentialsFolderIfNotExists(cfg.GetConfigPath()); err != nil {
			fmt.Println("Something went wrong creating the credentials folder.")
		}
		if err := viper.WriteConfig(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Config created")
		}
	},
}

var viewConfigCmd = &cobra.Command{
	Use:   "view",
	Short: "view your config file",
	Run: func(cmd *cobra.Command, args []string) {
		viper.ReadInConfig()
		url := viper.Get("url")
		fmt.Println("URL:")
		fmt.Println(url)
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

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(initConfigCmd)
	configCmd.AddCommand(viewConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	setConfigCmd.AddCommand(setUrlConfigCmd)

	initConfigCmd.Flags().StringVar(&url, "url", "", "Hiarc API URL (required)")
	initConfigCmd.MarkFlagRequired("url")
	initConfigCmd.Flags().StringVar(&adminKey, "adminKey", "", "Hiarc Admin Key (required)")
	initConfigCmd.MarkFlagRequired("adminKey")
}
