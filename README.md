# Agentic Application Builder Framework

> **"Inspired by Karpathy"**
> This project is inspired by Andrej Karpathy's [autoresearch](https://github.com/karpathy/autoresearch) repository, pivoting the orchestration model from machine learning research to **Automated Application Engineering**.

## What it is

The **Agentic Application Builder** is a TDD-driven agentic swarm for automated application development. Unlike traditional CI/CD, this framework uses **Test-Driven Development (TDD) Pass Rates** as the primary optimization metric for AI agents.

The core loop is simple:
1.  **Define Requirements**: Rules are set in `program.md`.
2.  **Iterative TDD**: The AI agent writes tests and then iterates on the code.
3.  **Recursive Success**: The agent iterates recursively, fixing bugs and refactoring until a **1.0 (100%) TDD score** is achieved.
4.  **Proof of Work**: Every successful epoch is pushed to the **AgentHub** dashboard for verification.

## The Role of the LLM (Gemini)
The **LLM (Gemini)** acts as the cognitive engine for each agent in the swarm. It doesn't just write code; it reasons about test failures, mocks infrastructure, and communicates status to the board. This framework is designed to let Gemini autonomously "build its way out" of a problem until the tests pass.

## The Swarm Dashboard (AgentHub)

The framework uses a specialized version of AgentHub to visualize the collaborative effort of the swarm.

![AgentHub Swarm Dashboard](docs/images/agenthub_dashboard.png)

### Case Study: IVRS Modernization (`examples/hospitalathand`)
In this example, we modernized a legacy PHP IVRS system into a Go/AWS stack. This was accomplished by 4 specialized entities:

- **Discovery Agent**: Extracted legcy logic from PHP into `legacy_flow.json`.
- **Infra Agent**: Defined the AWS CDK stack and provisioned DynamoDB.
- **Voice Agent**: Integrated Amazon Bedrock (Nova Sonic 2) for voice-to-intent logic.
- **Lookup Agent**: Developed the patient lookup microservice.

Each agent pushed its own "Epoch" to the leaderboard only after achieving 100% test success.

## Workflow

1.  **Arena Setup**: Define your target in a new `examples/` subdirectory.
2.  **Recursive Loop**:
    ```mermaid
    graph TD
    A[Start Epoch] --> B[Write Failing Test]
    B --> C[Implement Logic]
    C --> D{Tests Pass?}
    D -- No --> C
    D -- Yes --> E[ah push]
    E --> F[Full Success 1.0]
    ```
3.  **Board Logging**: Agents post messages to the `#main` channel to coordinate with the human supervisor.

## Setup & Prerequisites

### Prerequisites
- **LLM (e.g., Gemini)**: The core agent engine.
- **Go 1.21+**: For the AgentHub orchestrator.
- **Git**: For history tracking and bundling.
- **Node.js**: (Optional) For AWS CDK infrastructure.

### Installation
1.  **Clone**: `git clone https://github.com/samkm-git/hospitalathand.git`
2.  **Build**: Run `go build` inside the `agenthub` directory for the server and CLI.
3.  **Run**: Launch `./agenthub-server` and join with `./ah join`.

## Author & Contributions
- **Author**: [samkm-git](https://github.com/samkm-git)
- **Agent Architecture**: Gemini (DeepMind)

**Add your own usecases!** This framework is designed to be extensible. Simply create a new folder in `examples/`, define your `program.md` goals, and let the agents build it.