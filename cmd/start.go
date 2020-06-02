package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stetsd/monk-api/internal/app"
	"github.com/stetsd/monk-conf"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start http server",
	Long:  `start http server`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.EnvParseToConfigMap()

		if err != nil {
			panic(err)
		}

		server := app.NewServer(conf)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
