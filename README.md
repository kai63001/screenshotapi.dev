# ScreenshotAPI Business

**Capture any webpage on demand—sold as a subscription service**
SvelteKit dashboard · Go capture service · PocketBase · Stripe Billing

---

## 🚀 What it is

ScreenshotAPI Business lets companies program-matically grab fresh screenshots of websites, apps, or reports.
Customers create an account, choose a plan, and start hitting `/v1/screenshot?url=…`. All usage, credits, and invoices are handled for them.

### Core flow

1. **Signup / Login** — SvelteKit UI + PocketBase auth.
2. **Subscribe** — Stripe Checkout; credit quota linked to plan.
3. **Call the API** — Send a URL, get a PNG/JPEG/WEBP in milliseconds.
4. **Usage tracking** — Every successful render decrements the user’s balance; top-ups or plan upgrades via Stripe Customer Portal.

---

## 🧩 Tech stack

| Layer            | Tech                           | Role                                |
| ---------------- | ------------------------------ | ----------------------------------- |
| Frontend / Admin | **SvelteKit** + Tailwind       | Account, docs, usage graphs         |
| API Service      | **Go** (chromedp / rod)        | Headless Chrome renderer            |
| Auth / DB        | **PocketBase** (SQLite)        | Users, plans, usage, webhooks       |
| Payments         | **Stripe** (Checkout & Portal) | Recurring billing & metered credits |
| Queue (optional) | Redis or PocketBase realtime   | Burst handling                      |

---
