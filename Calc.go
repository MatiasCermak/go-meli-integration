package Calc

var SecretSwitch bool = false

type OpLog struct {
	Number1      float64
	Number2      float64
	Result       float64
	Username     string
	Operation    string
	SecretSwitch bool
}

func Add(n1, n2 float64, usr string) (OpLog) {
	if(SecretSwitch == true){
		n2+=5
	}
	return OpLog{
		Number1: n1,
		Number2: n2,
		Result: n1+n2,
		Username: usr,
		Operation: "Addition",
		SecretSwitch: SecretSwitch,
	}
}

func Subtract(n1, n2 float64, usr string) (OpLog) {
	if(SecretSwitch == true){
		n2+=5
	}
	return OpLog{                                          
		Number1: n1,
		Number2: n2,
		Result: n1-n2,
		Username: usr,
		Operation: "Subtraction",
		SecretSwitch: SecretSwitch,
	}
}

func Divide(n1, n2 float64, usr string) (OpLog) {
	if(SecretSwitch == true){
		n2+=5
	}
	return OpLog{
		Number1: n1,
		Number2: n2,
		Result: n1/n2,
		Username: usr,
		Operation: "Division",
		SecretSwitch: SecretSwitch,
	}
}

func Multiply(n1, n2 float64, usr string) (OpLog) {
	if(SecretSwitch == true){
		n2+=5
	}
	return OpLog{
		Number1: n1,
		Number2: n2,
		Result: n1*n2,
		Username: usr,
		Operation: "Multiplication",
		SecretSwitch: SecretSwitch,
	}
}

func Operate(n1, n2 float64, usr, op string) (OpLog){
	switch op {
		case "+": return Add(n1, n2, usr)
	    case "-": return Subtract(n1, n2, usr)
	    case "*": return Multiply(n1, n2, usr)
		case "/": return Divide(n1, n2, usr)
	    case op: return  OpLog{}
	}
	return OpLog{}
}