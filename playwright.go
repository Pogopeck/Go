package main

import (
    "fmt"
    "log"
    "github.com/playwright-community/playwright-go"
)

func main() {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("Could not start playwright: %v", err)
    }
    defer pw.Stop()

    browser, err := pw.Chromium.Launch()
    if err != nil {
        log.Fatalf("Could not launch browser: %v", err)
    }
    defer browser.Close()

    page, err := browser.NewPage()
    if err != nil {
        log.Fatalf("Could not create page: %v", err)
    }

    // Set viewport size for consistent screenshots
    if err = page.SetViewportSize(1920, 1080); err != nil {
        log.Fatalf("Could not set viewport size: %v", err)
    }

    response, err := page.Goto("https://whatmyuseragent.com/")
    if err != nil {
        log.Fatalf("Could not goto: %v", err)
    }

    // Print useful response information
    fmt.Printf("Status: %d\n", response.Status())
    fmt.Printf("URL: %s\n", response.URL())
    
    // Get the page title
    title, err := page.Title()
    if err != nil {
        log.Printf("Could not get page title: %v", err)
    } else {
        fmt.Printf("Page Title: %s\n", title)
    }

    // Take a screenshot
    screenshot, err := page.Screenshot(playwright.PageScreenshotOptions{
        Path:     playwright.String("screenshot.png"),  // Save to file
        //FullPage: playwright.Bool(true),              // Capture full scrollable page
        //Quality:  playwright.Int(90),                  // JPEG quality (0-100), only when format is jpeg
    })
    if err != nil {
        log.Fatalf("Could not create screenshot: %v", err)
    }

    fmt.Println("Screenshot saved as screenshot.png")

    // Optional: If you want to get the screenshot as bytes instead of saving to file
    // you can omit the Path option and use the screenshot bytes directly
    fmt.Printf("Screenshot size: %d bytes\n", len(screenshot))
}

