package main;
import "fmt";

type CurrentAccount struct {
owner string
agencyNumber int
accountNumber int 
money float64
}


func main(){

	var peterAccount = CurrentAccount{

		owner: "Pedro",
		agencyNumber:1232,
		accountNumber: 2323,
		money:200.50,
	}
	var agnesAccount = CurrentAccount{

		owner: "Agnes",
		agencyNumber:232,
		accountNumber: 1211,
		money:1000.50,
	}

    fmt.Println(agnesAccount)
	fmt.Println(peterAccount)

	var crisAccount * CurrentAccount
	crisAccount = new(CurrentAccount)
	crisAccount.owner = "Cris"
	crisAccount.money = 1983.50

	fmt.Println(crisAccount)


	
}