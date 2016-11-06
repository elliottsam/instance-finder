package main

import (
	"regexp"
	"strings"
)

func findAzureIPs(id string) (string, string) {
	var privateip, publicip, azPubIPID string
	for _, v := range resources {
		for k, v := range v {
			if strings.Contains(k, "azurerm_network_interface") {
				att := v.(map[string]interface{})["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
				interfaceid := v.(map[string]interface{})["primary"].(map[string]interface{})["id"].(string)
				if interfaceid == id {
					privateip = att["private_ip_address"].(string)
					for k, v := range att {
						if strings.Contains(k, "public_ip_address_id") {
							if v != "" {
								azPubIPID = v.(string)
							}
						}
					}
				}
			}
		}
	}
	if azPubIPID != "" {
		for _, v := range resources {
			for k, v := range v {
				if strings.Contains(k, "azurerm_public_ip") {
					att := v.(map[string]interface{})["primary"].(map[string]interface{})["attributes"].(map[string]interface{})
					pubipid := v.(map[string]interface{})["primary"].(map[string]interface{})["id"].(string)
					if pubipid == azPubIPID {
						publicip = att["ip_address"].(string)
					}
				}
			}
		}
	} else {
		publicip = ""
	}
	return privateip, publicip
}

func findAzureInstance(res map[string]interface{}) instance {
	const (
		netInterface string = "azurerm_network_interface."
		pubIP        string = "azurerm_public_ip."
	)
	var node instance
	att := res["primary"].(map[string]interface{})["attributes"].(map[string]interface{})

	node.name = att["name"].(string)

	regex := regexp.MustCompile("network_interface_ids.[0-9]{10}")
	for k, v := range att {
		if regex.MatchString(k) {
			node.privateip, node.publicip = findAzureIPs(v.(string))
		}
	}
	return node
}
