package main

import (
    "net/http"
    "fmt"
    "io"
    "os"
//  "encoding/json"
    "database/sql"
//  "github.com/gin-gonic/gin"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/go-sql-driver/mysql"
    "time"
    "strings"
    "github.com/dgrijalva/jwt-go"
    )
var (
	SIGN_NAME_SCERET = "aweQurt178BNI"
)
//用户结构
type User struct{
    Name        string  `json:"name"`
    Password    string  `json:"password"`
}

type Card struct{
    // Id          string  `json:"id"`
    Title       string  `json:"title"`
    Images      []string  `json:"images"`
    Video       string  `json:"video"`
    Message     string  `json:"message"`
}
//注册功能
func zhuce(c echo.Context) (err error) {

        // var userinfo User
        user:=new(User)
        if err = c.Bind(user); err != nil {
            return c.JSON(http.StatusOK, "数据错误")
        }
        fmt.Println(user.Name,user.Password)
        db,err:=opendb(c)      //连接数据库
        if err!=nil {
            return err
        }
        defer db.Close()    //关闭数据库
        id := time.Now().Unix()
        db.Exec("insert into userinfo(id,name,password) values (?,?,?)",id,user.Name,user.Password)     
        fmt.Println("注册成功！")
        resp :=map[string]string{"message": "注册成功！"}
        return c.JSON(http.StatusOK, resp)
    
}
//登陆功能
func denglu(c echo.Context) error{
   
    cookie,err :=c.Cookie("token")
    fmt.Println(cookie)
    if err != nil {
        user:=new(User)
        if err = c.Bind(user); err != nil {
            return c.JSON(http.StatusOK, "数据错误")
        }
        if user.Name!="" && user.Password !="" {
            db,err:=opendb(c)      //连接数据库
            if err!=nil {
            return err
            }
            defer db.Close()    //关闭数据库
        resp := map[string]string{"message": "登录失败！"}
        result:=db.QueryRow("select password,id,userimage from userinfo where name=?",user.Name)      //单行查询
        var password,id,userimage string
        result.Scan(&password,&id,&userimage)
        // fmt.Println(result)
        if user.Password == password {
            tokenString, err := createJwt(id)
        if err != nil {
            fmt.Println(err.Error())
            return c.JSON(http.StatusOK, "token生成失败")
        }   
        //因为跨域无法添加Cookie
        // cookie := new(http.Cookie)
        // cookie.Name = "token"
        // cookie.Value = tokenString
        // cookie.Expires = time.Now().Add(24 * time.Hour)
        // c.SetCookie(cookie)
        resp = map[string]string{"message": "登录成功！","name":user.Name,"userimage":userimage,"user_id":id,"token":tokenString}
        }
        
       return c.JSON(http.StatusOK, resp)
    }else{
        return c.JSON(http.StatusOK, "数据为空")
    }
    }else{
        fmt.Println(cookie)
        claims := parseJwt(cookie.Value)
        user_id  := claims["user_id"]
        if user_id =="expired" { 
            return delCookie(c)
        }
        db,err:=opendb(c)      //连接数据库
        if err!=nil {
        return err
        }
        defer db.Close()    //关闭数据库
        result:=db.QueryRow("select name,userimage from userinfo where id=?",user_id)      //单行查询
        var name,userimage string
        result.Scan(&name,&userimage)
        resp := map[string]string{"message": "登录成功！","user_id": user_id.(string),"name":name,"userimage":userimage}
       return c.JSON(http.StatusOK, resp)
    }
    
}
//上传图片功能
func uploadImage(c echo.Context) error{
    cookie,err :=c.Cookie("token")
    if err != nil {return c.JSON(http.StatusOK, "wrong")}else{
        claims := parseJwt(cookie.Value)
        user_id  := claims["user_id"]
        if user_id =="expired" { 
            return delCookie(c)
        }
    act :=c.QueryParam("act")
    if act =="avator" {
        file, err := c.FormFile("img")
        if err != nil {
            return err
        }
        src, err := file.Open()
        if err != nil {
            return err
        }
        defer src.Close()
    
        // Destination
        dst, err := os.Create("./avator/"+file.Filename)
        if err != nil {
            return err
        }
        defer dst.Close()
        // Copy
        if _, err = io.Copy(dst, src); err != nil {
            return err
        }else{
            db,err:=opendb(c)      //连接数据库
            if err!=nil {
            return err
            }
            defer db.Close()    //关闭数据库
            imageUrl :="/avator/"+file.Filename
            db.Exec("update userinfo set userimage=? where id=?",imageUrl,user_id)      
            return c.JSON(http.StatusOK, "上传成功！")
        }
    }else if act=="card"{
        form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["imgs"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create("./images/"+file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}else{
            return c.JSON(http.StatusOK, "上传成功！")
        }

	}
   }
   return c.JSON(http.StatusOK, "wrong")
    }
}
//删除cookie
func delCookie(c echo.Context) (err error){
    cookie := new(http.Cookie)      
    cookie.Name = "token"  
    cookie.MaxAge = -1        
    c.SetCookie(cookie)
    return c.JSON(http.StatusOK, "token已过期")
}
//发布card功能
func release(c echo.Context) (err error) {
    cookie,err :=c.Cookie("token")
    if err != nil {return c.JSON(http.StatusOK, "未登录")}else{
        claims := parseJwt(cookie.Value)
        user_id  := claims["user_id"]
        if user_id =="expired" { 
           return delCookie(c)
        }
        // var userinfo User
        card:=new(Card)
        if err = c.Bind(card); err != nil {
            return c.JSON(http.StatusOK, "数据错误")
        }
        fmt.Println(card.Title,card.Message,card.Images,card.Video)
        db,err:=opendb(c)      //连接数据库
        if err!=nil {
        return err
        }
        defer db.Close()    //关闭数据库
        id := time.Now().Unix()
        image :=""
        if card.Images != nil {
            i := 0
           for ;i < len(card.Images)-1 ; {
                image = image + card.Images[i]+"|"
                i++
           }
        }
        db.Exec("insert into cards(id,user_id,title,message,images,media) values (?,?,?,?,?,?)",id,user_id,card.Title,card.Message,image,card.Video)     
        fmt.Println("发布成功！")
        resp := map[string]string{"message": "发布成功！"}
        return c.JSON(http.StatusOK, resp)
    }
}
//上传视频资源
func uploadMedia(c echo.Context) error{
    cookie,err :=c.Cookie("token")
    if err != nil {return c.JSON(http.StatusOK, "wrong")}else{
        claims := parseJwt(cookie.Value)
        user_id  := claims["user_id"]
        if user_id =="expired" { 
            return delCookie(c)
        }
    file, err := c.FormFile("vid")
        if err != nil {
            return err
        }
        src, err := file.Open()
        if err != nil {
            return err
        }
        defer src.Close()
    
        // Destination
        dst, err := os.Create("./media/"+file.Filename)
        if err != nil {
            return err
        }
        defer dst.Close()
        // Copy
        if _, err = io.Copy(dst, src); err != nil {
            return err
        }else{
            return c.JSON(http.StatusOK, "上传成功！")
        }
    }
}
//获取信息接口
func showEverthing(c echo.Context) error{
    way := c.Param("where")
    id  := c.Param("id")
    resp := map[string]string{"message": "获取信息失败！"}
    //获取用户信息
    if way == "people" {
        db,err:=opendb(c)      //连接数据库
        if err!=nil {
        return err
        }    
    result:=db.QueryRow("select name,userimage from userinfo where id=?",id)      //单行查询用户基本信息
    var name,userimage string
    result.Scan(&name,&userimage)
    // fmt.Println(result)
    if name !="" {
        result,err := db.Query("select id,title,images,message,media from cards where user_id=?",id)
        if err!=nil {
            return err
            }
            defer db.Close()    //关闭数据库
    i :=0
    cards:=map[int]interface{}{}
    for result.Next(){        //循环显示所有的数据
        
        var id,title,images,message,media string
        result.Scan(&id,&title,&images,&message,&media)
        image_list := strings.Split(images,"|")
        card:=map[string]interface{}{"id":id,"title":title,"message":message,"images":image_list,"media":media}
        cards[i] = card
        i++ 
      
    }
    resp := map[string]interface{}{"name":name,"userimage":userimage,"cards":cards}
    return c.JSON(http.StatusOK,resp)
    }
    }
    if way =="card" {
        if id == "all" {
            db,err:=opendb(c)      //连接数据库
            if err!=nil {
            return err
            }
            result,err := db.Query("select id,title,images,message,media,user_id from cards ")
            if err!=nil {
                return err
                }
                defer db.Close()    //关闭数据库
            // fmt.Println(result)
        // columns, _ := result.Columns()
        // columnLength := len(columns)
        cards:=map[int]interface{}{}
        i :=0
        for result.Next(){        //循环显示所有的数据
         
            var id,title,images,message,media,user_id string
            result.Scan(&id,&title,&images,&message,&media,&user_id)
            image_list := strings.Split(images,"|")
            res:=db.QueryRow("select name from userinfo where id=?",user_id) 
            var name string
            res.Scan(&name)
            card:=map[string]interface{}{"id":id,"title":title,"message":message,"media":media,"images":image_list,"user_id":user_id,"name":name}
            cards[i] = card
            i++
        }
    // card:=map[string]interface{}{"id":id,"title":title,"message":message}
    // cards:=map[int]interface{}{1:card,2:card}
    resp := map[string]interface{}{"cards":cards}
    return c.JSON(http.StatusOK,resp)
        }
    }
    return c.JSON(http.StatusOK,resp)
}
//连接数据库
func opendb(c echo.Context) (*sql.DB ,error) {
    db,_:=sql.Open("mysql","root:xfz123456@(127.0.0.1:3306)/goblin") // 设置连接数据库的参数
      
        err:=db.Ping()      //连接数据库
        if err!=nil{
            fmt.Println("数据库连接失败")
            return db,c.JSON(http.StatusOK, "数据库连接失败")
        }
        return db,err

}

func main() {
    e := echo.New()
    //跨域中间件
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://127.0.0.1:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
        AllowCredentials: true,
      }))
    e.POST("/zhuce", zhuce)            //注册路由
    e.POST("/denglu", denglu)           //登陆路由
    e.POST("/uploadImage",uploadImage)   //上传图片路由
    e.POST("/uploadMedia",uploadMedia)   //上传视频路由
    e.POST("/release",release)           //发布card路由
    e.GET("/show/:where/:id",showEverthing)  //获取信息接口
    e.Static("/avator","./avator")    //用户头像资源
    e.Static("/images","./images")    //用户图片资源
    e.Static("/media","./media")    //用户视频资源
	e.Logger.Fatal(e.Start(":8000"))
    
}

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
        claims["user_id"]="expired" 
	}
	return claims
}
