package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/chazari-x/ningyotsukai/domain/fonts"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "fonts",
		Short: "fonts",
		Long:  "fonts",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := getConfig(cmd)

			log.Trace("fonts starting..")
			defer log.Trace("fonts stopped")

			if err := fonts.StartServer(&cfg.Fonts); err != nil {
				log.Fatalln(err)
			}
		},
	}
	rootCmd.AddCommand(cmd)
	PersistentConfigFlags(cmd)
}
