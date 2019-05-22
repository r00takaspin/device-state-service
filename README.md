# Device State Service

You have a fleet of devices, each publishes its state over MQTT. Implement a service with gRPC API which provides most recent device state and the last state change timestamp.

- Design gRPC API yourself.
- Device publishes JSON { "state": "started" } into devices/ID/state MQTT topic every second.
- Standalone binary with MQTT and gRPC connections CLI options.
- Process should run in foreground, no daemonization needed.
- No persistence needed, states can be kept in process memory.