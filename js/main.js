	// Declare global variable
    var playerName;
    var dataReceive;
    // Connect WebSocket
	var ws = new WebSocket("ws://localhost:12345/start");

	ws.onopen = function (){
        $("#statusTxt").text("Status : Connected");
	};
    ws.onclose = function (){
        $("#statusTxt").text("Status : Closed");
    };
    ws.onerror = function (){
        $("#statusTxt").text("Status : Connection error!");
    };
	ws.onmessage = function(msg) {
		dataReceive = JSON.parse(msg.data);
        $("img").attr("src","images/cupcake.png"); 
        $("#"+dataReceive.position).attr("src","images/gopher_in_cake.png"); 
	};

    $(document).ready(function(){
        $('img').on("click",function(){
            var clickedId = this.id;
            if(clickedId == dataReceive.position){
                $('#'+this.id).attr("src","images/gopher_break.png"); 
                    setTimeout(function() {
                        ws.send(JSON.stringify({"Action":"hit", "Position":clickedId, "Player": playerName}));
                    },200);
            }
        });
    });
    window.ondragstart = function() { return false; };
        $("#startBtn").on('click',function(){
            playerName = $("#login-name").val();
             if(playerName != "" && ws.readyState === 1){
                $("#ask-name").remove();
                ws.send(JSON.stringify({"Action":"newPlayer", "Player": playerName}));
             }else{
                $("#login-name").focus();
             }
        });