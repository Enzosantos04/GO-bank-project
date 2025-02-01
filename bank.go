package main

import (
	"fmt"
)


func main() {
 var balance float64 = 1000
 var withdraw float64
 var deposit float64
 var choice int 
fmt.Println("Welcome to Enzo's Bank")
    for {  
fmt.Println("What do you want to do?")
fmt.Println("1. Check balance")
fmt.Println("2. Deposit money")
fmt.Println("3. Withdraw money")
fmt.Println("4. Exit")
 fmt.Print("Your Choice: ")
 fmt.Scan(&choice)
 
 if choice == 1{
    fmt.Println("Your current balance is:", balance)
 }else if choice == 2 {
    fmt.Print("Deposit certain amount: ")
    fmt.Scan(&deposit) 
if deposit <= 0{
    fmt.Println("Please deposit a valid amount!") 
    continue
}
balance += deposit // += means the var balance will be sum with deposit
fmt.Println("Your new balance is:", balance)
 }else if choice == 3{
    fmt.Print("Withdraw certain amount: ")
    fmt.Scan(&withdraw)
    if withdraw > balance || withdraw <= 0{
        fmt.Println("This amount is not valid, please try again!")
        continue
    }
    balance -= withdraw
    fmt.Println("Your new Balance is:", balance)
  
 }else if choice == 4{
    fmt.Println("Exit successful, see you soon!")
    break
 }else{
    fmt.Println("Please choice between 1 and 4, try again!")
 }
    }


  

}