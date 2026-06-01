# Architectural Decision Records

## Workspace-based execution

**Context:**
Submitted code needs to be written to disk before execution. The question was whether to use a shared directory or isolate each request.

**Options considered:**

1. Use a single shared execution directory for all requests.
2. Create a fresh temporary workspace per request and delete it after.

**Decision:**
Create a temporary workspace for each execution request.

**Rationale:**
A shared directory creates race conditions under concurrent requests and makes cleanup ambiguous. Temporary workspaces are cleaned up deterministically after each run. This also prepares the architecture for future sandbox isolation — if we introduce nsjail or similar, the per-request workspace is already the right unit to jail.


## Python-only scope for Stage 1

**Context:**
The spec mentions multi-language support as a goal. The question was whether to add a second language before the Stage 1 deadline.

**Options considered:**

1. Add JavaScript or another language before submission.
2. Keep Python only and ensure it works correctly end to end.

**Decision:**
Python only for Stage 1.

**Rationale:**
Adding a second language without adequate testing coverage and documentation updates would introduce inconsistency between the code and the docs. A reviewer would notice. A single language done well is a stronger signal than two languages done partially.