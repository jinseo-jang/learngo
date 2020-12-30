package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinseo-jang/learngo/mydict"
)

var num int = 777
var x int = 0

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpperSecond(name string) (lenght int, uppercase string) {
	defer fmt.Println("I am done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
	input := words
	fmt.Println(reflect.TypeOf(input))
	fmt.Println(input[1])
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	return total
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true
	koreaAge := age + 2
	switch {
	case koreaAge <= 18:
		fmt.Println("No")
	case koreaAge > 18:
		fmt.Println("Yes")

	}
	return false

}

type person struct {
	name    string
	age     int
	favFood [][]string
	money   map[string]int
}

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	err := dictionary.Update("xxx", "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)

	deleteError := dictionary.Delete("cello")
	if deleteError == nil {
		fmt.Println("Deleted")
	} else {
		fmt.Println(deleteError)
	}
	fmt.Println(dictionary)

	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition: ", result)

	// err = dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// result, err := dictionary.Search("hello")
	// fmt.Println(result)

	// fmt.Println(dictionary)

	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// account := accounts.NewAccount("jinseo")
	// account.Deposit(10)

	// fmt.Println(account)

	// fmt.Println("Balance: ", account.Balance())
	// err := account.Withdraw(9878)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//log.Fatalln(err)
	// }
	// fmt.Println("withdrew: ", account.Balance())

	// fmt.Println(account.Owner())

	// account := banking.Account{owner: "jinseojang"}
	// fmt.Println(account)
	// account.owner = "kim"
	// fmt.Println(account)

	// food := [][]string{{"kimchi", "ramen"}, {"samgyoupsal"}}
	// money := map[string]int{"dollar": 100, "won": 15000}
	// jinseo := person{name: "jang", age: 77, favFood: food, money: money}
	// fmt.Println(jinseo)

	// fmt.Println(money)
	// fmt.Println(money["dollar"])

	// nico := map[string]string{"name": "nico", "age": "12"}
	// fmt.Println(nico)

	// nico["jinseo"] = "44"
	// delete(nico, "nico")
	// for key, value := range nico {
	// 	fmt.Println(key, value)
	// }
	// val, exists := nico["nico"]
	// fmt.Println(val, exists)

	// s := make([]string, 3)
	// fmt.Println("emp:", s)

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get", s[2])
	// fmt.Println("len", len(s))

	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("append:", s)

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c)

	// twoD := make([][]int, 3)
	// fmt.Println(twoD)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("twoD:", twoD)

	// names := [4]string{"jinseo", "min", "dal"}
	// fmt.Println(names)
	// names[3] = "alaa"
	// fmt.Println(names)

	// nicks := []string{"ba", "bo", "man"}
	// fmt.Println(nicks)
	// nicks = append(nicks, "check")
	// fmt.Println(nicks)
	// a := 2
	// b := &a
	// *b = 7
	// fmt.Println(a)
	// fmt.Println(canIDrink(16))
	// result := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(result)
	// fmt.Println("hello world")
	// var name string = "jinseo"
	// fmt.Println(name)

	// nickname := "supermand"
	// fmt.Println(nickname)

	// fmt.Println(num)

	// fmt.Println(multiply(2, 2))

	// totalLength, _ := lenAndUpper("jinseo")
	// fmt.Println(totalLength)

	// repeatMe("j", "jinseo", "kim", "check")

	// fmt.Println(lenAndUpperSecond("parkjin"))
}
