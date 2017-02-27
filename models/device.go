package models

type Device struct {
	MAC         string
	Vendor      string
	AccessPoint bool
	SeenBy      []string
}
