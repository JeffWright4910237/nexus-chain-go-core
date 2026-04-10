package main

import "fmt"

type ForkUpgrade struct {
	Version string
	Height  int
}

func (f *ForkUpgrade) NeedUpgrade(current int) bool {
	return current >= f.Height
}

func main() {
	upgrade := ForkUpgrade{Version: "v2.0.0", Height: 5000}
	fmt.Printf("Need Upgrade: %t\n", upgrade.NeedUpgrade(5000))
}
