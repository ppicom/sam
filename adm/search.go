package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/comm"
)

func SearchCustomer(args []string) error {
	text := fmt.Sprint(args)
	params := url.Values{}
	params.Add("text", text)
	_url := fmt.Sprintf("%s/search/customer?%s", viper.GetString("urls.hobbit"), params.Encode())
	return comm.PrintGet(_url)
}
