package main

import (
	"fmt"
	"testing"
)

type testIp struct {
	address string
	ipSl    []int
	valid   bool
	ipAsInt int
}

var testIps = []testIp{
	{"190.192.145.28", []int{190, 192, 145, 28}, true, 3200291100},
	{"0.0.0.0", []int{0, 0, 0, 0}, true, 0},
	{"1.552.24.65", []int{1, 0, 0, 0}, false, 0},
}

func TestCheckIp(t *testing.T) {
	for _, testval := range testIps {
		validity, sl := checkIp(testval.address)
		if testval.valid != validity {
			t.Error(
				"ip", testval.address,
				"validity should be: ", validity,
				"validity was: ", testval.valid,
			)
		}
		if len(testval.ipSl) != len(sl) {
			t.Error(
				"ip:", testval.address,
				"ip as slice length should be:", len(testval.ipSl),
				"ip as slice length was:", len(sl),
			)
			break
		}
		for i := range sl {
			if testval.ipSl[i] != sl[i] {
				t.Error(
					"ip:", testval.address,
					fmt.Sprintf("slice value at index %d should be: %d,",
						i, testval.ipSl[i]),
					fmt.Sprintf("slice value at index %d was: %d",
						i, sl[i]),
				)
			}
		}
	}
}

func TestConvertIpToInt(t *testing.T) {
	for _, testval := range testIps {
		ipAsInt, _ := convertIpToInt(testval.address)
		if testval.ipAsInt != ipAsInt {
			t.Error(
				"ip: ", testval.address,
				"ip as int shold be: ", testval.ipAsInt,
				"ip as int was: ", ipAsInt,
			)
		}
	}
}
