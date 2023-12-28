This document provides an overview of the key parameters and their functionalities in the Screenshot API.

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
- **Description:** Delay time in milliseconds before capturing the screenshot, allowing page elements to load after scrolling.
- **Usage:** Specify the delay time in milliseconds.
- **Example:** `scroll_delay=1000`

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
- **Description:** Time in milliseconds to wait before taking the screenshot after the page loads.
- **Usage:** Specify the delay time in milliseconds.
- **Example:** `delay=500`

## Timeout

- **Variable Name:** `timeout`
- **Description:** Maximum time in milliseconds to wait for the page to load before timing out.
- **Usage:** Set the timeout duration.
- **Example:** `timeout=3000`

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

## Response Type

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
- **Usage:** Specify the path and file name.
- **Example:** `path_file_name=/path/to/file.png`

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