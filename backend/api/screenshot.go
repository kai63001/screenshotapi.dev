package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	module "backend/module"
)

func TakeScreenshot(c echo.Context, db dbx.Builder, mongo *mongo.Collection, rdb *redis.Client) error {
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

	asyncChromeStr := c.QueryParam("async")
	asyncChrome, err := strconv.ParseBool(asyncChromeStr)
	if err != nil {
		asyncChrome = false
	}

	//get user_id from access_key
	userData := module.UserForKey{}
	errAccessKey := db.Select("user_id").From("access_keys").Where(dbx.NewExp("access_key = {:access_key}", dbx.Params{"access_key": access_key})).One(&userData)
	if errAccessKey != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is invalid",
		})
	}
	//check quota
	quotaData := module.GetQuotaScreenshot{}
	errCheckQuota := db.
		Select("screenshot_usage.screenshots_taken", "subscription_plans.name", "subscription_plans.included_screenshots", "subscription_plans.rate_limit_per_minute").
		From("screenshot_usage").
		InnerJoin("users", dbx.NewExp("users.id = screenshot_usage.user_id")).
		InnerJoin("subscription_plans", dbx.NewExp("subscription_plans.id = users.subscription_plan")).
		Where(dbx.NewExp("screenshot_usage.user_id = {:user_id}", dbx.Params{"user_id": userData.UserId})).
		One(&quotaData)
	if errCheckQuota != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is invalid",
		})
	}

	//check rate limit per minute with redis
	// Key to store the rate limit counter
	key := "rate_limit:" + userData.UserId

	// Increment the counter by 1
	count, err := rdb.Incr(context.Background(), key).Result()
	if err != nil {
		// Handle Redis error
		log.Println("err", err)
	}

	// Set the expiration time to 1 minute
	if count == 1 {
		_, err := rdb.Expire(context.Background(), key, time.Minute).Result()
		if err != nil {
			// Handle Redis error
			log.Println("err", err)
		}
	}

	// Check if the counter exceeds the rate limit
	if count > quotaData.SubscriptionPlans.RateLimitPerMinute {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Rate limit exceeded",
		})
	}

	if quotaData.ScreenshotUsage.ScreenshotTaken >= quotaData.SubscriptionPlans.IncludedScreenshots && quotaData.SubscriptionPlans.Name == "Free" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "You have reached your quota",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte
	if asyncChrome {
		go func() {
			// Execute chromedp tasks in a separate goroutine
			ctxAsync, cancel := chromedp.NewContext(context.Background())
			defer cancel()
			err := chromedp.Run(ctxAsync, screenshot(url, width, height, fullScreen, scrollDelay, noAds, noCookie, blockTracker, delay, &buf))
			if err != nil {
				log.Printf("Error taking screenshot: %v", err)
				// Handle error (e.g., send a notification or log the error)
			}

			// Process the screenshot (e.g., save to database or file system)
			// ...

			// Log or handle the completion of the screenshot task
			log.Println("Screenshot task completed")
		}()
	} else {
		if err := chromedp.Run(ctx, screenshot(url, width, height, fullScreen, scrollDelay, noAds, noCookie, blockTracker, delay, &buf)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}
	}

	//update screenshot_usage
	_, errUpdateScreenshotUsage := db.NewQuery(`
		UPDATE screenshot_usage
		SET screenshots_taken = screenshots_taken + 1
		WHERE user_id = {:user_id}
	`).Bind(dbx.Params{
		"user_id": userData.UserId,
	}).Execute()
	if errUpdateScreenshotUsage != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": errUpdateScreenshotUsage.Error(),
		})
	}

	fullUrl := c.Request().URL.String()
	result, _ := mongo.InsertOne(context.Background(), map[string]interface{}{
		"user_id":    userData.UserId,
		"access_key": access_key,
		"url":        url,
		"fullUrl":    fullUrl,
		"created":    time.Now(),
	})
	log.Println("resul", result)

	if asyncChrome {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   "Screenshot is being processed",
		})
	} else {
		return c.Blob(http.StatusOK, "image/png", buf)
	}
}

func screenshot(url string, width int64, height int64, fullScreen bool, scrollDelay int64, noAds bool, noCookie bool, blockTracker bool, delay int64, res *[]byte) chromedp.Tasks {
	var newHeight int64
	viewportDivID := "customViewportDiv"
	if fullScreen {
		//print log
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
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
					script += noAdsFun()
				}
				if noCookie {
					script += noCookieFunc()
				}

				if blockTracker {
					err := network.SetBlockedURLS(
						[]string{
							"https://*.doubleclick.net/*",
							"https://*.googleadservices.com/*",
							"https://*.googlesyndication.com/*",
							"https://*.google-analytics.com/*",
							"https://*.googletagmanager.com/*",
							"https://*.google.com/*",
						},
					).Do(ctx)
					if err != nil {
						return err
					}
				}

				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Sleep(time.Duration(delay) * time.Second),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	} else {
		return chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.EmulateViewport(width, height),
			chromedp.WaitVisible(`body`, chromedp.ByQuery),
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
					script += noAdsFun()
				}
				if noCookie {
					script += noCookieFunc()
				}
				if blockTracker {
					err := network.SetBlockedURLS(
						[]string{
							"https://*.doubleclick.net/*",
							"https://*.googleadservices.com/*",
							"https://*.googlesyndication.com/*",
							"https://*.google-analytics.com/*",
							"https://*.googletagmanager.com/*",
							"https://*.google.com/*",
						},
					).Do(ctx)
					if err != nil {
						return err
					}
				}
				return chromedp.Evaluate(script, nil).Do(ctx)
			}),
			chromedp.Sleep(time.Duration(delay) * time.Second),
			chromedp.Screenshot("#"+viewportDivID, res, chromedp.NodeVisible, chromedp.ByQuery),
		}
	}
}

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
