# Pastis AWSLambda Framework

Pasits is a "pseudo" Golang web framework to develep AWSLambda APIs. mainly inspired from the [Gin Web Framework](https://github.com/gin-gonic/gin).

## installation:

```
$ go get github.com/arouam/pastis
```
## usage:

1. import pastis in you code
```golang
import "github.com/arouam/pastis"
```
2. instantiate a pastis engine and add some routes
```golang
    func main(){
        r := pastis.New()
        r.GET("/helloworld", func(context *pastis.Context) {
            context.JSON(200, pastis.Object{
                "hello": "world",    
            })
        })
    }
```
## Path parameters:

```golang
    func main(){
        r := pastis.New()
        r.GET("/hello/:name", func(context *pastis.Context) {
            name := context.Param("name")
            context.JSON(200, pastis.Object{
                "hello": name,    
            })
        })
    }
```
## Query parameters:

```golang
    func main(){
        r := pastis.New()
        r.GET("/hello", func(context *pastis.Context) {
            name := context.Query("name")
            context.JSON(200, pastis.Object{
                "hello": name,    
            })
        })
    }
```
## Model binding:

```golang
   type User struct {
     Name string `json:"name"`
   }
   
   func main(){        
        r := pastis.New()
        r.POST("/users", func(context *pastis.Context) {
            var user User
            if err := context.BindJSON(&user); err != nil {
                context.JSON(http.StatusBadRequest, pastis.Object{
                  "err": err.Error(),
                })
            }
            context.JSON(201, user)
        })
    }
```
