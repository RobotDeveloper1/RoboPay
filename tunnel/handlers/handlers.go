package handlers

import "github.com/fabricfoundation/robot-tunnel-client/tunnel"

// RobotID returns a handler that responds with the robot's ID.
func RobotID(robotID string) tunnel.Handler {
	return func(method string, headers map[string]string, body []byte) (int, map[string]string, []byte) {
		return 200, map[string]string{"content-type": "text/plain"}, []byte(robotID)
	}
}
