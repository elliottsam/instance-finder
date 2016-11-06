package main

import (
	"encoding/json"
	"strings"
)

func findResource(res []map[string]interface{}, name string) map[string]interface{} {
	for _, v := range res {
		for k, v := range v {
			if k == name {
				return v.(map[string]interface{})
			}
		}
	}
	return nil
}

func findArrayInMap(m map[string]interface{}, key string) []interface{} {
	for k, v := range m {
		if k == key {
			return v.([]interface{})
		}
	}
	return nil
}

func findInMap(m map[string]interface{}, key string) []map[string]interface{} {
	var found []map[string]interface{}
	for k, v := range m {
		if k == key {
			found = append(found, v.(map[string]interface{}))
		}
	}
	return found
}

func findItemInArray(arr []interface{}, match string) string {
	for _, v := range arr {
		if strings.Contains(v.(string), match) {
			return v.(string)
		}
	}
	return ""
}

func instanceRetrieval(tfstate []byte) []instance {
	var instances []instance
	json.Unmarshal(tfstate, &data)
	modules := findArrayInMap(data, "modules")
	for _, v := range modules {
		resources = append(resources, findInMap(v.(map[string]interface{}), "resources")...)
	}
	for _, v := range resources {
		for k, v := range v {
			if strings.Contains(k, "aws_instance") {
				instances = append(instances, findAWSInstance(v.(map[string]interface{})))
			}
			if strings.Contains(k, "azurerm_virtual_machine") {
				instances = append(instances, findAzureInstance(v.(map[string]interface{})))
			}
		}
	}
	return instances
}
