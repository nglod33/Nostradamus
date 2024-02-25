package models

// Used to represent a player's performance during one game
type Player struct {
	Name         string
	MaxManpower  int
	ForceLimit   int
	TroopQuality int
	RealDev      int
	TotalIncome  int
}
