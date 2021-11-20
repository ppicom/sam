package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type ListConsumptions interface {
	ListOne(customerCode int) (string, error)
	List() (string, error)
}

type ListConsumptionsImpl struct {
	getManager util.HttpGetManager
}

func NewListConsumptions() ListConsumptions {
	return ListConsumptionsImpl{
		util.NewHttpGetManager(),
	}
}

func (l ListConsumptionsImpl) ListOne(customerCode int) (string, error) {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("%s/consum/%d", viper.GetString("urls.hobbit"), customerCode)
	return l.getManager.PrettyJson(url)
}

func (l ListConsumptionsImpl) List() (string, error) {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	url := fmt.Sprintf("%s/consum", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
