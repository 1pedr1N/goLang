package main;
import "fmt";

type CurrentAccount struct {
owner string
agencyNumber int
accountNumber int 
money float64
}

func (c *CurrentAccount) RetiredMoney(retiredValue float64) string {
	canRetired :=  retiredValue > 0 && retiredValue <= c.money
	if canRetired {
		c.money -= retiredValue
		return "Saque realizado com sucesso"
	} else {
		return "Dinheiro insuficiente"
	}

}


	


func (c *CurrentAccount) AddMoney(depositValue float64) (string, float64) {
    if depositValue > 0 {
        c.money += depositValue
        return "Deposito realizado com sucesso", c.money
    } else { 
        return "Valor do depósito menor que zero", c.money
    }
}


func main(){

silviaAccount := CurrentAccount{}
silviaAccount.owner = "Silvia"

silviaAccount.money = 900.50
fmt.Println(silviaAccount.money)
status, value := silviaAccount.AddMoney(350.00)
fmt.Println(status, value)
 
}