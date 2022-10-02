package listeners

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	logger ILogger
	conf   NatsConf

	conn              *nats.Conn
	natsConnCtx       context.Context
	natsConnCtxCancel context.CancelFunc
}

type NatsConf struct {
	Addr string
	Port int
}

func NewNats(logger ILogger, conf NatsConf) *Nats {
	return &Nats{
		logger: logger,
		conf:   conf,
	}
}

func (n *Nats) Connect() error {
	connStr := fmt.Sprintf("nats://%s:%d", n.conf.Addr, n.conf.Port)
	conn, err := nats.Connect(connStr)
	if err != nil {
		n.logger.Error("failed to connect to the NATS at: %s", connStr)
		return err
	}
	n.conn = conn
	n.natsConnCtx, n.natsConnCtxCancel = context.WithCancel(context.Background())
	go n.monitorConnection(n.natsConnCtx)
	n.logger.Info("connected to the NATS at: %s", connStr)
	return nil
}

func (n *Nats) Disconnect() {
	defer n.natsConnCtxCancel()
	if n.conn != nil {
		n.conn.Close()
		n.logger.Info("disconnected from the NATS")
	}
}

func (n *Nats) monitorConnection(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			n.logger.Debug("NATS connection monitoring stopped")
			return
		default:
			if !n.conn.IsConnected() {
				n.logger.Error("connection to the NATS is unavailable")
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func (n *Nats) StartListening() {
	panic("not implemented") // TODO: Implement
}

func (n *Nats) StopListening() {
	panic("not implemented") // TODO: Implement
}

func (n *Nats) DumpInfo() {
	panic("not implemented") // TODO: Implement
}

func (n *Nats) StreamTraffic() {
	panic("not implemented") // TODO: Implement
}

func (n *Nats) CleanData() {
	panic("not implemented") // TODO: Implement
}
