package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shortener/internal/app"
	"shortener/internal/commands/cmdargs"
	"shortener/internal/depgraph"
	"shortener/internal/repository"
)

func InitRunCommand() (*cobra.Command, error) {
	cmdArgs := cmdargs.RunArgs{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Starts server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			viper.AddConfigPath(".")
			viper.SetConfigName("local")
			viper.SetConfigType("env")

			err := viper.ReadInConfig()
			if err != nil {
				return err
			}
			err = viper.Unmarshal(&cmdArgs)
			if err != nil {
				return err
			}

			dg := depgraph.NewDepGraph(cmdArgs)

			logger, _ := dg.GetLogger()
			logger.Debugw(
				"Got config",
				"args", cmdArgs,
			)

			entrypoint := app.NewEntrypoint(repository.NewRepository(), cmdArgs)

			err = entrypoint.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&cmdArgs.DatabaseURL, "db-url", "d",
		"postgresql://localhost:5432@postgres:postgres/shortener", "Database URL")

	return cmd, nil
}
