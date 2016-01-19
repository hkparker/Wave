import { Router, Route, RouteHandler, Link, browserHistory} from 'react-router';
import { Nav, Navbar, NavItem, NavDropdown, MenuItem, Button, Input, Glyphicon } from 'react-bootstrap';
var React = require('react');
var ReactDOM = require('react-dom');

const App = React.createClass({
  render () {
    return (
      <Navbar inverse fixedTop>
        <Navbar.Header>
          <Navbar.Brand>
            <a href="#"><img className="logo" src="vendor/wave/wave.svg" /></a>
          </Navbar.Brand>
        </Navbar.Header>
        <Nav>
          <NavItem eventKey={1} href="#">Visualize</NavItem>
          <NavItem eventKey={2} href="#">Report<span className="badge pull-right report-count">42</span></NavItem>
          <NavItem eventKey={1} href="#">Investigate</NavItem>
	</Nav>
	<Nav>
          <Navbar.Form pullLeft>
	    <Input type="text" className="form-control navbar-search" placeholder="Search MAC, SSID, Manufacture..." />
	  </Navbar.Form>
        </Nav>
	<Nav pullRight>
          <NavDropdown eventKey={3} title="Hayden Parker" id="basic-nav-dropdown" noCaret>
            <MenuItem eventKey={3.3}>Settings<span className="glyphicon glyphicon-cog pull-right settings-icon"></span></MenuItem>
            <MenuItem divider />
            <MenuItem eventKey={3.3}>Logout<span className="glyphicon glyphicon-log-out pull-right logout-icon"></span></MenuItem>
          </NavDropdown>
        </Nav>
      </Navbar>
    )
  }
})

ReactDOM.render((
  <Router history={browserHistory}>
    <Route path="/" component={App}>
    </Route>
  </Router>
), document.getElementById('content'))
