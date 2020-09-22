package main

import (
	prom "github.com/ishangupta-ds/plot/pkg"
)

func main() {
	v, _ := prom.NewValidator("http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/v1/")
	v.ValidateAndFetch("../query.txt")
}
