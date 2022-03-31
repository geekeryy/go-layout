package main

import (
	"context"
	"log"

	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xpprof"
	"github.com/spf13/cobra"
)

func init() {
	xenv.Init(map[string]string{
		xenv.AppName:     "go-layout",
		xenv.AppVersion:  "v1.0",
		xenv.ApolloAppID: "go-layout",
	})
	log.Println("APP_ENV", xenv.GetEnv(xenv.AppEnv))
}

var rootCmd = &cobra.Command{
	Run: func(c *cobra.Command, args []string) {
		xpprof.Launch("localhost:" + xenv.GetEnv(xenv.PprofPort))

		logger := xlog.New(xlog.WithTrace(xenv.GetEnv(xenv.TraceName)))

		ctx, cancel := context.WithCancel(context.Background())

		app := InitApp(ctx, logger)

		if err := app.Run(cancel); err != nil {
			log.Println("Server Exit:", err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
