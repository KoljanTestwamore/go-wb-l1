package main

import "fmt"

type Set struct {
	Content map[string]bool
}

func (set *Set) Add(vals []string) {
	for _, elem := range vals {
		set.Content[elem] = true
	}
}

func NewSet() Set {
	return Set{
		Content: make(map[string]bool),
	}
}

func intersect(set1 Set, set2 Set) Set {
	result := NewSet()

	for elem := range set1.Content {
		_, ok := set2.Content[elem]
		if (ok) {
			result.Add([]string{elem})
		}
	}

	return result
}


func main() {
	set1 := NewSet()
	set2 := NewSet()

	set1.Add([]string{"a", "b", "c", "d"})
	//                 -    -    +    +
	set2.Add([]string{"e", "f", "c", "d"})

	set3 := intersect(set1, set2)

	for k := range set3.Content {
		fmt.Println(k)
	}
}