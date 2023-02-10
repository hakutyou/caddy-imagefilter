package watermark

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/caddyserver/caddy/v2"
	imagefilter "github.com/hakutyou/caddy-imagefilter"
)

type WatermarkFactory struct{}

type Watermark struct {
	WatermarkPath string `json:"watermark_path,omitempty"`
}

func (ff WatermarkFactory) Name() string { return "watermark" }

func (ff WatermarkFactory) New(args ...string) (imagefilter.Filter, error) {
	if len(args) < 1 {
		return nil, imagefilter.ErrTooFewArgs
	}

	if len(args) > 1 {
		return nil, imagefilter.ErrTooManyArgs
	}
	return Watermark{WatermarkPath: args[0]}, nil
}

func (ff WatermarkFactory) Unmarshal(data []byte) (imagefilter.Filter, error) {
	filter := Watermark{}
	err := json.Unmarshal(data, &filter)
	if err != nil {
		return nil, err
	}
	return filter, nil
}

func (f Watermark) Apply(repl *caddy.Replacer, img image.Image) (image.Image, error) {
	watermarkPath := repl.ReplaceAll(f.WatermarkPath, "")
	watermark_f, err := os.Open(watermarkPath)
	if err != nil {
		return img, fmt.Errorf("open watermark image failed: %w", err)
	}
	watermark_img, err := png.Decode(watermark_f)
	if err != nil {
		return img, fmt.Errorf("watermark image is not a valid png file: %w", err)
	}

	// watermark position
	offset := image.Pt(img.Bounds().Dx()-watermark_img.Bounds().Dx()-10, img.Bounds().Dy()-watermark_img.Bounds().Dy()-10)
	b := img.Bounds()
	new_image := image.NewRGBA(b)
	draw.Draw(new_image, b, img, image.ZP, draw.Src)
	draw.Draw(new_image, watermark_img.Bounds().Add(offset), watermark_img, image.ZP, draw.Over)
	return new_image, nil
}

func init() {
	imagefilter.Register(WatermarkFactory{})
}

var (
	_ imagefilter.FilterFactory = (*WatermarkFactory)(nil)
	_ imagefilter.Filter        = (*Watermark)(nil)
)
