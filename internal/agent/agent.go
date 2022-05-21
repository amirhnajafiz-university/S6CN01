package agent

type Agent struct {
	Util Util
}

func New() Agent {
	return Agent{
		Util: Util{},
	}
}

func (a Agent) Start() {
	go func() {

	}()
}
