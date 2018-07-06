package config

import (
	"time"

	"github.com/bio-routing/bio-rd/routingtable/locRIB"

	bnet "github.com/bio-routing/bio-rd/net"
	"github.com/bio-routing/bio-rd/routingtable"
	"github.com/bio-routing/bio-rd/routingtable/filter"
)

// Peer defines the configuration for a BGP session
type Peer struct {
	AdminEnabled            bool
	ReconnectInterval       time.Duration
	KeepAlive               time.Duration
	HoldTime                time.Duration
	LocalAddress            bnet.IP
	PeerAddress             bnet.IP
	LocalAS                 uint32
	PeerAS                  uint32
	Passive                 bool
	RouterID                uint32
	AddPathSend             routingtable.ClientOptions
	AddPathRecv             bool
	RouteServerClient       bool
	RouteReflectorClient    bool
	RouteReflectorClusterID uint32
	IPv4                    *AddressFamilyConfig
	IPv6                    *AddressFamilyConfig
}

// AddressFamilyConfig represents all configuration parameters specific for an address family
type AddressFamilyConfig struct {
	RIB          *locRIB.LocRIB
	ImportFilter *filter.Filter
	ExportFilter *filter.Filter
}
