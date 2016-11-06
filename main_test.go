package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	azuretf = []byte(`
        {"terraform_version":"0.7.7","modules":[{"path":["root","test1"],"resources":{"azurerm_network_interface.webserver-interface":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-if-webserver","attributes":{"ip_configuration.676093946.public_ip_address_id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/publicIPAddresses/azure-webserver","ip_configuration.676093946.subnet_id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/virtualNetworks/azure-mgt-net/subnets/azure-mgt","private_ip_address":"10.0.0.24"}}},"azurerm_public_ip.openvpn":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/publicIPAddresses/azure-webserver","attributes":{"ip_address":"255.255.255.255"}}},"azurerm_virtual_machine.webserver":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Compute/virtualMachines/azure-webserver","attributes":{"name":"azure-webserver","network_interface_ids.#":"1","network_interface_ids.3346436370":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-if-webserver"}}}}},{"path":["root","test2"],"resources":{"azurerm_network_interface.vm-interface.0":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos01","attributes":{"private_ip_address":"10.0.0.70"}}},"azurerm_network_interface.vm-interface.1":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos02","attributes":{"private_ip_address":"10.0.0.68"}}},"azurerm_network_interface.vm-interface.2":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos03","attributes":{"private_ip_address":"10.0.0.69"}}},"azurerm_virtual_machine.vm.0":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Compute/virtualMachines/azure-centos01","attributes":{"name":"azure-centos01","network_interface_ids.#":"1","network_interface_ids.1159325436":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos01"}}},"azurerm_virtual_machine.vm.1":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Compute/virtualMachines/azure-centos02","attributes":{"name":"azure-centos02","network_interface_ids.#":"1","network_interface_ids.3692083014":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos02"}}},"azurerm_virtual_machine.vm.2":{"primary":{"id":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Compute/virtualMachines/azure-centos03","attributes":{"name":"azure-centos03","network_interface_ids.#":"1","network_interface_ids.2870446032":"/subscriptions/<subscriptionID>/resourceGroups/azure-resource-group/providers/Microsoft.Network/networkInterfaces/azure-centos03"}}}}}]}
    `)

	awstf = []byte(`
		{"version":1,"serial":51,"modules":[{"path":["root","test1"],"resources":{"aws_instance.webserver":{"type":"aws_instance","primary":{"id":"i-25feb5af","attributes":{"private_ip":"10.0.0.24","public_ip":"255.255.255.255","tags.Name":"aws-webserver"}}}}},{"path":["root","test2"],"resources":{"aws_instance.vm.0":{"type":"aws_instance","primary":{"attributes":{"private_ip":"10.0.0.51","public_ip":"","tags.Name":"aws-centos01"}}},"aws_instance.vm.1":{"type":"aws_instance","primary":{"attributes":{"private_ip":"10.0.0.42","public_dns":"","public_ip":"","tags.Name":"aws-centos02"}}},"aws_instance.vm.2":{"type":"aws_instance","primary":{"attributes":{"private_ip":"10.0.0.50","public_ip":"","tags.Name":"aws-centos03"}}}}}]}
	`)
)

func TestAzureInstanceCollection(t *testing.T) {
	resources = nil
	azureInstances := instanceRetrieval(azuretf)
	assert.Len(t, azureInstances, 4, "Azure test should return 4 instances")
	for _, v := range azureInstances {
		if v.name == "azure-webserver" {
			assert.Equal(t, "10.0.0.24", v.privateip, "azure-webserver private ip should be 10.0.0.24")
			assert.Equal(t, "255.255.255.255", v.publicip, "azure-webserver private ip should be 255.255.255.255")
		}
		if strings.Contains(v.name, "azure-centos") {
			assert.Contains(t, v.privateip, "10.0.0", "Private IP should not be empty string")
			assert.Empty(t, v.publicip, "Public IP Address should be empty string")
		}
	}
}

func TestAWSInstanceCollection(t *testing.T) {
	resources = nil
	awsInstances := instanceRetrieval(awstf)
	assert.Len(t, awsInstances, 4, "AWS test should return 4 instances")
	for _, v := range awsInstances {
		if v.name == "aws-webserver" {
			assert.Equal(t, "10.0.0.24", v.privateip, "aws-webserver private ip should be 10.0.0.24")
			assert.Equal(t, "255.255.255.255", v.publicip, "aws-webserver private ip should be 255.255.255.255")
		}
		if strings.Contains(v.name, "aws-centos") {
			assert.Contains(t, v.privateip, "10.0.0", "Private IP should not be empty string")
			assert.Empty(t, v.publicip, "Public IP Address should be empty string")
		}
	}
}
