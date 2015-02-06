    $(document).ready(function(){
        // Declare global variable
        var playerName;
        var dataReceive;
        function genProcessBar(maxPoint){
            var sortable = [];
            for (var name in maxPoint)
                  sortable.push([name, maxPoint[name]])
            sortable.sort(function(a, b) {return b[1] - a[1]})
            var topTenPlayers = sortable.slice(0, 10);
            var sumTopTenPoint = 0;
            for (var playerRn in topTenPlayers){
                sumTopTenPoint += (topTenPlayers[playerRn])[1];
            }

        // Remove old chart
        $("#showPoint *").remove();
        // Get top ten player
        for (var playerRn in topTenPlayers){
                $("#showPoint").append(
                "<h4>Player : "+(topTenPlayers[playerRn])[0]+"</h4>"+
                "<h5>Point : "+(topTenPlayers[playerRn])[1]+"</h5>"+
                    "<div class='progress'>"+
                    "<div class='progress-bar green-color-bg' style='width:"+(((topTenPlayers[playerRn])[1]/sumTopTenPoint)*100)+"%'></div>"+
                "</div>"
                );
        }
    }
    // Connect WebSocket server
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
        console.log(dataReceive);
         if(dataReceive.username_already === undefined){
            $("#ask-name").remove();
            $("img").attr("src","images/cupcake.png"); 
            $("#"+dataReceive.position).attr("src","images/gopher_in_cake.png");
            genProcessBar(dataReceive.pointInfo);

        }else{
            $("#statusTxt").text("Status : User name already!");
            $("#login-name").focus();
        }
    };

        $('img').on("click",function(){
            var clickedId = this.id;
            if(clickedId == dataReceive.position){
                $('#'+this.id).attr("src","images/gopher_break.png"); 
                    setTimeout(function() {
                        ws.send(JSON.stringify({"Action":"hit", "Position":clickedId, "PlayerName": playerName}));
                    },200);
            }
        });

        window.ondragstart = function() { return false; };
        $("#startBtn").on('click',function(){
            playerName = $("#login-name").val();
             if(playerName != "" && ws.readyState === 1){
                ws.send(JSON.stringify({"Action":"newPlayer", "PlayerName": playerName}));
             }else{
                $("#login-name").focus();
             }
        });
    });
    