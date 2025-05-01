package personaldata

import (
	"fmt"
)

type Personal struct {
	Name   string  //имя пользователя
	Weight float64 //вес пользователя
	Height float64 // рост пользователя
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
