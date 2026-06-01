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