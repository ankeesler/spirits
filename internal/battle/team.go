package battle

import "github.com/ankeesler/spirits0/internal/spirit"

type Team struct {
	name    string
	spirits []*spirit.Spirit
}

func NewTeam(name string, spirits []*spirit.Spirit) *Team {
	return &Team{
		name:    name,
		spirits: spirits,
	}
}

func (t *Team) Name() string              { return t.name }
func (t *Team) Spirits() []*spirit.Spirit { return t.spirits }
