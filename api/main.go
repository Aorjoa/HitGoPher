package main

import (
    "net/http"
    "golang.org/x/net/websocket"
    "strconv"
    "fmt"
)

// Declare global variable
var MapPlayer map[string] int = map[string] int {}

type T struct {
    Action string
    Position string
    Player string
}

func randomGopher() (string) {
    return "A1"
}

func getCurrentPoint() (int) {
    return 200
}

func prepairData(position string, point int) (string) {
    return position + "|" + strconv.Itoa(point)
}

func newPlayer(name string) (){
    MapPlayer[name] = 0
}

func winnerAddPoint(name string, point int) {
    MapPlayer[name] += point
}
func sentToClient(ws *websocket.Conn) {
    //var position = randomGopher()
    //var point = getCurrentPoint()
    //var data = prepairData(position, point)
    var receiveData T
    for {
        err := websocket.JSON.Receive(ws,&receiveData)
        if err != nil {
            return
        }
        fmt.Println("Received")
        if receiveData.Action == "newPlayer" {
            newPlayer(receiveData.Player)
        }else{
            winnerAddPoint(receiveData.Player,1)
        }
        websocket.JSON.Send(ws, MapPlayer)
    }
}

func main() {
    http.Handle("/start", websocket.Handler(sentToClient))
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        panic("ListenAndServe: "+ err.Error())
    }
}