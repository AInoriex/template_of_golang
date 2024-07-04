package define

// 同步goim/api/protocol/operation.go
const (
	// OpHandshake handshake
	OpHandshake = int32(2)
	// OpHandshakeReply handshake reply
	OpHandshakeReply = int32(2)

	// OpHeartbeat heartbeat
	OpHeartbeat = int32(0)
	// OpHeartbeatReply heartbeat reply
	OpHeartbeatReply = int32(0)

	// OpSendMsg send message.
	OpSendMsg = int32(5)
	// OpSendMsgReply  send message reply
	OpSendMsgReply = int32(5)

	// OpDisconnectReply disconnect reply
	OpDisconnectReply = int32(6)

	// OpAuth auth connnect
	OpAuth = int32(1)
	// OpAuthReply auth connect reply
	OpAuthReply = int32(1)

	// OpRaw raw message
	OpRaw = int32(11)

	// OpProtoReady proto ready
	OpProtoReady = int32(13)
	// OpProtoFinish proto finish
	OpProtoFinish = int32(14)

	// OpChangeRoom change room
	OpRoomReady  = int32(12)
	OpChangeRoom = int32(2)
	// OpChangeRoomReply change room reply
	OpChangeRoomReply = int32(2)

	// for test
	OpTest      = int32(254)
	OpTestReplu = int32(255)

	/**************************** 以下新版本定义 部分重复定义 ******************************************/

	// OpSub subscribe operation
	//OpSub = int32(14)
	OpSub = int32(22)
	// OpSubReply subscribe operation
	OpSubReply = int32(15)

	// OpUnsub unsubscribe operation
	OpUnsub = int32(16)
	// OpUnsubReply unsubscribe operation reply
	OpUnsubReply = int32(17)

	// 同步历史消息
	OpSync      = int32(18)
	OpSyncReply = int32(19)

	// 消息偏移
	OpMessageAck      = int32(20)
	OpMessageAckReply = int32(21)
)
