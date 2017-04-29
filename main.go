package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var err error

func printIP(ip []int) {
	octets := make([]string, 4)
	for index := 0; index < len(ip); index++ {
		octets[index] = strconv.Itoa(ip[index])
	}
	fmt.Println(octets[0] + "." + octets[1] + "." + octets[2] + "." + octets[3])
}

func main() {
	// check args
	cidr := flag.String("range", "", "CIDR range")
	flag.Parse()
	if *cidr == "" {
		flag.Usage()
		return
	}

	// sperate mask and IP
	maskAndIP := strings.Split(*cidr, "/")

	// sanitize IP
	octets := make([]int, 4)
	for index, octet := range strings.Split(maskAndIP[0], ".") {
		octets[index], err = strconv.Atoi(octet)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if octets[index] < 0 {
			fmt.Println("[ERROR] " + octet + " is an invalid octet.")
			return
		}

	}
	if len(octets) < 4 {
		fmt.Println("[ERROR] CIDR range must include 4 octets.")
		return
	}

	// sanitize mask
	maskDigit, err := strconv.Atoi(maskAndIP[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if maskDigit < 16 || maskDigit > 32 {
		fmt.Println("[ERROR] Invalid mask => " + maskAndIP[1] + ".")
	}

	// calcuate number of ips
	numberOfIps := int(math.Pow(2, float64(32-maskDigit)))

	// print ips
	for index := 0; index < numberOfIps; index++ {
		if octets[0] > 255 {
			fmt.Println("[ERROR] Invalid address/mask specified: leftmost octet would be greater than 255.")
			return
		}
		printIP(octets)
		octets[3]++
		if octets[3] > 255 {
			octets[2]++
			octets[3] = 0
		}
		if octets[2] > 255 {
			octets[1]++
			octets[2] = 0
		}
		if octets[1] > 255 {
			octets[0]++
			octets[1] = 0
		}
	}
}
