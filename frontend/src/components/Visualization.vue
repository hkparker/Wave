<template>
  <h1>Visualization</h1>
</template>

<script>
  import ForceGraph3D from '3d-force-graph';

  export default {
    name: 'Visualization',
    data: function() {
      return {
        devicesByMAC: new Map(),
        associationsByKey: new Map(),
        isAssociated: new Map(),
        devices: [],
        associations: [],
        onlyShowAssociated: false,
        graph: ForceGraph3D(),
      }
    },
    methods: {
      updateDevice: function(device) {
        this.devicesByMAC.set(device.MAC, device)
        this.devices = []
        for (var member of this.devicesByMAC.values()) {
          this.devices.push(member)
        }
        this.graph.graphData({
            links: this.associations,
            nodes: this.devices,
        });
      },
      updateAssociation: function(association) {
	this.isAssociated.set(association.source, true)
	this.isAssociated.set(association.target, true)
        this.associationsByKey.set(association.Key, association)
        this.associations = []
        for (var member of this.associationsByKey.values()) {
          this.associations.push(member)
        }
        this.graph.graphData({
            links: this.associations,
            nodes: this.devices,
        });
      },
      nodeFilter: function(node) {
        if (this.onlyShowAssociated) {
          if (this.isAssociated.get(node.MAC)) {
            return true
          } else {
            return false
          }
        }
        return true
      },
    },
    mounted(){
      this.graph
        .nodeVisibility(this.nodeFilter)
        .nodeId("MAC")
        .nodeRelSize(6)
        .nodeOpacity(1)
        .linkOpacity(0.8)
        .linkWidth(3);
      this.graph(this.$el).graphData({nodes:this.devices, links: this.associations});

      var ws_protocol = "ws://"
      if (window.location.protocol == "https:") {
        ws_protocol = "wss://"
      }
      var socket = new WebSocket(ws_protocol + window.location.host + '/streams/visualizer')

      var context = this
      socket.onmessage = function (event) {
        var msg = JSON.parse(event.data)
        if (msg.type == "NewDevice") {
          context.updateDevice(msg)
        } else if (msg.type == "NewAssociation") {
          context.updateAssociation(msg)
        }
      }
    }
  }
</script>

<style scoped>
  h1 {
    text-align: center;
  }
</style>
