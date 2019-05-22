package service

import (
	service "device-state-service/grpc_api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

var ErrDeviceNotFound = status.Error(codes.NotFound, "device not found")

type DeviceState struct {
	timestamp int64
	status service.DeviceStatus
}

func (ds *DeviceState) Copy() *DeviceState {
	return &DeviceState{timestamp: ds.timestamp, status: ds.status}
}

type State struct {
	m sync.Mutex

	devices map[string]*DeviceState
}

func NewState() *State {
	return &State{m: sync.Mutex{}, devices: map[string]*DeviceState{}}
}

func (s *State) GetStatus(deviceID string) (*DeviceState, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if state, ok := s.devices[deviceID]; ok {
		return state.Copy(), nil
	}

	return nil, ErrDeviceNotFound
}

func (s *State) SetStatus(deviceID string, status service.DeviceStatus) {
	s.m.Lock()
	s.devices[deviceID] = &DeviceState{timestamp: time.Now().UnixNano(), status: status}
	s.m.Unlock()
}
