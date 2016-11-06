package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

type instance struct {
	name      string
	privateip string
	publicip  string
}

var (
	resources []map[string]interface{}
	data      map[string]interface{}
)

func findAWSInstance(res map[string]interface{}) instance {
	var node instance
	att := res["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
	node.name = att["tags.Name"].(string)
	node.privateip = att["private_ip"].(string)
	node.publicip = att["public_ip"].(string)
	return node
}

func outputTable(data []instance) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "PrivateIP", "PublicIP"})
	for _, v := range data {
		table.Append([]string{v.name, v.privateip, v.publicip})
	}
	table.Render()
}

func main() {
	var instances []instance

	file := flag.String("file", "terraform.tfstate", "filename for tfstate file")
	flag.Parse()

	jsonBlob, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	instances = instanceRetrieval(jsonBlob)

	if len(instances) == 0 {
		fmt.Printf("No instances found if tfstate file\n")
		os.Exit(0)
	}
	outputTable(instances)
}
