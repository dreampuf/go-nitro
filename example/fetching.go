package main

import (
	"github.com/dreampuf/go-nitro/netscaler"
	"log"
)


func main() {
	client, err := netscaler.NewNitroClientFromEnv()
	if err != nil {
		log.Fatal("Could not create a client: ", err)
	}
	lbvserver, err := client.FindAllResources("lbvserver")
	if err != nil {
		log.Fatal("could not find all lbvserver:", err)
	}
	for _, lb := range lbvserver {
		log.Println(lb["name"])
	}
	bounds, err := client.FindAllBoundResources("lbvserver", "", "service")
	if err != nil {
		log.Fatal("could not find all lbvserver:", err)
	}
	for _, bound := range bounds {
		log.Println(bound["servicename"])
	}
}
