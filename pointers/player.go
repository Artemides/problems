package pointers

import "fmt"

type Player struct {
	health uint16
	name   string
}

func (player *Player) SetName(name string) {
	player.name = name
}

func (player *Player) Name() string {
	return player.name
}

func (player *Player) ReceiveSwordDmg(dmg uint16) {
	player.health -= dmg
}

func Battle() {
	player1 := Player{health: 100, name: "luz"}
	// fmt.Printf("Before battle %+v\n", player1)
	// player1.receiveSwordDmg(50)
	// fmt.Printf("After battle %+v\n", player1)
	fmt.Printf("name: %s", player1.name)
	player1.SetName("jose")
	fmt.Printf("name: %s", player1.name)

}
