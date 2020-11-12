package cmd

import (
	"github.com/d3ta-go/ms-geolocation-restapi/cmd/db"
	"github.com/d3ta-go/ms-geolocation-restapi/cmd/server"
)

func init() {
	RootCmd.AddCommand(db.DBCmd)
	RootCmd.AddCommand(server.ServerCmd)
	// Add your custom command
}
