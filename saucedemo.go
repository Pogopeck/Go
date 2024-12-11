package main

import (
    "fmt"
    "log"
    "time"
    "github.com/playwright-community/playwright-go"
)

func main() {
    pw, err := playwright.Run()
    if err != nil {
        log.Fatalf("Could not start playwright: %v", err)
    }
    defer pw.Stop()

    browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(true),
    })
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

    // Navigate to the login page
    if _, err = page.Goto("https://www.saucedemo.com/v1/"); err != nil {
        log.Fatalf("Could not navigate to page: %v", err)
    }

    // Fill in login credentials
    if err = page.Fill("#user-name", "standard_user"); err != nil {
        log.Fatalf("Could not fill username: %v", err)
    }

    if err = page.Fill("#password", "secret_sauce"); err != nil {
        log.Fatalf("Could not fill password: %v", err)
    }

    // Click the login button
    if err = page.Click("#login-button"); err != nil {
        log.Fatalf("Could not click login button: %v", err)
    }

    // Wait for navigation and content to load
    // Adding a small delay to ensure all elements are loaded
    time.Sleep(2 * time.Second)

    // Verify we're logged in by checking for an element that appears after login
    // Properly handle both return values from WaitForSelector
    _, err = page.WaitForSelector(".inventory_list", playwright.PageWaitForSelectorOptions{
        State: playwright.WaitForSelectorStateVisible,
    })
    if err != nil {
        log.Fatalf("Login might have failed: %v", err)
    }

    // Take a screenshot of the products page
    screenshot, err := page.Screenshot(playwright.PageScreenshotOptions{
        Path:     playwright.String("saucedemo_logged_in.png"),
        //FullPage: playwright.Bool(true),
        //Quality:  playwright.Int(100),
		Timeout:  playwright.Float(5000),
    })
    if err != nil {
        log.Fatalf("Could not create screenshot: %v", err)
    }

    fmt.Println("Login successful!")
    fmt.Println("Screenshot saved as saucedemo_logged_in.png")
    fmt.Printf("Screenshot size: %d bytes\n", len(screenshot))

    // Optional: Take additional screenshots of specific elements
    // For example, screenshot just the inventory container
    element, err := page.QuerySelector(".inventory_list")
    if err == nil {
        _, err = element.Screenshot(playwright.ElementHandleScreenshotOptions{
            Path: playwright.String("inventory_list.png"),
        })
        if err == nil {
            fmt.Println("Inventory list screenshot saved as inventory_list.png")
        }
    }
}
