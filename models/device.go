package models

type Device struct {
	MAC         string
	Vendor      string
	AccessPoint bool
	Probing     bool
	ProbedFor   []Network
	Online      bool
}

func (device *Device) VisualData() map[string]string {
	is_ap := "false"
	if device.AccessPoint {
		is_ap = "true"
	}
	probing := "false"
	if device.Probing {
		probing = "true"
	}
	probed_for := ""
	for i, net := range device.ProbedFor {
		probed_for += net.SSID
		if i < len(device.ProbedFor)-1 {
			probed_for += ","
		}
	}
	return map[string]string{
		"type":      "NewDevice",
		"MAC":       device.MAC,
		"Vendor":    device.Vendor,
		"IsAP":      is_ap,
		"Probing":   probing,
		"ProbedFor": probed_for,
	}
}
