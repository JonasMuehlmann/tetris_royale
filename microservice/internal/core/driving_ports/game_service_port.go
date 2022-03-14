package drivingPorts

import "net"

type GameServicePort interface {
	StartGrpcServer(net.Listener) error
	ConnectPlayer(userID string, connection interface{}) error
}
