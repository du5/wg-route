package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	fs, err := os.Open("./china_ip_list.txt")
	if err != nil {
		log.Panicln(err.Error())
	}
	defer fs.Close()

	ipList := []string{"0.0.0.0/8", "10.0.0.0/8", "100.64.0.0/10", "127.0.0.0/8", "169.254.0.0/16", "172.16.0.0/12", "192.0.0.0/24", "192.0.2.0/24", "192.88.99.0/24", "192.168.0.0/16", "198.18.0.0/15", "198.51.100.0/24", "203.0.113.0/24", "224.0.0.0/4", "233.252.0.0/24", "240.0.0.0/4", "255.255.255.255/32"}

	br := bufio.NewReader(fs)

	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		ip, ipNet, err := net.ParseCIDR(string(line))
		if err != nil || ip == nil || ipNet == nil {
			continue
		}
		ipList = append(ipList, ipNet.String())
	}

	routes := []string{"", ""}
	files := []string{"add.sh", "del.sh"}
	for _, v := range ipList {
		routes[0] += fmt.Sprintf("ip r add %s via $(ip r show | grep default -m 1 | awk '{print $3}')\n", v)
		routes[1] += fmt.Sprintf("ip r del %s\n", v)
	}

	for k := range routes {
		f, err := os.Create(files[k])
		if err != nil {
			log.Panicln(err.Error())
		}
		defer f.Close()
		_, _ = f.WriteString(routes[k])
		_ = f.Sync()
		log.Printf("%s created", files[k])
	}

}
