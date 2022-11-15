/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as SpiritsV1Meta from "./meta.pb"
import * as SpiritsV1Spirit from "./spirit.pb"

export enum BattleState {
  BATTLE_STATE_UNSPECIFIED = "BATTLE_STATE_UNSPECIFIED",
  BATTLE_STATE_PENDING = "BATTLE_STATE_PENDING",
  BATTLE_STATE_STARTED = "BATTLE_STATE_STARTED",
  BATTLE_STATE_WAITING = "BATTLE_STATE_WAITING",
  BATTLE_STATE_FINISHED = "BATTLE_STATE_FINISHED",
  BATTLE_STATE_CANCELLED = "BATTLE_STATE_CANCELLED",
  BATTLE_STATE_ERROR = "BATTLE_STATE_ERROR",
}

export enum BattleTeamSpiritIntelligence {
  BATTLE_TEAM_SPIRIT_INTELLIGENCE_UNSPECIFIED = "BATTLE_TEAM_SPIRIT_INTELLIGENCE_UNSPECIFIED",
  BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN = "BATTLE_TEAM_SPIRIT_INTELLIGENCE_HUMAN",
  BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM = "BATTLE_TEAM_SPIRIT_INTELLIGENCE_RANDOM",
}

export type CreateBattleRequest = {
}

export type CreateBattleResponse = {
  battle?: Battle
}

export type ListBattlesRequest = {
}

export type ListBattlesResponse = {
  battles?: Battle[]
}

export type WatchBattleRequest = {
  id?: string
}

export type WatchBattleResponse = {
  battle?: Battle
}

export type AddBattleTeamRequest = {
  battleId?: string
  teamName?: string
}

export type AddBattleTeamResponse = {
  battle?: Battle
}

export type AddBattleTeamSpiritRequest = {
  battleId?: string
  teamName?: string
  spiritId?: string
  intelligence?: BattleTeamSpiritIntelligence
  seed?: string
}

export type AddBattleTeamSpiritResponse = {
  battle?: Battle
}

export type StartBattleRequest = {
  id?: string
}

export type StartBattleResponse = {
  battle?: Battle
}

export type CancelBattleRequest = {
  id?: string
}

export type CancelBattleResponse = {
  battle?: Battle
}

export type CallActionRequest = {
  battleId?: string
  spiritId?: string
  turn?: string
  actionName?: string
  targetSpiritIds?: string[]
}

export type CallActionResponse = {
}

export type Battle = {
  meta?: SpiritsV1Meta.Meta
  state?: BattleState
  errorMessage?: string
  teams?: BattleTeam[]
  inBattleTeams?: BattleTeam[]
  nextSpiritIds?: string[]
  turns?: string
}

export type BattleTeam = {
  name?: string
  spirits?: BattleTeamSpirit[]
}

export type BattleTeamSpirit = {
  spirit?: SpiritsV1Spirit.Spirit
  intelligence?: BattleTeamSpiritIntelligence
  seed?: string
}

export class BattleService {
  static CreateBattle(req: CreateBattleRequest, initReq?: fm.InitReq): Promise<CreateBattleResponse> {
    return fm.fetchReq<CreateBattleRequest, CreateBattleResponse>(`/spirits.v1.BattleService/CreateBattle`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static WatchBattle(req: WatchBattleRequest, entityNotifier?: fm.NotifyStreamEntityArrival<WatchBattleResponse>, initReq?: fm.InitReq): Promise<void> {
    return fm.fetchStreamingRequest<WatchBattleRequest, WatchBattleResponse>(`/spirits.v1.BattleService/WatchBattle`, entityNotifier, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ListBattles(req: ListBattlesRequest, initReq?: fm.InitReq): Promise<ListBattlesResponse> {
    return fm.fetchReq<ListBattlesRequest, ListBattlesResponse>(`/spirits.v1.BattleService/ListBattles`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static AddBattleTeam(req: AddBattleTeamRequest, initReq?: fm.InitReq): Promise<AddBattleTeamResponse> {
    return fm.fetchReq<AddBattleTeamRequest, AddBattleTeamResponse>(`/spirits.v1.BattleService/AddBattleTeam`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static AddBattleTeamSpirit(req: AddBattleTeamSpiritRequest, initReq?: fm.InitReq): Promise<AddBattleTeamSpiritResponse> {
    return fm.fetchReq<AddBattleTeamSpiritRequest, AddBattleTeamSpiritResponse>(`/spirits.v1.BattleService/AddBattleTeamSpirit`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static StartBattle(req: StartBattleRequest, initReq?: fm.InitReq): Promise<StartBattleResponse> {
    return fm.fetchReq<StartBattleRequest, StartBattleResponse>(`/spirits.v1.BattleService/StartBattle`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CancelBattle(req: CancelBattleRequest, initReq?: fm.InitReq): Promise<CancelBattleResponse> {
    return fm.fetchReq<CancelBattleRequest, CancelBattleResponse>(`/spirits.v1.BattleService/CancelBattle`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CallAction(req: CallActionRequest, initReq?: fm.InitReq): Promise<CallActionResponse> {
    return fm.fetchReq<CallActionRequest, CallActionResponse>(`/spirits.v1.BattleService/CallAction`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}