package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

// golang生成pdf
func main() {
	// 创建一个新的PDF文档
	pdf := gofpdf.New("P", "mm", "A4", "")

	// 添加一页
	pdf.AddPage()

	// 设置字体和大小
	pdf.SetFont("Arial", "", 14)

	// 写入文本
	pdf.Text(10, 15, "hello, pom!")

	// 保存PDF文件
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PDF file created successfully.")
}

// 图片生成pdf
//func main() {
//	err := GeneratePdf("hello.pdf")
//	if err != nil {
//		panic(err)
//	}
//}
//
//// GeneratePdf generates our pdf by adding text and images to the page
//// then saving it to a file (name specified in params).
//func GeneratePdf(filename string) error {
//
//	pdf := gofpdf.New("P", "mm", "A4", "")
//	pdf.AddPage()
//	pdf.SetFont("Arial", "B", 16)
//
//	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
//	pdf.CellFormat(190, 7, "Welcome to https://nateshao.gitlab.io/", "0", 0, "CM", false, 0, "")
//
//	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
//	pdf.ImageOptions(
//		"pom.png",
//		80, 20,
//		0, 0,
//		false,
//		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
//		0,
//		"",
//	)
//
//	return pdf.OutputFileAndClose(filename)
//}
