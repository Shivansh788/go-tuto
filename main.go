package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// struct
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("need to fuel up frist!")
	}
}

func main() {
	fmt.Println("Hello shivansh !")

	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello \nshivansh"
	// we can also add srings
	var myString1 string = "Hello" + " " + "shivansh"

	fmt.Println(myString)
	fmt.Println(myString1)

	// to find no. of bytes string is taking
	fmt.Println("K")

	//to find length of string
	fmt.Println(utf8.RuneCountInString("abs"))

	//ACSII no.
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	//other type to initailize variables
	var myVar = "text"
	myVar1 := "text"

	fmt.Println(myVar)
	fmt.Println(myVar1)

	//initailize many variables at a time
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// Function
	var printValue string = "Hello Shivansh"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v and remainder is %v", result, remainder)
	}

	//ARRAY
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//Adress of index 0,1,2
	fmt.Println((&intArr[0]))
	fmt.Println((&intArr[1]))
	fmt.Println((&intArr[2]))

	//other method to initialize array
	var intArr1 [3]int32 = [3]int32{1, 2, 3}
	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr1)
	fmt.Println(intArr2)

	//opration on array
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n The lenght is %v ,with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// map (like dictnaries)
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) //if not in map value = 0

	//if & else if & else
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	//loop
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Name: %v, Age: %v \n", i, v)
	}

	// Understanding of string and rune
	var myString4 = []rune("résumé")
	var indexed = myString4[1]
	fmt.Printf("%v,%T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of 'myString' is %v", len(myString4))

	var myRune1 = 'a'
	fmt.Printf("\nmyRune = %v", myRune1)

	var strSlice = []string{"s", "h", "i", "v", "a", "n", "s", "h"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

	//stuct (like class)
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
