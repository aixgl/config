# config
config like php.ini
解析配置文件，解析文本格式文件；文本格式样式如php.ini,目前仅支持二维
支持多个配置文件解析

# 安装(Install)

```
    $ go get github.com/aixgl/config
```

##  解析的文本格式如下

```
    ;dfdfdfd

    db = a
    [DEBUG]
    log = true
    cms = dfdf d

    [nothing]
     best = bests
     m = a
```

# 例子example

```
package main

import (
  "fmt"
  "github.com/aixgl/config"
)

func main() {

  //设置并解析  第一个参数可以任意设置，获取的时候要用设置的配置文件别名；支持多个配置文件的解析
  conf := config.C("G", "config.env")
  confE := config.C()
  fmt.Println(confE.Get())
  fmt.Println(conf.GetAll())

  //根据配置文件键值获取
  fmt.Println(confE.Get())
  debuglog, _ := confE.Get("DEBUG.log").(string)
  fmt.Println(debuglog)
  fmt.Println(confE.Get("DEBUG", "log"))
  fmt.Println(confE.Get("DEBUG", "cms"))
  debug, _ := confE.Get("DEBUG").(map[string]string)
  fmt.Println(debug)
}
```
