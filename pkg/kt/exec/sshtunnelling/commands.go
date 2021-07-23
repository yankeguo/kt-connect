package sshtunnelling

import (
	"os/exec"
)

// AddRoute add route to kubernetes network.
func (s *Cli) AddRoute(cidr string) *exec.Cmd {
	// run command: route add -net 192.168.0.3 dev tun1
	cmd := exec.Command("route",
		"add",
		"-net",
		cidr,
		s.TunName,
	)
	return cmd
}

// AddDevice add a tun device on machine
func (s *Cli) AddDevice() *exec.Cmd {
	// run command: ip tuntap add dev tun0 mod tun
	return exec.Command("ip",
		"tuntap",
		"add",
		"dev",
		s.TunName,
		"mod",
		"tun",
	)
}

// SetupDeviceIP set the ip of tun device
func (s *Cli) SetupDeviceIP() *exec.Cmd {
	// run command: ifconfig tun0 10.1.1.1 10.1.1.2 netmask 255.255.255.252
	return exec.Command("ifconfig",
		s.TunName,
		s.SourceIP,
		s.DestIP,
		"netmask",
		"255.255.255.252",
	)
}

func (s *Cli) RemoveDevice() *exec.Cmd {
	// run command: ip l d tun0
	return exec.Command("ip",
		"l",
		"d",
		s.TunName,
	)
}
