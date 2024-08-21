package accounts

import (
	"errors"
	"fmt"
)

// 2-0. Account + NewAccount
// Account struct
type Account struct {
	// export 하려면 대문자여야 한다.
	owner   string
	balance int
}

// NewAccount creates Account
// Go에서 constructor를 만드는 방법
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 새로운 object를 만들어 반환한다.
}

// 2-1. methods 1 / 2-2. methods
// Deposit x amount on your account
// method
// 복사본을 만들지 않고 account의 balance를 증가시켜야 한다.
// a *Account => 누군가 account.Deposit()을 호출한다면 account를 복사하지 않고 Deposit mothod를 호출한 account를 사용한다.
func (a *Account) Deposit(amount int) { // receiver(a Account)를 갖는다. (receiver -> struct의 첫 글자를 딴 소문자)
	// balance와 owner를 변경할 수 있다.
	a.balance += amount
}

// balance of your account
func (a Account) Balance() int { // a는 복사본이지만 신경쓰지 않는다. (필요한 건 balance)
	return a.balance
}

// 2-2. methods
var errNoMoney = errors.New("Can't withdraw. ") // error르 위한 변수 선언, 새로운 에러 생성
// withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	// a.balance < amount일 경우 인출해서는 안된다.
	// error를 반환해 직접 체크한다.
	// account의 balance가 없다면 error를 발생시킨다.
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // null, none
}

// 2-3. finishing up
// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string { // 값 변경 X, 복사본을 사용해도 된다.
	return a.owner
}
func (a Account) String() string { // 내부적으로(자동으로) 호출하는 method를 사용한다.
	return fmt.Sprint(a.Owner(), "'s account has ", a.Balance())
}
