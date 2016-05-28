package common

//Role is a description and name for a role a seated player will have
type Role struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	MaximumRatio int    `json:"maximumRatio"` //0 indicates only one of this role may be present, every other number is assumed out of 10.
}

//RoleInteraction is coupled with a Role to determine outcomes based on certain
//EventIDs
type RoleInteraction interface {
	onEvent(eventIDs []int)
}

//EventIDEaten and other EventIDs are passed to RoleInteraction
const (
	EventIDEaten = iota
	EventIDHung
	EventIDWon
	EventIDLost
)

//Player is a seat in a lobby that will be playing the game
type Player struct {
	role Role
}

//Screen is an observer in a lobby that will be watching the game
type Screen struct {
}

//Settings are passed to the API to create a Lobby.
type Settings struct {
	Roles       map[Role]int
	PlayerCount int
}

//Lobby is where players are grouped to play a game. It can be joined.
type Lobby struct {
	Screens []Screen
	Players []Player
}
