package ecspos

import (
	"github.com/disintegration/imaging"
	"image"
	"io"
)

type Ecspos struct {
	printer    io.Writer
	content    [][]byte
	imageWidth int
}

func New(printer io.Writer) *Ecspos {
	return &Ecspos{
		printer:    printer,
		imageWidth: 350,
	}
}

// SetImageWidth coloca el ancho para la impresion de las imagenes
func (e *Ecspos) SetImageWidth(w int) *Ecspos {
	e.imageWidth = w
	return e
}

func (e *Ecspos) AlignCenter() *Ecspos {
	e.content = append(e.content, []byte{0x1B, 0x61, 0x01})
	return e
}

func (e *Ecspos) AlignLeft() *Ecspos {
	e.content = append(e.content, []byte{0x1B, 0x61, 0x00})
	return e
}

func (e *Ecspos) AlignRight() {
	e.content = append(e.content, []byte{0x1B, 0x61, 0x02})
}

func (e *Ecspos) FontBold() *Ecspos {
	e.content = append(e.content, []byte{0x1B, 0x21, 0x08})
	return e
}

func (e *Ecspos) FontNormal() *Ecspos {
	e.content = append(e.content, []byte{0x1B, 0x21, 0x00})
	return e
}

func (e *Ecspos) Text(txt string) *Ecspos {
	e.content = append(e.content, []byte(txt+"\n"))
	return e
}

func (e *Ecspos) TextBold(txt string) *Ecspos {
	e.FontBold()
	e.Text(txt)
	return e
}

func (e *Ecspos) TextCenter(txt string) *Ecspos {
	e.AlignCenter()
	e.Text(txt)
	return e
}

func (e *Ecspos) TextLeft(txt string) {
	e.AlignLeft()
	e.Text(txt)
}

func (e *Ecspos) TextRight(txt string) {
	e.AlignRight()
	e.Text(txt)
}

func (e *Ecspos) Return() *Ecspos {
	e.content = append(e.content, []byte{0x0D})
	return e
}

func (e *Ecspos) ResetFormat() *Ecspos {
	e.FontNormal()
	e.AlignLeft()
	return e
}

func (e *Ecspos) Image(img image.Image) *Ecspos {
	imageGray := imaging.Grayscale(img)
	imageReduced := imaging.Resize(imageGray, e.imageWidth, 0, imaging.Linear)
	xL, xH, yL, yH, data := getImagePrintableData(imageReduced)
	e.content = append(
		e.content,
		[]byte{0x1D, 'v', 48, 0, xL, xH, yL, yH},
		data,
	)

	return e
}

func (e *Ecspos) Print() {
	e.printer.Write([]byte{0x1B, 0x40})
	for _, c := range e.content {
		e.printer.Write(c)
	}
	// line feed
	for i := 0; i < 4; i++ {
		e.printer.Write([]byte{0x0A})
	}

	// limpiamos el contenido
	e.content = [][]byte{}
}
