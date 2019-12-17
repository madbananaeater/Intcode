package intcode

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type command struct {
	do     func(tab *[]int, args []int)
	argNum int
}

//Pointer of intcode computer
var Pointer int
var flagPointer = false

func itoa(a int) []int {
	var temp []int
	for _, a := range strconv.Itoa(a) {
		temp = append(temp, int(a)-48)
	}
	return temp
}

func atoi(a string) int {
	temp, _ := strconv.Atoi(a)
	return temp
}

//Input takes some number and passes it to Intcode
func Input() int {
	var i int
	fmt.Println("uin")
	fmt.Scan(&i)
	return i
}

//Output send outwards value
func Output(out int) {
	fmt.Println("Debug: ", out)
}

var commands = map[int]command{
	1: command{func(tab *[]int, args []int) { (*tab)[args[2]] = args[0] + args[1] }, 3},
	2: command{func(tab *[]int, args []int) { (*tab)[args[2]] = args[0] * args[1] }, 3},
	3: command{func(tab *[]int, args []int) { (*tab)[args[0]] = Input() }, 1},
	4: command{func(tab *[]int, args []int) { Output(args[0]) }, 1},
	5: command{func(tab *[]int, args []int) {
		if args[0] != 0 {
			Pointer = args[1]
			flagPointer = true
		}
	}, 2},
	6: command{func(tab *[]int, args []int) {
		if args[0] == 0 {
			Pointer = args[1]
			flagPointer = true
		}
	}, 2},
	7: command{func(tab *[]int, args []int) {
		if args[0] < args[1] {
			(*tab)[args[2]] = 1
		} else {
			(*tab)[args[2]] = 0
		}
	}, 3},
	8: command{func(tab *[]int, args []int) {
		if args[0] == args[1] {
			(*tab)[args[2]] = 1
		} else {
			(*tab)[args[2]] = 0
		}
	}, 3},
}

func parseCommand(input string) []int {
	var tab []int
	for _, i := range strings.Split(input, ",") {
		a, _ := strconv.Atoi(i)
		tab = append(tab, a)
	}
	return tab
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}

func evalArgs(tab []int, ptr int) (int, []int) {
	var args []int
	comm := tab[ptr]%10 + tab[ptr]%100 - tab[ptr]%10
	temp := itoa(tab[ptr] / 100)
	for _, a := range tab[ptr+1 : ptr+1+commands[comm].argNum] {
		args = append(args, a)
	}
	for len(temp) < len(args) {
		temp = prependInt(temp, 0)
	}
	reverse(temp)
	fmt.Print(temp)
	for i := range args {
		if temp[i] == 0 && i != len(args)-1 {
			args[i] = tab[args[i]]
		}
	}
	if (comm == 4) && temp[0] == 0 {
		args[0] = tab[args[0]]
	}

	if (comm == 5 || comm == 6) && temp[1] == 0 {
		args[1] = tab[args[1]]
	}

	return comm, args
}
func executeCommand(tab *[]int) {
	for Pointer = 0; (*tab)[Pointer] != 99; {
		comm, args := evalArgs(*tab, Pointer)
		fmt.Println(comm, args)
		commands[comm].do(tab, args)
		if flagPointer == true {
			flagPointer = false
		} else {
			Pointer += len(args) + 1
		}

	}
}

//Intcode executes command string
func Intcode(command string) []int {
	temp := parseCommand(command)
	executeCommand(&temp)
	return temp
}

/*
func main() {
	input := `3,225,1,225,6,6,1100,1,238,225,104,0,1002,43,69,224,101,-483,224,224,4,224,1002,223,8,223,1001,224,5,224,1,224,223,223,1101,67,60,225,1102,5,59,225,1101,7,16,225,1102,49,72,225,101,93,39,224,101,-98,224,224,4,224,102,8,223,223,1001,224,6,224,1,224,223,223,1102,35,82,225,2,166,36,224,101,-4260,224,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,102,66,48,224,1001,224,-4752,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1001,73,20,224,1001,224,-55,224,4,224,102,8,223,223,101,7,224,224,1,223,224,223,1102,18,41,224,1001,224,-738,224,4,224,102,8,223,223,101,6,224,224,1,224,223,223,1101,68,71,225,1102,5,66,225,1101,27,5,225,1101,54,63,224,1001,224,-117,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1,170,174,224,101,-71,224,224,4,224,1002,223,8,223,1001,224,4,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1007,226,226,224,1002,223,2,223,1006,224,329,1001,223,1,223,1007,226,677,224,102,2,223,223,1006,224,344,1001,223,1,223,108,677,677,224,102,2,223,223,1005,224,359,1001,223,1,223,1007,677,677,224,1002,223,2,223,1006,224,374,101,1,223,223,8,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,7,226,226,224,1002,223,2,223,1005,224,404,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,419,1001,223,1,223,8,226,677,224,1002,223,2,223,1005,224,434,101,1,223,223,1008,226,677,224,102,2,223,223,1006,224,449,1001,223,1,223,7,226,677,224,1002,223,2,223,1006,224,464,1001,223,1,223,108,677,226,224,102,2,223,223,1005,224,479,101,1,223,223,108,226,226,224,1002,223,2,223,1006,224,494,101,1,223,223,8,226,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1107,677,226,224,102,2,223,223,1005,224,524,1001,223,1,223,1107,226,226,224,102,2,223,223,1005,224,539,1001,223,1,223,1108,677,677,224,1002,223,2,223,1006,224,554,101,1,223,223,107,226,677,224,102,2,223,223,1005,224,569,1001,223,1,223,1108,226,677,224,1002,223,2,223,1005,224,584,1001,223,1,223,1107,226,677,224,1002,223,2,223,1005,224,599,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,614,101,1,223,223,107,226,226,224,102,2,223,223,1006,224,629,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,644,101,1,223,223,107,677,677,224,1002,223,2,223,1005,224,659,101,1,223,223,1108,677,226,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226`
	Intcode(input)
}*/
