package mark

import (
	"encoding/json"
	"fmt"
	imgtype "github.com/shamsher31/goimgtype"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

const MaxWidth float64 = 600

type configuration struct {
	Enabled bool
	Path    string
}

type Config struct {
	Path string `json:"path"`
	Mark string `json:"mark"`
	SaveAsNewFolder int `json:"saveAsNewFolder"`
	SaveAsNewFolderAddress string `json:"saveAsNewFolderAddress"`
	ImageSizeAllow int `json:"imageSizeAllow"`
}

func NewJsonStruct() *Config {
	return &Config{}
}

func (jst *Config) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func getFileList(path string, mark string) {
	fs,_:= ioutil.ReadDir(path)

	wmb, ok := os.Open(mark)
	if ok != nil{
		fmt.Print("图2读取失败！")
		return
	}
	datatype, err2 := imgtype.Get(mark)
	if err2 != nil {
		fmt.Print("水印图片文件类型获取失败！文件类型",datatype)
		return
	}
	watermark, ok := png.Decode(wmb)
	if ok != nil{
		fmt.Print("水印图片解析失败！文件类型",datatype)
		return
	}
	defer wmb.Close()

	for _,file:=range fs{
		if file.IsDir(){
			getFileList(path+file.Name()+"/", mark)
		}else{
			// 获取图片的类型
			datatype, err2 := imgtype.Get(path+file.Name())
			if err2 != nil {
				continue
			} else{
				imgb, ok := os.Open(path+file.Name())
				if ok != nil{
					fmt.Print("图1读取失败！")
					return
				}

				var b image.Rectangle

				// 根据文件类型执行响应的操作
				switch datatype {
				case `image/jpeg`:
					img, ok := jpeg.Decode(imgb)
					if ok != nil{
						fmt.Print("jpeg图片解析失败")
						return
					}

					defer imgb.Close()
					b = img.Bounds()
					offset := image.Pt(0, 0)
					m := image.NewRGBA(b)
					draw.Draw(m, b, img, image.ZP, draw.Src)
					draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

					imgw, _ := os.Create(path+file.Name())


					jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
					defer imgw.Close()
				case `image/png`:
					img, ok := png.Decode(imgb)
					if ok != nil{
						fmt.Print("png图片解析失败")
						return
					}
					defer imgb.Close()
					b = img.Bounds()
					offset := image.Pt(0, 0)
					m := image.NewRGBA(b)
					draw.Draw(m, b, img, image.ZP, draw.Src)
					draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

					imgw, _ := os.Create(path+file.Name())


					jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
					defer imgw.Close()
				}
			}
		}
	}
}

func Mark() {
	JsonParse := NewJsonStruct()
	v := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("G:/999/config.json", &v)

	getFileList(v.Path, v.Mark)
}
