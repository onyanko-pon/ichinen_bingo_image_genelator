package image_genelator

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/chromedp/chromedp"
)

type ImageGenalator struct {
	ImageQuority int
	ImageHeight  int64
	ImageWidth   int64
}

func stateHTML(html string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})
}

func (genelator ImageGenalator) GenImage(ctx context.Context, html string) (io.Reader, error) {
	ctx, cancel := chromedp.NewContext(ctx) // 引数追加
	defer cancel()
	ts := httptest.NewServer(stateHTML(html))
	defer ts.Close()

	var buf []byte

	// TODO エラーハンドリング
	chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.EmulateViewport(genelator.ImageWidth, genelator.ImageHeight),
		chromedp.FullScreenshot(&buf, genelator.ImageQuority),
	)
	return bytes.NewReader(buf), nil
}
