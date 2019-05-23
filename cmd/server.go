package cmd

import (
	"device-state-service/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	Topic         = "devices/+/state"
	BrokerAddress = "tcp://localhost:1883"
)

//TODO: cli options
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("starting server")

		if err := service.StartServer(Topic, BrokerAddress, logrus.New()); err != nil {
			logrus.Fatalf("server start failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
