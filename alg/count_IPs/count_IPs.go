package count_ips

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// * With input "10.0.0.0", "10.0.0.50"  => return   50
// * With input "10.0.0.0", "10.0.1.0"   => return  256
// * With input "20.0.0.10", "20.0.1.0"  => return  246

func IPsBetween(start, end string) uint32 {
	var numIPs uint32 // Number between IPs

	separator := "."
	startParts := strings.Split(start, separator)
	endParts := strings.Split(end, separator)

	startPartsUint32 := strArrToUintArr(startParts)
	endPartsUint32 := strArrToUintArr(endParts)

	lenIPs := len(endPartsUint32) // len of IP slice

	fmt.Println(startPartsUint32)
	fmt.Println(endPartsUint32)

	for level := range startPartsUint32 {
		var numOnLevel uint32 //sum numbers on the level between IPs

		if startPartsUint32[level] == endPartsUint32[level] {
			continue
		}

		numOnLevel = ((endPartsUint32[level]) - startPartsUint32[level]) * uint32(math.Pow(256, float64(lenIPs-level-1)))
		println(level, numOnLevel)

		numIPs += numOnLevel

	}
	return numIPs
}

func strArrToUintArr(strArr []string) []uint32 {
	uintArr := make([]uint32, len(strArr))

	for i, str := range strArr {
		num, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			log.Print("Error convert str array to uint array")
			return nil
		}
		numUint32 := uint32(num)

		uintArr[i] = numUint32
	}
	return uintArr
}
