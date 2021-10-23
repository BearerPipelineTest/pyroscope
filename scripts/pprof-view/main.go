package main

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"

	"os"

	"google.golang.org/protobuf/proto"

	"github.com/pyroscope-io/pyroscope/pkg/agent/pprof"
)

// the idea here is that you can run it via go main like so:
//   cat heap.pprof.gz | go run scripts/pprof-view/main.go
// and the script will print a json version of a given profile
func main() {
	profile := &pprof.Profile{}
	g, err := gzip.NewReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(g)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(buf, profile); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		panic(err)
	}

	log.Println(string(b))
}
