package ArpPoison

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

// AttackArpPoison :
// arp-scan -l
// arpspoof -i <network-interface> -t <victim-ip> <router-ip>
func AttackArpPoison(ifname string) {
	fmt.Println("Attack!!!")
}

func DetectArpPoison(ifname string) {
	fmt.Println("Detect")
	version := pcap.Version()
	fmt.Println(version)
}
