<template>
  <h1>Visualization</h1>
</template>

<script>
  import ForceGraph3D from '3d-force-graph';

  export default {
    name: 'Visualization',
    data: function() {
      return {
      }
    },
    mounted(){
      var visualization = ForceGraph3D();
      visualization(this.$el).graphData({nodes:[], links: []});

      var ws_protocol = "ws://"
      if (window.location.protocol == "https:") {
        ws_protocol = "wss://"
      }
      var socket = new WebSocket(ws_protocol + window.location.host + '/streams/visualizer')

      socket.onmessage = function (event) {
        var msg = JSON.parse(event.data)
        var { nodes, links } = visualization.graphData();
        if (msg.type == "NewDevice") {
          visualization.graphData({
              links: links,
              nodes: [...nodes, { "id": msg.MAC }],
          });
        } else if (msg.type == "NewAssociation") {
          visualization.graphData({
              links: [...links, { "target": msg.MAC1, "source": msg.MAC2}],
              nodes: nodes,
          });
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
