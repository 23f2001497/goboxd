# Issues Encountered

## 2026-06-01 · Python execution failing on Windows

**What we were trying to do:**
Run submitted Python code using exec.Command as part of the execution pipeline.

**What went wrong:**
The implementation called python3, which is the standard binary name on Linux and macOS. The Windows environment used during development only exposed python, so the command failed to find the interpreter.

**How we resolved it:**
Identified the binary name difference and adjusted the execution config. Verified output against the expected response format using the available Python installation.

**What we learned:**
Don't assume interpreter binary names are consistent across platforms. This is worth noting in the architecture doc as something that will need a proper resolution path before deployment on Linux infrastructure.


## 2026-06-01 · Test output mismatch caused by line endings

**What we were trying to do:**
Validate execution output in automated tests by comparing against expected strings.

**What went wrong:**
Python on Windows produces CRLF line endings in some output contexts. The test expected Unix LF endings, so the comparison failed even when the output was semantically correct.

**How we resolved it:**
Identified the platform difference. Documented that output normalisation (stripping or standardising line endings) will be needed for cross-platform test reliability.

**What we learned:**
String equality on output is fragile across platforms. Normalise before comparing, or test for content rather than exact byte matches.

## 2026-06-06 · Timeout logic existed but was not enforced

**What we were trying to do:**
Prevent user programs from running forever.

**What went wrong:**
The implementation created a context using context.WithTimeout(), but execution still used exec.Command() instead of exec.CommandContext().

This meant the timeout context existed but was never attached to the Python process.

Programs such as:

```python
while True:
    pass
```

continued running indefinitely.

**How we resolved it:**
Replaced exec.Command() with exec.CommandContext() and attached the timeout context directly to the executed process.

Verified by running an infinite loop and confirming the API returned:

```json
{
  "error": "time_exceeded"
}
```

**What we learned:**
Creating a timeout context is not enough. The executed process must explicitly use that context.

## 2026-06-07 · Cross-platform test failure caused by Windows line endings

**What we were trying to do:**
Validate execution output in automated tests.

**What went wrong:**
Python on Windows produced CRLF line endings (\r\n) while the test expected LF (\n).

The execution result was correct but the test still failed.

**How we resolved it:**
Adjusted the test comparison to account for platform-specific line endings.

The production code was left unchanged because the issue existed only in test assertions.

**What we learned:**
Tests should validate behavior, not platform-specific formatting details.

## 2026-06-08 · Timeout implementation existed but was ineffective

**What we were trying to do:**
Terminate long-running programs after a fixed timeout.

**What went wrong:**
The implementation created a timeout context but executed programs using exec.Command() rather than exec.CommandContext().

As a result, the timeout was never attached to the process.

**How we resolved it:**
Replaced exec.Command() with exec.CommandContext().

**What we learned:**
Timeout contexts must be attached directly to the process being executed.
