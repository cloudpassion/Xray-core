package tagged

import (
	"context"

	"github.com/cloudpassion/xray-core/common/net"
	"github.com/cloudpassion/xray-core/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
