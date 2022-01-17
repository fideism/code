package nginx

import (
	"fmt"
	"testing"
)

func Test_NginxSmoothWeight(t *testing.T) {
	fmt.Println(`test`)

	servers := new(Servers)
	servers.Add("80")
	servers.Add("81")
	servers.Add("82")
	servers.Add("83")

	fmt.Println(servers.Next())
	fmt.Println(servers.Next())
	fmt.Println(servers.Next())
	fmt.Println(servers.Next())
	fmt.Println(servers.Next())
	fmt.Println(servers.Next())
}
