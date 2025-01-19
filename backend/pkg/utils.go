package pkg

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

func LoadFiglet() {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
		figlet4go.ColorBlue,
		figlet4go.ColorRed,
	}
	renderStr, _ := ascii.RenderOpts("Redis    Docker    Cosmos", options)
	fmt.Println(renderStr)
}
