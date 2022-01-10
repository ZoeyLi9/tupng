package main

import (
	"github.com/golang/glog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"tupng/id"
)

var (
	root       string //输入文件的路径
	outputPath string //输出文件的路径
	nameIs     int    //用户是否维护原文件名
	width      int    //可以自定义输出图像宽度
	quality    int    //控制输出图片的质量

)

//获取文件路径，将其存入channel
func getFile(root string) (value chan string, err chan error) {
	err = make(chan error, 1)
	value = make(chan string)
	//开一个goroutine遍历根目录文件
	go func() {
		//记得结束后关闭channel
		defer close(value)
		//遍历根目录
		err <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error { //函数作为参数，嵌套使用函数
			if err != nil {
				return err
			}
			//判断文件信息是否合法，若不合法，返回空，即不记录其path
			if !info.Mode().IsRegular() {
				return nil
			}
			//将path的值传给value
			value <- path
			return nil
		})
	}()
	return
}

//接收文件并将其发送到channel中处理

//生成图片的唯一id，用于分辨批量处理的各文件
func genID1() string {
	u, err := id.NewUUID(id.Version1, nil)
	if err != nil {
		glog.Error(err) //将错误信息记录为日志，可用于调试
	}
	glog.V(1).Info("the id is generated by UUID")
	return u.String()
}

//默认使用雪花算法生成id
func genID2() string {
	s, err := id.NewSnowFlake(66) //本机的节点号
	if err != nil {
		glog.Error(err)
	}
	glog.V(1).Info("the id is generated by SnowFlake")
	return strconv.FormatInt(s.GetID(), 10) //返回10进制的字符串，整数转换为字符串
}

//获取转换后文件后缀名
//JPEG格式文件不做处理，过滤
func getName(name string) string {
	name = strings.ToLower(name) //将文件名全部转换为小写
	v1 := name[len(name)-4:]
	v2 := name[len(name)-3:]
	if v1 == "jpeg" {
		return v1
	}
	if v2 == "jpg" || v2 == "png" || v2 == "gif" {
		return v1
	}
	return ""
}
