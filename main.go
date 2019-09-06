package main

import (
	"fmt"
	"github.com/vapor-ware/ksync/pkg/cli"
	"github.com/vapor-ware/ksync/pkg/ksync"
	"github.com/vapor-ware/ksync/pkg/ksync/doctor"
	"log"
)

func main() {
	//tempDir := os.TempDir()
	//err := syncthing.Fetch(path.Join(tempDir, "syncthing"))
	//if err != nil {
	//	panic(err)
	//}

	initLocal()

	//syncthing.NewServer()
}

// Copied function from KSync.
func initLocal() {
	if err := doctor.DoesSyncthingExist(); err == nil {
		return
	}

	fmt.Println("==== Local Environment ====")

	if err := cli.TaskOut(
		"Fetching extra binaries", ksync.NewSyncthing().Fetch); err != nil {
		log.Fatal()
	}

	fmt.Println()
}
