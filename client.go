package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){

	/*resp,_ := http.Get("http://127.0.0.1:8000/test3/name=TAO/passwd=123")
	helpRead(resp)

	// POST传参数,使用gin的Param解析格式: /test3/:name/:passwd
	resp1,_ := http.Post("http://127.0.0.1:8000/test4/name=PT/passwd=456", "",strings.NewReader(""))
	helpRead(resp1)

	// 注意Param中':'和'*'的区别
	resp2,_ := http.Get("http://127.0.0.1:8000/test5/name=TAO/passwd=789")
	helpRead(resp2)
	resp3,_ := http.Get("http://127.0.0.1:8000/test5/name=TAO/")
	helpRead(resp3)

	resp4,_ := http.Get("http://127.0.0.1:8000/test6?name=BBB&passwd=CCC")
	helpRead(resp4)
	resp5,_ := http.Post("http://127.0.0.1:8000/test7?name=DDD&passwd=EEE", "",strings.NewReader(""))
	helpRead(resp5)

	resp6,_ := http.Post("http://127.0.0.1:8000/test8","application/x-www-form-urlencoded",strings.NewReader("message=88888&extra=99999"))
	helpRead(resp6)*/

	/*buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, _ := w.CreateFormFile("uploadFile","images.png")
	fd, _ := os.Open("images.png")
	defer fd.Close()
	io.Copy(fw,fd)
	w.Close()
	resp7, _ := http.Post("http://127.0.0.1:8000/upload", w.FormDataContentType(), buf)
	helpRead(resp7)*/
	/*resp8, _ := http.Post("http://127.0.0.1:8000/bindJSON","application/json", strings.NewReader("{\"user\":\"TAO\",\"password\":\"123\"}"))
	helpRead(resp8)

	resp9,_ := http.Post("http://127.0.0.1:8000/bindForm","application/x-www-form-urlencoded",
		strings.NewReader("user=TAO&password=123"))
	helpRead(resp9)

	resp10,_ := http.Get("http://127.0.0.1:8000/someJSON")
	helpRead(resp10)
	resp11,_ := http.Get("http://127.0.0.1:8000/moreJSON")
	helpRead(resp11)
	resp12,_ := http.Get("http://0.0.0.0:8000/someXML")
	helpRead(resp12)

	resp13,_ := http.Get("http://127.0.0.1:8000/g1/read1")
	helpRead(resp13)
	resp14,_ := http.Get("http://127.0.0.1:8000/g1/read2")
	helpRead(resp14)
	resp15,_ := http.Post("http://127.0.0.1:8000/g2/write1", "", strings.NewReader(""))
	helpRead(resp15)
	resp16,_ := http.Post("http://127.0.0.1:8000/g2/write2", "", strings.NewReader(""))
	helpRead(resp16)

	resp17, _ := http.Get("http://127.0.0.1:8000/index")
	helpRead(resp17)*/

	// 下面测试使用中间件
	resp18,_ := http.Get("http://127.0.0.1:8000/logger")
	helpRead(resp18)

	// 测试验证权限中间件BasicAuth
	resp19,_ := http.Get("http://127.0.0.1:8000/admin/secrets")
	helpRead(resp19)


}
// 用于读取resp的body
func helpRead(resp *http.Response)  {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR2!: ", err)
	}
	fmt.Println(string(body))
}