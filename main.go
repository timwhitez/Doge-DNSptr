package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/alitto/pond"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"time"
)

var dnsr = ""

func main() {

	threads := 10
	filename := ""
	fastUsage := "lookup PTR threads"
	flag.IntVar(&threads, "t", 10, fastUsage)
	flag.StringVar(&dnsr, "r", "", "input dns host")
	flag.StringVar(&filename, "f", "", "input cidr targets filename")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 && filename == "" {
		log.Fatal("Required argument \"NET/MASK\" missing")
	}
	pool := pond.New(threads, 0, pond.MinWorkers(1))
	if len(args) == 1 {
		block := args[0]
		//fmt.Printf("Input Cidr: %s\n", block)
		ip, ipnet, err := net.ParseCIDR(block)
		if err != nil {
			log.Fatal(err)
		}
		for ip0 := ip.Mask(ipnet.Mask); ipnet.Contains(ip0); incIP(ip0) {
			myip := ip0.String()
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
	}

	if filename != "" {
		cidrs := getContent(filename)
		for _, block := range cidrs {
			if block == "" {
				continue
			}
			//fmt.Printf("Input Cidr: %s\n", block)
			ip, ipnet, err := net.ParseCIDR(block)
			if err != nil {
				log.Fatal(err)
			}
			for ip0 := ip.Mask(ipnet.Mask); ipnet.Contains(ip0); incIP(ip0) {
				myip := ip0.String()
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
		}

	}
	pool.StopAndWait()
}

func getContent(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err == nil {
		str := strings.Replace(string(content), "\r\n", "\n", -1)
		lines := strings.Split(str, "\n")
		return lines
	}
	return nil
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
