package main

import (
	"fmt"
	"go101/entity"
)

func main() {
	redMine := entity.NewBase("Land-1", 0)
	for i := 0; i < 5; i++ {
		redMine.Spawn(fmt.Sprintf("Car-%d", i))
	}
	// redMine.Spawn("Car-1")
	// redMine.Spawn("Car-2")
	fmt.Println(redMine.GetMiner())
	// redMine.SendToMining("Car-1")
	redMine.ListenMiner()
}
