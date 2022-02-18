package spirit

// Action is some contribution a Spirit makes to a battle.Battle.
//
// For example, a Spirit might perform an Action that attacks all opponents, or heals all allies.
type Action interface {
	Run(*ActionContext)
}

// ActionContext provides the details of the battle.Battle in which an Action is being performed.
type ActionContext struct {
	// Me is the Spirit performing the Action.
	Me *Spirit
	// Us is the team.Team of the Spirit performing the Action.
	Us []*Spirit
	// Them is the list of Spirit's on opposing team.Team's from the Spirit performing the Action.
	Them [][]*Spirit
}
