package main

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestConnectToRouter11(t *testing.T) {
	// Command to list available containers
	psCommand := "sudo docker ps"

	// Execute the command to list available containers
	psOutput, err := exec.Command("sh", "-c", psCommand).Output()
	if err != nil {
		t.Fatalf("Error executing docker ps command: %v", err)
	}

	// Print the output of the docker ps command
	fmt.Println("Available containers:")
	fmt.Println(string(psOutput))

	// Command to log in to bash shell of the container
	loginCommand := "sudo docker exec clab-frrlab-router1 /bin/bash"

	// Execute the command to log in to the bash shell of the container
	loginOutput, err := exec.Command("sh", "-c", loginCommand).CombinedOutput()
	if err != nil {
		t.Fatalf("Error logging in to the container: %v", err)
	}

	// Print the output of the login command
	fmt.Println("Login output:")
	fmt.Println(string(loginOutput))

	// Command to log in to vtysh
	vtyshCommand := "vtysh"

	// Execute the command to log in to vtysh within the container's environment
	vtyshOutput, err := exec.Command("sh", "-c", loginCommand+" -c '"+vtyshCommand+"'").CombinedOutput()
	if err != nil {
		t.Fatalf("Error logging in to vtysh: %v", err)
	}

	// Print the output of the vtysh command
	fmt.Println("vtysh output:")
	fmt.Println(string(vtyshOutput))

	// Command to log in to config mode
	configCommand := "config"

	// Execute the command to log in to config mode within the container's environment
	configOutput, err := exec.Command("sh", "-c", loginCommand+" -c '"+vtyshCommand+" -c '"+configCommand+"'").CombinedOutput()
	if err != nil {
		t.Fatalf("Error logging in to config mode: %v", err)
	}

	// Print the output of the config mode command
	fmt.Println("Config mode output:")
	fmt.Println(string(configOutput))

	// Command to log in to OSPF mode
	ospfCommand := "router ospf"

	// Execute the command to log in to OSPF mode within the container's environment
	ospfOutput, err := exec.Command("sh", "-c", loginCommand+" -c '"+vtyshCommand+" -c '"+configCommand+" -c '"+ospfCommand+"'").CombinedOutput()
	if err != nil {
		t.Fatalf("Error logging in to OSPF mode: %v", err)
	}

	// Print the output of the OSPF mode command
	fmt.Println("OSPF mode output:")
	fmt.Println(string(ospfOutput))

	// Command to configure OSPF network
	networkCommand := "network 192.168.1.0/24 area 0"

	// Execute the command to configure OSPF network within the container's environment
	networkOutput, err := exec.Command("sh", "-c", loginCommand+" -c '"+vtyshCommand+" -c '"+configCommand+" -c '"+ospfCommand+" -c '"+networkCommand+"'").CombinedOutput()
	if err != nil {
		t.Fatalf("Error configuring OSPF network: %v", err)
	}

	// Print the output of the OSPF network configuration command
	fmt.Println("OSPF network configuration output:")
	fmt.Println(string(networkOutput))

	// Command to show OSPF neighbors
	showNeighborsCommand := "do show ip ospf neighbor"

	// Execute the command to show OSPF neighbors within the container's environment
	showNeighborsOutput, err := exec.Command("sh", "-c", loginCommand+" -c '"+vtyshCommand+" -c '"+showNeighborsCommand+"'").CombinedOutput()
	if err != nil {
		t.Fatalf("Error showing OSPF neighbors: %v", err)
	}

	// Print the output of the show OSPF neighbors command
	fmt.Println("OSPF neighbors output:")
	fmt.Println(string(showNeighborsOutput))
}
