/* Convert IP and/or list of IPv4 IPs to their integer representation
syncore <syncore@syncore.org> */

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIp(ip string) (bool, []int) {
	a := make([]int, 4)
	if strings.Count(ip, ".") != 3 {
		return false, a
	}
	strArr := strings.Split(ip, ".")

	for i := 0; i < len(strArr); i++ {
		num, err := strconv.Atoi(strArr[i])
		if err != nil {
			return false, a
		}
		if num < 0 || num > 255 {
			return false, a
		}
		a[i] = num
	}
	return true, a
}

func convertIpToInt(ip string) (int, error) {
	isValidIp, i := checkIp(ip)

	if !isValidIp {
		return 0, errors.New("Not a valid IP")
	}
	return ((i[0] * 16777216) + (i[1] * 65536) +
		(i[2] * 256) + (i[3])), nil
}

func showConversionResult(ip string) {
	result, err := convertIpToInt(ip)
	if err != nil {
		fmt.Printf("Unable to convert IP: %s to integer: %s\n", ip, err)
	} else {
		fmt.Printf("%s -> %d\n", ip, result)
	}
}

func main() {
	var fVar string
	flag.StringVar(&fVar, "f", "",
		"File containing IP addresses, one per line")
	flag.Parse()
	if len(flag.Args()) == 0 && fVar == "" {
		fmt.Println("IPV4 to Integer: [-f ipfile.txt] <ip1 ip2 ipN>")
		return
	}
	if len(flag.Args()) != 0 {
		for _, val := range flag.Args() {
			showConversionResult(val)
		}
	}
	if fVar != "" {
		file, err := os.Open(fVar)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			showConversionResult(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file: ", err)
		}
	}
}
