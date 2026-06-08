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

## Timeout enforcement using context-aware process execution

**Context:**
User-submitted programs may contain infinite loops.

**Options considered:**

1. Kill processes manually using OS signals.
2. Use context.WithTimeout with exec.CommandContext.

**Decision:**
Use context.WithTimeout combined with exec.CommandContext.

**Rationale:**
This approach integrates directly with Go's process execution model and keeps timeout logic simple and testable.

## Passing stdin through strings.NewReader

**Context:**
Programs using input() could not receive user input.

**Options considered:**

1. Hardcode input values.
2. Pass request stdin directly into the process.

**Decision:**
Attach stdin using strings.NewReader.

**Rationale:**
This allows arbitrary user input while keeping the execution pipeline simple.

## Execution Status Classification

**Context:**
The API originally returned stdout and error fields only.

**Options considered:**

1. Infer execution state from stdout and error fields.
2. Add an explicit execution status field.

**Decision:**
Add a status field with values accepted, runtime_error and time_exceeded.

**Rationale:**
This makes execution outcomes explicit and aligns better with online judge systems.
