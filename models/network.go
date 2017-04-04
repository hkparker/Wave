package models

type Network struct {
	SSID         string
	Clients      []Device
	AccessPoints []Device
}

func (network *Network) VisualData() []map[string]string {
	set := make([]map[string]string, 0)
	for _, ap := range network.AccessPoints {
		set = append(set, map[string]string{
			"type": "UpdateAccessPoint",
			"MAC":  ap.MAC,
			"IsAP": "true",
			"SSID": network.SSID,
		})
	}
	// add the client associations
	return set
}
