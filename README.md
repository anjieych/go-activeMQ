# go-activeMQ
activeMQ client for go based on github.com/go-stomp/stomp

# usage:
```
//Send
if err := utils.NewActiveMQ("localhost:61613").Send("/queue/test-1", "test from 1"); err != nil {
    fmt.Println("AMQ ERROR:", err)

//this func will handle the messges get from activeMQ server.
handler := func(err error, msg string) { fmt.Println("AMQ MSG:", err, msg) }
if err := utils.NewActiveMQ("localhost:61613").Subscribe("/queue/test-1", handler); err != nil {
    fmt.Println("AMQ ERROR:", err)
}
```
# Contact Me:
Email:anjieych@126.com
QQ:   272348197
