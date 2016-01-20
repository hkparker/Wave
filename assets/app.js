'use strict';

var _reactRouter = require('react-router');

var _reactBootstrap = require('react-bootstrap');

var React = require('react');
var ReactDOM = require('react-dom');
require("./wave.css");

var App = React.createClass({
  displayName: 'App',
  render: function render() {
    return React.createElement(
      _reactBootstrap.Navbar,
      { inverse: true, fixedTop: true },
      React.createElement(
        _reactBootstrap.Navbar.Header,
        null,
        React.createElement(
          _reactBootstrap.Navbar.Brand,
          null,
          React.createElement(
            'a',
            { href: '#' },
            React.createElement('img', { className: 'logo', src: 'vendor/wave/wave.svg' })
          )
        )
      ),
      React.createElement(
        _reactBootstrap.Nav,
        null,
        React.createElement(
          _reactBootstrap.NavItem,
          { eventKey: 1, href: '#' },
          'Visualize'
        ),
        React.createElement(
          _reactBootstrap.NavItem,
          { eventKey: 2, href: '#' },
          'Report',
          React.createElement(
            'span',
            { className: 'badge pull-right report-count' },
            '42'
          )
        ),
        React.createElement(
          _reactBootstrap.NavItem,
          { eventKey: 1, href: '#' },
          'Investigate'
        )
      ),
      React.createElement(
        _reactBootstrap.Nav,
        null,
        React.createElement(
          _reactBootstrap.Navbar.Form,
          { pullLeft: true },
          React.createElement(_reactBootstrap.Input, { type: 'text', className: 'form-control navbar-search', placeholder: 'Search MAC, SSID, Manufacture...' })
        )
      ),
      React.createElement(
        _reactBootstrap.Nav,
        { pullRight: true },
        React.createElement(
          _reactBootstrap.NavDropdown,
          { eventKey: 3, title: 'Hayden Parker', id: 'basic-nav-dropdown', noCaret: true },
          React.createElement(
            _reactBootstrap.MenuItem,
            { eventKey: 3.3 },
            'Settings',
            React.createElement('span', { className: 'glyphicon glyphicon-cog pull-right settings-icon' })
          ),
          React.createElement(_reactBootstrap.MenuItem, { divider: true }),
          React.createElement(
            _reactBootstrap.MenuItem,
            { eventKey: 3.3 },
            'Logout',
            React.createElement('span', { className: 'glyphicon glyphicon-log-out pull-right logout-icon' })
          )
        )
      )
    );
  }
});

ReactDOM.render(React.createElement(
  _reactRouter.Router,
  { history: _reactRouter.browserHistory },
  React.createElement(_reactRouter.Route, { path: '/', component: App })
), document.getElementById('content'));