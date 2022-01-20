package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net"
	"os"
	"time"
)

var test = make([]int64, 0)
var c = int64(0)

func main() {
	t := time.Now()
	app := fiber.New()
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	ip := ""
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = ipv4.String()
			fmt.Println("IPv4: ", ipv4)
		}
	}

	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.SendString(fmt.Sprintf("Time:%s, IP:%s", t.String(), ip))
	})
	app.Get("test-1", func(ctx *fiber.Ctx) error {
		fmt.Println(ip)
		for i := 0; i < 1_000_000; i++ {
			//test = append(test, time.Now().UnixMilli())
			test = append(test, time.Now().UnixMilli())
		}

		return ctx.SendString(fmt.Sprintf("Length: %d, IP:%s", len(test), ip))
	})
	app.Get("test-2", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Test 2")
	})

	err := app.Listen(":4000")
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return
	}
}
