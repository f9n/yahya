package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/pleycpl/yahya/ArpPoison"
	"github.com/pleycpl/yahya/ArpTable"
	"github.com/pleycpl/yahya/Gateway"
	"github.com/pleycpl/yahya/Util"
)

var (
	toolVersion = 0.1
)

func main() {
	versionPtr := flag.Bool("version", false, "get tool version")
	attackPtr := flag.String("attack", "", "attacking ip address with arp method.")
	gatewayPtr := flag.Bool("check-gateway", false, "checking gateway change")
	arpTablePtr := flag.Bool("check-arp-table", false, "checking arp table changes")
	notificationPtr := flag.String("notification", "", "send notification")
	espeakPtr := flag.String("espeak", "", "send sound notification")
	detectArpPoisonPtr := flag.String("detect-arp-poison", "", "listening network, then detecting arp reply")
	webPtr := flag.Bool("web", false, "display arp information")

	flag.Parse()
	switch {
	case *versionPtr:
		fmt.Println(toolVersion)
	case *attackPtr != "":
		fmt.Println("Attacking!")
	case *gatewayPtr:
		Gateway.CheckGateway()
	case *arpTablePtr:
		ArpTable.CheckArpTable()
	case *notificationPtr != "":
		Util.Notification(*notificationPtr)
	case *espeakPtr != "":
		Util.Espeak(*espeakPtr)
	case *detectArpPoisonPtr != "":
		ArpPoison.DetectArpPoison(*detectArpPoisonPtr)
	case *webPtr:
		fmt.Println("Serving localhost:8181")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Yahya Tool is running in Web!")
		})
		http.ListenAndServe(":8181", nil)
	default:
		fmt.Println("Usage of this tool:")
		flag.PrintDefaults()
		//fmt.Println(flag.Args())
	}
}
