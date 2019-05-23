package service

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type StatusRequest struct {
	State string
}

func (sr *StatusRequest) Valid() bool {
	return sr.State != ""
}

func StartServer(topic string, brokerAddr string, logger *logrus.Logger) error {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	state := NewState()
	s := &Server{state: state, logger: logger}

	opts := mqtt.NewClientOptions().AddBroker(brokerAddr)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("connect to mqtt: %s", token.Error())
		return token.Error()
	}

	// TODO: check how to disconnect correctly
	// TODO: graceful shutdown
	defer client.Disconnect(0)

	go startSubscriber(topic, client, s.state, s.logger)

	go func() {
		sig := <-sigs
		logger.Warnf("SIG: %s", sig)
		done <- true
	}()

	//TODO: graceful shutdown
	_, err := StartGrpcServer(state, logger)
	if err != nil {
		logger.Panicf("start grpc server", err)
	}

	logger.Println("awaiting signal")
	<-done
	logger.Println("exiting")

	return nil
}

func startSubscriber(topic string, client mqtt.Client, state *State, logger *logrus.Logger) {
	logger.Infof("starting listening topic: %s", topic)

	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		deviceID, err := getDeviceID(msg.Topic())
		if err != nil {
			logger.Errorf("wrong topic: %s", err)
			return
		}

		//TODO: add device id validation, for example check uuid

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
		logger.Error("subscription error %s", token.Error())
	}
}

func getDeviceID(topicStr string) (string, error) {
	topicParts := strings.Split(topicStr, "/")

	if len(topicParts) != 3 || topicParts[0] != "devices" || topicParts[1] == "" || topicParts[2] != "state" {
		return "", errors.New("invalid route")
	}

	return topicParts[1], nil
}
