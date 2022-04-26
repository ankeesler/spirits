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
*ActionsApi* | [**createSessionBattleSpiritActions**](docs/ActionsApi.md#createsessionbattlespiritactions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
*BattlesApi* | [**createSessionBattleSpiritActions**](docs/BattlesApi.md#createsessionbattlespiritactions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
*BattlesApi* | [**createSessionBattles**](docs/BattlesApi.md#createsessionbattles) | **POST** /sessions/{sessionName}/battles | 
*BattlesApi* | [**deleteSessionBattles**](docs/BattlesApi.md#deletesessionbattles) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
*BattlesApi* | [**getSessionBattleSpirits**](docs/BattlesApi.md#getsessionbattlespirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
*BattlesApi* | [**getSessionBattles**](docs/BattlesApi.md#getsessionbattles) | **GET** /sessions/{sessionName}/battles/{battleName} | 
*BattlesApi* | [**listSessionsBattles**](docs/BattlesApi.md#listsessionsbattles) | **GET** /sessions/{sessionName}/battles | 
*BattlesApi* | [**listSessionsBattlesSpirits**](docs/BattlesApi.md#listsessionsbattlesspirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
*DefaultApi* | [**rootGet**](docs/DefaultApi.md#rootget) | **GET** / | 
*SessionsApi* | [**createSessionBattleSpiritActions**](docs/SessionsApi.md#createsessionbattlespiritactions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
*SessionsApi* | [**createSessionBattles**](docs/SessionsApi.md#createsessionbattles) | **POST** /sessions/{sessionName}/battles | 
*SessionsApi* | [**createSessionTeamSpirits**](docs/SessionsApi.md#createsessionteamspirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SessionsApi* | [**createSessionTeams**](docs/SessionsApi.md#createsessionteams) | **POST** /sessions/{sessionName}/teams | 
*SessionsApi* | [**createSessions**](docs/SessionsApi.md#createsessions) | **POST** /sessions | 
*SessionsApi* | [**deleteSessionBattles**](docs/SessionsApi.md#deletesessionbattles) | **DELETE** /sessions/{sessionName}/battles/{battleName} | 
*SessionsApi* | [**deleteSessionTeamSpirits**](docs/SessionsApi.md#deletesessionteamspirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SessionsApi* | [**deleteSessionTeams**](docs/SessionsApi.md#deletesessionteams) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
*SessionsApi* | [**deleteSessions**](docs/SessionsApi.md#deletesessions) | **DELETE** /sessions/{sessionName} | 
*SessionsApi* | [**getSessionBattleSpirits**](docs/SessionsApi.md#getsessionbattlespirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
*SessionsApi* | [**getSessionBattles**](docs/SessionsApi.md#getsessionbattles) | **GET** /sessions/{sessionName}/battles/{battleName} | 
*SessionsApi* | [**getSessionTeamSpirits**](docs/SessionsApi.md#getsessionteamspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SessionsApi* | [**getSessionTeams**](docs/SessionsApi.md#getsessionteams) | **GET** /sessions/{sessionName}/teams/{teamName} | 
*SessionsApi* | [**getSessions**](docs/SessionsApi.md#getsessions) | **GET** /sessions/{sessionName} | 
*SessionsApi* | [**listSessions**](docs/SessionsApi.md#listsessions) | **GET** /sessions | 
*SessionsApi* | [**listSessionsBattles**](docs/SessionsApi.md#listsessionsbattles) | **GET** /sessions/{sessionName}/battles | 
*SessionsApi* | [**listSessionsBattlesSpirits**](docs/SessionsApi.md#listsessionsbattlesspirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
*SessionsApi* | [**listSessionsTeams**](docs/SessionsApi.md#listsessionsteams) | **GET** /sessions/{sessionName}/teams | 
*SessionsApi* | [**listSessionsTeamsSpirits**](docs/SessionsApi.md#listsessionsteamsspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SessionsApi* | [**updateSessionTeamSpirits**](docs/SessionsApi.md#updatesessionteamspirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SessionsApi* | [**updateSessionTeams**](docs/SessionsApi.md#updatesessionteams) | **PUT** /sessions/{sessionName}/teams/{teamName} | 
*SessionsApi* | [**updateSessions**](docs/SessionsApi.md#updatesessions) | **PUT** /sessions/{sessionName} | 
*SpiritsApi* | [**createSessionBattleSpiritActions**](docs/SpiritsApi.md#createsessionbattlespiritactions) | **POST** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName}/actions | 
*SpiritsApi* | [**createSessionTeamSpirits**](docs/SpiritsApi.md#createsessionteamspirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SpiritsApi* | [**deleteSessionTeamSpirits**](docs/SpiritsApi.md#deletesessionteamspirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SpiritsApi* | [**getSessionBattleSpirits**](docs/SpiritsApi.md#getsessionbattlespirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits/{spiritName} | 
*SpiritsApi* | [**getSessionTeamSpirits**](docs/SpiritsApi.md#getsessionteamspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*SpiritsApi* | [**listSessionsBattlesSpirits**](docs/SpiritsApi.md#listsessionsbattlesspirits) | **GET** /sessions/{sessionName}/battles/{battleName}/spirits | 
*SpiritsApi* | [**listSessionsTeamsSpirits**](docs/SpiritsApi.md#listsessionsteamsspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
*SpiritsApi* | [**updateSessionTeamSpirits**](docs/SpiritsApi.md#updatesessionteamspirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*TeamsApi* | [**createSessionTeamSpirits**](docs/TeamsApi.md#createsessionteamspirits) | **POST** /sessions/{sessionName}/teams/{teamName}/spirits | 
*TeamsApi* | [**createSessionTeams**](docs/TeamsApi.md#createsessionteams) | **POST** /sessions/{sessionName}/teams | 
*TeamsApi* | [**deleteSessionTeamSpirits**](docs/TeamsApi.md#deletesessionteamspirits) | **DELETE** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*TeamsApi* | [**deleteSessionTeams**](docs/TeamsApi.md#deletesessionteams) | **DELETE** /sessions/{sessionName}/teams/{teamName} | 
*TeamsApi* | [**getSessionTeamSpirits**](docs/TeamsApi.md#getsessionteamspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*TeamsApi* | [**getSessionTeams**](docs/TeamsApi.md#getsessionteams) | **GET** /sessions/{sessionName}/teams/{teamName} | 
*TeamsApi* | [**listSessionsTeams**](docs/TeamsApi.md#listsessionsteams) | **GET** /sessions/{sessionName}/teams | 
*TeamsApi* | [**listSessionsTeamsSpirits**](docs/TeamsApi.md#listsessionsteamsspirits) | **GET** /sessions/{sessionName}/teams/{teamName}/spirits | 
*TeamsApi* | [**updateSessionTeamSpirits**](docs/TeamsApi.md#updatesessionteamspirits) | **PUT** /sessions/{sessionName}/teams/{teamName}/spirits/{spiritName} | 
*TeamsApi* | [**updateSessionTeams**](docs/TeamsApi.md#updatesessionteams) | **PUT** /sessions/{sessionName}/teams/{teamName} | 


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

