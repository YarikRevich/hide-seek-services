package printlogform

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

func Print(vals ...string)error{
	renderer := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorGreen,
	}
	
	for _, val := range vals{
		if _, err := renderer.RenderOpts(val, options); err != nil{
			return fmt.Errorf("hideseek service errror: print: %w", err)
		}
	}
	return nil
}