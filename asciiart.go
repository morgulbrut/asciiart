package asciiart

import (
	"fmt"
	"strconv"
	"strings"
)

type PNM struct {
	MagicNumber string
	Comments    []string
	Width       int
	Height      int
	NoOfCols    int
	ImageData   []int
}

func parsePNM(pnm string) PNM {
	var ret PNM
	var tok []string
	// Sanitize input
	file := strings.ReplaceAll(pnm, "\r", "\n")
	file = strings.ReplaceAll(file, "\t", " ")
	file = strings.ReplaceAll(file, "  ", " ")
	file = strings.ReplaceAll(file, "\n\n", "\n")
	lines := strings.Split(file, "\n")
	for _, l := range lines {
		tokens := strings.Split(l, "#")
		if len(tokens) > 1 {
			for _, t := range tokens[1:] {
				ret.Comments = append(ret.Comments, t)
			}
		}
		tns := strings.Split(tokens[0], " ")
		for _, t := range tns {
			if t != "" {
				tok = append(tok, t)
			}
		}
	}
	ret.MagicNumber = tok[0]
	ret.Width, _ = strconv.Atoi(tok[1])
	ret.Height, _ = strconv.Atoi(tok[2])
	if ret.MagicNumber != "P1" {
		ret.NoOfCols, _ = strconv.Atoi(tok[3])
		for _, s := range tok[4:] {
			val, _ := strconv.Atoi(s)
			ret.ImageData = append(ret.ImageData, val)
		}
	} else {
		ret.NoOfCols = 2
		for _, s := range tok[3:] {
			val, _ := strconv.Atoi(s)
			ret.ImageData = append(ret.ImageData, val)
		}
	}
	return ret
}

func Info(pnm string) {
	info := parsePNM(pnm)
	fmt.Println()
	fmt.Printf("MagicNumber:\n\t%s\n", info.MagicNumber)
	fmt.Printf("Comments:\n\t%s\n", info.Comments)
	fmt.Printf("Width:\n\t%d\n", info.Width)
	fmt.Printf("Height:\n\t%d\n", info.Height)
	fmt.Printf("NoOfCols:\n\t%d\n", info.NoOfCols)
	fmt.Printf("ImageData:\n\t%d\n", info.ImageData)
}

func RenderPNMBW(pnm string) string {
	img := parsePNM(pnm)
	if img.MagicNumber == "P2" {
		return renderP2BW(img)
	}
	if img.MagicNumber == "P1" {
		return renderP2BW(img)
	}
	return "not valid PGM"
}

func renderP2BW(img PNM) string {
	var sb strings.Builder
	colors := 5
	debug := false
	sb.WriteString("\n")

	for i := 0; i < img.Height; i++ {
		for j := 0; j < img.Width; j++ {
			if img.ImageData[j+i*img.Width] == 0 {
				sb.WriteRune(' ')
			} else if img.ImageData[j+i*img.Width] == img.NoOfCols {
				sb.WriteRune('█')
			} else if img.ImageData[j+i*img.Width] < img.NoOfCols/colors {
				sb.WriteRune(' ')
			} else if img.ImageData[j+i*img.Width] < img.NoOfCols/colors*2 {
				sb.WriteRune('░')
			} else if img.ImageData[j+i*img.Width] < img.NoOfCols/colors*3 {
				sb.WriteRune('▒')
			} else if img.ImageData[j+i*img.Width] < img.NoOfCols/colors*4 {
				sb.WriteRune('▓')
			} else if img.ImageData[j+i*img.Width] < colors {
				sb.WriteRune('█')
			} else {
				sb.WriteRune(' ')
			}
		}
		if debug {
			sb.WriteString("#\n")
		} else {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func Blockelems() string {
	return "▀ ▁ ▂ ▃ ▄ ▅ ▆ ▇ █ ▉ ▊ ▋ ▌ ▍ ▎ ▏ ▐ ░ ▒ ▓ ▔ ▕ ▖ ▗ ▘ ▙ ▚ ▛ ▜ ▝ ▞ ▟"
}
