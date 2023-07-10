package game

import (
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/poker/fonts"
	"golang.org/x/image/font"
)

const (
	_fontSize = 60
	_two      = 2
)

type StartScene struct {
	bus  *Bus
	face font.Face
	ui   *ebitenui.UI
}

// nolint: gomnd
func NewStart(bus *Bus) *StartScene {
	start := &StartScene{bus: bus, face: fonts.Body(_fontSize)}
	buttonImage, _ := loadButtonImage()
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(35),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(30)),
			),
		),
	)

	button1 := start.newButton("主机", buttonImage, nil)
	rootContainer.AddChild(button1)

	button2 := start.newButton("加入", buttonImage, nil)
	rootContainer.AddChild(button2)
	// construct a standard textinput widget
	standardTextInput := start.newInput()
	rootContainer.AddChild(standardTextInput)
	widget.NewRadioGroup(
		widget.RadioGroupOpts.Elements(button1, button2),
		widget.RadioGroupOpts.ChangedHandler(func(args *widget.RadioGroupChangedEventArgs) {
			// fmt.Println(args.Active.(*widget.Button).Text().Label)
			if args.Active == button1 {
				standardTextInput.GetWidget().Disabled = true
			} else {
				standardTextInput.GetWidget().Disabled = false
				standardTextInput.Focus(true)
			}
		}),
	)

	button := start.newButton("确定", buttonImage, func(args *widget.ButtonClickedEventArgs) {
		logs.D.Println("button clicked")
	})
	rootContainer.AddChild(button)

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	start.ui = ui

	return start
}

// nolint: gomnd
func (p *StartScene) newInput() *widget.TextInput {
	input := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),

		widget.TextInputOpts.Image(&widget.TextInputImage{
			Idle:     image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
			Disabled: image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 100, A: 255}),
		}),

		widget.TextInputOpts.Face(p.face),

		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          color.NRGBA{254, 255, 255, 255},
			Disabled:      color.NRGBA{R: 200, G: 200, B: 200, A: 255},
			Caret:         color.NRGBA{254, 255, 255, 255},
			DisabledCaret: color.NRGBA{R: 200, G: 200, B: 200, A: 255},
		}),

		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),

		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(p.face, 2),
		),

		widget.TextInputOpts.Placeholder("192.168.1.234:3880"),

		widget.TextInputOpts.SubmitHandler(func(args *widget.TextInputChangedEventArgs) {
			logs.D.Println("Text Submitted: ", args.InputText)
		}),

		widget.TextInputOpts.ChangedHandler(func(args *widget.TextInputChangedEventArgs) {
			logs.D.Println("Text Changed: ", args.InputText)
		}),
	)

	input.InputText = "192.168.1.234:3880"

	return input
}

// nolint: gomnd
func (p *StartScene) newButton(
	label string,
	buttonImage *widget.ButtonImage,
	handler widget.ButtonClickedHandlerFunc,
) *widget.Button {
	opts := []widget.ButtonOpt{
		widget.ButtonOpts.WidgetOpts(

			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			}),
		),

		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.Text(label, p.face, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
	}

	if handler != nil {
		opts = append(opts, widget.ButtonOpts.ClickedHandler(handler))
	}

	return widget.NewButton(opts...)
}

func (p *StartScene) Update() error {
	p.ui.Update()

	return nil
}

func (p *StartScene) Keys() map[ebiten.Key]func() { return nil }

func (p *StartScene) Draw(screen *ebiten.Image) {
	p.ui.Draw(screen)

	max := lo.MaxBy(p.bus.Start, func(a, b string) bool { return len(a) > len(b) })
	width, height := p.bus.Layout()
	left := (width - len(max)*_fontSize/_two) / _two
	top := height/_two - (_fontSize*len(p.bus.Start))/_two

	for index, txt := range p.bus.Start {
		text.Draw(screen, txt, p.face, left, top+index*_fontSize, color.RGBA{0xdf, 0xd0, 0x00, 0xff})
	}
}

// nolint: gomnd
func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
