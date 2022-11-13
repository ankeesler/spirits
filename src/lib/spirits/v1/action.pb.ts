/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as SpiritsV1Meta from "./meta.pb"

type Absent<T, K extends keyof T> = { [k in Exclude<keyof T, K>]?: undefined };
type OneOf<T> =
  | { [k in keyof T]?: undefined }
  | (
    keyof T extends infer K ?
      (K extends string & keyof T ? { [k in K]: T[K] } & Absent<T, K>
        : never)
    : never);
export type CreateActionRequest = {
  action?: Action
}

export type CreateActionResponse = {
  action?: Action
}

export type GetActionRequest = {
  id?: string
}

export type GetActionResponse = {
  action?: Action
}

export type ListActionsRequest = {
}

export type ListActionsResponse = {
  actions?: Action[]
}

export type UpdateActionRequest = {
  action?: Action
}

export type UpdateActionResponse = {
  action?: Action
}

export type DeleteActionRequest = {
  id?: string
}

export type DeleteActionResponse = {
  action?: Action
}


type BaseAction = {
  meta?: SpiritsV1Meta.Meta
  description?: string
}

export type Action = BaseAction
  & OneOf<{ script: string }>

export class ActionService {
  static CreateAction(req: CreateActionRequest, initReq?: fm.InitReq): Promise<CreateActionResponse> {
    return fm.fetchReq<CreateActionRequest, CreateActionResponse>(`/spirits.v1.ActionService/CreateAction`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAction(req: GetActionRequest, initReq?: fm.InitReq): Promise<GetActionResponse> {
    return fm.fetchReq<GetActionRequest, GetActionResponse>(`/spirits.v1.ActionService/GetAction`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ListActions(req: ListActionsRequest, initReq?: fm.InitReq): Promise<ListActionsResponse> {
    return fm.fetchReq<ListActionsRequest, ListActionsResponse>(`/spirits.v1.ActionService/ListActions`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateAction(req: UpdateActionRequest, initReq?: fm.InitReq): Promise<UpdateActionResponse> {
    return fm.fetchReq<UpdateActionRequest, UpdateActionResponse>(`/spirits.v1.ActionService/UpdateAction`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteAction(req: DeleteActionRequest, initReq?: fm.InitReq): Promise<DeleteActionResponse> {
    return fm.fetchReq<DeleteActionRequest, DeleteActionResponse>(`/spirits.v1.ActionService/DeleteAction`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}