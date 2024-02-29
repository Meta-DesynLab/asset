package main

import (
  "context"
  "fmt"
  "log"

  "github.com/ethpandaops/beacon/pkg/beacon"
)

func main() {
  // Create options
  opts := *beacon.DefaultOptions()

  // Create beacon node instance
  beaconNode := beacon.NewNode(e.log, &beacon.Config{
    Addr: "localhost:5052",
    Name: "beacon node",
  }, "eth", opts)

  // Register a callback that will be called when the beacon node is ready.
  beaconNode.OnBlock(func(ctx context.Context, event *v1.BlockEvent) error {
    fmt.Println(block)

    return nil
  })

  // Start the beacon node. StartAsync will start the beacon node in the background.
  if err := beaconNode.StartAsync(context.Background()); err != nil {
    log.Fatal(err)
  }
}