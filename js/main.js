	// body...
	var ws = new WebSocket("ws://localhost:12345/start");

	ws.onopen = function (){
            ws.send(JSON.stringify({"Action":"newPlayer", "Player":"aorjoa"}));
			ws.send(JSON.stringify({"Action":"hit", "Position":"A1", "Player":"aorjoa"}));
	}
	ws.onmessage = function(msg) {
		var dataFirstLv = JSON.parse(msg.data);
        $("#"+dataFirstLv.position).attr("src","images/gopher_in_cake.png");
	}
