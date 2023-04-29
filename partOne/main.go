package main;
import (
	"fmt";
	"partOne/accounts";
)


func main(){

silviaAccount := accounts.CurrentAccount{Owner: "Silvia" , Money:2000.00}
gustavoAccount := accounts.CurrentAccount{Owner: "Gustavo" , Money:300.00}
status := silviaAccount.Transfer(200.00 , &gustavoAccount)
fmt.Println(status)
fmt.Println(silviaAccount)
fmt.Println(gustavoAccount)
}