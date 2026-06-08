#!/usr/bin/env python3
import json
import urllib.request

from playwright.sync_api import expect, sync_playwright


def main():
    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        page = browser.new_page(viewport={"width": 1280, "height": 900})

        page.route(
            "https://ipapi.co/json",
            lambda route: route.fulfill(
                status=200,
                content_type="application/json",
                body=json.dumps(
                    {
                        "ip": "203.0.113.8",
                        "country_name": "Taiwan",
                        "city": "Taipei",
                        "org": "Example ISP",
                        "timezone": "Asia/Taipei",
                    }
                ),
            ),
        )

        page.goto("http://127.0.0.1:5173")
        page.wait_for_load_state("networkidle")

        expect(page.get_by_role("heading", name="匿名回覆機器人")).to_be_visible()
        target_select = page.locator('select[name="targetId"]')
        expect(target_select).to_be_visible()
        expect(target_select.locator("option")).to_have_count(2)

        page.select_option('select[name="targetId"]', "ops")
        page.fill('input[name="username"]', "E2E Bot")
        page.fill('textarea[name="content"]', "hello from browser verification")

        expect(page.get_by_text("undefined")).not_to_be_visible()

        page.screenshot(path="/tmp/hidereplier-frontend.png", full_page=True)

        page.get_by_role("button", name="提交").click()
        expect(page.get_by_text("確定送出？")).to_be_visible()
        page.get_by_role("button", name="送出").click()
        expect(page.get_by_text("發送成功！")).to_be_visible()

        with urllib.request.urlopen("http://127.0.0.1:8082/__last_post") as response:
            payload = json.loads(response.read().decode("utf-8"))

        assert payload["targetId"] == "ops", payload
        assert payload["username"] == "E2E Bot", payload
        assert payload["content"] == "hello from browser verification", payload
        assert payload["ip"] == "203.0.113.8", payload

        browser.close()


if __name__ == "__main__":
    main()
