package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	ma "github.com/nrtkbb/go-mayaascii"
)

func main() {
	// Read file Maya Ascii.
	fp, err := os.Open("basic.ma")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)

	// Unmarshal Maya Ascii File to 'mo' (Maya Ascii Object).
	mo, err := ma.Unmarshal(reader)
	if err != nil {
		log.Fatal(err)
	}

	// requires command parsed data.
	for _, r := range mo.Requires {
		fmt.Printf("%s version is %s. %d nodeTypes, %d dataTypes, %d Plugin's nodes.",
			r.Name, r.Version, len(r.NodeTypes), len(r.DataTypes), len(r.Nodes))
	}

	// Print all nodes.
	for _, n := range mo.Nodes {
		fmt.Printf("%d : %s\n", n.Lineno, n.Name)
	}

	// Specify node name.
	persp, err := mo.GetNode("persp")
	if err != nil {
		log.Fatal(err)
	}

	// Get attribute (Must be short name) and cast to string.
	ow, err := persp.Attr(".ow").asString()  // or .asInt() or .asFloat() etc..
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s.ow is %s", persp.Name, ow)

	// Print Node's all attrs.
	for _, a := range persp.Attrs {
		fmt.Printf("%s.%s is %s\n", persp.Name, a.Name, a.Value.String())
	}

	// Print Node's all children.
	for _, c := range persp.Children {
		fmt.Printf("%s child is %s\n", persp.Name, c.Name)
	}

	// Print Node's parent node.
	if persp.Parent != nil {
		fmt.Printf("%s parent is %s\n", persp.Name, persp.Parent.Name)
	}

	// Get nodes by nodeType.
	transforms, err := mo.GetNodes("transform") // Specify node type.
	if err != nil {
		log.Fatal(err)
	}

	// Print transform nodes.
	for _, t := range transforms {
		fmt.Printf("%d : %s\n", t.Lineno, t.Name)
	}

	// Get specified source connection nodes.
	srcNodes, err := persp.Src(&ma.ConnectInfo{
		Name: "topShape",  // (Optional)
		Attr: ".ow",       // (Optional) Must be specified with a short name
		Type: "camera",    // (Optional)
	})

	// Get all destination connection nodes.
	dstNodes, err := persp.Dst(nil)  // nil will return all connection nodes.
}