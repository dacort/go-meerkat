package main

import (
	"fmt"
	"sort"

	"github.com/dacort/go-meerkat/meerkat"
)

// ByScore implements sort.Interface for []FollowerProfile based on
// the Score field.
type ByScore []meerkat.FollowerProfile

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func main() {
	meerkat := meerkat.NewClient(nil)

	f, err := meerkat.Profiles.Followers("550099452400006f00a5277f")
	if err != nil {
		panic(err)
	}

	sort.Sort(ByScore(*f))

	for i := 0; i < 10; i++ {
		fmt.Println((*f)[i].Username, "-", (*f)[i].Score)
	}
}
