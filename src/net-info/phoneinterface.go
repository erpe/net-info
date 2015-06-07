package main

import (
	"fmt"
	"net"
	"strings"
)

type PhoneInterface struct {
	iface net.Interface
}

func (pi *PhoneInterface) Name() string {
	return pi.iface.Name
}
func (pi *PhoneInterface) Mac() string {
	return pi.iface.HardwareAddr.String()
}

func (pi *PhoneInterface) Addrs() []string {
	var arrStr []string
	addrs, _ := pi.iface.Addrs()
	for _, addr := range addrs {
		arrStr = append(arrStr, addr.String())
	}
	return arrStr
}

func (pi *PhoneInterface) setInterface(ifc net.Interface) {
	pi.iface = ifc
}

func (pi *PhoneInterface) hasAddress() bool {
	addrs, _ := pi.iface.Addrs()
	if len(addrs) > 0 {
		return true
	} else {
		return false
	}
}

func GetPhoneInterfaces() []PhoneInterface {
	ifaces, err := net.Interfaces()

	if err != nil {
		fmt.Println("error...")
	}

	phoneIntfs := make([]PhoneInterface, 0)

	for _, iface := range ifaces {

		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		} else if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		} else {
			newIntfs := new(PhoneInterface)
			newIntfs.iface = iface
			phoneIntfs = append(phoneIntfs, *newIntfs)
		}
	}
	return phoneIntfs
}

func PhoneInterfacesString() string {
  ifcs := GetPhoneInterfaces()
  str := ""
  for _, pi := range ifcs {
    if pi.hasAddress() {
      str += pi.Name() + ":\n" + strings.Join(pi.Addrs(), ",\n") + "\nMAC: " + pi.Mac() + "\n\n"
    }
  }
  return str
}

//func printPhoneInterfaces(pis []PhoneInterface) {
//	for _, pi := range pis {
//		if pi.hasAddress() {
//			fmt.Println("interface " + pi.Name() + ": " + strings.Join(pi.Addrs(), ",\n ") + "\nMAC: " + pi.Mac())
//		}
//	}
//
