/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as SpiritsV1Action from "./action.pb"
import * as SpiritsV1Meta from "./meta.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type CreateSpiritRequest = {
  spirit?: Spirit
}

export type CreateSpiritResponse = {
  spirit?: Spirit
}

export type GetSpiritRequest = {
  id?: string
}

export type GetSpiritResponse = {
  spirit?: Spirit
}

export type ListSpiritsRequest = {
  name?: string
}

export type ListSpiritsResponse = {
  spirits?: Spirit[]
}

export type UpdateSpiritRequest = {
  spirit?: Spirit
}

export type UpdateSpiritResponse = {
  spirit?: Spirit
}

export type DeleteSpiritRequest = {
  id?: string
}

export type DeleteSpiritResponse = {
  spirit?: Spirit
}

export type Spirit = {
  meta?: SpiritsV1Meta.Meta
  name?: string
  stats?: SpiritStats
  actions?: SpiritAction[]
}

export type SpiritStats = {
  health?: string
  physicalPower?: string
  physicalConstitution?: string
  mentalPower?: string
  mentalConstitution?: string
  agility?: string
}


type BaseSpiritAction = {
  name?: string
}

export type SpiritAction = BaseSpiritAction
  & OneOf<{ actionId: string; inline: SpiritsV1Action.Action }>

export class SpiritService {
  static CreateSpirit(req: CreateSpiritRequest, initReq?: fm.InitReq): Promise<CreateSpiritResponse> {
    return fm.fetchReq<CreateSpiritRequest, CreateSpiritResponse>(`/spirits/v1/spirits`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetSpirit(req: GetSpiritRequest, initReq?: fm.InitReq): Promise<GetSpiritResponse> {
    return fm.fetchReq<GetSpiritRequest, GetSpiritResponse>(`/spirits/v1/${req["idspirits"]}?${fm.renderURLSearchParams(req, ["idspirits"])}`, {...initReq, method: "GET"})
  }
  static ListSpirits(req: ListSpiritsRequest, initReq?: fm.InitReq): Promise<ListSpiritsResponse> {
    return fm.fetchReq<ListSpiritsRequest, ListSpiritsResponse>(`/spirits/v1/spirits?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static UpdateSpirit(req: UpdateSpiritRequest, initReq?: fm.InitReq): Promise<UpdateSpiritResponse> {
    return fm.fetchReq<UpdateSpiritRequest, UpdateSpiritResponse>(`/spirits/v1/${req["spiritMetaIdspirits"]}`, {...initReq, method: "PUT", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteSpirit(req: DeleteSpiritRequest, initReq?: fm.InitReq): Promise<DeleteSpiritResponse> {
    return fm.fetchReq<DeleteSpiritRequest, DeleteSpiritResponse>(`/spirits/v1/${req["idspirits"]}`, {...initReq, method: "DELETE"})
  }
}