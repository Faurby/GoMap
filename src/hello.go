package main

import (
	"fmt"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/glaslos/go-osm"
)

func main() {
	p := filepath.Join("data", "map.osm")
	m, err := osm.DecodeFile(p)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(m.Nodes[0].Lat, m.Nodes[0].Lng)

	dc := gg.NewContext(100, 100)
	dc.DrawCircle(50, 50, 40)
}
