package header

import (
	"mem-ws/core/stomp"
)

const (
	DestinationHeader       = "simpDestination"
	MessageTypeHeader       = "simpMessageType"
	SessionIdHeader         = "simpSessionId"
	SessionAttributes       = "simpSessionAttributes"
	SubscriptionIdHeader    = "simpSubscriptionId"
	UserHeader              = "simpUser"
	ConnectMessageHeader    = "simpConnectMessage"
	DisconnectMessageHeader = "simpDisconnectMessage"
	HeartBeatHeader         = "simpHeartbeat"
)

type Accessor[T interface{}] struct {
	headers map[string]T
}

func (accessor *Accessor[interface{}]) SetDestination(destination string) {
	accessor.headers[DestinationHeader] = destination
}

func (accessor *Accessor[interface{}]) GetDestination() string {
	return accessor.headers[DestinationHeader].(string)
}

func (accessor *Accessor[interface{}]) SetMessageType(messageType stomp.Command) {
	accessor.headers[MessageTypeHeader] = messageType
}

func (accessor *Accessor[interface{}]) GetMessageType() stomp.Command {
	return accessor.headers[MessageTypeHeader].(stomp.Command)
}

/*************************************************** Static methods ***************************************************/

func GetDestination(headers map[string]interface{}) string {
	return headers[DestinationHeader].(string)
}

func GetMessageType(headers map[string]interface{}) stomp.Command {
	return headers[MessageTypeHeader].(stomp.Command)
}
