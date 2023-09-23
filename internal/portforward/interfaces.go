package portforward

import "context"

type Service interface {
	Start(ctx context.Context) (runError <-chan error, err error)
	Stop() (err error)
	GetPortForwarded() (port uint16)
}

type PortAllower interface {
	SetAllowedPort(ctx context.Context, port uint16, intf string) (err error)
	RemoveAllowedPort(ctx context.Context, port uint16) (err error)
}

type Logger interface {
	Info(s string)
	Warn(s string)
	Error(s string)
}
