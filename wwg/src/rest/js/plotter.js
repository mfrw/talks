const  baseURL = 'http://localhost:8080/data';
var ctx = document.getElementById('myChart').getContext('2d');
var chart = new Chart(ctx, {
	// The type of chart we want to create
	type: 'line',

	// The data for our dataset
	data: {
		labels: [],
		datasets: [{
			label: "rand.Float()*50.0",
			backgroundColor: 'rgb(255, 99, 132)',
			borderColor: 'rgb(255, 99, 132)',
			data: [],
		}]
	},

	// Configuration options go here
	options: {}
});

function removeData(chart) {
	chart.data.labels.shift();
	chart.data.datasets.forEach((dataset) => {
		dataset.data.shift();
	});
	chart.update();
}

function addData(chart, label, data) {
	chart.data.labels.push(label);
	chart.data.datasets.forEach((dataset) => {
		dataset.data.push(data);
	});
	chart.update();
}

function modifyData(chart, label, data) {
	addData(chart, label, data);
	if (chart.data.labels.length > 60) {
		removeData(chart);
	}
}


setInterval(function() { fetch(baseURL).then(response => response.json()).then(json => { modifyData(chart, json.delayedPriceTime, json.delayedPrice) }) }, 1000);
