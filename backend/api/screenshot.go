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
	//get query url
	url := c.QueryParam("url")
	widthStr := c.QueryParam("width")
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		width = 1280
	}
	heightStr := c.QueryParam("height")
	height, err := strconv.ParseInt(heightStr, 10, 64)
	if err != nil {
		height = 1024
	}
	fullScreenStr := c.QueryParam("fullScreen")
	fullScreen, err := strconv.ParseBool(fullScreenStr)
	if err != nil {
		fullScreen = false
	}
	noAdsStr := c.QueryParam("noAds")
	noAds, err := strconv.ParseBool(noAdsStr)
	if err != nil {
		noAds = false
	}
	//delay
	delayStr := c.QueryParam("delay")
	delay, err := strconv.ParseInt(delayStr, 10, 64)
	if err != nil {
		delay = 2
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, screenshot(url, width, height, fullScreen, noAds, delay, &buf)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Blob(http.StatusOK, "image/png", buf)
}

func screenshot(url string, width int64, height int64, fullScreen bool, noAds bool, delay int64, res *[]byte) chromedp.Tasks {
	var newHeight int64
	viewportDivID := "customViewportDiv"
	if fullScreen {
		//print log
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(time.Duration(delay) * time.Second), // Wait for page to render
			chromedp.ActionFunc(func(ctx context.Context) error {
				// Execute JavaScript to get the total height of the content
				err := chromedp.Evaluate(`document.documentElement.scrollHeight`, &newHeight).Do(ctx)
				if err != nil {
					return err
				}
				return nil
			}),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// Set the viewport with static width and dynamic height
				return chromedp.EmulateViewport(width, newHeight).Do(ctx)
			}),
			chromedp.ActionFunc(func(ctx context.Context) error {
				for i := 0; i < 2; i++ { // You might need to adjust the number of iterations
					if i == 0 {
						err := chromedp.Evaluate(`window.scrollTo(0, window.innerHeight)`, nil).Do(ctx)
						if err != nil {
							return err
						}
						time.Sleep(2 * time.Second) // Wait for content to load; adjust the delay as needed
					} else {
						// back to top
						log.Println("back to top")
						err := chromedp.Evaluate(`window.scrollTo(0,0)`, nil).Do(ctx)
						if err != nil {
							return err
						}
						time.Sleep(2 * time.Second) // Wait for content to load; adjust the delay as needed
					}
				}

				return nil
			}),
			chromedp.ActionFunc(func(ctx context.Context) error {
				script := `
                var div = document.createElement('div');
                div.id = '` + viewportDivID + `';
                div.style.position = 'absolute';
                div.style.top = '0';
                div.style.left = '0';
                div.style.width = '` + strconv.FormatInt(width, 10) + `px';
				div.style.height = document.documentElement.scrollHeight + 'px';
                document.body.appendChild(div);
            `
				if noAds {
					script += `
				var style = document.createElement('style');
				style.innerHTML = 'div[id^="google_ads_iframe"] { display: none; }';
				document.head.appendChild(style);
				//ezo_ad 
				var style2 = document.createElement('style');
				style2.innerHTML = '.ezo_ad { display: none; }';
				document.head.appendChild(style2);
				`
				}
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.EmulateViewport(width, height),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(time.Duration(delay) * time.Second),
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
				if noAds {
					script += `
			var style = document.createElement('style');
			style.innerHTML = 'div[id^="google_ads_iframe"] { display: none; }';
			document.head.appendChild(style);
			//ezo_ad 
			var style2 = document.createElement('style');
			style2.innerHTML = '.ezo_ad { display: none; }';
			document.head.appendChild(style2);
			`
				}
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}
