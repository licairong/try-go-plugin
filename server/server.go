package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/licairong/try-go-plugin/shared"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//log.SetOutput(os.Stdout)
	pluginClientConfig := &plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Cmd:              exec.Command("./helloPlugin"),
		Plugins:          map[string]plugin.Plugin{"main": &shared.GRPCHelloPlugin{}},
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	}

	client := plugin.NewClient(pluginClientConfig)
	pluginClientConfig.Reattach = client.ReattachConfig()
	protocol, err := client.Client()
	if err != nil {
		log.Fatalln(err)
	}
	raw, err := protocol.Dispense("main")
	if err != nil {
		log.Fatalln(err)
	}
	service := raw.(shared.IHelloService)
	res, err := service.Hello("sxy")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
