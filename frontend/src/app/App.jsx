import { useEffect, useMemo, useRef, useState } from "react";
import { Room, RoomEvent } from "livekit-client";

const API_BASE = import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/api/v1";
const TOKEN_STORAGE_KEY = "oris_access_token";

async function api(path, method = "GET", body, token) {
  const res = await fetch(`${API_BASE}${path}`, {
    method,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: body ? JSON.stringify(body) : undefined,
  });
  const data = await res.json().catch(() => ({}));
  if (!res.ok) throw new Error(data.error || "Request failed");
  return data;
}

function roleRank(role) {
  const order = { member: 1, moderator: 2, admin: 3, owner: 4 };
  return order[role] || 0;
}

function formatTime(value) {
  if (!value) return "--:--";
  const d = new Date(value);
  if (Number.isNaN(d.getTime())) return "--:--";
  return d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
}

function shortId(value) {
  if (!value) return "unknown";
  if (value.length <= 8) return value;
  return `${value.slice(0, 4)}...${value.slice(-4)}`;
}

export default function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [accessToken, setAccessToken] = useState(() => localStorage.getItem(TOKEN_STORAGE_KEY) || "");
  const [currentUser, setCurrentUser] = useState(null);

  const [communityName, setCommunityName] = useState("");
  const [channelName, setChannelName] = useState("");
  const [channelType, setChannelType] = useState("text");
  const [messageText, setMessageText] = useState("");

  const [communities, setCommunities] = useState([]);
  const [channels, setChannels] = useState([]);
  const [members, setMembers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [onlineUsers, setOnlineUsers] = useState([]);

  const [selectedCommunity, setSelectedCommunity] = useState(null);
  const [selectedChannel, setSelectedChannel] = useState(null);

  const [voiceToken, setVoiceToken] = useState("");
  const [voiceState, setVoiceState] = useState("idle");
  const [voiceParticipants, setVoiceParticipants] = useState(0);
  const [micEnabled, setMicEnabled] = useState(false);

  const [logs, setLogs] = useState([]);
  const [inviteCodeCreated, setInviteCodeCreated] = useState("");
  const [joinInviteCode, setJoinInviteCode] = useState("");

  const roomRef = useRef(null);
  const messagesRef = useRef(null);

  const loggedIn = useMemo(() => Boolean(accessToken), [accessToken]);

  const myMembership = useMemo(() => {
    if (!currentUser) return null;
    return members.find((m) => m.userId === currentUser.id) || null;
  }, [members, currentUser]);

  const usernameByUserId = useMemo(() => {
    const map = {};
    for (const m of members) {
      map[m.userId] = m.username;
    }
    if (currentUser?.id && currentUser?.username) {
      map[currentUser.id] = currentUser.username;
    }
    return map;
  }, [members, currentUser]);

  const orderedMessages = useMemo(() => {
    const items = [...messages];
    items.sort((a, b) => new Date(a.createdAt).getTime() - new Date(b.createdAt).getTime());
    return items;
  }, [messages]);

  const canModerate = useMemo(() => roleRank(myMembership?.role) >= roleRank("moderator"), [myMembership]);

  useEffect(() => {
    if (accessToken) {
      localStorage.setItem(TOKEN_STORAGE_KEY, accessToken);
    } else {
      localStorage.removeItem(TOKEN_STORAGE_KEY);
    }
  }, [accessToken]);

  useEffect(() => {
    if (!loggedIn) return;
    loadMe();
    loadCommunities();
  }, [loggedIn]);

  useEffect(() => {
    if (!selectedCommunity) return;
    loadChannels(selectedCommunity.id);
    loadMembers(selectedCommunity.id);
  }, [selectedCommunity]);

  useEffect(() => {
    if (roomRef.current && (!selectedChannel || roomRef.current.name !== selectedChannel.id)) {
      leaveVoice();
    }
    if (!selectedChannel) return;
    loadMessages(selectedChannel.id);
    loadPresence(selectedChannel.id);
  }, [selectedChannel]);

  useEffect(() => {
    if (!accessToken || !selectedChannel) return;
    const wsBase = API_BASE.replace("http://", "ws://").replace("https://", "wss://").replace("/api/v1", "");
    const ws = new WebSocket(`${wsBase}/api/v1/ws/${selectedChannel.id}?token=${encodeURIComponent(accessToken)}`);

    ws.onmessage = (event) => {
      try {
        const payload = JSON.parse(event.data);
        if (payload?.type === "message.created" && payload?.data) {
          setMessages((prev) => [...prev, payload.data]);
        }
        if (payload?.type === "presence.updated") {
          setOnlineUsers(payload?.data?.users || []);
        }
      } catch {
        // ignore malformed payload
      }
    };

    ws.onopen = () => pushLog("Realtime connected");
    ws.onclose = () => pushLog("Realtime disconnected");
    return () => ws.close();
  }, [selectedChannel, accessToken]);

  useEffect(() => {
    if (!messagesRef.current) return;
    messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
  }, [orderedMessages.length, selectedChannel?.id]);

  useEffect(() => {
    return () => {
      if (roomRef.current) {
        roomRef.current.disconnect();
        roomRef.current = null;
      }
    };
  }, []);

  function pushLog(line) {
    setLogs((prev) => [line, ...prev].slice(0, 10));
  }

  function updateVoiceParticipants(room) {
    const count = room ? room.remoteParticipants.size + (room.state === "connected" ? 1 : 0) : 0;
    setVoiceParticipants(count);
  }

  async function loadMe() {
    try {
      const data = await api("/me", "GET", null, accessToken);
      setCurrentUser(data.user || null);
    } catch (e) {
      setAccessToken("");
      setCurrentUser(null);
      pushLog(e.message);
    }
  }

  async function register() {
    try {
      const data = await api("/auth/register", "POST", { username, password });
      setAccessToken(data.accessToken);
      pushLog("Account created");
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function login() {
    try {
      const data = await api("/auth/login", "POST", { username, password });
      setAccessToken(data.accessToken);
      pushLog("Logged in");
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function loadCommunities() {
    try {
      const data = await api("/communities", "GET", null, accessToken);
      const items = data.items || [];
      setCommunities(items);
      if (items.length && (!selectedCommunity || !items.find((c) => c.id === selectedCommunity.id))) {
        setSelectedCommunity(items[0]);
      }
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function loadChannels(communityId) {
    try {
      const data = await api(`/communities/${communityId}/channels`, "GET", null, accessToken);
      const items = data.items || [];
      setChannels(items);
      if (items.length && (!selectedChannel || !items.find((ch) => ch.id === selectedChannel.id))) {
        setSelectedChannel(items[0]);
      }
      if (!items.length) {
        setSelectedChannel(null);
        setMessages([]);
      }
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function loadMembers(communityId) {
    try {
      const data = await api(`/communities/${communityId}/members`, "GET", null, accessToken);
      setMembers(data.items || []);
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function loadMessages(channelId) {
    try {
      const data = await api(`/channels/${channelId}/messages`, "GET", null, accessToken);
      setMessages(data.items || []);
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function loadPresence(channelId) {
    try {
      const data = await api(`/channels/${channelId}/presence`, "GET", null, accessToken);
      setOnlineUsers(data.users || []);
    } catch {
      setOnlineUsers([]);
    }
  }

  async function createCommunity() {
    if (!communityName.trim()) return;
    try {
      await api("/communities", "POST", { name: communityName, safeModeEnabled: true }, accessToken);
      setCommunityName("");
      await loadCommunities();
      pushLog("Community created");
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function createChannel() {
    if (!selectedCommunity || !channelName.trim()) return;
    try {
      await api(`/communities/${selectedCommunity.id}/channels`, "POST", { name: channelName, type: channelType }, accessToken);
      setChannelName("");
      await loadChannels(selectedCommunity.id);
      pushLog(`Channel ${channelType} created`);
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function sendMessage() {
    if (!selectedChannel || !messageText.trim()) return;
    try {
      await api(`/channels/${selectedChannel.id}/messages`, "POST", { content: messageText }, accessToken);
      setMessageText("");
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function createInvite() {
    if (!selectedCommunity) return;
    try {
      const data = await api(`/communities/${selectedCommunity.id}/invites`, "POST", { hours: 24, maxUses: 25 }, accessToken);
      setInviteCodeCreated(data.code);
      pushLog(`Invite created: ${data.code}`);
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function joinByInvite() {
    if (!joinInviteCode.trim()) return;
    try {
      await api(`/invites/${joinInviteCode.trim().toUpperCase()}/join`, "POST", {}, accessToken);
      setJoinInviteCode("");
      await loadCommunities();
      pushLog("Joined community by invite");
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function moderate(targetUserId, action) {
    if (!selectedCommunity) return;
    try {
      await api(`/communities/${selectedCommunity.id}/moderation/${action}`, "POST", { targetUserId }, accessToken);
      await loadMembers(selectedCommunity.id);
      pushLog(`${action} applied on member`);
    } catch (e) {
      pushLog(e.message);
    }
  }

  async function prepareVoiceToken() {
    if (!selectedChannel) return;
    const data = await api(`/voice/token`, "POST", { channelId: selectedChannel.id }, accessToken);
    setVoiceToken(data.token || "");
    return data;
  }

  async function joinVoice() {
    if (!selectedChannel || selectedChannel.type !== "voice") return;
    if (roomRef.current) return;
    try {
      setVoiceState("connecting");
      const data = await prepareVoiceToken();
      const room = new Room();

      room.on(RoomEvent.Connected, () => {
        setVoiceState("connected");
        updateVoiceParticipants(room);
        pushLog("Voice connected");
      });
      room.on(RoomEvent.Disconnected, () => {
        setVoiceState("idle");
        setMicEnabled(false);
        setVoiceParticipants(0);
        pushLog("Voice disconnected");
      });
      room.on(RoomEvent.ParticipantConnected, () => updateVoiceParticipants(room));
      room.on(RoomEvent.ParticipantDisconnected, () => updateVoiceParticipants(room));
      room.on(RoomEvent.TrackSubscribed, (track) => {
        if (track.kind === "audio") {
          const el = track.attach();
          el.style.display = "none";
          document.body.appendChild(el);
        }
      });
      room.on(RoomEvent.TrackUnsubscribed, (track) => {
        track.detach().forEach((el) => el.remove());
      });

      await room.connect(data.serverUrl, data.token);
      await room.localParticipant.setMicrophoneEnabled(true);
      setMicEnabled(true);
      roomRef.current = room;
    } catch (e) {
      setVoiceState("error");
      pushLog(`Voice error: ${e.message}`);
    }
  }

  async function leaveVoice() {
    if (!roomRef.current) return;
    roomRef.current.disconnect();
    roomRef.current = null;
    setVoiceState("idle");
    setMicEnabled(false);
    setVoiceParticipants(0);
  }

  async function toggleMic() {
    if (!roomRef.current) return;
    const next = !micEnabled;
    await roomRef.current.localParticipant.setMicrophoneEnabled(next);
    setMicEnabled(next);
  }

  async function exportAudit() {
    if (!selectedCommunity) return;
    try {
      const url = `${API_BASE}/communities/${selectedCommunity.id}/exports?format=csv`;
      const res = await fetch(url, { headers: { Authorization: `Bearer ${accessToken}` } });
      const csv = await res.text();
      pushLog(`Audit CSV ready (${csv.length} chars)`);
    } catch (e) {
      pushLog(e.message);
    }
  }

  function resetSessionState() {
    setCurrentUser(null);
    setCommunities([]);
    setChannels([]);
    setMembers([]);
    setMessages([]);
    setOnlineUsers([]);
    setSelectedCommunity(null);
    setSelectedChannel(null);
    setVoiceToken("");
    setVoiceState("idle");
    setVoiceParticipants(0);
    setMicEnabled(false);
    setInviteCodeCreated("");
    setJoinInviteCode("");
  }

  async function logout() {
    try {
      await api("/auth/logout", "POST", {}, accessToken);
    } catch {
      // best effort
    } finally {
      await leaveVoice();
      resetSessionState();
      setAccessToken("");
      pushLog("Logged out");
    }
  }

  if (!loggedIn) {
    return (
      <main className="auth-screen">
        <div className="auth-card">
          <h1>Oris</h1>
          <p>Open-source community voice and text with safe defaults.</p>
          <input value={username} onChange={(e) => setUsername(e.target.value)} placeholder="username" />
          <input value={password} type="password" onChange={(e) => setPassword(e.target.value)} placeholder="password" />
          <div className="row">
            <button onClick={register}>Create account</button>
            <button className="ghost" onClick={login}>Login</button>
          </div>
        </div>
      </main>
    );
  }

  return (
    <main className="app-shell">
      <aside className="servers-panel">
        <h3>Communities</h3>
        <div className="server-list">
          {communities.map((c) => (
            <button
              key={c.id}
              className={selectedCommunity?.id === c.id ? "server-btn active" : "server-btn"}
              onClick={() => setSelectedCommunity(c)}
              title={c.name}
            >
              {c.name.slice(0, 2).toUpperCase()}
            </button>
          ))}
        </div>
        <input value={communityName} onChange={(e) => setCommunityName(e.target.value)} placeholder="new community" />
        <button onClick={createCommunity}>Create</button>
      </aside>

      <aside className="channels-panel">
        <h3>{selectedCommunity?.name || "No community"}</h3>
        <div className="channel-list">
          {channels.map((ch) => (
            <button
              key={ch.id}
              className={selectedChannel?.id === ch.id ? "channel-btn active" : "channel-btn"}
              onClick={() => setSelectedChannel(ch)}
            >
              #{ch.name} <span>{ch.type}</span>
            </button>
          ))}
        </div>
        <input value={channelName} onChange={(e) => setChannelName(e.target.value)} placeholder="new channel" />
        <select value={channelType} onChange={(e) => setChannelType(e.target.value)}>
          <option value="text">text</option>
          <option value="voice">voice</option>
        </select>
        <button onClick={createChannel}>Add channel</button>

        <div className="invite-block">
          <h4>Invites</h4>
          <div className="row">
            <button onClick={createInvite} disabled={!selectedCommunity}>Create invite</button>
          </div>
          {inviteCodeCreated && <p className="small">Code: <strong>{inviteCodeCreated}</strong></p>}
          <div className="row">
            <input value={joinInviteCode} onChange={(e) => setJoinInviteCode(e.target.value)} placeholder="invite code" />
            <button onClick={joinByInvite}>Join</button>
          </div>
        </div>
      </aside>

      <section className="chat-panel">
        <header>
          <h2>{selectedChannel ? `#${selectedChannel.name}` : "Select a channel"}</h2>
          <div className="actions">
            <span className="small">{onlineUsers.length} online</span>
            <button onClick={exportAudit} disabled={!selectedCommunity}>Export audit</button>
          </div>
        </header>

        {selectedChannel?.type === "voice" ? (
          <div className="voice-panel">
            <p>Voice room for <strong>{selectedChannel.name}</strong></p>
            <p>Status: {voiceState}</p>
            <p>Participants: {voiceParticipants}</p>
            <div className="row">
              <button onClick={joinVoice} disabled={voiceState === "connected" || voiceState === "connecting"}>Join voice</button>
              <button onClick={leaveVoice} disabled={voiceState !== "connected"}>Leave</button>
              <button onClick={toggleMic} disabled={voiceState !== "connected"}>{micEnabled ? "Mute mic" : "Unmute mic"}</button>
            </div>
            <p className="small">Voice token: {voiceToken ? "ready" : "not ready"}</p>
          </div>
        ) : (
          <>
            <div className="messages" ref={messagesRef}>
              {orderedMessages.map((m) => {
                const authorId = m.authorUserId || m.authorID;
                const authorName = usernameByUserId[authorId] || shortId(authorId);
                return (
                  <div key={m.id} className="msg">
                    <div className="msg-head">
                      <strong>{authorName}</strong>
                      <span>{formatTime(m.createdAt)}</span>
                    </div>
                    <p>{m.content}</p>
                  </div>
                );
              })}
            </div>

            <footer>
              <input value={messageText} onChange={(e) => setMessageText(e.target.value)} placeholder="Type a message..." />
              <button onClick={sendMessage} disabled={!selectedChannel || selectedChannel.type !== "text"}>Send</button>
            </footer>
          </>
        )}
      </section>

      <aside className="status-panel">
        <div className="status-header">
          <h3>{currentUser?.username || "User"}</h3>
          <button className="ghost" onClick={logout}>Logout</button>
        </div>

        <h3>Online In Channel</h3>
        <div className="members-list">
          {onlineUsers.map((u) => (
            <div key={u.userId} className="member-row">
              <div>
                <strong>{u.username}</strong>
                <small>{u.userId}</small>
              </div>
            </div>
          ))}
        </div>

        <h3>Members</h3>
        <div className="members-list">
          {members.map((m) => {
            const disabledAction = !canModerate || !currentUser || m.userId === currentUser.id || roleRank(m.role) >= roleRank(myMembership?.role);
            return (
              <div key={m.userId} className="member-row">
                <div>
                  <strong>{m.username}</strong>
                  <small>{m.role}{m.muted ? " • muted" : ""}{m.banned ? " • banned" : ""}</small>
                </div>
                <div className="member-actions">
                  <button disabled={disabledAction} onClick={() => moderate(m.userId, "mute")}>Mute</button>
                  <button disabled={disabledAction} onClick={() => moderate(m.userId, "kick")}>Kick</button>
                  <button disabled={disabledAction} onClick={() => moderate(m.userId, "ban")}>Ban</button>
                </div>
              </div>
            );
          })}
        </div>

        <h4>Activity</h4>
        <ul>
          {logs.map((l, i) => <li key={i}>{l}</li>)}
        </ul>
      </aside>
    </main>
  );
}
