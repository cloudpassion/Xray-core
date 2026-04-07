package outbound

import (

	"github.com/cloudpassion/xray-core/common/net"
	"github.com/cloudpassion/xray-core/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}
