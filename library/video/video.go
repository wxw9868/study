package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName, ffmpeg.KwArgs{"ss": "00:00:30"}).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg", "qscale:v": 2}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func main() {
	filename := "./test/video/lc.mp4"
	filename = "C:\\Users\\wxw9868\\go\\src\\video\\assets\\video\\123118_790_まんチラの誘惑_〜欲求不満な友達のママ〜 _古瀬玲.mp4"
	reader := ExampleReadFrameAsJpeg(filename, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = imaging.Save(img, "./test/video/out.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	//err := ffmpeg.Input("C:\\Users\\wxw9868\\go\\src\\video\\assets\\video\\123118_790_まんチラの誘惑_〜欲求不満な友達のママ〜 _古瀬玲.mp4", ffmpeg.KwArgs{"ss": "1"}).
	//	Output("C:\\Users\\wxw9868\\go\\src\\study\\test\\video\\out2.gif", ffmpeg.KwArgs{"pix_fmt": "rgb24", "t": "3", "r": "3", "vf": "split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"}).
	//	OverWriteOutput().ErrorToStdOut().Run()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//reader := ExampleReadFrameAsJpeg("C:\\Users\\wxw9868\\go\\src\\study\\test\\video\\lc.mp4", 5)
	//img, err := imaging.Decode(reader)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = imaging.Save(img, "C:\\Users\\wxw9868\\go\\src\\study\\test\\video\\out1.jpeg")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//_, err := GetSnapshot("C:\\Users\\wxw9868\\go\\src\\study\\test\\video\\lc.mp4", "C:\\Users\\wxw9868\\go\\src\\study\\test\\video\\test", 1)
	//if err != nil {
	//	return
	//}
}
