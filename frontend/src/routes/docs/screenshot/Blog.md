This document provides an overview of the key parameters and their functionalities in the Screenshot API.


## Try it out in the Playground

Experience the power and flexibility of the Screenshot API in our interactive playground. The playground allows you to experiment with different parameters and see the results in real-time.

To access the playground:

1. Navigate to the [Screenshot API Playground](/playground).
2. Specify the `url` of the webpage you want to capture.
3. Experiment with different settings like `full_screen`, `scroll_delay`, `v_width`, and others to see how they affect the screenshot.
4. Click on the `Capture` button to take a screenshot with your specified settings.

Remember, the playground is a great place to test and understand the functionalities of the Screenshot API before integrating it into your application.

## URL

- **Variable Name:** `url`
- **Description:** The URL of the webpage to capture.
- **Usage:** Specify the target webpage URL.
- **Example:** `url=https://example.com`

## Access Key

- **Variable Name:** `access_key`
- **Description:** Authentication key for API access.
- **Usage:** Include this key in your API request for authorization.
- **Example:** `access_key=YOUR_ACCESS_KEY`

## Full Screen Capture

- **Variable Name:** `full_screen`
- **Description:** Option to capture the entire screen of the webpage.
- **Usage:** Set to `true` for full-screen capture.
- **Example:** `full_screen=true`

## Scroll Delay

- **Variable Name:** `scroll_delay`
- **Description:** Delay time in seconds before capturing the screenshot, allowing page elements to load after scrolling.
- **Usage:** Specify the delay time in seconds.
- **Example:** `scroll_delay=2`

## Viewport Width

- **Variable Name:** `v_width`
- **Description:** Width of the viewport in pixels for the screenshot.
- **Usage:** Set the width of the capturing viewport.
- **Example:** `v_width=1200`

## Viewport Height

- **Variable Name:** `v_height`
- **Description:** Height of the viewport in pixels for the screenshot.
- **Usage:** Set the height of the capturing viewport.
- **Example:** `v_height=800`

## Delay

- **Variable Name:** `delay`
- **Description:** Time in seconds to wait before taking the screenshot after the page loads.
- **Usage:** Specify the delay time in seconds.
- **Example:** `delay=5`

## Timeout

- **Variable Name:** `timeout`
- **Description:** Maximum time in seconds to wait for the page to load before timing out.
- **Usage:** Set the timeout duration.
- **Example:** `timeout=30`

## No Ads

- **Variable Name:** `no_ads`
- **Description:** Option to remove ads from the page before taking the screenshot.
- **Usage:** Set to `true` to remove ads.
- **Example:** `no_ads=true`

## No Cookie Banner

- **Variable Name:** `no_cookie_banner`
- **Description:** Option to remove cookie banners from the page before taking the screenshot.
- **Usage:** Set to `true` to remove cookie banners.
- **Example:** `no_cookie_banner=true`

## Block Tracker

- **Variable Name:** `block_tracker`
- **Description:** Option to block tracking scripts on the page.
- **Usage:** Set to `true` to block trackers.
- **Example:** `block_tracker=true`

## Custom

- **Variable Name:** `custom`
- **Description:** Custom settings for the screenshot capture defined in the database.
- **Usage:** Specify custom settings.
- **Example:** `custom=modal`

## Response Type [BETA]

- **Variable Name:** `response_type`
- **Description:** Format of the API response.
- **Usage:** Specify the response format (e.g., JSON, XML).
- **Example:** `response_type=json`

## Save to S3

- **Variable Name:** `save_to_s3`
- **Description:** Option to save the screenshot directly to Amazon S3.
- **Usage:** Set to `true` and provide S3 credentials in the custom settings.
- **Example:** `save_to_s3=true`

## Path File Name

- **Variable Name:** `path_file_name`
- **Description:** The path and file name for saving the screenshot.
- **Usage:** Specify the path and file name without format file
- **Example:** `path_file_name=/path/to/file`

## Format

- **Variable Name:** `format`
- **Description:** The format of the screenshot image.
- **Usage:** Specify the image format (e.g., PNG, JPG).
- **Example:** `format=png`

## Quality

- **Variable Name:** `quality`
- **Description:** The quality of the screenshot image (relevant for formats like JPG).
- **Usage:** Specify the image quality as a percentage.
- **Example:** `quality=80`


## Element HTML
- **Variable Name:** `element`
- **Description:** Use the `element` variable to specify the HTML element you want to capture in the screenshot. To select a specific element, you can right-click on it and choose 'Copy Selector' in the inspector. Then, set the value of `element` to the copied selector. This allows you to focus on capturing a specific element rather than the entire page.
- **Usage:** Specify the image quality as a percentage.
- **Usage for element:** Use the `element` variable to specify the HTML element you want to capture in the screenshot. To select a specific element, you can right-click on it and choose 'Copy Selector' in the inspector. Then, set the value of `element` to the copied selector. This allows you to focus on capturing a specific element rather than the entire page.
- **Example:** `element=body>div>main>section`