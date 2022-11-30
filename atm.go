package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "math/rand"
    "log"
    //"strconv"
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
  fmt.Println("Nome: ",a.name)
  fmt.Println("Cognome: ",a.surname)
  fmt.Println("PIN: ", a.pin)
  fmt.Println("Numero Account: ",a.aNumber)
  fmt.Printf("Bilancio: %d $\n",a.balance)

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
      // Create two account for testing
      var acc1 = CreateAccount()
      var acc2 = CreateAccount()
      r := bufio.NewReader(os.Stdin)
      for {
          fmt.Println("---------------------")
          fmt.Println("1 - Mostra account   ")
          fmt.Println("2 - Deposito         ")
          fmt.Println("3 - Ritira           ")
          fmt.Println("4 - Trasferimento    ")
          fmt.Println("5 - Esci             ")
          fmt.Println("---------------------")
          scelta, err := r.ReadString('\n')
          if err != nil {
            log.Fatal(err)
          }
          switch scelta[:len(scelta)-1]{

          case "1":
            fmt.Printf("1 - Account %s %s \n", acc1.name, acc1.surname)
            fmt.Printf("2 - Account %s %s \n", acc2.name, acc2.surname)
            n, err := r.ReadString('\n')
            if err != nil{
              log.Fatal(err)
            }
            if n[:len(n)-1] == "1"{
              acc1.ShowAccount()
            }else if n[:len(n)-1] == "2" {
              acc2.ShowAccount()
            }else{
              fmt.Println("Account non esiste")
            }
          case "2":
            fmt.Printf("1 - Account %s %s \n", acc1.name, acc1.surname)
            fmt.Printf("2 - Account %s %s \n", acc2.name, acc2.surname)
            n, err := r.ReadString('\n')
            if err != nil{
              log.Fatal(err)
            }
            if n[:len(n)-1] == "1"{
              fmt.Println("Quanto vuoi depositare? ")
              var amount int
              _, err := fmt.Scanf("%d",&amount)
              if err != nil{
                log.Fatal(err)
              }
              acc1 = Deposit(acc1, amount)
            }else if n[:len(n)-1] == "2" {
              fmt.Println("Quanto vuoi depositare? ")
              var amount int
              _, err := fmt.Scanf("%d",&amount)
              if err != nil{
                log.Fatal(err)
              }
              acc2 = Deposit(acc2, amount)
            }else{
              fmt.Println("Account non esiste")
            }
          case "3":
            fmt.Printf("1 - Account %s %s \n", acc1.name, acc1.surname)
            fmt.Printf("2 - Account %s %s \n", acc2.name, acc2.surname)
            n, err := r.ReadString('\n')
            if err != nil{
              log.Fatal(err)
            }
            if n[:len(n)-1] == "1"{
              fmt.Println("Quanto vuoi depositare? ")
              var amount int
              _, err := fmt.Scanf("%d",&amount)
              if err != nil{
                log.Fatal(err)
              }
              acc1 = Withdraw(acc1, amount)
            }else if n[:len(n)-1] == "2" {
              fmt.Println("Quanto vuoi depositare? ")
              var amount int
              _, err := fmt.Scanf("%d",&amount)
              if err != nil{
                log.Fatal(err)
              }
              acc2 = Withdraw(acc2, amount)
            }else{
              fmt.Println("Account non esiste")
            }
          case "4":
            fmt.Println("Quanto vuoi trasferire? ")
            var amount int
            _, err := fmt.Scanf("%d",&amount)
            if err != nil{
              log.Fatal(err)
            }
            acc1, acc2 = Transfert(acc1, acc2, amount)
          case "5":
            os.Exit(3)
          default:
            fmt.Println("Il comando non esiste")
          }
      }






}

