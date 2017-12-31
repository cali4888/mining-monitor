package monitor

type Monitor interface {
	Start(c chan error, stop chan bool)
}
