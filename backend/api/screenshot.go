package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/labstack/echo/v5"
)

func TakeScreenshot(c echo.Context) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, screenshot(`https://www.thairath.co.th/news/politic/2739898`, 1920, 1080, false, &buf)); err != nil {
		log.Fatal(err)
	}

	return c.Blob(http.StatusOK, "image/png", buf)
}

func screenshot(url string, width int64, height int64, fullScreen bool, res *[]byte) chromedp.Tasks {
	viewportDivID := "customViewportDiv"
	if fullScreen {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second), // Wait for page to render
			chromedp.Screenshot(`body`, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.EmulateViewport(width, height),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.ActionFunc(func(ctx context.Context) error {
				script := `
                var div = document.createElement('div');
                div.id = '` + viewportDivID + `';
                div.style.position = 'absolute';
                div.style.top = '0';
                div.style.left = '0';
                div.style.width = '` + strconv.FormatInt(width, 10) + `px';
                div.style.height = '` + strconv.FormatInt(height, 10) + `px';
                document.body.appendChild(div);
            `
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}

func screenshotStyleCustom(url string, width int64, height int64, fullScreen bool, res *[]byte) chromedp.Tasks {
	viewportDivID := "customViewportDiv"
	if fullScreen {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second), // Wait for page to render
			chromedp.ActionFunc(func(ctx context.Context) error {
				script := `
				var style = document.createElement('style');
				style.innerHTML = 'div[id^="google_ads_iframe"] { display: none; }';
				document.head.appendChild(style);
				//ezo_ad 
				var style2 = document.createElement('style');
				style2.innerHTML = '.ezo_ad { display: none; }';
				document.head.appendChild(style2);
				`
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),

			chromedp.Screenshot(`body`, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.EmulateViewport(width, height),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(2 * time.Second),
			chromedp.ActionFunc(func(ctx context.Context) error {
				script := `
                var div = document.createElement('div');
                div.id = '` + viewportDivID + `';
                div.style.position = 'absolute';
                div.style.top = '0';
                div.style.left = '0';
                div.style.width = '` + strconv.FormatInt(width, 10) + `px';
                div.style.height = '` + strconv.FormatInt(height, 10) + `px';
                document.body.appendChild(div);
				var style = document.createElement('style');
				style.innerHTML = 'div[id^="google_ads_iframe"] { display: none; }';
				document.head.appendChild(style);
				//ezo_ad 
				var style2 = document.createElement('style');
				style2.innerHTML = '.ezo_ad { display: none; }';
				document.head.appendChild(style2);
            `
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}
