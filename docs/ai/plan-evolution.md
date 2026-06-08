# Plan Evolution

## 2026-05-24 · What I thought the project was

**What we thought we'd do:**
Initially read the spec and assumed Stage 1 was mainly about scaffolding — a Go HTTP server that accepted requests and returned a placeholder response, with the real execution logic coming later.

**What we actually did:**
After reading more carefully and looking at the evaluation criteria, we realised the expectation was a working execution pipeline: receive code, write it to disk, run it, return real output.

**Why it changed:**
The spec was clearer on re-reading than on first pass. The phrase "working prototype" in the evaluation section was the key signal. Scaffolding with stubbed responses wouldn't satisfy that.


## 2026-05-27 · Realising documentation was a scored deliverable, not an afterthought

**What we thought we'd do:**
Write docs at the end once the code was done.

**What we actually did:**
Started writing architecture, API and security docs in parallel with implementation rather than after.

**Why it changed:**
The maintainer's discussion post (#5) made it explicit that SDLC discipline — including documentation — is part of the review criteria, not separate from it. Leaving docs to the last day would mean rushed, shallow coverage that a reviewer would notice.


## 2026-06-01 · Narrowing scope before the deadline

**What we thought we'd do:**
Add at least one more language beyond Python before Stage 1 submission.

**What we actually did:**
Kept Python only and used the time to make the documentation complete.

**Why it changed:**
A second language would have required testing and doc updates to stay consistent. The risk of introducing a half-working feature outweighed the benefit. A clean single-language implementation with accurate docs felt more aligned with what Stage 1 actually rewards.

## 2026-06-06 · Moving from simple execution to interactive execution

**What we thought we'd do:**
Only execute Python programs that print output.

**What we actually did:**
Added stdin support so programs using input() could run correctly.

**Why it changed:**
Most programming problems depend on runtime input. Without stdin support the service behaved more like a script runner than a coding platform.

## 2026-06-07 · Discovering that timeout implementation was incomplete

**What we thought we'd done:**
Timeout support was complete after adding context.WithTimeout.

**What we actually discovered:**
The timeout context was never attached to the executed process.

**Why it changed:**
Testing an infinite loop revealed that the process was not being terminated.

The implementation was corrected using exec.CommandContext.

## 2026-06-08 · Moving from a code runner toward a judge engine

**What we thought we'd do:**
Return stdout and stderr only.

**What we actually did:**
Added execution status classification and stdin support.

**Why it changed:**
These features are required for real programming problem execution and prepare the service for future multi-test judging.
