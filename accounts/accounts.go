package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// error처리
var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
// Java에서 get, set 패턴과 비슷해 보인다.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit : 계조의 잔고에서 입금하는 메소드__type이 Account인 메소드 작성
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance : 계좌의 잔고를 내보내는 메소드
func (a Account) Balance() int {
	return a.balance
}

//Withdraw : 계좌의 잔고에서 출금하는 메소드__error처리 포함
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner : 고객의 이름을 바꾸는 것
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner : 고객의 이름을 return 하는 것
func (a Account) Owner() string {
	return a.owner
}

//String : Java에서 출력값 지정해주듯 지정할 수 있다.
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
