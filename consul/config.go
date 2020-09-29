package consul

import (
	"fmt"
	"reflect"
	"time"
)

type Config struct {
	ServiceName string
	ServiceID   string
	Downstream  Downstream
	Upstreams   []Upstream
}

type Upstream struct {
	Name             string
	LocalBindAddress string
	LocalBindPort    int
	Protocol         string
	ConnectTimeout   time.Duration
	ReadTimeout      time.Duration

	TLS

	Nodes []UpstreamNode
}

func (n Upstream) Equal(o Upstream) bool {
	return n.LocalBindAddress == o.LocalBindAddress &&
		n.LocalBindPort == o.LocalBindPort &&
		n.TLS.Equal(o.TLS)
}

type UpstreamNode struct {
	Host   string
	Port   int
	Weight int
}

func (n UpstreamNode) ID() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n UpstreamNode) Equal(o UpstreamNode) bool {
	return n == o
}

type Downstream struct {
	LocalBindAddress string
	LocalBindPort    int
	Protocol         string
	TargetAddress    string
	TargetPort       int
	ConnectTimeout   time.Duration
	ReadTimeout      time.Duration

	EnableForwardFor  bool
	AppNameHeaderName string

	TLS
}

func (d Downstream) Equal(o Downstream) bool {
	return reflect.DeepEqual(d, o)
}

type TLS struct {
	Cert []byte
	Key  []byte
	CAs  [][]byte
}

func (t TLS) Equal(o TLS) bool {
	return reflect.DeepEqual(t, o)
}
