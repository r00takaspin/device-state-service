package service

import (
	"context"
	grpc_api "device-state-service/grpc_api"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	state  *State
	logger *logrus.Logger
}

func StartGrpcServer(state *State, port int, logger *logrus.Logger) (*Server, error) {
	srv := &Server{state: state, logger: logger}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	grpc_api.RegisterDeviceStateServiceServer(s, srv)

	logger.Infof("grpc server listening on: %s", lis.Addr().String())

	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			panic("cannot start grpc server")
		}
	}()

	return srv, nil
}

func (s *Server) GetDeviceState(ctx context.Context, req *grpc_api.DeviceStateRequest) (*grpc_api.DeviceStateResponse, error) {
	deviceState, err := s.state.GetStatus(req.DeviceId)
	if err != nil {
		return nil, err
	}

	return &grpc_api.DeviceStateResponse{State: deviceState.status, Timestamp: deviceState.timestamp}, nil
}
