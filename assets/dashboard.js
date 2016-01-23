"use strict";

var React = require('react');

var AlertsGraph = React.createClass({
  displayName: "AlertsGraph",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-9" },
      React.createElement(
        "div",
        { className: "dash-unit" },
        React.createElement(
          "dtitle",
          null,
          "Alert History",
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "All"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1y"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "6m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "3m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "2w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "3d"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1d"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "12h"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "6h"
          )
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { id: "space" },
          "line chat of alert severity count by time"
        ),
        React.createElement("hr", null),
        React.createElement(
          "h2",
          null,
          "up/down 15% by severity"
        )
      )
    );
  }
});

var NearbyDanger = React.createClass({
  displayName: "NearbyDanger",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-3" },
      React.createElement(
        "div",
        { className: "half-unit" },
        React.createElement(
          "dtitle",
          null,
          "Nearby Danger"
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { className: "clockcenter" },
          React.createElement(
            "font",
            { size: 4 },
            "1 nearby threat"
          ),
          React.createElement("br", null),
          React.createElement(
            "h4",
            null,
            "0 ongoing attacks"
          )
        )
      ),
      React.createElement(
        "div",
        { className: "half-unit" },
        React.createElement(
          "dtitle",
          null,
          "ElasticSearch Size"
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { className: "clockcenter" },
          React.createElement(
            "h3",
            null,
            "268.3MB"
          )
        )
      )
    );
  }
});

var CollectorActivity = React.createClass({
  displayName: "CollectorActivity",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit" },
        React.createElement(
          "dtitle",
          null,
          "Collector Activity"
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { id: "space" },
          "line graph each collector framerate last 1 min"
        ),
        React.createElement("hr", null),
        React.createElement(
          "center",
          null,
          React.createElement(
            "h3",
            null,
            "3457 frames/sec"
          ),
          React.createElement("center", null)
        )
      )
    );
  }
});

var AlertsByMac = React.createClass({
  displayName: "AlertsByMac",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit" },
        React.createElement(
          "dtitle",
          null,
          "Alerts by MAC",
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1y"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "2w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1d"
          )
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { id: "space" },
          "pie chart of alerts this week by mac address, macs link to device view"
        )
      )
    );
  }
});

var ChannelUtil = React.createClass({
  displayName: "ChannelUtil",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit" },
        React.createElement(
          "dtitle",
          null,
          "Channel Util",
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1d"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "5m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "30s"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "5s"
          )
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { id: "space" },
          "pie chart"
        )
      )
    );
  }
});

var GenericPieChart = React.createClass({
  displayName: "GenericPieChart",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-3 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit" },
        React.createElement(
          "dtitle",
          null,
          "pie chart"
        ),
        React.createElement("hr", null),
        React.createElement(
          "div",
          { id: "space" },
          "chart"
        )
      )
    );
  }
});

var LeftView = React.createClass({
  displayName: "LeftView",

  render: function render() {
    return React.createElement(
      "div",
      { "class": "col-sm-3 col-lg-6 unit-container" },
      React.createElement(CollectorActivity, null),
      React.createElement(AlertsByMac, null),
      React.createElement(ChannelUtil, null),
      React.createElement(GenericPieChart, null)
    );
  }
});

var ActiveAPs = React.createClass({
  displayName: "ActiveAPs",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-6 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit double-height" },
        React.createElement(
          "dtitle",
          null,
          "Active APs",
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "Proximity"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "Activity"
          )
        ),
        React.createElement("hr", { className: "accordion-hr" }),
        React.createElement("div", { id: "accordion2", className: "accordion" })
      )
    );
  }
});

var DataHogs = React.createClass({
  displayName: "DataHogs",

  render: function render() {
    return React.createElement(
      "div",
      { className: "col-sm-6 col-lg-3 half-width" },
      React.createElement(
        "div",
        { className: "dash-unit double-height" },
        React.createElement(
          "dtitle",
          null,
          "Data Hogs",
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1w"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "1d"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "5m"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "30s"
          ),
          React.createElement(
            "span",
            { className: "unit-option pull-right" },
            "5s"
          )
        ),
        React.createElement("hr", { className: "accordion-hr" }),
        React.createElement("div", { id: "accordion2", className: "accordion" })
      )
    );
  }
});

var RightView = React.createClass({
  displayName: "RightView",

  render: function render() {
    return React.createElement(
      "div",
      { "class": "col-sm-3 col-lg-6 unit-container" },
      React.createElement(ActiveAPs, null),
      React.createElement(DataHogs, null)
    );
  }
});

var Dashboard = React.createClass({
  displayName: "Dashboard",

  render: function render() {
    return React.createElement(
      "div",
      { className: "container dashboard" },
      React.createElement(AlertsGraph, null),
      React.createElement(NearbyDanger, null),
      React.createElement(LeftView, null),
      React.createElement(RightView, null)
    );
  }
});