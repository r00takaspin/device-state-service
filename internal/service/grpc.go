package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	service "device-state-service/grpc_api"
)

type Server struct {
	state *State

	logger *logrus.Logger
}

func StartGrpcServer(state *State, port int, logger *logrus.Logger) (*Server, error) {
	srv := &Server{state: state, logger: logger}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	service.RegisterDeviceStateServiceServer(s, srv)

	logger.Infof("grpc server listening on: %s", lis.Addr().String())

	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic("cannot start grpc server")
		}
	}()

	return srv, nil
}

func (s *Server) GetDeviceState(ctx context.Context, req *service.DeviceStateRequest) (*service.DeviceStateResponse, error) {
	deviceState, err := s.state.GetStatus(req.DeviceId)
	if err != nil {
		return nil, err
	}

	return &service.DeviceStateResponse{State: deviceState.status, Timestamp: deviceState.timestamp}, nil
}
