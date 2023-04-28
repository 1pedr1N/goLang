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

	fmt.Println(peterAccount)


	
}