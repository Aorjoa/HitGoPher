	// body...
	var ws = new WebSocket("ws://localhost:12345/start");

	ws.onopen = function (){
			ws.send(JSON.stringify({"Action":"newPlayer", "Position":"A1", "Player":"nong"}));
			ws.send(JSON.stringify({"Action":"hit", "Position":"A1", "Player":"nong"}));
	}
	ws.onmessage = function(msg) {
		console.log(msg.data);
	}
