# spirits Bash client

## Overview

This is a Bash client script for accessing spirits service.

The script uses cURL underneath for making all REST calls.

## Usage

```shell
# Make sure the script has executable rights
$ chmod u+x 

# Print the list of operations available on the service
$ ./ -h

# Print the service description
$ ./ --about

# Print detailed information about specific operation
$ ./ <operationId> -h

# Make GET request
./ --host http://<hostname>:<port> --accept xml <operationId> <queryParam1>=<value1> <header_key1>:<header_value2>

# Make GET request using arbitrary curl options (must be passed before <operationId>) to an SSL service using username:password
 -k -sS --tlsv1.2 --host https://<hostname> -u <user>:<password> --accept xml <operationId> <queryParam1>=<value1> <header_key1>:<header_value2>

# Make POST request
$ echo '<body_content>' |  --host <hostname> --content-type json <operationId> -

# Make POST request with simple JSON content, e.g.:
# {
#   "key1": "value1",
#   "key2": "value2",
#   "key3": 23
# }
$ echo '<body_content>' |  --host <hostname> --content-type json <operationId> key1==value1 key2=value2 key3:=23 -

# Make POST request with form data
$  --host <hostname> <operationId> key1:=value1 key2:=value2 key3:=23

# Preview the cURL command without actually executing it
$  --host http://<hostname>:<port> --dry-run <operationid>

```

## Docker image

You can easily create a Docker image containing a preconfigured environment
for using the REST Bash client including working autocompletion and short
welcome message with basic instructions, using the generated Dockerfile:

```shell
docker build -t my-rest-client .
docker run -it my-rest-client
```

By default you will be logged into a Zsh environment which has much more
advanced auto completion, but you can switch to Bash, where basic autocompletion
is also available.

## Shell completion

### Bash

The generated bash-completion script can be either directly loaded to the current Bash session using:

```shell
source .bash-completion
```

Alternatively, the script can be copied to the `/etc/bash-completion.d` (or on OSX with Homebrew to `/usr/local/etc/bash-completion.d`):

```shell
sudo cp .bash-completion /etc/bash-completion.d/
```

#### OS X

On OSX you might need to install bash-completion using Homebrew:

```shell
brew install bash-completion
```

and add the following to the `~/.bashrc`:

```shell
if [ -f $(brew --prefix)/etc/bash_completion ]; then
  . $(brew --prefix)/etc/bash_completion
fi
```

### Zsh

In Zsh, the generated `_` Zsh completion file must be copied to one of the folders under `$FPATH` variable.

## Documentation for API Endpoints

All URIs are relative to **

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**rootGet**](docs/DefaultApi.md#rootget) | **GET** / | 
*SessionApi* | [**createSession**](docs/SessionApi.md#createsession) | **POST** /sessions | 
*SessionApi* | [**deleteSession**](docs/SessionApi.md#deletesession) | **DELETE** /sessions/{sessionName} | 
*SessionApi* | [**getSession**](docs/SessionApi.md#getsession) | **GET** /sessions/{sessionName} | 
*SessionApi* | [**listSessions**](docs/SessionApi.md#listsessions) | **GET** /sessions | 
*SessionApi* | [**updateSession**](docs/SessionApi.md#updatesession) | **PUT** /sessions/{sessionName} | 
*SessionBattleApi* | [**createSessionBattle**](docs/SessionBattleApi.md#createsessionbattle) | **POST** /sessions/{sessionName}/battles | 
*SessionBattleApi* | [**deleteSessionBattle**](docs/SessionBattleApi.md#deletesessionbattle) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
*SessionBattleApi* | [**getSessionBattle**](docs/SessionBattleApi.md#getsessionbattle) | **GET** /sessions/{sessionName}/battles/{battleName} | 
*SessionBattleApi* | [**listSessionBattles**](docs/SessionBattleApi.md#listsessionbattles) | **GET** /sessions/{sessionName}/battles | 
*SessionBattleSpiritApi* | [**getSessionBattleSpirit**](docs/SessionBattleSpiritApi.md#getsessionbattlespirit) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
*SessionBattleSpiritApi* | [**listSessionBattleSpirits**](docs/SessionBattleSpiritApi.md#listsessionbattlespirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
*SessionBattleSpiritActionApi* | [**createSessionBattleSpiritAction**](docs/SessionBattleSpiritActionApi.md#createsessionbattlespiritaction) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
*SessionTeamApi* | [**createSessionTeam**](docs/SessionTeamApi.md#createsessionteam) | **POST** /sessions/{sessionName}/teams | 
*SessionTeamApi* | [**deleteSessionTeam**](docs/SessionTeamApi.md#deletesessionteam) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
*SessionTeamApi* | [**getSessionTeam**](docs/SessionTeamApi.md#getsessionteam) | **GET** /sessions/{sessionName}/teams/{teamName} | 
*SessionTeamApi* | [**listSessionTeams**](docs/SessionTeamApi.md#listsessionteams) | **GET** /sessions/{sessionName}/teams | 
*SessionTeamApi* | [**updateSessionTeam**](docs/SessionTeamApi.md#updatesessionteam) | **PUT** /sessions/{sessionName}/teams/{teamName} | 
*SessionTeamSpiritApi* | [**createSessionTeamSpirit**](docs/SessionTeamSpiritApi.md#createsessionteamspirit) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SessionTeamSpiritApi* | [**deleteSessionTeamSpirit**](docs/SessionTeamSpiritApi.md#deletesessionteamspirit) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SessionTeamSpiritApi* | [**getSessionTeamSpirit**](docs/SessionTeamSpiritApi.md#getsessionteamspirit) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SessionTeamSpiritApi* | [**listSessionTeamSpirits**](docs/SessionTeamSpiritApi.md#listsessionteamspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SessionTeamSpiritApi* | [**updateSessionTeamSpirit**](docs/SessionTeamSpiritApi.md#updatesessionteamspirit) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 


## Documentation For Models

 - [Action](docs/Action.md)
 - [Battle](docs/Battle.md)
 - [Session](docs/Session.md)
 - [SessionAuth](docs/SessionAuth.md)
 - [SessionAuthOidc](docs/SessionAuthOidc.md)
 - [Spirit](docs/Spirit.md)
 - [SpiritStats](docs/SpiritStats.md)
 - [Team](docs/Team.md)


## Documentation For Authorization

 All endpoints do not require authorization.

