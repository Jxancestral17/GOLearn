package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "math/rand"
    "log"
    //"encoding/xml"
	   //"io/ioutil"
)



//Account ATM

type Account struct{

      name string
      surname string
      pin int
      aNumber int
      balance int

}

func CreatePin()(int){
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(9999 - 1000) + 1000
}

func CreateANumber()(int){
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(999999999 - 100000000) + 100000000
}



func (a Account)ShowAccount(){
  fmt.Println("Name: ",a.name)
  fmt.Println("Surname: ",a.surname)
  fmt.Println("PIN: ", a.pin)
  fmt.Println("Account Number: ",a.aNumber)
  fmt.Printf("Balance: %d $\n",a.balance)

}

func Transfert(accountStart Account, accountEnd Account, amount int)(Account, Account){
    if accountStart.balance < amount {
        fmt.Printf("Transferimento bloccato: %s %s non ha abbastanza fondi per procedere\n",accountStart.surname, accountStart.name)
    }else {
      accountStart.balance -= amount
      fmt.Println(accountStart.balance)
      accountEnd.balance += amount
      fmt.Println(accountEnd.balance)
      fmt.Printf("Transferimento effettuato con successo! %d$ è stato trasferito all'account %s %s \n", amount, accountEnd.surname, accountEnd.name)
    }
    return accountStart, accountEnd
}


func Deposit(account Account, amount int)(Account){
  if amount < 0 {
      fmt.Printf("Deposito bloccato: accetta solo numeri positivi \n")
  }else {
    account.balance += amount
    fmt.Printf("Deposito effettuato con successo! %d$ è stato trasferito all'account %s %s \n", amount, account.surname, account.name)
  }
  return account
}

func Withdraw(account Account, amount int)(Account){
  if amount > account.balance || amount < 0 {
      fmt.Printf("Ritiro bloccato:  saldo insufficente o accetta solo numeri positivi \n")
  }else {
    account.balance -= amount
    fmt.Printf("Ritiro effettuato con successo! %d$ è stato prelevato all'account %s %s \n", amount, account.surname, account.name)
  }
  return account
}

func CreateAccount()(Account){
    r := bufio.NewReader(os.Stdin)
    fmt.Println("Il tuo nome: ")
    name, err := r.ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Il tuo cognome: ")
    surname, err := r.ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    return Account{name: name[:len(name)-1], surname: surname[:len(surname)-1], pin: CreatePin(), aNumber: CreateANumber(), balance: 1000}
}

func main(){

      var acc1 = CreateAccount()
      acc1.ShowAccount()
      var acc2 = CreateAccount()
      acc2.ShowAccount()

      acc1, acc2 = Transfert(acc1, acc2, 100)
      acc2.ShowAccount()
      acc1 = Deposit(acc1, 20000)
      acc1 = Withdraw(acc1, 19000)
      acc1.ShowAccount()





}
