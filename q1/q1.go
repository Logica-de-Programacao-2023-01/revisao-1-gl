package q1

import (
	"fmt"
)

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	//definição do erro
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("valor de compra inválido.")

	}

	//calculo da media
	var soma float64

	for _, h := range purchaseHistory {
		soma += h
	}
	media := 0.0
	media = soma / float64(len(purchaseHistory))

	//condições de desconto
	var discount float64

	if currentPurchase == 0 && media == 0 {
		discount = currentPurchase * 0.1
	} else if currentPurchase > 1000.00 {
		discount = currentPurchase * 0.1
	} else if currentPurchase <= 1000.00 {
		if media > 1000.00 {
			discount = currentPurchase * 0.2
		} else {
			media = currentPurchase * 0.05
		}
	} else if currentPurchase <= 500.00 {
		if media > 1000.00 {
			discount = currentPurchase * 0.2
		} else {
			media = currentPurchase * 0.02
		}
	}

	return discount, nil

}

func main() {
	currentPurchase := 800.00
	purchaseHistory := []float64{500.00, 1000.00, 1500.00}

	descontoFinal, err := CalculateDiscount(currentPurchase, purchaseHistory)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Seu desconto é de:", descontoFinal)
	}

}
