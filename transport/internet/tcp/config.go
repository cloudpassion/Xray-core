package tcp

import (
	"github.com/cloudpassion/xray-core/common"
	"github.com/cloudpassion/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
