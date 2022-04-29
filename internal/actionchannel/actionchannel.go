package actionchannel

type ActionChannel struct {
	// battleName -> spiritName -> spiritGeneration -> actionName
	m map[string]map[string]map[string]string
}

func New() *ActionChannel {
	return &ActionChannel{
		m: make(map[string]map[string]map[string]string),
	}
}

func (ac *ActionChannel) Post(battleName, spiritName, spiritGeneration, actionName string) error {
	return nil
}

func (ac *ActionChannel) Pend(battleName, spiritName, spiritGeneration string) (string, error) {
	return "", nil
}
