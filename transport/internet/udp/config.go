package udp

import (
	"github.com/cloudpassion/Xray-core/common"
	"github.com/cloudpassion/Xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
