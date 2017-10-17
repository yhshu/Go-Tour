package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (v IPAddr) String() string { // 接口IPAddr的 方法String()
	return fmt.Sprintf("%v.%v.%v.%v", v[0], v[1], v[2], v[3])
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for IPname, IPAddr := range addrs {
		fmt.Printf("%v:%v\n", IPname, IPAddr)
	}
}
