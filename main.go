package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/tucnak/steampipe-plugin-minislack/slack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: slack.Plugin})
}
