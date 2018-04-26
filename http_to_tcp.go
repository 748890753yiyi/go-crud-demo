package main

import (
	"net"
	"log"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type connmsg struct {
	Id string
}
var conn_gw map[string]net.Conn

func startServer()  {
	tcpAddr, err := net.ResolveTCPAddr("tcp4",":8987")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("start success!")
	conn_gw = make(map[string]net.Conn)
	for  {
		conn, err := listener.Accept()
		checkError(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn)  {
	connmsgs := connmsg{}
	fmt.Println("client connect")
	buf := make([]byte,100)
	n,_ := conn.Read(buf)
	jsonerr := json.Unmarshal(buf[0:n], &connmsgs)
	if jsonerr != nil {
		fmt.Println("error: ",jsonerr)
		return
	}
	conn_gw[connmsgs.Id] = conn
	_, err := conn.Write([]byte("{\"errcode\":\"0\",\"errmsg\":\"success\"}\n"))
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
	}
	fmt.Println(connmsgs.Id)
}

func start_httpserver()  {
	router := gin.Default()
	router.POST("/test1", func(c *gin.Context) {
		post_id := c.PostForm("id")
		post_msg := c.PostForm("msg")
		post_type := "0"
		conn, ok := conn_gw[post_id]

		if ok {
			fmt.Println(post_id)
			charg_json := "{\"msg\":\"" + post_msg + "\","
			charg_json = charg_json + "\"msgtype\":\"" + post_type + "\","
			charg_json = charg_json + "\"id\":\"" + post_id + "\"}\n"
			_,err := conn.Write([]byte(charg_json))
			if err != nil {
				fmt.Println("Can't resolve address: ", err)
			}
			fmt.Println(charg_json)
			c.String(http.StatusOK,"ok")
		} else {
			c.String(http.StatusOK, "can not find gwid,please check gwid")
		}
	})
	router.Run(":8083")
}

//检查错误
func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}

func main()  {
	conn_gw = make(map[string]net.Conn)
	go start_httpserver()
	startServer()
}