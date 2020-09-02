package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"crypto/rand"
	"net"
	"unsafe"

	"golang.org/x/sys/unix"
)

var showBool, changeBool bool
var netIface, macAddr string

//mac addr struct
type hwaddr struct {
	Name         [unix.IFNAMSIZ]byte
	HardwareAddr unix.RawSockaddr
}

//reading flags
func init() {
	flag.BoolVar(&showBool, "s", false, "print current mac address")
	flag.BoolVar(&changeBool, "c", false, "change mac address")

	flag.StringVar(&netIface, "w", "", "insert device")
	flag.StringVar(&macAddr, "m", "", "insert new mac address")
}

func main() {

	flag.Parse()

	if netIface == "" {
		fmt.Println("[-]Please enter a device.")
		os.Exit(1)
	}

	switch {

	case showBool:
		wlan, err := net.InterfaceByName(netIface)
		checkerr(err)
		fmt.Println(wlan.HardwareAddr)

	case changeBool:

		wlan, err := net.InterfaceByName(netIface)
		checkerr(err)

		var parsedMac net.HardwareAddr
		if macAddr == "" {
			fmt.Println("[-]Please insert a valid mac address.")
			break
		} else if macAddr == "random" {
			parsedMac, err = RandomMacAddress()
			checkerr(err)
		} else {
			parsedMac, err = net.ParseMAC(macAddr)
			checkerr(err)
		}

		//changing mac address
		err = ChangeMac(wlan, parsedMac)
		checkerr(err)

		fmt.Println("[+]Done.")
	}

}

func ChangeMac(device *net.Interface, mac net.HardwareAddr) error {

	var obj hwaddr

	copy(obj.Name[:], append([]byte(device.Name), 0))
	obj.HardwareAddr.Family = unix.ARPHRD_ETHER

	if len(mac) > len(obj.HardwareAddr.Data) {
		mac = mac[:len(obj.HardwareAddr.Data)]
	}

	//changing original mac addr with the new mac addr
	for i, _ := range mac {
		obj.HardwareAddr.Data[i] = int8(mac[i])
	}

	sock, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}

	defer unix.Close(sock)
	return unix.IoctlSetInt(sock, unix.SIOCSIFHWADDR, int(uintptr(unsafe.Pointer(&obj))))

}

func RandomMacAddress() (net.HardwareAddr, error) {

	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	buf[0] |= 2

	mac := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	parsedMac, err := net.ParseMAC(string(mac))
	if err != nil {
		return nil, err
	}

	return parsedMac, nil
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
