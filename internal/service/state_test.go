package service

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestState_GetStatus(t *testing.T) {
	r := require.New(t)

	s := NewState()
	deviceID1 := uuid.New().String()

	s.SetStatus(deviceID1, "STARTED")
	res, err := s.GetStatus(deviceID1)
	r.NoError(err)
	r.Equal(res.status, "STARTED")

	res, err = s.GetStatus(uuid.New().String())
	r.EqualValues(err, ErrDeviceNotFound)
}
