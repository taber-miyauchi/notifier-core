# notifier-core

## Project Structure

```
├── notifier-core/          ← The foundation (you are here)
│   ├── go.mod
│   ├── notifier.go         ← Notifier interface
│   ├── message.go          ← Message struct, Priority enum
│   └── README.md
│
├── notifier-email/         ← Implementation
│   ├── go.mod              ← depends on notifier-core
│   ├── email.go            ← EmailNotifier implements Notifier
│   └── README.md
│
└── notifier-service/       ← Consumer/API
    ├── go.mod              ← depends on both
    ├── main.go             ← HTTP server using both packages
    └── README.md
```

## Overview

Core library defining the `Notifier` interface and shared types for the notification system.

## Types

- **`Notifier`** - Interface that all notification implementations must satisfy
- **`Message`** - The notification payload (recipient, subject, body, priority)
- **`Priority`** - Enum for Low, Normal, High urgency

## Usage

This package is imported by:
- `notifier-email` - Implements the `Notifier` interface
- `notifier-service` - Uses the interface and types to send notifications

## Testing Precise Code Navigation

Open this repo in Sourcegraph and try:

### 1. Find Implementations (cross-repo)

- In `notifier.go`, click on `Notifier` interface (line 8) → **Find Implementations**
- → Jumps to `EmailNotifier.Send()` in `notifier-email/email.go`

### 2. Find References (cross-repo)

- In `message.go`, click on `Message` struct (line 14) → **Find References**
- → Shows usages across all 3 repos: `notifier-core`, `notifier-email`, and `notifier-service`
