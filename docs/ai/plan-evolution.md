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