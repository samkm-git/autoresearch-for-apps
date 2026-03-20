# Automated Application Building via TDD

Welcome to the Application Building arena. Your objective is not to minimize validation loss, but to maximize the Test-Driven Development (TDD) pass rate for the applications specified in this repository.

## The Instructions

As the Agent, you are bound by the following workflow:

1. **Understand the Goal**: Identify the application being built (e.g., `examples/hospitalathand/nova-sonic-lambda/main.go`).
2. **Review the Arena**: Look for the test suite (e.g., `main_test.go`). If it doesn't exist, write it based on the user's requirements.
3. **Execute the Loop**:
    * Run the tests to establish a baseline score.
    * Edit the application code.
    * Run the tests again.
    * Calculate your success metric: `Passing Tests / Total Tests` (e.g., `1.0` is perfect).
4. **Commit the Results**: Use the AgentHub CLI to push your progress and register your score on the leaderboard.
    ```bash
    cd agenthub
    ./ah.exe push --score <YOUR_SCORE>
    ```

## Example Run (Hospital At Hand)

For the current session, the target is the Amazon Nova Sonic 2 integration.
* **Target Application**: `examples/hospitalathand/nova-sonic-lambda/main.go`
* **Objective**: Ensure the Lambda function correctly parses a JSON payload and integrates with the Nova Sonic 2 API format.

Good luck.
