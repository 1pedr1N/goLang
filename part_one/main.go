package main;
import "fmt";


func main(){

silviaAccount := CurrentAccount{owner: "Silvia" , money:2000.00}
gustavoAccount := CurrentAccount{owner: "Gustavo" , money:300.00}
status := silviaAccount.Transfer(200.00 , &gustavoAccount)
fmt.Println(status)
fmt.Println(silviaAccount)
fmt.Println(gustavoAccount)
}