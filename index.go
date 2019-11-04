package main

import (
	 "net/http"
     "fmt"
    //  "io/ioutil"
    //  "encoding/json"
     "database/sql"
    // "github.com/gin-gonic/gin"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/go-sql-driver/mysql"
    "time"
    // "strings"
    "github.com/dgrijalva/jwt-go"
)
var (
	SIGN_NAME_SCERET = "aweQurt178BNI"
)

type User struct{
    Name        string  `json:"name"`
    Password    string  `json:"password"`
}
func zhuce(c echo.Context) error {

        // var userinfo User
        name := c.FormValue("name")
        password := c.FormValue("password")
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/goblin") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return c.JSON(http.StatusOK, "数据库连接失败")
        }
        id := time.Now().Unix()
        db.Exec("insert into userinfo(id,name,password) values (?,?,?)",id,name,password)     
        fmt.Println("注册成功！")
        resp := map[string]string{"message": "注册成功！"}
        return c.JSON(http.StatusOK, resp)
    
	
    
}

func denglu(c echo.Context) error{
   
    cookie,err :=c.Cookie("denglu")
    if err != nil {
	    name := c.FormValue("name")
        user_password := c.FormValue("password")
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/goblin") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return c.JSON(http.StatusOK, "数据库连接失败")
        }
        // id := time.Now().Unix()
        resp := map[string]string{"message": "登录失败！"}
        result:=db.QueryRow("select password,id,userimage from userinfo where name=?",name)      //单行查询
        var password,id,userimage string
        result.Scan(&password,&id,&userimage)
        // fmt.Println(result)
        if user_password == password {
            tokenString, err := createJwt(id)
        if err != nil {
            fmt.Println(err.Error())
            return c.JSON(http.StatusOK, "数据库连接失败")
        }
            resp = map[string]string{"message": "登录成功！","name":name,"token":tokenString,"userimage":userimage}
        }
        
       return c.JSON(http.StatusOK, resp)
   
    }else{
        fmt.Println(cookie)
        claims := parseJwt(cookie.Value)
        user_id  := claims["user_id"]
        db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/goblin") // 设置连接数据库的参数
        defer db.Close()    //关闭数据库
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return c.JSON(http.StatusOK, "数据库连接失败")
        }
        // id := time.Now().Unix()
        
        result:=db.QueryRow("select name,userimage from userinfo where id=?",user_id)      //单行查询
        var name,userimage string
        result.Scan(&name,&userimage)
        resp := map[string]string{"message": "登录成功！","user_id": user_id.(string),"name":name,"userimage":userimage}
       return c.JSON(http.StatusOK, resp)
    }
    
}
func main() {
    	
   
    e := echo.New()
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://127.0.0.1:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowCredentials: true,
      }))
    e.POST("/zhuce", zhuce)
    e.POST("/denglu", denglu)
	e.Logger.Fatal(e.Start(":8000"))
    
}
//跨域处理
// func CORS() echo.HandlerFunc {
// 	return func(c echo.Context) {
// 		method := http.Request.Method	//请求方法
		
// 		//放行所有OPTIONS方法
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "Options Request!")
// 		}
// 		// 处理请求
		
// 	}
// }
//创建 token
func createJwt(id string) (string, error) {
	//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//		"foo": "bar",
	//		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),

	//	})

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user_id"] = id 
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SIGN_NAME_SCERET))
	return tokenString, err
}
//解析 token
func parseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SIGN_NAME_SCERET), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
        // fmt.Println(claims["user_id"], claims["nbf"])
        return claims
	} else {
		fmt.Println(err)
	}

	return claims
}
