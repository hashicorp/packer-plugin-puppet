package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	puppetmasterlessProv "github.com/hashicorp/packer-plugin-puppet/provisioner/puppet-masterless"
	puppetserverProv "github.com/hashicorp/packer-plugin-puppet/provisioner/puppet-server"
	"github.com/hashicorp/packer-plugin-puppet/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner("masterless", new(puppetmasterlessProv.Provisioner))
	pps.RegisterProvisioner("server", new(puppetserverProv.Provisioner))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
