package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alitto/pond"
	"log"
	"net"
	"time"
)

var dnsr = ""

func main() {

	threads := 10
	fastUsage := "lookup PTR threads"
	flag.IntVar(&threads, "t", 10, fastUsage)
	flag.StringVar(&dnsr, "r", "", "input dns host")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Required argument \"NET/MASK\" missing")
	}
	block := args[0]

	ip, ipnet, err := net.ParseCIDR(block)
	if err != nil {
		log.Fatal(err)
	}

	pool := pond.New(threads, 0, pond.MinWorkers(1))

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		myip := ip.String()
		if dnsr == "" {
			pool.Submit(func() {
				revIP(myip)
			})
		} else {
			pool.Submit(func() {
				CusrevIP(myip)
			})
		}

	}
	pool.StopAndWait()
}

func CusrevIP(ip string) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(5000),
			}
			return d.DialContext(ctx, network, dnsr+":53")
		},
	}

	var ptr []string
	var e error
	ptr, e = r.LookupAddr(context.Background(), ip)

	if e == nil {
		ptrs := ""
		for _, v := range ptr {
			ptrs += "," + v
		}
		fmt.Printf("%s%s\n", ip, ptrs)
	}
}

func revIP(ip string) {
	var ptr []string
	var e error
	ptr, e = net.LookupAddr(ip)

	if e == nil {
		ptrs := ""
		for _, v := range ptr {
			ptrs += "," + v
		}
		fmt.Printf("%s%s\n", ip, ptrs)
	}
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
