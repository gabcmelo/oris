# Desktop Plan (Post-MVP)

This project is web-first for MVP. Desktop support will be added after MVP stability.

## Strategy
1. Use Electron with:
   - `main` process for window/lifecycle and updater hooks.
   - `preload` for secure IPC boundaries.
   - `renderer` pointing to frontend build output.
2. Keep backend self-host model unchanged.
3. Desktop client connects to existing Oris server endpoints.

## Planned milestones
1. Basic shell app with login and community join.
2. Installer packaging (Windows/macOS/Linux).
3. Stable channel auto-update for desktop client.
4. First-run setup wizard for non-technical users.

## Not in current scope
1. Electron dependency integration in MVP branch.
2. Native modules or platform-specific voice stack customization.
