package hysteria

import (
	"github.com/cloudpassion/Xray-core/transport/internet/hysteria/padding"
)

var (
	tcpRequestPadding = padding.Padding{Min: 64, Max: 512}
	// tcpResponsePadding = padding.Padding{Min: 128, Max: 1024}
)
