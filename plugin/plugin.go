package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/licairong/try-go-plugin/shared"
        r "github.com/licairong/avago"
)

/*
type PluginService struct{}

func (p PluginService) Hello(name string) (string, error) {
        cmd := exec.Command("pwd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out), nil
}
*/

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"PrintPlugin": &shared.GRPCHelloPlugin{Impl: &r.PluginService{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
