var http = require('http'); // HL
var fs = require('fs');
var index = fs.readFileSync('index.html');

var server = http.createServer(function (req, res) { // HL
	  res.writeHead(200, {'Content-Type': 'text/plain'});
	  res.end(index);
});

server.listen(8080); // HL

console.log("[+] Server listening on port 8080");
