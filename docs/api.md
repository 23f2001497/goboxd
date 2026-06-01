# API Documentation

## Overview

GoBoxD exposes HTTP endpoints for health monitoring and source code execution.

Current Stage 1 implementation supports Python code execution.

---

## GET /healthz

### Description

Health check endpoint used to verify that the service is running.

### Response

```json
{
  "status": "ok"
}
```

---

## POST /run

### Description

Accepts source code, executes it, and returns execution results.

### Request

```json
{
  "language": "python",
  "source": "print(\"hello\")"
}
```

### Response

```json
{
  "stdout": "hello\n"
}
```

### Supported Language

* Python

### Validation

The request must contain:

* language
* source

Invalid requests return HTTP 400.
