package conn

import (
	"net/rpc"

	log "github.com/Sirupsen/logrus"
)

type Connector struct {
	pool map[string](*rpc.Client)
}

func NewConnector() (connector *Connector) {
	connector = new(Connector)
	connector.pool = make(map[string](*rpc.Client))
	return
}

func (this *Connector) Connect(host string) (client *rpc.Client, err error) {

	client, yes = this.pool[host]
	if yes {
		return
	}
	log.DebugInfo()
	client, err = rpc.DialHTTP(PROTO, host)
	if err != nil {
		return
	}
}
