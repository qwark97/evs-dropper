package listeners

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

const (
	NATS_CH_BUFF         = 10
	SUBSCRIPTION_PATTERN = ">"
)

var (
	errListeningAlreadyOn  = errors.New("listening already on")
	errListeningAlreadyOff = errors.New("listening already off")
)

type Nats struct {
	logger ILogger
	conf   NatsConf

	conn *nats.Conn

	eventsChan   chan *nats.Msg
	subscription *nats.Subscription

	isListeningState     bool
	listeningSwitchMutex *sync.Mutex
}

type NatsConf struct {
	Addr string
	Port int
}

func (n *Nats) setListeningOn() error {
	n.listeningSwitchMutex.Lock()
	defer n.listeningSwitchMutex.Unlock()
	if n.isListeningState {
		return errListeningAlreadyOn
	}
	n.isListeningState = true
	return nil
}

func (n *Nats) setListeningOff() error {
	n.listeningSwitchMutex.Lock()
	defer n.listeningSwitchMutex.Unlock()
	if !n.isListeningState {
		return errListeningAlreadyOff
	}
	n.isListeningState = false
	return nil
}

func msgHandler(msg *nats.Msg) {
	fmt.Println(string(msg.Data))
}

func NewNats(logger ILogger, conf NatsConf) *Nats {
	eventsChan := make(chan *nats.Msg, NATS_CH_BUFF)
	n := &Nats{
		logger:               logger,
		conf:                 conf,
		eventsChan:           eventsChan,
		listeningSwitchMutex: new(sync.Mutex),
	}
	return n
}

func (n *Nats) Connect() error {
	connStr := fmt.Sprintf("nats://%s:%d", n.conf.Addr, n.conf.Port)
	conn, err := nats.Connect(connStr)
	if err != nil {
		n.logger.Error("failed to connect to the NATS at: %s", connStr)
		return err
	}
	n.conn = conn
	n.logger.Info("connected to the NATS at: %s", connStr)
	return nil
}

func (n *Nats) Disconnect() {
	if n.conn != nil {
		n.conn.Close()
		n.logger.Info("disconnected from the NATS")
	}
}

func (n *Nats) StartListening() error {
	if err := n.setListeningOn(); err != nil {
		return nil // deliberately do not return an error
	}
	s, err := n.conn.Subscribe(SUBSCRIPTION_PATTERN, msgHandler)
	if err != nil {
		return err
	}
	n.subscription = s
	return nil
}

func (n *Nats) StopListening() error {
	if err := n.setListeningOff(); err != nil {
		return nil // deliberately do not return an error
	}
	defer func() {
		n.subscription = nil
	}()

	err := n.subscription.Unsubscribe()
	if err != nil {
		n.logger.Error(err.Error())
	}

	return nil
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
