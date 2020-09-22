package main

import (
	"github.com/ishangupta-ds/plot/pkg"
)

func main() {
	v, _ := newValidator("http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/prom/api/v1/")
	v.validateAndFetch("query.txt")
}
