package main

import (
    "net/http"
    "golang.org/x/net/websocket"
    "strconv"
    "fmt"
)

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

func sentToClient(ws *websocket.Conn) {
    //var position = randomGopher()
    //var point = getCurrentPoint()
    //var data = prepairData(position, point)
    var receiveData T
    for {
        websocket.JSON.Receive(ws,&receiveData)
        fmt.Println(receiveData.Action)
        fmt.Println(receiveData.Position)
        fmt.Println(receiveData.Player)
        websocket.Message.Send(ws,receiveData.Action)
    }
}

func main() {
    http.Handle("/action", websocket.Handler(sentToClient))
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        panic("ListenAndServe: "+ err.Error())
    }
}