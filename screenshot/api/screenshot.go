package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"screenshot/lib"
	"screenshot/module"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/labstack/echo/v4"
)

func TakeScreenshot(c echo.Context) error {
	//get query url
	url := c.QueryParam("url")
	if url == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "url is required",
		})
	}
	access_key := c.QueryParam("access_key")
	if access_key == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is required",
		})
	}
	widthStr := c.QueryParam("v_width")
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		width = 1280
	}
	heightStr := c.QueryParam("v_height")
	height, err := strconv.ParseInt(heightStr, 10, 64)
	if err != nil {
		height = 1024
	}
	fullScreenStr := c.QueryParam("full_screen")
	fullScreen, err := strconv.ParseBool(fullScreenStr)
	if err != nil {
		fullScreen = false
	}

	scrollDelayStr := c.QueryParam("scroll_delay")
	scrollDelay, err := strconv.ParseInt(scrollDelayStr, 10, 64)
	if err != nil {
		scrollDelay = 1
	}

	noAdsStr := c.QueryParam("no_ads")
	noAds, err := strconv.ParseBool(noAdsStr)
	if err != nil {
		noAds = false
	}

	noCookieStr := c.QueryParam("no_cookie_banner")
	noCookie, err := strconv.ParseBool(noCookieStr)
	if err != nil {
		noCookie = false
	}

	//delay
	delayStr := c.QueryParam("delay")
	delay, err := strconv.ParseInt(delayStr, 10, 64)
	if err != nil {
		delay = 0
	}

	blockTrackerStr := c.QueryParam("block_trackers")
	blockTracker, err := strconv.ParseBool(blockTrackerStr)
	if err != nil {
		blockTracker = false
	}

	timeoutStr := c.QueryParam("timeout")
	timeout, err := strconv.ParseInt(timeoutStr, 10, 64)
	if err != nil {
		timeout = 60
	}

	body := c.Request().Body
	customData := module.CustomSet{}
	if body != nil {
		json_map := make(map[string]interface{})
		bodyBytes, _ := ioutil.ReadAll(body)
		validJson := strings.Replace(string(bodyBytes), "\r", "", -1)
		err := json.NewDecoder(strings.NewReader(validJson)).Decode(&json_map)
		if err != nil {
			log.Println("err 1", err)
		}
		customJson, err := json.Marshal(json_map["custom"])
		if err != nil {
			log.Println("err 2", err)
		}
		//customJson to struct customData
		err = json.Unmarshal(customJson, &customData)
		if err != nil {
			log.Println("err 3", err)
		}
	}

	pathFileName := c.QueryParam("path_file_name")
	if pathFileName == "" {
		//randomString
		pathFileName = lib.GenerateRandomString(10)
	}

	responseType := c.QueryParam("response_type")
	if responseType == "" {
		responseType = "image"
	}

	imageQualityStr := c.QueryParam("quality")
	imageQuality, err := strconv.Atoi(imageQualityStr)
	if err != nil || imageQuality < 0 || imageQuality > 100 {
		imageQuality = 100 // Default quality
	}

	imageFormat := c.QueryParam("format")
	if imageFormat == "" {
		imageFormat = "png" // Default format
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte
	errCh := chromedp.Run(ctx, screenshot(url, width, height, fullScreen, scrollDelay, noAds, noCookie, blockTracker, delay, customData, &buf))
	if errCh != nil {
		log.Printf("Error taking screenshot: %v", errCh)
	}

	log.Println("imageFormat", imageFormat)
	//format image
	if imageFormat != "png" {
		err := lib.FormatImage(&buf, imageFormat)
		if err != nil {
			log.Println("err", err)
		}
	}
	//image quality
	if imageQuality != 100 {
		err := lib.ImageQuality(&buf, imageQuality)
		if err != nil {
			log.Println("err", err)
		}
	}

	imageType := http.DetectContentType(buf)

	//blob
	return c.Blob(http.StatusOK, imageType, buf)
}

func screenshot(url string, width int64, height int64, fullScreen bool, scrollDelay int64, noAds bool, noCookie bool, blockTracker bool, delay int64, customData module.CustomSet, res *[]byte) chromedp.Tasks {
	var newHeight int64
	viewportDivID := "customViewportDiv"
	header := map[string]interface{}{}
	//headers customeData
	if customData.Headers != "" {
		//convert it to network.Headers
		err := json.Unmarshal([]byte(customData.Headers), &header)
		if err != nil {
			log.Println("err", err)
		}

	}
	//check if custom user agent
	if customData.UserAgent != "" {
		header["User-Agent"] = customData.UserAgent
	}
	networkBlockedURLs := []string{}
	if noAds {
		networkBlockedURLs = append(networkBlockedURLs,
			"https://*.doubleclick.net/*",
			"https://*.googleadservices.com/*",
			"https://*.googlesyndication.com/*",
			"https://*.google-analytics.com/*",
			"https://*.googletagmanager.com/*",
			"https://*.google.com/*",
			//ezoic
			"https://*.ezoic.net/*",
			"https://*.ezoic.com/*",
		)
	}
	if blockTracker {
		networkBlockedURLs = append(networkBlockedURLs,
			"https://*.google-analytics.com/*",
			"https://*.googletagmanager.com/*",
			"https://*.facebook.com/*",
			"https://*.facebook.net/*",
			"https://*.twitter.com/*",
			"https://*.scorecardresearch.com/*",
			"https://*.quantserve.com/*",
			"https://*.adnxs.com/*",
			"https://*.adsrvr.org/*",
			"https://*.adroll.com/*",
			"https://*.taboola.com/*",
			"https://*.outbrain.com/*",
		)
	}

	if fullScreen {
		//print log
		return chromedp.Tasks{
			network.Enable(),
			network.SetExtraHTTPHeaders(header),
			network.SetBlockedURLS(networkBlockedURLs),
			chromedp.Navigate(url),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// use mainCustomScript
				err := mainCustomScript(customData, ctx)
				if err != nil {
					return err
				}
				return nil
			}),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.Sleep(time.Duration(delay) * time.Second),
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
						time.Sleep(time.Duration(scrollDelay) * time.Second) // Wait for content to load; adjust the delay as needed
					} else {
						// back to top
						err := chromedp.Evaluate(`window.scrollTo(0,0)`, nil).Do(ctx)
						if err != nil {
							return err
						}
					}
				}

				return nil
			}),
			chromedp.ActionFunc(func(ctx context.Context) error {
				script, err := mainScript(noAds, noCookie, blockTracker, viewportDivID, width, newHeight, customData, ctx)
				if (err) != nil {
					return err
				}
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			network.Enable(),
			network.SetExtraHTTPHeaders(header),
			network.SetBlockedURLS(networkBlockedURLs),
			chromedp.Navigate(url),
			chromedp.ActionFunc(func(ctx context.Context) error {
				// use mainCustomScript
				err := mainCustomScript(customData, ctx)
				if err != nil {
					return err
				}
				return nil
			}),
			chromedp.EmulateViewport(width, height),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
			chromedp.ActionFunc(func(ctx context.Context) error {
				script, err := mainScript(noAds, noCookie, blockTracker, viewportDivID, width, height, customData, ctx)
				if (err) != nil {
					return err
				}
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Sleep(time.Duration(delay) * time.Second),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}

func mainCustomScript(customData module.CustomSet, ctx context.Context) error {
	if customData.Cookies != "" {
		err := chromedp.Evaluate(fmt.Sprintf(`document.cookie = "%s";`, customData.Cookies), nil).Do(ctx)
		if err != nil {
			return err
		}
	}

	if customData.LocalStorage != "" {
		localStorageScript := ""
		for _, pair := range strings.Split(customData.LocalStorage, ";") {
			parts := strings.SplitN(pair, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				localStorageScript += fmt.Sprintf(`localStorage.setItem("%s", "%s");`, key, value)
			}
		}
		if localStorageScript != "" {
			err := chromedp.Evaluate(localStorageScript, nil).Do(ctx)
			if err != nil {
				return err
			}
		}
	}

	if customData.Cookies != "" || customData.LocalStorage != "" {
		err := chromedp.Reload().Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func mainScript(noAds bool, noCookie bool, blockTracker bool, viewportDivID string, width int64, height int64, customData module.CustomSet, ctx context.Context) (string, error) {
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
		script += noAdsFun()
	}
	if noCookie {
		script += noCookieFunc()
	}

	if customData.CSS != "" {
		// Escape single quotes in CSS
		css := strings.ReplaceAll(customData.CSS, "'", "\\'")

		// Create the script to inject CSS
		styleCss := "var styleCss = document.createElement('style');"
		styleCss += "styleCss.innerHTML = `" + css + "`;"
		styleCss += "document.head.appendChild(styleCss);"
		script += styleCss
	}

	if customData.JavaScript != "" {
		var scriptCustom string
		if strings.Contains(customData.JavaScript, "\n") {
			js, _ := json.Marshal(customData.JavaScript)
			scriptCustom = fmt.Sprintf(`
				var scriptCustom = document.createElement('script');
				scriptCustom.appendChild(document.createTextNode(%s));
				document.head.appendChild(scriptCustom);
			`, js)
		} else {
			scriptCustom = fmt.Sprintf(`
				var scriptCustom = document.createElement('script');
				scriptCustom.innerHTML = '%s';
				document.head.appendChild(scriptCustom);
			`, strings.ReplaceAll(customData.JavaScript, "'", "\\'"))
		}
		script += scriptCustom
	}

	return script, nil
}

// noCookieFunc is a function that removes cookie banners and clicks on accept buttons to simulate accepting cookies.
// It uses JavaScript code to manipulate the DOM and remove cookie-related elements from the webpage.
// The function does not take any parameters and returns a string containing JavaScript code.
func noCookieFunc() string {
	return `
	var style = document.createElement('style');
	style.innerHTML = 'div[id^="cookie"] { display: none; }';
	document.head.appendChild(style);
	var acceptButtons = ['button[aria-label="Accept cookies"]', 'button[id*="cookie"]', 'button[class*="cookie"]', '.cookie-accept', '.cookie-agree', '.cookie-consent-accept'];
	acceptButtons.forEach(function(selector) {
		var btn = document.querySelector(selector);
		if(btn) {
			btn.click();
		}
	});
	var allButtons = document.querySelectorAll('button');
	allButtons.forEach(function(btn) {
		if(btn.textContent.toLowerCase().includes('cookie')) {
			btn.click();
		}
	});
	var cookieBanner = document.querySelector('div[id*="cookie"]');
	if (cookieBanner) {
		cookieBanner.remove();
	}
	var cookieBanner2 = document.querySelector('div[id*="Cookie"]');
	if (cookieBanner2) {
		cookieBanner2.remove();
	}
	var cookieBanner3 = document.querySelector('div[name*="cookie"]');
	if (cookieBanner3) {
		cookieBanner3.remove();
	}
	var cookieBanner4 = document.querySelector('div[name*="Cookie"]');
	if (cookieBanner4) {
		cookieBanner4.remove();
	}
	`
}

// noAdsFun is a function that returns a string containing JavaScript code to hide ads on a web page.
func noAdsFun() string {
	return `
	var style = document.createElement('style');
	style.innerHTML = 'div[id^="google_ads_iframe"] { display: none; }';
	document.head.appendChild(style);
	//ezo_ad 
	var style2 = document.createElement('style');
	style2.innerHTML = '.ezo_ad { display: none; }';
	document.head.appendChild(style2);
	`
}
