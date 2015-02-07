package main

import (
    "net/http"
    "golang.org/x/net/websocket"
    "time"
    "fmt"
    "math/rand"
)

// Declare global variable
type PlayerInfo struct{
    PlayerName string
    PlayerPoint int
}
type SendBackData struct {
    Action string
    Position string
    PlayerName string
}
var golbalNumber string
var WebSocketMaps map[*websocket.Conn] PlayerInfo = make(map[*websocket.Conn] PlayerInfo)

func randomGopher() (string) {
    var position = []string{"A1","A2","A3",
                            "B1","B2","B3",
                            "C1","C2","C3",}

    rand.Seed(time.Now().UnixNano())
    randNumber := rand.Intn(8)
    return position[randNumber]
}

func winnerAddPoint(player PlayerInfo,point int) PlayerInfo {
    player.PlayerPoint += point
    return player
}

func processReceive(ws *websocket.Conn) {
    var receiveData SendBackData
    for {
        err := websocket.JSON.Receive(ws,&receiveData)
        if err != nil {
            fmt.Print("Destroy player : ")
            fmt.Println(WebSocketMaps[ws].PlayerName)
            delete(WebSocketMaps,ws)
            return
        } 
        fmt.Println("======Received======")
        fmt.Println(receiveData)
            if receiveData.Action == "newPlayer" {
                if checkUserNameNotUsed(receiveData.PlayerName) {
                    WebSocketMaps[ws] = PlayerInfo{PlayerName:receiveData.PlayerName,PlayerPoint:0}
                    fmt.Print("Player list : ")
                    fmt.Println(WebSocketMaps)
                    sendDataLoop()
                }else{
                    fmt.Println("User name in used!")
                    websocket.JSON.Send(ws,map[string] string {
                            "username_already" : "true",
                        });
                }
            }

            if golbalNumber == receiveData.Position {
                WebSocketMaps[ws] = winnerAddPoint(WebSocketMaps[ws],1)
                golbalNumber = randomGopher()
                sendDataLoop()
                fmt.Println("Position : " + golbalNumber)
                fmt.Println(WebSocketMaps)
            }
    }
}

func checkUserNameNotUsed(checkName string) bool {
    for _,value := range WebSocketMaps {
        if value.PlayerName == checkName {
            return false
        }
    }
    return true
}

func sendDataLoop() {
    var allPlayerScore map[string]int = make(map[string] int)
    for _,value := range WebSocketMaps {
        allPlayerScore[value.PlayerName] = value.PlayerPoint
    }

    for key,_ := range WebSocketMaps {
        websocket.JSON.Send(key, map[string] interface{} {
            "position" : golbalNumber,
            "pointInfo" : allPlayerScore,
        });
    }
}
func main() {
    golbalNumber = randomGopher()
    fmt.Println(golbalNumber)
    http.Handle("/start", websocket.Handler(processReceive))
    err := http.ListenAndServe(":5535", nil)
    if err != nil {
        panic("ListenAndServe: "+ err.Error())
    }
}