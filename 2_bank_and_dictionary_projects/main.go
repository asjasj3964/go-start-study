package main

import (
	"fmt"

	"github.com/asjasj3964/learngo/2_bank_and_dictionary_projects/accounts"
	"github.com/asjasj3964/learngo/2_bank_and_dictionary_projects/mydict"
)

func main() {
	// 2-0. Account + NewAccount
	/* // 대문자로 public하게 만들어 누구든지 export 할 수 있게 만든다.
	account := accounts.Account{Owner: "seongjin", Balance: 1000}
	fmt.Println(account)
	account = accounts.Account{Owner: "seongjin"}
	fmt.Println(account)
	account.Owner = "heejin"
	account.Balance = 2000
	fmt.Println(account)
	fmt.Println("--------------------") */

	// 2-1. methods 1
	account := accounts.NewAccount("seongjin") // (constructor)
	fmt.Println(account)                       // address 반환(복사본이 아닌 object)
	// account.balance = 3000 // private하므로 작동 X
	account.Deposit(10) // account가 Deposit func을 가진다.
	// 실제 account가 아닌 account 복사본이다.
	fmt.Println(account)
	fmt.Println(account.Balance())
	fmt.Println("--------------------")

	// 2-2. methods 2
	account.Withdraw(20) // error를 반환하므로 Go가 관여하지 않는다.
	fmt.Println(account.Balance())
	err1 := account.Withdraw(20)
	if err1 != nil {
		// log.Fatalln(err) // 프로그램을 종료시킨다.
		fmt.Println(err1, account.Balance())
	}
	fmt.Println("--------------------")

	// 2-3. finishing up
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("beomgeun")
	fmt.Println(account)
	fmt.Println("--------------------")

	// 2-4. dictionary 1
	dictionary := mydict.Dictionary{} // 빈 dictionary 생성
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
	dictionary = mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary["first"])
	definition, err2 := dictionary.Search("first")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}
	fmt.Println("--------------------")

	// 2-5. add method
	wordToAdd := "hello"
	defToAdd := "Greeting"
	err3 := dictionary.Add(wordToAdd, defToAdd)
	if err3 != nil {
		fmt.Println(err3)
	}
	searchDef, _ := dictionary.Search(wordToAdd)
	fmt.Println("found:", wordToAdd, "/ definition:", searchDef)
	err5 := dictionary.Add(wordToAdd, defToAdd)
	if err5 != nil {
		fmt.Println(err5)
	}
	fmt.Println("--------------------")

	// 2-6. update delete
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err6 := dictionary.Update(baseWord, "Second")
	if err6 != nil {
		fmt.Println(err6)
	}
	wordToFind, _ := dictionary.Search(baseWord)
	fmt.Println(wordToFind)
	fmt.Println(dictionary)
	err7 := dictionary.Delete(baseWord)
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(dictionary)
	_, err8 := dictionary.Search(baseWord)
	if err8 != nil {
		fmt.Println(err8)
	}
}
