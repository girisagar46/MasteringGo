package main

import "fmt"

/**
Notes:

- Go does not have classes
- A method is a function with a special receiver argument
	- func (recv <type>) fn() // for example. <type> can be a struct
	- Tye type to which the method belongs is known as the receiver
- <type> has to be declared in the same package
- Pointer receivers are used when metho		d needs to change the values to with the receiver points
- Pointer only save memory because they only pass the memory addresses around
*/

type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSSLNode() *SLLNode {
	return new(SLLNode) // new allocates memory and returns pointer to data
}

type level int

func main() {
	node := NewSSLNode()
	node.SetValue(5)
	fmt.Println("Node is of value: ", node.GetValue())

	sl := new(level)
	sl.raiseShieldLevel(4)
	sl.raiseShieldLevel(5)
	fmt.Println(*sl)
}

// we're attaching raiseShieldLevel to level type
func (lv *level) raiseShieldLevel(i int) {
	if *lv == 0 {
		*lv = 1
	}
	*lv = (*lv) * level(i)
}
