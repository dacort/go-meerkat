package main

import (
	"fmt"

	"github.com/dacort/go-meerkat/meerkat"
)

func main() {
	meerkat := meerkat.NewClient(nil)
	p, err := meerkat.Profiles.Get("550099452400006f00a5277f")
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Info.Username)

	b, err := meerkat.Broadcasts.Get("36aab946-661a-47af-a249-5d395659fd2c")
	if err != nil {
		panic(err)
	}
	fmt.Println(b.Caption)
}
