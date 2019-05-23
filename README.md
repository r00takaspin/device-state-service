# Device State Service

You have a fleet of devices, each publishes its state over MQTT. Implement a service with gRPC API which provides most recent device state and the last state change timestamp.

- Design gRPC API yourself.
- Device publishes JSON { "state": "started" } into devices/ID/state MQTT topic every second.
- Standalone binary with MQTT and gRPC connections CLI options.
- Process should run in foreground, no daemonization needed.
- No persistence needed, states can be kept in process memory.

Usage:
```bash
# run mosquitto docker image
docker run -it -p 1883:1883 -p 9001:9001 eclipse-mosquitto

# protoc
protoc -I proto/ proto/service.proto --go_out=plugins=grpc:grpc_api/

# build and run server in separate terminal
go build && ./device-state-service server

# publish state message
./mqttcli_darwin_amd64.dms pub -t "devices/123/state" -m "{\"state\":\"STARTED\"}"

# read state from grpc 
grpc_cli call localhost:59027 service.DeviceStateService/GetDeviceState 'device_id:"123"' --proto_path proto --protofiles service.proto

```