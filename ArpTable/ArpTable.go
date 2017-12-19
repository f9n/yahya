package ArpTable

import (
	"github.com/pleycpl/yahya/Util"
)

// ClearArpTable : sudo ip -s -s neigh flush all
func ClearArpTable() {
	Util.RunCommandOnBash("sudo ip -s -s neigh flush all")
}

// GetArpTable : cat /proc/net/arp || arp -an  || ip neigh
func GetArpTable() string {
	result := Util.RunCommandOnBashReturnResult("arp -an")
	return result
}

// GetMacs : arp -an | awk '{ print $4}' | sort | uniq -c
func GetMacs() string {
	result := Util.RunCommandOnBashReturnResult("arp -an | awk '{ print $4}' | sort | uniq -c")
	return result
}

// GetDuplicateMacs : arp -an | awk '{ print $4}' | sort | uniq -c | awk '$1>0'
func GetDuplicateMacs() string {
	result := Util.RunCommandOnBashReturnResult("arp -an | awk '{ print $4}' | sort | uniq -c | awk '$1>0'")
	return result
}

// CheckArpTable : Checking Duplication Mac addresses
func CheckArpTable() {
	for {
		r := GetDuplicateMacs()
		if r != "" {
			Util.Notification("Ohh shit, You affected arp poisoning!")
			Util.Espeak("Ohh shit, You affected arp poisoning!")
			return
		}
	}
}
