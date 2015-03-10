package gocmdo

import (
	"log"
	"os"
	"io"
	"code.google.com/p/draw2d/draw2d"
	"code.google.com/p/freetype-go/freetype/truetype"
	"github.com/ingar/igo"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"strconv"
	"strings"
	"fmt"
)

var gobanColor = color.RGBA{0xFF, 0xCB, 0x00, 0xFF}

type boardContext struct {
	boardSize, imageSize int
	borderWidth          float64
	gridUnit             float64
	ctx *draw2d.ImageGraphicContext
	image                draw.Image
}

func newBoardContext(boardSize, imageSize int) (*boardContext) {
	var img draw.Image
	var ctx *draw2d.ImageGraphicContext
	rect := image.Rect(0, 0, imageSize, imageSize)
	img = image.NewRGBA(rect)
	ctx = draw2d.NewGraphicContext(img)

	ctx.Save()
	ctx.SetFillColor(color.White)
	ctx.ClearRect(0, 0, imageSize, imageSize)
	ctx.Restore()

	borderWidth := float64(imageSize) * 0.1
	scale := float64(imageSize) - 2.0 * borderWidth
	ctx.Translate(float64(borderWidth), float64(borderWidth))
	ctx.Scale(scale, scale)
	unit := 1.0 / float64(boardSize)

	bc := &boardContext{boardSize, imageSize, borderWidth, unit, ctx, img}
	bc.drawGoban()

	return bc
}

func (self *boardContext) Png(w io.Writer) {
	if err := png.Encode(w, self.image); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (self *boardContext) drawGrid() {
	ctx := self.ctx
	ctx.Save()
	ctx.Scale(self.gridUnit, self.gridUnit)
	ctx.Translate(0.5, 0.5)
	ctx.SetStrokeColor(color.Black)
	ctx.SetLineWidth(0.05)

	for i := 0; i < self.boardSize; i++ {
		ctx.MoveTo(float64(i), 0)
		ctx.LineTo(float64(i), float64(self.boardSize - 1))
		ctx.MoveTo(0, float64(i))
		ctx.LineTo(float64(self.boardSize - 1), float64(i))
	}
	ctx.Stroke()
	ctx.Restore()
}

func (self *boardContext) drawHoshiPoints() {
	if self.boardSize < 13 {
		return
	}

	self.drawHoshi(3, 3)
	self.drawHoshi(self.boardSize - 1 - 3, 3)
}

func (self *boardContext) drawCoordinates() {
	fontSize := 0.5

	ctx := self.ctx

	ctx.Save()
	ctx.Scale(self.gridUnit, self.gridUnit)
	ctx.Translate(0.5, 0.5)

	ctx.SetFontData(draw2d.FontData{"SourceCodePro-Semibold", draw2d.FontFamilyMono, draw2d.FontStyleNormal})
	ctx.SetFontSize(fontSize)
	ctx.SetFillColor(color.Black)

	letters := "ABCDEFGHJKLMNOPQRST"[0:self.boardSize]
	for idx, letter := range letters {
		left, _, right, _ := ctx.GetStringBounds(string(letter))
		x := float64(idx) - (right - left) * 0.5
		ctx.FillStringAt(string(letter), x, -1.0)
		ctx.FillStringAt(string(letter), x, float64(self.boardSize) + fontSize)
	}

	for i := 0; i < self.boardSize; i++ {
		number := i + 1
		s := string(fmt.Sprintf("%d", number))
		left, top, right, bottom := ctx.GetStringBounds(s)
		x := -float64(right - left) - 0.5 - 0.6
		y := float64(self.boardSize - number) + (bottom - top) / 2
		ctx.FillStringAt(s, x, y)

		x = float64(self.boardSize - 1) + .5 + .5
		ctx.FillStringAt(s, x, y)
	}

	ctx.Stroke()
	ctx.Restore()
}

func (self *boardContext) drawGoban() {
	ctx := self.ctx
	ctx.Save()
	border := (1.0 / float64(self.boardSize)) * 0.25
	ctx.SetFillColor(gobanColor)
	draw2d.Rect(ctx, -border, -border, 1.0+border, 1.0+border)
	ctx.Fill()

	self.drawGrid()
	self.drawCoordinates()

	ctx.Restore()
}

// Not sure how to make strokes appear on the inside of paths, so just fill twice to get the border
func (self *boardContext) drawStone(x, y int, black bool) {
	outlineWidth := 0.05

	ctx := self.ctx
	ctx.Save()
	ctx.Translate(0.0, 1.0)
	ctx.Scale(1.0, -1.0)
	ctx.Translate((float64(x)+0.5)*self.gridUnit, (float64(y)+0.5)*self.gridUnit)

	// draw outline
	ctx.SetFillColor(image.Black)
	draw2d.Circle(ctx, 0, 0, self.gridUnit*0.5)
	ctx.Fill()

	if !black {
		ctx.SetFillColor(image.White)
	}
	draw2d.Circle(ctx, 0, 0, self.gridUnit*(0.5-outlineWidth))
	ctx.Fill()
	ctx.Restore()
}

func (self *boardContext) drawHoshi(x, y int) {
	ctx := self.ctx
	ctx.Save()
	ctx.Translate(0.0, 1.0)
	ctx.Scale(1.0, -1.0)
	ctx.Translate((float64(x)+0.5)*self.gridUnit, (float64(y)+0.5)*self.gridUnit)
	ctx.SetFillColor(image.Black)
	draw2d.Circle(ctx, 0, 0, self.gridUnit*0.25)
	ctx.Fill()
	ctx.Restore()
}

func boardHandler(w http.ResponseWriter, req *http.Request) {
	boardSize, _ := strconv.Atoi(req.FormValue("size"))

	imageSize := 600
	bc := newBoardContext(boardSize, imageSize)

	black := true
	for _, moves := range []string{req.FormValue("black"), req.FormValue("white")} {
		matches := igo.A1Regexp.FindAllStringSubmatch(strings.ToLower(moves), -1)

		for _, match := range matches {
			a1coords := match[0]
			coords, err := igo.A1toXY(a1coords)

			if err != nil {
				panic(fmt.Sprintf("Bad A1-style coordinate: %v", a1coords))
			}

			if coords.X >= boardSize ||
				coords.Y >= boardSize ||
				coords.X < 0 ||
				coords.Y < 0 {
				panic(fmt.Sprintf("Coordinates out of range for board size %v: %v", boardSize, a1coords))
			}

			bc.drawStone(coords.X, coords.Y, black)
		}
		black = !black
	}
	bc.Png(w)
}

var boardFont *truetype.Font

func loadFont() (err error) {
	draw2d.SetFontFolder("./fonts/")

	file, err := os.Open("./fonts/Courier New.ttf")
	if err != nil {
		log.Fatal(err)
		return
	}

	data := make([]byte, 1024*1024)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	boardFont, err = truetype.Parse(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func StartGobanServer() {
	loadFont()
	http.HandleFunc("/board", boardHandler)
	http.ListenAndServe(":8000", nil)
}
