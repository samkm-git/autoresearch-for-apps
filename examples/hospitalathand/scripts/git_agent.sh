#!/bin/bash

# Git Swarm Agent
# Logic: Monitors the hospitalathand workspace and pushes frequent updates to GitHub

REMOTE_URL="https://github.com/samkm-git/hospitalathand.git"
BRANCH="main"

echo "----------------------------------------------------"
echo "🤖 Git Swarm Agent Started"
echo "Target: $REMOTE_URL"
echo "----------------------------------------------------"

# Ensure origin is set
git remote add origin "$REMOTE_URL" 2>/dev/null || git remote set-url origin "$REMOTE_URL"

while true; do
    # 1. Check for changes
    if [[ -n $(git status -s) ]]; then
        echo "发现变更 (Changes detected). Preparing to sync..."
        
        # 2. Add and Commit
        git add .
        COMMIT_MSG="Auto-commit from Agent Swarm at $(date +'%Y-%m-%d %H:%M:%S')"
        git commit -m "$COMMIT_MSG"
        
        # 3. Push
        echo "🚀 Pushing to GitHub..."
        git push origin "$BRANCH"
        
        if [ $? -eq 0 ]; then
            echo "✅ Sync successful."
            # Post to AgentHub Board
            ../agenthub/ah.exe post main "Git Agent: Successfully pushed latest workspace changes to GitHub ($BRANCH)."
        else
            echo "⚠️ Push failed. Waiting for repository creation or connectivity..."
        fi
    else
        echo "No changes in workspace. Sleeping..."
    fi

    # Wait for 5 minutes before next check
    sleep 300
done
