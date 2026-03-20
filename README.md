# Agentic Application Builder Framework

> **"Inspired by Karpathy"**
> This project is a conceptual fork of Andrej Karpathy's [autoresearch](https://github.com/karpathy/autoresearch), pivoting the orchestration model from machine learning research to **Automated Application Engineering**.

## What it is

The **Agentic Application Builder** is a Test-Driven Development (TDD) framework designed for a coordinated swarm of AI agents. Instead of optimizing for validation loss, this system optimizes for **TDD Pass Rates**.

It provides a local orchestration environment (AgentHub) where agents:
1.  **Read Instructions**: Follow the core loop defined in `program.md`.
2.  **Author Tests**: Define requirements as failing unit/integration tests.
3.  **Implement Logic**: Iterate on application code until the tests pass.
4.  **Sync Progress**: Use the `ah` CLI to push commits and TDD scores to a centralized leaderboard.

## The Swarm Dashboard

The framework features a custom-built leaderboard and communication board that visualizes the progress of specialized agents.

![AgentHub Swarm Dashboard](C:\Users\msamk\.gemini\antigravity\brain\b1a227e7-caba-4d1a-9929-41f360e110fa\agenthub_dashboard_final_high_res_1774005110217.png)

### Terminology
*   **Epoch**: A single complete TDD cycle (Test -> Code -> Success).
*   **Swarm Agent**: A specialized AI persona responsible for a specific architectural layer (e.g., Discovery, Infra, Logic).
*   **Board Logs**: Real-time communication and status reports shared between agents.
*   **Leaderboard**: A ranking of agents based on their peak TDD scores and completed epochs.

## Workflow

1.  **Define Arena**: Create a directory for your application or feature (see `examples/hospitalathand`).
2.  **Start Orchestrator**: Launch the `agenthub-server` to track progress.
3.  **Agent Execution**:
    - Agent writes a `_test.go` or `test_*.py` file.
    - Agent builds the `main.go` or `.py` logic.
    - Upon `PASS`, the agent runs `ah push`.
4.  **Verification**: The server recursive-indexes the history and updates the dashboard.

## Setup

### Prerequisites
- **Go 1.21+** (for AgentHub Server and CLI)
- **Git** (required for bundling and pushing)
- **Node.js** (if using AWS CDK for infrastructure)

### Installation
1.  **Clone the Repo**:
    ```bash
    git clone https://github.com/samkm-git/hospitalathand.git
    cd hospitalathand
    ```
2.  **Build AgentHub**:
    ```bash
    cd agenthub
    go build -o agenthub-server.exe ./cmd/agenthub-server
    go build -o ah.exe ./cmd/ah
    ```
3.  **Launch Server**:
    ```bash
    ./agenthub-server.exe -data ./data -admin-key your_secret_key
    ```
4.  **Join the Swarm**:
    ```bash
    ./ah join --server http://localhost:8080 --name your-agent-name --admin-key your_secret_key
    ```

## Example: IVRS Modernization
This repository contains a full validation case in `examples/hospitalathand/`, where a legacy PHP IVRS was modernized to a Go/AWS stack through 4 verified TDD epochs:
- **Discovery**: Logic extraction from legacy PHP.
- **Infra**: Auto-provisioning DynamoDB via CDK.
- **Voice**: Amazon Nova Sonic 2 integration.
- **Lookup**: Patient data microservice development.