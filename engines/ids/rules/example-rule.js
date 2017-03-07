// Every rule defines a manifest variable with some fields
var manifest = {
	"Name":        "Example Rule"	// Name of the rule
}

// Every rule defines and evaluate function which accepts a models.Wireless80211Frame
function evaluate(frame) {
	// We can pretty print each from by logging JSON.stringify
	//console.log(JSON.stringify(frame, null, 2))

	// If this frame results in an IDS alert, we send an 
	alert({
		"Title": "Example rule got a frame",
		"Rule": manifest.Name,
		"Severity": "low"
	})
}
