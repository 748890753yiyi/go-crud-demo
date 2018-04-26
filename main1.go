package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"time"
	"database/sql"
	_ "github.com/bmizerany/pq"
	"io/ioutil"
	"encoding/json"

	"strconv"
	"strings"
	"math/rand"
	"reflect"
	"os"
	"log"
	"io"
)
/*type httpServer struct {

}*/
/*func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers","Action, Module")
	}
	if r.Method == "OPTIONS" {
		return
	}
}*/
/*
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	//添加中间件
	router.Use(Middleware)
	router.GET("/simple/server/get", GetHandler)
	router.POST("/simple/server/post", PostHandler)
	router.PUT("/simple/server/put", PutHandler)
	router.DELETE("/simple/server/delete", DeleteHandler)

	http.ListenAndServe(":8005", router)

	*/
/*router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})
	router.Run(":8000")*//*

}

func Middleware(c *gin.Context)  {
	fmt.Println("this is a middleware")
}
func GetHandler(c *gin.Context)  {
	value, exist := c.GetQuery("key")

	if !exist {
		value = "the key is not exist"
	}
	c.Data(http.StatusOK, "text/plain",[]byte(fmt.Sprintf("get success! %s\n",value)))
	return
}
func PostHandler(c *gin.Context)  {
	type JsonHolder struct {
		Id 			int		`json:"id"`
		Name 		string 	`json:"name"`
	}
	holder := JsonHolder{Id:1, Name:"yang"}
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context)  {
	c.Data(http.StatusOK,"text/plain",[]byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context)  {
	c.Data(http.StatusOK,"text/plain",[]byte("delete success!\n"))
	return
}*/
/*func func1(c *gin.Context)  {
	c.String(http.StatusOK, "test1 ok")
	return
}
func func2(c *gin.Context)  {
	c.String(http.StatusOK, "test2 ok")
	return
}
func func3(c *gin.Context)  {
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK,"参数:%s %s test3 ok", name, passwd)
	return
}
func func4(c *gin.Context)  {
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK,"参数:%s %s test4 ok", name, passwd)
	return
}
func func5(c *gin.Context)  {
	name := c.Param("name")
	passwd := c.Param("passwd")
	c.String(http.StatusOK,"参数:%s %s test5 ok", name, passwd)
	return
}
func func6(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test6 OK", name, passwd)
}
func func7(c *gin.Context)  {
	// 回复一个200 OK
	// 获取传入的参数
	name := c.Query("name")
	passwd := c.Query("passwd")
	c.String(http.StatusOK, "参数:%s %s  test7 OK", name, passwd)
}
func func8(c *gin.Context)  {
	message := c.PostForm("message")
	extra := c.PostForm("extra")
	nick := c.DefaultPostForm("nick","anonymous")

	c.JSON(200,gin.H{
		"status": "test8:posted",
		"message": message,
		"nick": nick,
		"extra": extra,
	})
}
func func9(c *gin.Context)  {
	file, header, err := c.Request.FormFile("uploadFile")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("copy_"+filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err1 := io.Copy(out, file)
	if err != nil {
		log.Fatal(err1)
	}
	c.String(http.StatusOK, "upload file success")
}

type Login struct {
	User 		string 	`form:"user" json:"user" binding:"required"`
	Password 	string 	`form:"password" json:"password" binding:"required"`
}

func funcBindJSON(c *gin.Context)  {
	var json Login
	if c.BindJSON(&json) == nil {
		if json.User == "TAO" && json.Password == "123" {
			c.JSON(http.StatusOK,gin.H{
				"status":"success",
			})
		}else {
			c.JSON(http.StatusOK,gin.H{
				"status":"failed",
			})
		}
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status":"failed",
		})
	}
}

// 下面测试bind FORM数据
func funcBindForm(c *gin.Context) {
	var form Login
	// 本质是将c中的request中的BODY数据解析到form中

	// 方法一: 对于FORM数据直接使用Bind函数, 默认使用使用form格式解析,if c.Bind(&form) == nil
	// 方法二: 使用BindWith函数,如果你明确知道数据的类型
	if c.BindWith(&form, binding.Form) == nil{
		if form.User == "TAO" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"FORM=== status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"FORM=== status": "unauthorized"})
		}
	} else {
		c.JSON(404, gin.H{"FORM=== status": "binding FORM error!"})
	}
}
// router GROUP - GET测试
func func10(c *gin.Context)  {
	c.String(http.StatusOK, "test10 OK")
	return
}
func func11(c *gin.Context)  {
	c.String(http.StatusOK, "test11 OK")
	return
}

// router GROUP - POST测试
func func12(c *gin.Context)  {
	c.String(http.StatusOK, "test12 OK")
	return
}
func func13(c *gin.Context)  {
	c.String(http.StatusOK, "test13 OK")
	return
}
func Logger() gin.HandlerFunc  {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example","123456")
		c.Next()
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}*/
var db1 *sql.DB
type User struct {
	Id 			 string
	Username     string
}
func sqlOpen1()  {
	db1, _ = sql.Open("postgres","port=5432 user=postgres password=root dbname=userdemo sslmode=disable")
}

func sqlClose1()  {
	db1.Close()
}

func sqlInsert1(name string)  {

	rand.Seed(time.Now().UnixNano())
	rs := make([]string, 6)
	for start := 0; start < 6; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(6)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	id := strings.Join(rs, "")
	fmt.Println("name: ",name)
	fmt.Println("insert id: ",id)
	fmt.Println("type:", reflect.TypeOf(id))
	stmt, _ := db1.Prepare("INSERT INTO public.user(id,username) VALUES($1,$2)")
	res,_ := stmt.Exec(id,name)
	affect , _ := res.RowsAffected()
	fmt.Println("insert row affected: ",affect)
}

func sqlQuery1()  string  {
	rows, _ := db1.Query("SELECT * FROM public.user")
	userList := []User{}
	for rows.Next() {
		var id 			string
		var username 	string
		_ = rows.Scan(&id,&username)
		var user User
		user.Id = id
		user.Username = username
		userList = append(userList, user)
	}
	b, err := json.Marshal(userList)
	fmt.Println("b: ", string(b))
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(b)
}

func sqlUpdate1(name string,id string) int64 {
	stmt, _ := db1.Prepare("UPDATE public.user SET username=$1 WHERE id=$2")
	res,_ := stmt.Exec(name,id)
	affect, _ := res.RowsAffected()
	fmt.Println("update rows affected: ",affect)
	return affect
}

func sqlDelete1(id string) int64 {
	stmt,_ := db1.Prepare("DELETE FROM public.user WHERE id=$1")
	res,_ := stmt.Exec(id)
	affect,_ := res.RowsAffected()
	fmt.Println("delete rows affected: ",affect)
	return affect
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type,Content-Length, Authorization, Accept,X-Requested-With,pragma")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200,gin.H{
				"status":"success",
			})
		}else{
			c.Next()
		}

	}
}

func add(c *gin.Context)  {
	/*var json User
	fmt.Println(json)*/

	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)
	fmt.Println("获取json中的username:", user["username"].(string)) //转字符串通过len(password)!=0判断长度
	name := user["username"].(string)
	if name != "" {
		sqlOpen1()
		sqlInsert1(name)
		sqlClose1()
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status":"failed",
		})
	}

}

func update(c *gin.Context)  {
	/*var json User
	fmt.Println(json)*/

	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)
	fmt.Println("获取json中的username:", user["username"].(string))
	fmt.Println("获取json中的id:", user["id"].(string))
	name := user["username"].(string)
	id := user["id"].(string)
	if name != "" {
		sqlOpen1()
		res := sqlUpdate1(name,id)
		sqlClose1()
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
			"affect":res,
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status":"failed",
		})
	}

}

func list(c *gin.Context)  {
	sqlOpen1()
	b := sqlQuery1()
	sqlClose1()
	//fmt.Println("b: ",b)
	c.String(200,b)
	/*c.JSON(http.StatusOK,gin.H{
		"status":"success",
	})*/

}

func delete(c *gin.Context)  {
	id := c.Query("id")
	fmt.Println("id: ",id)
	sqlOpen1()
	b := sqlDelete1(id)
	sqlClose1()
	fmt.Println("删除: ",b)
	c.JSON(http.StatusOK,gin.H{
		"status":"success",
		"affect":b,
	})
}

func uploadFile(c *gin.Context)  {
	file, header, err := c.Request.FormFile("uploadFile")
	filename := header.Filename
	fmt.Println(&file)
	fmt.Println(header.Filename)
	out, err := os.Create("copy_"+filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err1 := io.Copy(out, file)
	if err != nil {
		log.Fatal(err1)
	}
	c.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}

/*func downloadFile(fileFullPath string, res *restful.Response) {
	file, err := os.Open(fileFullPath)

	if err != nil {
		res.WriteEntity(_dto.ErrorDto{Err: err})
		return
	}

	defer file.Close()
	fileName := path.Base(fileFullPath)
	fileName = url.QueryEscape(fileName) // 防止中文乱码
	res.AddHeader("Content-Type", "application/octet-stream")
	res.AddHeader("content-disposition", "attachment; filename=\""+fileName+"\"")
	_, error := io.Copy(res.ResponseWriter, file)
	if error != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}*/
func main()  {
	router := gin.Default()
	/*router.GET("/test1", func1)
	router.POST("/test2", func2)
	router.GET("/test3/:name/:passwd", func3)
	router.POST("/test4/:name/:passwd", func4)
	router.GET("/test5/:name/*passwd", func5)
	router.GET("/test6", func6)
	router.POST("/test7", func7)
	router.POST("/test8", func8)
	router.POST("/upload",func9)

	router.POST("/bindJSON", funcBindJSON)
	router.POST("/bindForm", funcBindForm)

	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey, budy", "status": http.StatusOK})
	})

	router.GET("moreJSON", func(c *gin.Context) {
		var msg struct{
			Name 		string 	`json:"user"`
			Message 	string
			Number 	int
		}
		msg.Name = "TAO"
		msg.Message = "hey, budy"
		msg.Number = 123
		c.JSON(http.StatusOK,msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK,gin.H{
			"name":"TAO",
			"message":"hey budy",
			"status": http.StatusOK,
		})
	})

	group1 := router.Group("/g1")
	group1.GET("/read1", func10)
	group1.GET("/read2", func11)

	group2 := router.Group("/g2")
	group2.POST("/write1",func12)
	group2.POST("/write2",func13)

	router.StaticFS("/showDir", http.Dir("."))
	router.Static("/files", "/bin")
	router.StaticFile("/image","./assets/1.png")

	router.LoadHTMLGlob("templates/*");
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"html template",
		})
	})

	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})*/
	/*router.Use(Logger())
	router.GET("/logger", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
	}))
	// 请求URL: 0.0.0.0:8888/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})*/
	//router.Run(":8000")
	//http.ListenAndServe(":8000",router)
	router.Use(Cors())
	router.POST("/uploadFile", uploadFile)
	router.POST("/add", add)
	router.GET("/list", list)
	router.GET("/del", delete)
	router.POST("/update", update)
	router.LoadHTMLGlob("templates/*");
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",gin.H{
			"title":"html template",
		})
	})
	router.StaticFile("/assets/jquery.min.js","./assets/jquery.min.js")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	server := &http.Server{
		Addr:				":8000",
		Handler: 			router,
		ReadHeaderTimeout: 	10*time.Second,
		WriteTimeout: 		10*time.Second,
		MaxHeaderBytes: 	1<<20,
	}
	server.ListenAndServe()
}
