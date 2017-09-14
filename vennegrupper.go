package main

import(
	"fmt"
	"math/rand"
	"time"
	stack "github.com/eirikbell/vennegrupper/stack"
)

func main() {

	aj := randomStack([]string{"Henriette", "Alexandra", "Nikoline", "Gabriela"})

	ag := randomStack([]string{"Nathaniel", "Sebastian", "William", "Noah", "Casper", "Joakim", "Vebjørn"})

	bj := randomStack([]string{"Sekina", "Emma", "Natalie", "Katarina", "Thea", "Malene", "Vanessa", "Christina"})

	bg := randomStack([]string{"Ariander", "Chrisander"})

	// 12 jenter, 9 gutter, 11 a, 10 bg
	// 1aj, 1ag, 1bj, 1bg
	// 2aj, 2ag, 2bj, 2bg
	// 3aj, 3ag, 3bj, 4bj
	// 4ag, 5ag, 5bj, 6bj
	// 4aj, 6ag, 7ag, 7bj, 8bj
	
	total := int(aj.GetDepth() + ag.GetDepth() + bj.GetDepth() + bg.GetDepth())
	additional := total % 4
	groupCount := (total - additional) / 4

	fmt.Println("total", total, "groupCount", groupCount)

	for i := 1; i <= groupCount; i++ {
		antall := 4;
		if i > groupCount - additional {
			antall = 5
		}
		fmt.Println("\nGruppe", i)
		medlemmer := []string{}

		if allPresent(aj, ag, bj, bg) {
			medlemmer = append(medlemmer, aj.Pop().(string))
			medlemmer = append(medlemmer, ag.Pop().(string))
			medlemmer = append(medlemmer, bj.Pop().(string))
			medlemmer = append(medlemmer, bg.Pop().(string))
		} else if bg.Empty() {
			if allPresent(aj, ag, bj) {
				if bj.GetDepth() > 1 {
					medlemmer = append(medlemmer, bj.Pop().(string))
					medlemmer = append(medlemmer, bj.Pop().(string))
					medlemmer = append(medlemmer, ag.Pop().(string))

					if aj.GetDepth() == 2 {
						medlemmer = append(medlemmer, aj.Pop().(string))
					} else if ag.GetDepth() > aj.GetDepth() {
						medlemmer = append(medlemmer, ag.Pop().(string))
					} else {
						medlemmer = append(medlemmer, aj.Pop().(string))
					}

				}
			} else {

			}			
		}		

		if antall == 5 {
			medlemmer = append(medlemmer, ag.Pop().(string))
		}	

		for _, medlem := range medlemmer {
			fmt.Println("\t", medlem)
		}
	}
}

func randomStack(items []string) *stack.Stack {
	var result *stack.Stack = stack.New()
	for _, value := range Shuffle(items) {
		result.Push(value)
	}
	return result
}

func allPresent(s ...*stack.Stack) bool {
	for _, stack := range s {
		if stack.Empty() {
			return false
		}
	}

	return true
}

func nextGroup(aj, ag, bj, bg, medlemmer []string) ([]string, []string, []string, []string, []string) {
	medlem, ajPop := aj[0], aj[1:]
	medlemmerAppend := append(medlemmer, medlem)

	medlem, agPop := ag[0], ag[1:]
	medlemmerAppend = append(medlemmerAppend, medlem)

	medlem, bjPop := bj[0], bj[1:]
	medlemmerAppend = append(medlemmerAppend, medlem)

	medlem, bgPop := bg[0], bg[1:]
	medlemmerAppend = append(medlemmerAppend, medlem)

	return ajPop, agPop, bjPop, bgPop, medlemmerAppend
}

func compare(x, y []string) int {
	return len(x) - len(y)
}

func Shuffle(array []string) []string {
    random := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := len(array) - 1; i > 0; i-- {
        j := random.Intn(i + 1)
        array[i], array[j] = array[j], array[i]
    }

	return array
}