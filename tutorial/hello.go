package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// An Alias
var pl = fmt.Println

func main() {
	pl("Hello Go")

	// Input
	pl("What is your name ?")
	/*
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		if err == nil {
			pl("Hello", name)
		} else {
			log.Fatal(err)
		} */

	// Data Type

	var vName string = "Azlan"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4
	v4 = 5.4
	pl(vName)
	pl(v1, v2, v3, v4)

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf("World"))
	pl(reflect.TypeOf(true))

	// Casting

	cV1 := 1.5
	cV2 := int(cV1)
	pl("cV1 Type", reflect.TypeOf(cV1))
	pl("cV2 Type", reflect.TypeOf(cV2))

	cV3 := "5000000"
	cV4, err := strconv.Atoi(cV3) // ASCII To Integer
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := "3.14"
	if cV6, err := strconv.ParseFloat(cV5, 64); err == nil {
		pl(cV6)
	}
	cV7 := fmt.Sprintf("%f", 3.14)
	// %f float, %d integers, %v string
	pl("cV7", cV7)

	// Conditionals

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("Important Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday Too")
	} else {
		pl("Not Important Birthday")
	}
	pl("!true =", !true)

	// Strings

	sV1 := "A word"
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)
	pl("Length: ", len(sV2))
	pl("Contains Another: ", strings.Contains(sV2, "Another"))
	pl("o index: ", strings.Index(sV2, "o"))
	pl("Replace: ", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\n Some Words \n"
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)
	pl("Split: ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(sV2))
	pl("Lower: ", strings.ToUpper(sV2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))

	// pointer
	i := 32
	fmt.Println(&i)
	// pointer to the variable
	p := &vName
	fmt.Println(p)

	// rune

	rStr := "abcdefg"
	pl("Rune Count: ", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf(" %d : %#U : %c \n", i, runeVal, runeVal)
	}

	// time

	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// random value

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50)
	pl("Random: ", randNum)

	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)

	pl("Sin(90) = ", math.Sin(r90))

	// Print Format

	fmt.Printf("Stirng %s Integer %d Char %c Float %f Boolean %t Base8 %o Base16 %x GuessType %v\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1, "Test")

	// For Loops

	for x := 1; x <= 5; x++ {
		fmt.Print(" ", x)
	}
	pl("")
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// Conditions

	seedSecs = time.Now().Unix()
	rand.Seed(seedSecs)
	randNum2 := rand.Intn(50)
	for true {
		fmt.Print("Guess a number between 0 and 50 :")
		pl("Random Number is :", randNum2)
		break
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum2 {
			pl("Pick a Lower value")
		} else if iGuess < randNum2 {
			pl("Pick a Upper value")
		} else {
			pl("You Guessed it")
			break
		}
	}

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)

		// ARRAYS

		var arr1 [5]int
		arr1[0] = 1
		arr2 := [5]int{1, 2, 3, 4, 5}
		pl("Index 0:", arr2[0])
		pl("Arr Length :", len(arr2))
		for i := 0; i < len(arr2); i++ {
			pl(arr2[i])
		}
		for i, v := range arr2 {
			fmt.Printf("%d : %d\n", i, v)
		}
		arr3 := [2][2]int{
			{1, 2},
			{3, 4},
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				pl(arr3[i][j])
			}
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I am a string:", bStr)

	// SLICES

	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "sumulated"
	sl1[4] = "universe"
	pl("Slice Size: ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("1si 3 values", sArr[:3])
	pl("Last 3: ", sArr[2:])
	sArr[0] = 10
	pl("sl3 :", sl3)
	sl3[0] = 1
	pl("sArr :", sArr)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	sl4 := make([]string, 6)
	pl("sl4 ;", sl4)
	pl("sl4[0] :", sl4[0])

	// Functions

	sayHello()
	pl(getSum(4, 5))
	pl(getTwo(10))

	// Errors

	pl(getQuotient(5, 0))

	pl(getSum2(1, 2, 3, 4, 5))

	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)

	// Pointers

	f4 := 5
	pl("f4 before func :", f4)
	changeVal2(&f4)
	pl("f4 after func :", f4)

	// Array and Pointers

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))

	// File IO

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{1, 2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
		f, err = os.Open("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		scan1 := bufio.NewScanner(f)
		for scan1.Scan() {
			pl("Prime :", scan1.Text())
		}
		if err := scan1.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(os.Args)

	pl("\nEND OF ROUTINE")
	// Closing Bracket
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("Error: You cant divide by zero")
	} else {
	}
	return x / y, nil
}

func sayHello() {
	pl("Hello")
}

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}
