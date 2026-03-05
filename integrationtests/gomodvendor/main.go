package main

import "github.com/MahamShakir/quic-go-patch/http3"

// The contents of this script don't matter.
// We just need to make sure that quic-go is imported.
func main() {
	_ = http3.Server{}
}
