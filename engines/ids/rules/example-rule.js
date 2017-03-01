// Every rule defines a manifest variable with some fields
var manifest = {
	"Name":        "Example Rule"	// Name of the rule
}

// Every rule defines and evaluate function which accepts a models.Wireless80211Frame
function evaluate(frame) {
	// We can pretty print each from by logging JSON.stringify
	//console.log(JSON.stringify(frame, null, 2))
}
