package Gateway

import (
	"time"

	"github.com/pleycpl/yahya/Util"
)

func GetGatewayInLinux() string {
	result := Util.RunCommandOnBashReturnResult("route -n | awk '$1==\"0.0.0.0\" { print $2}'")
	return result
}

func CheckGateway() {
	CurrentGateway := GetGatewayInLinux()
	for {
		NextGateway := GetGatewayInLinux()
		if NextGateway != CurrentGateway {
			Util.Notification("Your Gateway Changed!")
			time.Sleep(time.Second * 2)
			return
		}
	}
}
