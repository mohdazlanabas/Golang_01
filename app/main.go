package main

import (
	"fmt"
	stuff "golang/app/mypackage"
	"log"
	"regexp"
	"sync"
	"time"
)

// An Alias
var pl = fmt.Println

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

type customer struct {
	name    string
	address string
	bal     float64
}

type Tsp float64
type TBs float64
type ML float64

type MyConstraint interface {
	int | float64
}

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("Cat attacks its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hiss")
}

func (c Cat) HappySound() {
	pl("Cat says Purr")
}

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tsp TBs) ML {
	return ML(tsp * 14.79)
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}
func newCustAdd(c *customer, address string) {
	c.address = address
}

func printTo10() {
	for i := 1; i <= 4; i++ {
		pl("Func 1:", i)
	}
}

func printTo15() {
	for i := 1; i <= 4; i++ {
		pl("Func 2:", i)
	}
}

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

type Account struct {
	balance int
	lock    sync.Mutex
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough balance to withdraw")
	} else {
		fmt.Printf("%d withdrawn: Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {

	pl("Hello, World!")

	// HASH MAP

	var heroes map[string]string
	heroes = make(map[string]string)
	villians := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luthor"
	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}
	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip :", superPets[3])
	_, ok := superPets[3]
	pl("Is there a 3rd pet; ", ok)
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// GENERICS

	pl("5 + 4", getSumGen(5, 4))
	pl("5.6 + 4.7", getSumGen(5.6, 4.7))

	// Struct

	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 main st"
	tS.bal = 234.56
	getCustInfo(tS)
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Composition

	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plunbing",
		"234 North St",
		con1,
	}
	bus1.info()

	// Types

	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("\n3 tsps = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 4.92)
	fmt.Printf("3 tBs = %.2f ML\n", ml2)
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// Getter Setter

	date := stuff.Date{}

	err := date.SetDay(11)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day : %d/%d/%d\n",
		date.Day(), date.Month(), date.Year())

	// Interfaces

	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())

	// Concurrency

	go printTo10()
	go printTo15()

	time.Sleep(2 * time.Second)

	// Channels
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel1)
	pl(<-channel2)

	// mutx or lock

	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())
	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)

	// Closures

	intSum := func(x, y int) int { return x + y }
	pl("5 + 4", intSum(5, 4))

	sampl := 1
	changeVar := func() {
		sampl += 1
	}
	changeVar()
	pl("sampl =", sampl)

	// Recursions

	pl("Factorial 4 =", factorial(4))

	// CLOSING BRCKET
}
