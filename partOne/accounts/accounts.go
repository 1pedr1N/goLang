package accounts

type CurrentAccount struct {
	Owner string
	AgencyNumber int
	AccountNumber int 
	Money float64
	}
	
	func (c *CurrentAccount) RetiredMoney(retiredValue float64) string {
		canRetired :=  retiredValue > 0 && retiredValue <= c.Money
		if canRetired {
			c.Money -= retiredValue
			return "Saque realizado com sucesso"
		} else {
			return "Dinheiro insuficiente"
		}
	
	}
	
	
		
	
	
	func (c *CurrentAccount) AddMoney(depositValue float64) (string, float64) {
		if depositValue > 0 {
			c.Money += depositValue
			return "Deposito realizado com sucesso", c.Money
		} else { 
			return "Valor do depósito menor que zero", c.Money
		}
	}
	func (c *CurrentAccount) Transfer(transferValue float64, destiny *CurrentAccount) bool {
	if transferValue < c.Money && transferValue > 0 {
		c.Money -= transferValue
		destiny.AddMoney(transferValue)
		return true
	} else {
		return false
	}
	}
	