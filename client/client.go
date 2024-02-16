package main

import (
    "crypto/tls"
    "flag"
    "github.com/gorilla/websocket"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "net/url"
    "os"
    "os/signal"
    "time"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    flag.Parse()
    log.SetFlags(0)

    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    host := "ec2-3-27-86-30.ap-southeast-2.compute.amazonaws.com:443"
    //host := fmt.Sprintf("127.0.0.1:%v", os.Getenv("PORT"))
    var addr = flag.String("addr", host, "http service address")
    u := url.URL{Scheme: "wss", Host: *addr, Path: "/api/chat/ws/v1/echo"}
    log.Printf("connecting to %s", u.String())

    dialer := *websocket.DefaultDialer
    dialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    h := http.Header{}
    h.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGdtYWlsLmNvbSIsImV4cCI6MTcwODEwMjg1Miwibm8iOiI1MzQ5NTZkZS03OGFmLTQ0YjUtYmRmMS00NWFhNTA5NDg2MTgifQ.Kl6vcpDiELIdy4B8Dyt96ZfmjQsNh8hn2t27KAzFgLw")

    c, _, err := dialer.Dial(u.String(), h)
    if err != nil {
        log.Fatal("dial:", err)
    }
    defer c.Close()

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                return
            }
            log.Printf("recv: %s", message)
        }
    }()

    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-done:
            return
        case t := <-ticker.C:
            err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
            if err != nil {
                log.Println("write:", err)
                return
            }
        case <-interrupt:
            log.Println("interrupt")

            // Cleanly close the connection by sending a close message and then
            // waiting (with timeout) for the server to close the connection.
            err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
            if err != nil {
                log.Println("write close:", err)
                return
            }
            select {
            case <-done:
            case <-time.After(time.Second):
            }
            return
        }
    }
}
