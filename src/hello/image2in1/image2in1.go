package image2in1

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

const MaxWidth float64 = 600

func main() {
	file1, _ := os.Open("G:/999/主图.png") //打开图片1
	file2, _ := os.Open("G:/999/001.jpg") //打开图片2
	defer file1.Close()
	defer file2.Close()

	// image.Decode 图片
	var (
		img1, img2 image.Image
		err        error
	)
	if img1, _, err = image.Decode(file1); err != nil {
		log.Fatal(err)
		return
	}
	if img2, _, err = image.Decode(file2); err != nil {
		log.Fatal(err)
		return
	}
	m1 := img1.Bounds()
	m2 := img2.Bounds()

	// 将两个图片合成一张
	newWidth := m1.Bounds().Max.X       //新宽度 = 随意一张图片的宽度
	newHeight := m1.Bounds().Max.Y + m2.Bounds().Max.Y // 新图片的高度为两张图片高度的和
	newImg := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight)) //创建一个新RGBA图像
	draw.Draw(newImg, newImg.Bounds(), m1, m1.Bounds().Min, draw.Over) //画上第一张缩放后的图片
	draw.Draw(newImg, newImg.Bounds(), m2, m2.Bounds().Min.Sub(image.Pt(0, m1.Bounds().Max.Y)), draw.Over) //画上第二张缩放后的图片（这里需要注意Y值的起始位置）

	// 保存文件
	imgfile, _ := os.Create("003.jpg")
	defer imgfile.Close()
	jpeg.Encode(imgfile, newImg, &jpeg.Options{100})
}
