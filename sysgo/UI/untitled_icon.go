package UI

import (
	"strings"
	
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/cuu/gogame/draw"	
	"github.com/cuu/gogame/color"
)


type UntitledIcon struct {

	PosX int
	PosY int
	Width int
	Height int
	Words [2]string
	FontObj *ttf.Font
	BG *sdl.Surface
	Color *color.Color
	BlankPng string
	Text *sdl.Surface
}

func NewUntitledIcon() *UntitledIcon {
	u := &UntitledIcon{}
	u.Width = 80
	u.Height = 80
	u.Words = [2]string{"G","s"}

	u.FontObj = Fonts["varela40"]

	u.Color = &color.Color{83,83,83,255}

	u.BlankPng = SkinMap("gameshell/blank.png")
}

func (self *UntitledIcon) Init() {
	self.BG = image.Load(self.BlankPng)
)

func (self *UntitledIcon) SetWords( TwoWords ...string) {
	if len(TwoWords) == 1 {
		self.Words[0] = strings.ToUpper(TwoWords[0])
	}
	if len(TwoWords) == 2 {
		self.Words[0] = strings.ToUpper( TwoWords[0])
		self.Words[1] = strings.ToLower( TwoWords[1] )

		self.Text = font.Render(self.FontObj, strings.Join(self.Words,""),true,self.Color, nil)
	}
}

func (self *UntitledIcon) Draw() {
	if self.BG != nil {
		w_ := self.Text.W
		h_ := self.Text.H
		
		surface.Blit(self.BG,self.Text,draw.MidRect(self.Width/2, self.Height/2,w_,h_, self.Width, self.Height),nil)
	}
}

func (self *UntitledIcon) Surface() *sdl.Surface {
	self.Draw()
	return self.BG
}


		