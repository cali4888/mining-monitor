package monitor

type Strategy interface {
	Run(m Monitor)
}
