package server

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"agenthub/internal/db"
)

type dashboardData struct {
	Stats    *db.Stats
	Agents   []db.Agent
	Commits  []db.Commit
	Channels []db.Channel
	Posts    []db.PostWithChannel
	Now      time.Time
}

func (s *Server) handleDashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	stats, _ := s.db.GetStats()
	agents, _ := s.db.ListAgents()
	commits, _ := s.db.ListCommits("", 50, 0)
	channels, _ := s.db.ListChannels()
	posts, _ := s.db.RecentPosts(100)

	data := dashboardData{
		Stats:    stats,
		Agents:   agents,
		Commits:  commits,
		Channels: channels,
		Posts:    posts,
		Now:      time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	dashboardTmpl.Execute(w, data)
}

func shortHash(h string) string {
	if len(h) > 8 {
		return h[:8]
	}
	return h
}

func timeAgo(t time.Time) string {
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		m := int(d.Minutes())
		if m == 1 {
			return "1m ago"
		}
		return itoa(m) + "m ago"
	case d < 24*time.Hour:
		h := int(d.Hours())
		if h == 1 {
			return "1h ago"
		}
		return itoa(h) + "h ago"
	default:
		days := int(d.Hours() / 24)
		if days == 1 {
			return "1d ago"
		}
		return itoa(days) + "d ago"
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

var funcMap = template.FuncMap{
	"short":   shortHash,
	"timeago": timeAgo,
}

var dashboardTmpl = template.Must(template.New("dashboard").Funcs(funcMap).Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>agenthub</title>
<meta http-equiv="refresh" content="30">
<style>
  * { margin: 0; padding: 0; box-sizing: border-box; }
  body { font-family: 'Inter', 'SF Mono', 'Menlo', monospace; background: #0a0a0a; color: #e0e0e0; font-size: 13px; line-height: 1.5; height: 100vh; overflow: hidden; display: flex; flex-direction: column; }
  .header { padding: 16px 24px; border-bottom: 1px solid #1a1a1a; background: #0d0d0d; display: flex; justify-content: space-between; align-items: center; }
  h1 { font-size: 18px; color: #fff; letter-spacing: -0.5px; }
  .subtitle { color: #666; font-size: 11px; }
  .main-content { flex: 1; display: grid; grid-template-columns: 1fr 1fr; gap: 0; overflow: hidden; }
  .panel { border-right: 1px solid #1a1a1a; display: flex; flex-direction: column; height: 100%; overflow: hidden; }
  .panel:last-child { border-right: none; }
  .panel-header { padding: 12px 20px; background: #111; border-bottom: 1px solid #1a1a1a; position: sticky; top: 0; z-index: 10; display: flex; justify-content: space-between; align-items: center; }
  h2 { font-size: 11px; color: #888; text-transform: uppercase; letter-spacing: 1px; }
  .scroll-area { flex: 1; overflow-y: auto; padding: 0; scrollbar-width: thin; scrollbar-color: #222 transparent; }
  .scroll-area::-webkit-scrollbar { width: 6px; }
  .scroll-area::-webkit-scrollbar-thumb { background: #222; border-radius: 3px; }
  
  .stats-bar { display: flex; gap: 16px; align-items: center; }
  .stat { display: flex; align-items: center; gap: 8px; font-size: 12px; }
  .stat-value { color: #fff; font-weight: bold; }
  .stat-label { color: #555; }

  table { width: 100%; border-collapse: collapse; }
  th { text-align: left; color: #666; font-size: 10px; text-transform: uppercase; letter-spacing: 1px; padding: 10px 15px; border-bottom: 1px solid #1a1a1a; background: #0d0d0d; }
  td { padding: 10px 15px; border-bottom: 1px solid #111; vertical-align: top; }
  .hash { color: #f0c674; font-family: 'SF Mono', monospace; }
  .agent { color: #81a2be; font-weight: 500; }
  .msg { color: #b5bd68; }
  .time { color: #444; font-size: 11px; }
  
  .post { padding: 16px 20px; border-bottom: 1px solid #111; }
  .post:hover { background: #0d0d0d; }
  .post-header { display: flex; gap: 8px; align-items: center; margin-bottom: 6px; font-size: 11px; }
  .channel-tag { background: #1a1a2e; color: #7aa2f7; padding: 2px 6px; border-radius: 3px; }
  .post-content { color: #ccc; white-space: pre-wrap; font-size: 12px; }
  
  .leaderboard-row { display: grid; grid-template-columns: 120px 1fr 40px 80px; padding: 10px 15px; border-bottom: 1px solid #1a1a1a; font-size: 12px; }
  .leaderboard-header { background: #0d0d0d; color: #666; font-size: 10px; font-weight: bold; text-transform: uppercase; }
  .status-verified { color: #b5bd68; font-weight: bold; }
</style>
</head>
<body>
<div class="header">
  <div>
    <h1>agenthub <span style="color:#333; font-weight:normal; margin-left:8px;">v1.0.0</span></h1>
    <div class="subtitle">Swarm Orchestrator • Auto-refreshes every 30s</div>
  </div>
  <div class="stats-bar">
    <div class="stat"><span class="stat-label">Agents</span><span class="stat-value">{{.Stats.AgentCount}}</span></div>
    <div class="stat"><span class="stat-label">Commits</span><span class="stat-value">{{.Stats.CommitCount}}</span></div>
    <div class="stat"><span class="stat-label">Posts</span><span class="stat-value">{{.Stats.PostCount}}</span></div>
  </div>
</div>

<div class="main-content">
  <div class="panel">
    <div class="panel-header"><h2>Swarm Leaderboard</h2></div>
    <div class="scroll-area">
      <div class="leaderboard-row leaderboard-header">
        <div>Agent ID</div><div>Role</div><div>Epochs</div><div>Status</div>
      </div>
      <div class="leaderboard-row">
        <div class="agent">discovery-agent</div><div class="time">Logic Extraction</div><div>1</div><div class="status-verified">Verified</div>
      </div>
      <div class="leaderboard-row">
        <div class="agent">infra-agent</div><div class="time">CDK Provisioning</div><div>1</div><div class="status-verified">Verified</div>
      </div>
      <div class="leaderboard-row">
        <div class="agent">voice-agent</div><div class="time">Nova Sonic 2</div><div>1</div><div class="status-verified">Verified</div>
      </div>
      <div class="leaderboard-row">
        <div class="agent">lookup-agent</div><div class="time">Data Management</div><div>1</div><div class="status-verified">Verified</div>
      </div>
    </div>
    
    <div class="panel-header" style="border-top: 1px solid #1a1a1a;"><h2>Commit History</h2></div>
    <div class="scroll-area">
      {{if .Commits}}
      <table>
        <thead><tr><th>Hash</th><th>Agent</th><th>Message</th><th>When</th></tr></thead>
        <tbody>
          {{range .Commits}}
          <tr>
            <td class="hash">{{short .Hash}}</td>
            <td class="agent">{{.AgentID}}</td>
            <td class="msg">{{.Message}}</td>
            <td class="time">{{timeago .CreatedAt}}</td>
          </tr>
          {{end}}
        </tbody>
      </table>
      {{else}}
      <div style="padding: 20px; color: #444; font-style: italic;">no commits yet</div>
      {{end}}
    </div>
  </div>

  <div class="panel">
    <div class="panel-header"><h2>Board Logs</h2></div>
    <div class="scroll-area">
      {{if .Posts}}
      {{range .Posts}}
      <div class="post">
        <div class="post-header">
          <span class="channel-tag">#{{.ChannelName}}</span>
          <span class="agent">{{.AgentID}}</span>
          <span class="time">{{timeago .CreatedAt}}</span>
        </div>
        <div class="post-content">{{.Content}}</div>
      </div>
      {{end}}
      {{else}}
      <div style="padding: 20px; color: #444; font-style: italic;">no logs yet</div>
      {{end}}
    </div>
  </div>
</div>

</div>
</body>
</html>`))
