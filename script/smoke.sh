#!/usr/bin/env bash

set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Run from root of repo
cd "${MY_DIR}/.."

note() {
  echo ">>> note: $*"
}

client() {
  host=${HOST:-"127.0.0.1:8080"}
  note "$@"
  ./script/generated/cli/client.sh -i --host "$host" "$@" &2>/dev/null
}

sessions=("a" "b")
teams=("0", "1")
spirits=("A", "B")
check() {
  note "check"
  client listSessions
  for session in "${sessions[@]}"; do
    client listSessionTeams sessionName="session-${session}"
    for team in "${teams[@]}"; do
      client listSessionTeamSpirits sessionName="session-${session}" teamName="team-${team}"
    done
  done
}

note "no sessions"
check

note "create sessions"
for session in "${sessions[@]}"; do
  client createSession name=="session-${session}"
done
check

note "create teams"
for session in "${sessions[@]}"; do
  for team in "${teams[@]}"; do
    client createSessionTeam sessionName="session-${session}" name=="team-${team}"
  done
done

note "create spirits"
for session in "${sessions[@]}"; do
  for team in "${teams[@]}"; do
    for spirit in "${spirits[@]}"; do
      client createSessionTeamSpirit sessionName="session-${session}" teamName="team-${team}" name=="spirit-${spirit}"
    done
  done
done
check
