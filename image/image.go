package image

import (
	"context"
	"fmt"
	"image"
	"net/http"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/buckket/go-blurhash"
	"go.opentelemetry.io/otel/api/global"
)

func GetBlurHash(ctx context.Context, url string) (string, error) {
	ctx, span := global.Tracer("image").Start(ctx, "image.GetBlurHash")
	defer span.End()

	image, err := GetFromURL(ctx, url)
	if err != nil {
		return "", err
	}

	return blurhash.Encode(4, 3, image)
}

func GetFromURL(ctx context.Context, url string) (image.Image, error) {
	ctx, span := global.Tracer("image").Start(ctx, "image.GetFromURL")
	defer span.End()
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get image %s %w:", url, err)
	}
	defer resp.Body.Close()

	m, format, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to get image %s %w:", url, err)
	}
	span.SetAttribute("image.format", format)
	return m, err

}
