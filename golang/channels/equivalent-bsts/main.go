package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func same(t1, t2 *Tree) bool {
	chanOne, chanTwo := make(chan int), make(chan int)

	go func() {
		walk(t1, chanOne)
		close(chanOne)
	}()

	go func() {
		walk(t2, chanTwo)
		close(chanTwo)
	}()

	for {
		nodeOne, chanOneOpen := <-chanOne
		nodeTwo, chanTwoOpen := <-chanTwo

		if !chanOneOpen && !chanTwoOpen {
			break
		}

		if chanOneOpen != chanTwoOpen {
			fmt.Println("one of the channels closed")
			return false
		}

		fmt.Println(nodeOne, nodeTwo)
		if nodeOne != nodeTwo {
			fmt.Println("nodes are not the same")
			return false
		}
	}
	return true
}

func main() {
	treeOne, treeTwo := getTrees()

	fmt.Println("Trees have equivalent values:", same(&treeOne, &treeTwo))

	// go func() {
	// 	for {
	// 		fmt.Println(<-chanOne)
	// 	}
	// }()
	//
	// walk(&treeOne, chanOne)
}

func getTrees() (treeOne, treeTwo Tree) {
	//        3
	//      /   \
	//     1     8
	//   /  \   / \
	//  1    2 5  13
	treeOne = Tree{
		&Tree{
			&Tree{
				nil,
				1,
				nil,
			},
			1,
			&Tree{
				nil,
				2,
				nil,
			},
		},
		3,
		&Tree{
			&Tree{
				nil,
				5,
				nil,
			},
			8,
			&Tree{
				nil,
				13,
				nil,
			},
		},
	}

	//          8
	//        /   \
	//       3    13
	//     /  \
	//    1    5
	//   / \
	//  1   2
	treeTwo = Tree{
		&Tree{
			&Tree{
				&Tree{
					nil,
					1,
					nil,
				},
				1,
				&Tree{
					nil,
					2,
					nil,
				},
			},
			3,
			&Tree{
				nil,
				5,
				nil,
			},
		},
		8,
		&Tree{
			nil,
			13,
			&Tree{
				nil,
				42,
				nil,
			},
		},
	}
	return
}
