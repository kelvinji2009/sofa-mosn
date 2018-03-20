package proxy

import (
	"gitlab.alipay-inc.com/afe/mosn/pkg/api/v2"
	"gitlab.alipay-inc.com/afe/mosn/pkg/server"
	"gitlab.alipay-inc.com/afe/mosn/pkg/types"
	"gitlab.alipay-inc.com/afe/mosn/pkg/stream/tcp"
)

type TcpProxyFilterConfigFactory struct {
	Proxy *v2.TcpProxy
}

func (tpcf *TcpProxyFilterConfigFactory) CreateFilterFactory(clusterManager types.ClusterManager) server.NetworkFilterFactoryCb {
	return func(manager types.FilterManager) {
		manager.AddReadFilter(tcp.NewProxy(tpcf.Proxy, clusterManager))
	}
}
