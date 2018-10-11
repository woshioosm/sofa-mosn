package sofarpc

import (
	"github.com/alipay/sofa-mosn/pkg/types"
	"github.com/alipay/sofa-mosn/pkg/protocol/rpc"
)

var (
	sofarpcEngine = rpc.NewMixedEngine()
)

func Register(protocolCode byte, encoder types.Encoder, decoder types.Decoder) {
	sofarpcEngine.Register(protocolCode, encoder, decoder)
}

// TODO: should be replaced with configure specify(e.g. downstream_protocol: rpc, sub_protocol:[boltv1,boltv2])
func Engine() types.ProtocolEngine {
	return sofarpcEngine
}
