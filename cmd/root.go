package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"os"
)

var (
	//rootCmd represents the base command when called without any sub commands
	rootCmd = &cobra.Command{
		Use:   "go_clean_arch",
		Short: "go_clean_arch",
		Long:  "go_clean_arch",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvPrefix("clean")
	_ = viper.BindEnv("env")
	_ = viper.BindEnv("consul_url")
	_ = viper.BindEnv("consul_path")

	consulURL := viper.GetString("consul_url")
	consulPATH := viper.GetString("consul_path")

	if consulURL == ""{
		log.Println("CONSUL_URL missing")
		os.Exit(1)
	}

	if consulPATH == ""{
		log.Println("CONSUL_PATH missing")
		os.Exit(1)
	}

	_ = viper.AddRemoteProvider("consul", consulURL, consulPATH)
	viper.SetConfigType("yml")
	err := viper.ReadRemoteConfig()
	if err != nil{
		log.Fatal(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPATH))
	}
}
