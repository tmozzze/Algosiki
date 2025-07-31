package main

import (
	count_ips "alg/alg/count_IPs"
	"fmt"
)

func main() {
	//find_phrase.FindPhrase("Для")
	//fmt.Println(duplicate_encoder.DuplicateEncode("din"))
	fmt.Println(count_ips.IPsBetween("20.0.0.10", "20.0.1.0"))
}
