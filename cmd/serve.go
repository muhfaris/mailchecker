package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/mailchecker/gateway/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "mailchecker",
		Short: "A verifier email",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := router.Init()

			r := fiber.New()
			app.Router(r)
			r.Listen(fmt.Sprintf(":%d", app.Port))
			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	cobra.OnInitialize(initConfig)
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("./config")
		viper.SetConfigType("toml")
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
