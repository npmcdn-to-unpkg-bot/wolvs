package common

type Role interface {
}
type Player interface {
}
type Settings struct {
	Roles       []Role
	PlayerCount int
}
type Lobby struct {
	Players []Player
}
