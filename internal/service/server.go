package service

import (
	"encoding/json"
	"os"
	"os/signal"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type StatusRequest struct {
	State string
}

func (sr *StatusRequest) Valid() bool {
	return sr.State != ""
}

func StartServer(topic string, brokerAddr string, listenersNum int, port int, logger *logrus.Logger) error {
	if listenersNum < 0 {
		return errors.New("invalid listeners number")
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	state := NewState()
	s := &Server{state: state, logger: logger}

	logger.Infof("successfully connected to broker: %v", brokerAddr)

	for i := 0; i < listenersNum; i++ {
		go startSubscriber(topic, brokerAddr, s.state, s.logger)
	}

	_, err := StartGrpcServer(state, port, logger)
	if err != nil {
		logger.Panicf("start grpc server: %v", err)
	}

	go func() {
		sig := <-sigs
		logger.Warnf("SIG: %s", sig)
		done <- true
	}()

	logger.Println("awaiting signal")
	<-done
	logger.Println("exiting")

	return nil
}

func startSubscriber(topic string, brokerAddr string, state *State, logger *logrus.Logger) {
	client := mqtt.NewClient(mqtt.NewClientOptions().AddBroker(brokerAddr))
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Errorf("connect to broker: %v", token.Error())
	}

	logger.Infof("starting listening topic: %s", topic)

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		deviceID, err := getDeviceID(msg.Topic())
		if err != nil {
			logger.Errorf("wrong topic: %s", err)
			return
		}

		jsonReq := &StatusRequest{}
		if err := json.Unmarshal(msg.Payload(), jsonReq); err != nil {
			logger.Errorf("unmarshal json: %s", err)
			return
		}
		if !jsonReq.Valid() {
			logger.Error("invalid request")
			return
		}

		logger.Infof("setting device %s state %s", deviceID, jsonReq.State)
		state.SetStatus(deviceID, jsonReq.State)

	}); token.Wait() && token.Error() != nil {
		logger.Errorf("subscription error %s", token.Error())
	}
}

func getDeviceID(topicStr string) (string, error) {
	topicParts := strings.Split(topicStr, "/")

	if len(topicParts) != 3 || topicParts[0] != "devices" || topicParts[1] == "" || topicParts[2] != "state" {
		return "", errors.New("invalid route")
	}

	return topicParts[1], nil
}
