var React = require('react');

var AlertsGraph = (props) => (
	<div className="col-sm-3 col-lg-9">
		<div className="dash-unit">
			<dtitle>
				Alert History
				<span className="unit-option pull-right">All</span>
				<span className="unit-option pull-right">1y</span>
				<span className="unit-option pull-right">6m</span>
				<span className="unit-option pull-right">3m</span>
				<span className="unit-option pull-right">1m</span>
				<span className="unit-option pull-right">2w</span>
				<span className="unit-option pull-right">1w</span>
				<span className="unit-option pull-right">3d</span>
				<span className="unit-option pull-right">1d</span>
				<span className="unit-option pull-right">12h</span>
				<span className="unit-option pull-right">6h</span>
			</dtitle>
			<hr />
			<div id="space">line chat of alert severity count by time</div>
			<hr />
			<h2>up/down 15% by severity</h2>
		</div>
	</div>
)

var NearbyDanger = (props) => (
	<div className="col-sm-3 col-lg-3">
		<div className="half-unit">
			<dtitle>Nearby Danger</dtitle>
			<hr />
			<div className="clockcenter">
				<font size={4}>1 nearby threat</font><br />
				<h4>0 ongoing attacks</h4>
			</div>
		</div>
		<div className="half-unit">
			<dtitle>ElasticSearch Size</dtitle>
			<hr />
			<div className="clockcenter"><h3>268.3MB</h3></div>
		</div>
	</div>
)


var CollectorActivity = (props) => (
	<div className="col-sm-3 col-lg-3 half-width">
		<div className="dash-unit">
			<dtitle>Collector Activity</dtitle>
			<hr />
			<div id="space">line graph each collector framerate last 1 min</div>
			<hr />
			<center><h3>3457 frames/sec</h3></center>
		</div>
	</div>
)

var AlertsByMac = (props) => (
	<div className="col-sm-3 col-lg-3 half-width">
		<div className="dash-unit">
			<dtitle>
				Alerts by MAC
				<span className="unit-option pull-right">1y</span>
				<span className="unit-option pull-right">1m</span>
				<span className="unit-option pull-right">2w</span>
				<span className="unit-option pull-right">1w</span>
				<span className="unit-option pull-right">1d</span>
			</dtitle>
			<hr />
			<div id="space">pie chart of alerts this week by mac address, macs link to device view</div>
		</div>
	</div>
)

var ChannelUtil = (props) => (
	<div className="col-sm-3 col-lg-3 half-width">
		<div className="dash-unit">
			<dtitle>
				Channel Util
				<span className="unit-option pull-right">1w</span>
				<span className="unit-option pull-right">1d</span>
				<span className="unit-option pull-right">5m</span>
				<span className="unit-option pull-right">30s</span>
				<span className="unit-option pull-right">5s</span>
			</dtitle>
			<hr />
			<div id="space">pie chart</div>
		</div>
	</div>
)

var GenericPieChart = (props) => (
	<div className="col-sm-3 col-lg-3 half-width">
		<div className="dash-unit">
			<dtitle>pie chart</dtitle>
			<hr />
			<div id="space">chart</div>
		</div>
	</div>
)

var LeftView = (props) => (
	<div class="col-sm-3 col-lg-6 unit-container">
		<CollectorActivity />
		<AlertsByMac />
		<ChannelUtil />
		<GenericPieChart />
	</div>
)

var ActiveAPs = (props) => (
	<div className="col-sm-6 col-lg-3 half-width">
		<div className="dash-unit double-height">
			<dtitle>
				Active APs
				<span className="unit-option pull-right">Proximity</span>
				<span className="unit-option pull-right">Activity</span>
			</dtitle>
			<hr className="accordion-hr" />
			<div id="accordion2" className="accordion"></div>
		</div>
	</div>
)


var DataHogs = (props) => (
	<div className="col-sm-6 col-lg-3 half-width">
		<div className="dash-unit double-height">
			<dtitle>
				Data Hogs
				<span className="unit-option pull-right">1w</span>
				<span className="unit-option pull-right">1d</span>
				<span className="unit-option pull-right">5m</span>
				<span className="unit-option pull-right">30s</span>
				<span className="unit-option pull-right">5s</span>
			</dtitle>
			<hr className="accordion-hr" />
			<div id="accordion2" className="accordion"></div>
		</div>
	</div>
)

var RightView = (props) => (
	<div class="col-sm-3 col-lg-6 unit-container">
		<ActiveAPs />
		<DataHogs />
	</div>
)

var Dashboard = (props) => (
	<div className="container dashboard">
		<AlertsGraph />
		<NearbyDanger />
		<LeftView />
		<RightView />
	</div>
)
