package main

import (
	"logger"
	"job"
	"router"
	"db/redis"
)

func main() {
	//e := echo.New()
	//
	//e.POST("/user/add", handler.AddUser)
	//e.GET("/user/get/:id", handler.GetUserById)
	//
	//e.Start(":88")

	//gin.SetMode(gin.ReleaseMode)

	logger.Init()
	//redis.Client()
	//redis.Set("sex", "boy")
	redis.Init()
	job.CreateLogfileByday()
	router.Listen(":88")

	//http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
	//	defer r.Body.Close()
	//	body, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		fmt.Println("error")
	//	}
	//	user := new(model.User)
	//	err = json.Unmarshal(body, &user)
	//	if err != nil {
	//		fmt.Println("err")
	//	}
	//
	//	w.Write([]byte("ok"))
	//})
	//
	//http.ListenAndServe("10.221.100.138:88", nil)

	//net.Listen("tcp", ":88")
}
