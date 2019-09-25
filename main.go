package main

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"syscall/js"
)

//缩放图片
//支持格式: png, jpeg
//必须参数: imageUint8Array, width, height
//如需保持比例,参数width和height请只设置一个,另一个设为0.
func scaleImage(this js.Value, args []js.Value) interface{} {
	//读取js传入的图片字节
	srcBytes := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(srcBytes, args[0])

	log.Println("参数:图片字节长度:", len(srcBytes))
	log.Println("参数:宽度:", args[1].Int())
	log.Println("参数:高度:", args[2].Int())

	//创建图片
	//注意:必须导入对应格式,比如 import _ "image/png"
	reader := bytes.NewReader(srcBytes)
	img, format, err := image.Decode(reader)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("图片格式:", format)

	//缩放图片
	resizeImg := resize.Resize(uint(args[1].Int()), uint(args[2].Int()), img, resize.Lanczos3)

	//返回缩放后的图片字节给js
	buffer := new(bytes.Buffer)
	if format == "png" {
		err = png.Encode(buffer, resizeImg)
		if err != nil {
			log.Println(err)
			return nil
		}
	} else if format == "jpeg" {
		err = jpeg.Encode(buffer, resizeImg, nil)
		if err != nil {
			log.Println(err)
			return nil
		}
	} else {
		log.Println("不支持图片格式:", format)
		return nil
	}
	outputBytes := buffer.Bytes()
	jsArray := js.Global().Get("Uint8Array").New(len(outputBytes))
	js.CopyBytesToJS(jsArray, outputBytes)

	return jsArray
}

func main() {
	//创建空的通道
	c := make(chan struct{}, 0)

	//注册方法
	js.Global().Set("go_scaleImage", js.FuncOf(scaleImage))

	log.Println("Go WebAssembly:缩放图片")

	//阻塞进程
	<-c
}
