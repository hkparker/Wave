package models

//
// An Assocation contains all the data about
// interactions between two MAC addresses
//
type Association struct {
	Source             string
	Target             string
	DataToSource       int64
	DataToTarget       int64
	SourceTransmitting bool
	TargetTransmitting bool
	Direct             bool
}
