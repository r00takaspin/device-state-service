package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestState_GetStatus(t *testing.T) {
	r := require.New(t)

	s := NewState()
	deviceID1 := uuid.New().String()

	s.SetStatus(deviceID1, "STARTED")
	res1, err := s.GetStatus(deviceID1)
	r.NoError(err)
	r.Equal(res1.status, "STARTED")

	_, err = s.GetStatus(uuid.New().String())
	r.EqualValues(err, ErrDeviceNotFound)
}
