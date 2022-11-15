/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufTimestamp from "../../google/protobuf/timestamp.pb"
export type Meta = {
  id?: string
  createdTime?: GoogleProtobufTimestamp.Timestamp
  createdBy?: Identity
  updatedTime?: GoogleProtobufTimestamp.Timestamp
  updatedBy?: Identity
}

export type Identity = {
  principle?: string
}