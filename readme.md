# 基于golang的日期转时间戳,时间戳转日期

## 使用须知
* linux 系统请把编译后的文件放入bin目录(用户bin目录或全局都可)
* win请加入环境变量


> go build -o ./自己起名        //这里将程序编译为你想要的名字的可执行文件(名字可以自己改,自己看顺眼就可以,目录名为t的文件可以直接用)

> sudo cp ./macos/t /Users/yongnan_wong/bin/        //复制文件到bin目录(如果是window系统请复制windows目录下的.exe文件)

~~~golang
./t 1576022400                //仅支持10位时间戳
本地时间 20191211 08:00:00
Utc时间 20191211 00:00:00
输入的时间戳为 1576022400

./t 20191211                  //日期格式请遵从 20191211 00:00:00 方式
LocalTimeStamp = 1575993600
UtcTimeStamp = 1576022400
输入的时间为 20191211 00:00:00

./t 20191211 08:01:02         //日期格式请遵从 20191211 00:00:00 方式
LocalTimeStamp = 1576022462
UtcTimeStamp = 1576051262
输入的时间为 20191211 08:01:02
~~~
