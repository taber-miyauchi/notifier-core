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

Open this repo in Sourcegraph and try the following:

### 1. Find Implementations (cross-repo interface)

Discover all implementations of an interface across separate repositories.

- In `notifier.go`, click on `Send` method (line 8) → **Find Implementations**
- → Highlights `Send` function in `notifier-email/email.go` (line 28)

**Benefit:** Instantly see which repos provide concrete implementations of your interface—critical for understanding the full scope of a plugin or adapter pattern across a microservices architecture.

### 2. Find References (cross-repo struct)

Locate all usages of a shared type across multiple repositories.

- In `message.go`, click on `Message` struct (line 13) → **Find References**
- → Highlights `Message` (line 21) in `notifier-core/message.go`
- → Highlights `Message` (line 8) in `notifier-core/notifier.go`
- → Highlights `core.Message` (line 28) in `notifier-email/email.go`

**Benefit:** Understand the impact radius of changing a shared data structure—see every repo and every call site that would be affected before making breaking changes.

### 3. Find References (cross-repo function)

Track usage of exported functions across repository boundaries.

- In `message.go`, click on `NewMessage` func (line 21) → **Find References**
- → Highlights `NewMessage` (line 38) in `notifier-service/main.go`

**Benefit:** Identify all consumers of your API functions across the organization—essential for deprecation planning and understanding adoption.
