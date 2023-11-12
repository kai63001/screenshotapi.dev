package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"go.mongodb.org/mongo-driver/mongo"

	module "backend/module"
)

func TakeScreenshot(c echo.Context, db dbx.Builder, mongo *mongo.Collection) error {
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
	fullScreenStr := c.QueryParam("full_screen")
	fullScreen, err := strconv.ParseBool(fullScreenStr)
	if err != nil {
		fullScreen = false
	}
	noAdsStr := c.QueryParam("no_ads")
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
		Select("screenshot_usage.screenshots_taken", "subscription_plans.name", "subscription_plans.included_screenshots").
		From("screenshot_usage").
		InnerJoin("subscription_plans", dbx.NewExp("subscription_plans.id = screenshot_usage.subscription_plan")).
		Where(dbx.NewExp("screenshot_usage.user_id = {:user_id}", dbx.Params{"user_id": userData.UserId})).
		One(&quotaData)
	if errCheckQuota != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "access_key is invalid",
		})
	}

	if quotaData.ScreenshotUsage.ScreenshotTaken >= quotaData.SubscriptionPlans.IncludedScreenshots && quotaData.SubscriptionPlans.Name == "FREE" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "You have reached your quota",
		})
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
