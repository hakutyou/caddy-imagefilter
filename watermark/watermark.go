package watermark

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

type WatermarkFactory struct{}

type Watermark struct {
	WatermarkPath string `json:"watermark_path,omitempty"`
}

func (ff WatermarkFactory) Name() string { return "watermark" }

func (ff WatermarkFactory) New(args ...string) (imageFilter.Filter, error) {
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
	watermarkPathRepl := repl.ReplaceAll(f.WatermarkPath, "")
	watermarkPath, err := strconv.Atoi(watermarkPathRepl)
	if err != nil {
		return img, fmt.Errorf("invalid watermark image path: %w", err)
	}

	watermark_f, err := os.Open(watermarkPath)
	if err != nil {
		return img, fmt.Errorf("open watermark image failed: %w", err)
	}
	watermark_img, err := png.Decode(wmb_f)
	if err != nil {
		return img, fmt.Errorf("watermark image is not a valid png file: %w", err)
	}

	// watermark position
	offset := image.Pt(img.Bounds().Dx()-watermark_img.Bounds().Dx()-10, img.Bounds().Dy()-watermark_img.Bounds().Dy()-10)
	b := img.Bounds()
	new_image := image.NewRGBA(b)
	draw.Draw(new_image, b, img, image.ZP, draw.Src)
	draw.Draw(new_image, water_img.Bounds().Add(offset), water_img, image.ZP, draw.Over)
	return new_image, nil
}
