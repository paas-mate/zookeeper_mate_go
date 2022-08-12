package config

import "github.com/paas-mate/gutil"

// zk config
var (
	ClusterEnable   bool
	TlsEnable       bool
	RemoteMode      bool
	QuorumTlsEnable bool
)

func init() {
	ClusterEnable = gutil.GetEnvBool("CLUSTER_ENABLE", false)
	TlsEnable = gutil.GetEnvBool("ZOOKEEPER_TLS_ENABLE", false)
	RemoteMode = gutil.GetEnvBool("REMOTE_MODE", true)
	QuorumTlsEnable = gutil.GetEnvBool("QUORUM_TLS_ENABLE", false)
}
