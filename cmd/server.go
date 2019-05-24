package cmd

import (
	"device-state-service/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	topic                 = "devices/+/state"
	defaultBrokerAddress  = "tcp://localhost:1883"
	defaultGrpcServerPort = 5678
	defaultListenersNum   = 10
)

type Config struct {
	broker       string
	port         int
	listenersNum int
}

var conf Config

var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Infof("starting server with %d listeners", conf.listenersNum)

		if err := service.StartServer(topic, conf.broker, conf.listenersNum, conf.port, logger); err != nil {
			logrus.Fatalf("server start failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVar(&conf.broker, "b", defaultBrokerAddress, "mqtt broker address")
	serverCmd.Flags().IntVar(&conf.port, "p", defaultGrpcServerPort, "grpc server port")
	serverCmd.Flags().IntVar(&conf.listenersNum, "n", defaultListenersNum, "topic listeners count")
}
