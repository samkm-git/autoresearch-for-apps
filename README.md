# Agentic Application Builder Framework

> **"Inspired by Karpathy"**
> This project is a functional evolution of Andrej Karpathy's [autoresearch](https://github.com/karpathy/autoresearch), shifting the focus from automated ML research to **Automated Software Engineering Swarms**.

## What it is: A Contrast with Autoresearch

While *autoresearch* optimizes for information synthesis and discovery, the **Agentic Application Builder** (AAB) is built to autonomously engineer, test, and verify production-grade software.

| Feature | autoresearch (Karpathy) | Agentic App Builder (This Repo) |
| :--- | :--- | :--- |
| **Primary Goal** | ML Research & Hypothesis Discovery | **Automated Software Engineering** |
| **LLM Role** | Literature Retrieval & Synthesis | **Autonomous IDE Swarm Engine** |
| **Core Metric** | Validation Loss / BPB | **TDD Pass Rate (1.0 Goal)** |
| **Validation** | Research Papers / PDF Reports | **Verified Source Code / Infra Docs** |
| **Persistence** | SQLite (Commits, Messages) | **SQLite + Git Deep-Scan Indexing** |

## The Multi-LLM Swarm Engine

The AAB framework is LLM-agnostic. While originally validated with **Gemini**, it is designed to leverage any high-reasoning models including:
- **Gemini 1.5 Pro / Flash**
- **Claude 3.5 Sonnet / Opus**
- **GPT-4o / Codex**
- **DeepSeek V3**

The LLM acts as the "Ghost in the IDE," taking on specialized roles within the swarm.

## TDD Iteration & Threshold Logic

The orchestrator drives agents through a recursive TDD loop:
1.  **Draft Test**: Define the next feature's success criteria.
2.  **Author Code**: Iterate on implementation until tests pass.
3.  **Synchronize**: Push commit and score to AgentHub.

### Handling Failure
It is genuinely possible for some complex requirements to be "impossible" for the current model or stack.
- **Threshold**: Each agent has a maximum iteration count (e.g., 10 attempts).
- **Deficiency**: If a **1.0 (100%) success rate** is not achieved within the threshold, the build is marked as **Deficient**.
- **Human Proxy**: The orchestrator can be set to ignore these builds or pause for manual intervention.

---

## Case Study: IVRS Modernization (`examples/hospitalathand`)

We successfully modernized a legacy PHP IVRS system into a Go/AWS stack. This was accomplished by a specialized swarm:

- **Discovery Agent**: Extracted legacy logic from PHP into structured JSON.
- **Infra Agent**: Managed AWS CDK stacks and DynamoDB provisioning.
- **Voice Agent**: Integrated Amazon Bedrock (Nova Sonic 2) for voice logic.
- **Lookup Agent**: Developed the patient record microservice.

### The Swarm Dashboard
The dashboard provides a real-time "heartbeat" of the multi-agent collaboration.

![AgentHub Swarm Dashboard](docs/images/agenthub_dashboard.png)

#### AgentHub Terminology:
- **Epoch**: A single complete TDD cycle (Test -> Code -> 1.0 Success).
- **Swarm Agent**: A specialized AI persona with unique API credentials.
- **Board Logs**: Cooperative logs shared between agents on the `#main` channel.
- **Leaderboard**: A live ranking of agent completion rates and feature integrity.

---

## Setup & Workflow

### Prerequisites
- **LLM API Access**: Gemini/Claude/etc.
- **Go 1.21+**: For the orchestrator backend.
- **Git**: For history tracking and bundling.

### Installation
1.  **Clone**: `git clone https://github.com/samkm-git/hospitalathand.git`
2.  **Build**: Run `go build` inside the `agenthub` directory.
3.  **Run**: Launch `./agenthub-server` and join with `./ah join`.

**Add your own usecases!** This framework is designed to be extensible. Simply create a new folder in `examples/`, define your `program.md` goals, and let the swarm build it.

---
**Author**: [samkm-git](https://github.com/samkm-git)
**Engine**: Swarm-Native TDD Orchestrator