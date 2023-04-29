package main;
import "fmt";

type CurrentAccount struct {
owner string
agencyNumber int
accountNumber int 
money float64
}

func (c *CurrentAccount) RetiredMoney(retiredValue float64) string {
	canRetired := retiredValue <= c.money
	if canRetired {
		c.money -= retiredValue
		return "Saque realizado com sucesso"
	} else {
		return "Dinheiro insuficiente"
	}

}

func main(){

silviaAccount := CurrentAccount{}
silviaAccount.owner = "Silvia"

silviaAccount.money = 900.50
fmt.Println(silviaAccount.money)
/*
retireMoney := 200.00
silviaAccount.money = silviaAccount.money - retireMoney */
fmt.Println(silviaAccount.RetiredMoney(200.00))
fmt.Println(silviaAccount.money)
}