# Visualizer

## Events

Events are streamed from a websocket connection to `/streams/visualizer`.  All events are JSON representations of the Go type `map[string]string`.

### Adding Items

The following events are recieved when a new item should be added to the visualization.

##### NewDevice

This event instructs views to add a new device to the visualization.

```js
{
  "type": "NewDevice",
  "MAC": "11:22:33:44:55:66",
  "VendorByes": "",
  "IsAP": "false",
}
```

##### NewAssociation

This event instructs views to associate two devices on the graph.

```js
{
  "type": "NewAssociation",
  "MAC1": "11:22:33:44:55:66",
  "MAC2": "11:22:33:44:55:66",
}
```

### Updating Items

##### UpdateAssociation

```js
{
  "type": "UpdateAssociation",
  "MAC1": "11:22:33:44:55:66",
  "MAC2": "11:22:33:44:55:66",
  "MAC1Tx": "16325"
}
```

### Drawing Events

These events are one time animations that do not impact the state of the graph.

##### AnimateNullProbeRequest

This animation indicates a device has probed for any available network.

```js
{
  "type": "AnimateNullProbeRequest",
  "MAC1": "11:22:33:44:55:66",
}
```

##### AnimateProbeRequest

This animation indicates a device has probed for a specific SSID.

```js
{
  "type": "AnimateProbeRequest",
  "MAC1": "11:22:33:44:55:66",
  "SSID": "wifi"
}
```

### Removing Items
