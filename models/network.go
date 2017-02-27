package models

type Network struct {
	SSID         string
	Clients      []Device
	AccessPoints []Device
}
