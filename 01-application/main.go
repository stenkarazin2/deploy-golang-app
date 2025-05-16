package main

import (
    "os"
    "fmt"
    "log"
    "net"
    "net/http"
)

func GetLocalIPs() ([]net.IP, error) {
    var ips []net.IP
    addresses, err := net.InterfaceAddrs()
    if err != nil {
        return nil, err
    }

    for _, addr := range addresses {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ips = append(ips, ipnet.IP)
            }
        }
    }
    return ips, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
        log.Println("Error!")
		return
	}
	fmt.Fprintln(w, "hostname:", hostname)

    ips, err := GetLocalIPs()
    if err != nil {
        log.Println("Error!")
		return
    }
    fmt.Fprintln(w, "IP address:", ips[0])

    author, exists := os.LookupEnv("AUTHOR")
	if !exists {
        log.Println("Env var AUTHOR not exists!")
		return
	}
	fmt.Fprintln(w, "author:", author)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}
