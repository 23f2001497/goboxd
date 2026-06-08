# AI Prompt Log

## 2026-05-24 · Understanding what Stage 1 actually evaluates

**Prompt:**
I read the GoBoxD spec and I'm not sure what "working prototype" means in the context of Stage 1. Does the evaluation care more about features like sandboxing and multi-language, or about having a clean working pipeline even if limited? What should I prioritise?

**Response summary:**
AI said Stage 1 rewards a functioning end-to-end pipeline over breadth of features. It suggested getting code submission, execution and response working first, then layering documentation and test coverage. It also mentioned that SDLC signals — commits, branching, issue tracking — are part of how reviewers assess engineering discipline.

**What we used / didn't use:**
Used the recommendation to narrow scope to a working Python execution pipeline first. Ignored the suggestion to add multi-language support in Stage 1 — that felt premature and the spec didn't require it.


## 2026-05-25 · Structuring the architecture documentation

**Prompt:**
I need to write an architecture doc for a code execution service. The service receives source code over HTTP, writes it to a temporary workspace, executes it, and returns stdout/stderr. What components should the doc cover and how should they be organised so a new engineer can follow request flow?

**Response summary:**
AI suggested organising the doc around: request entry point, input validation, workspace lifecycle, language execution, response formatting. It also suggested including a sequence diagram or numbered flow to make execution order explicit.

**What we used / didn't use:**
Used the component breakdown as the structure for our execution flow doc. Didn't use the sequence diagram suggestion — described the flow in numbered prose instead since our implementation is straightforward enough that a diagram would add little.


## 2026-05-30 · Writing the security roadmap — understanding isolation concepts

**Prompt:**
I'm documenting planned security improvements for a sandboxed code execution service. I understand we're not implementing these in Stage 1 but I want the roadmap to be technically accurate. Can you explain how Linux namespaces and cgroups work and how nsjail uses them, so I can describe what isolation would actually involve?

**Response summary:**
AI explained that namespaces isolate process view of resources (PID, network, filesystem, user) while cgroups enforce resource consumption limits (CPU, memory, pids). nsjail combines both to create a confined jail for process execution. It also mentioned that dropping capabilities is a separate concern from namespace isolation.

**What we used / didn't use:**
Used the namespace/cgroup distinction when writing the security roadmap section. Added capability dropping as a separate line item based on that — hadn't considered it before. Did not use the nsjail config examples AI generated since we aren't implementing it in Stage 1 and including config without implementation would be misleading.


## 2026-06-01 · Documenting the benchmark roadmap

**Prompt:**
What metrics matter most for a code execution service that runs user-submitted code? I want to write a benchmark roadmap that covers the right things — not just "measure latency" but specific things worth tracking given the nature of the service.

**Response summary:**
AI suggested: execution latency per language, workspace creation and cleanup time separately from execution time, memory usage per run, concurrency under load (what happens at N simultaneous requests), and cold start vs warm path differences. It also mentioned tracking error rates by failure category.

**What we used / didn't use:**
Used the idea of separating workspace overhead from actual execution time — that's a meaningful distinction for our architecture. Used the concurrency angle since we don't have a request slot limiter yet and that's worth flagging. Didn't use the cold start framing since we don't have any caching layer that would make that distinction meaningful.

## 2026-06-06 · Verifying timeout implementation

**Prompt:**
I added context.WithTimeout to my executor. Does this guarantee that infinite loops are terminated?

**Response summary:**
AI explained that exec.CommandContext must be used. Creating a timeout context alone is insufficient.

**What we used / didn't use:**
Used the recommendation to switch to exec.CommandContext.
Ignored alternative approaches involving manual process termination because they added unnecessary complexity.

## 2026-06-07 · Designing stdin support

**Prompt:**
How should I pass user input into executed Python programs while keeping the API simple?

**Response summary:**
AI suggested adding a stdin field to the request model and attaching it to the process using strings.NewReader.

**What we used / didn't use:**
Used the stdin request field and strings.NewReader approach.
Did not implement interactive streaming input because it was unnecessary for Stage 1.
