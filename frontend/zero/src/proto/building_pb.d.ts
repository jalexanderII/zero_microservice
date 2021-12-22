import * as jspb from 'google-protobuf'

import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as content_pb from './content_pb';


export class Building extends jspb.Message {
  getId(): number;
  setId(value: number): Building;

  getName(): string;
  setName(value: string): Building;

  getFullAddress(): string;
  setFullAddress(value: string): Building;

  getStreet(): string;
  setStreet(value: string): Building;

  getCity(): string;
  setCity(value: string): Building;

  getState(): string;
  setState(value: string): Building;

  getZipCode(): number;
  setZipCode(value: number): Building;

  getNeighborhood(): string;
  setNeighborhood(value: string): Building;

  getLat(): number;
  setLat(value: number): Building;

  getLng(): number;
  setLng(value: number): Building;

  getDescription(): string;
  setDescription(value: string): Building;

  getAmenitiesList(): Array<string>;
  setAmenitiesList(value: Array<string>): Building;
  clearAmenitiesList(): Building;
  addAmenities(value: string, index?: number): Building;

  getUploadIdsList(): Array<string>;
  setUploadIdsList(value: Array<string>): Building;
  clearUploadIdsList(): Building;
  addUploadIds(value: string, index?: number): Building;

  getRealtorRef(): number;
  setRealtorRef(value: number): Building;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Building.AsObject;
  static toObject(includeInstance: boolean, msg: Building): Building.AsObject;
  static serializeBinaryToWriter(message: Building, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Building;
  static deserializeBinaryFromReader(message: Building, reader: jspb.BinaryReader): Building;
}

export namespace Building {
  export type AsObject = {
    id: number,
    name: string,
    fullAddress: string,
    street: string,
    city: string,
    state: string,
    zipCode: number,
    neighborhood: string,
    lat: number,
    lng: number,
    description: string,
    amenitiesList: Array<string>,
    uploadIdsList: Array<string>,
    realtorRef: number,
  }
}

export class CreateBuildingRequest extends jspb.Message {
  getBuilding(): Building | undefined;
  setBuilding(value?: Building): CreateBuildingRequest;
  hasBuilding(): boolean;
  clearBuilding(): CreateBuildingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateBuildingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateBuildingRequest): CreateBuildingRequest.AsObject;
  static serializeBinaryToWriter(message: CreateBuildingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateBuildingRequest;
  static deserializeBinaryFromReader(message: CreateBuildingRequest, reader: jspb.BinaryReader): CreateBuildingRequest;
}

export namespace CreateBuildingRequest {
  export type AsObject = {
    building?: Building.AsObject,
  }
}

export class GetBuildingRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetBuildingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetBuildingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetBuildingRequest): GetBuildingRequest.AsObject;
  static serializeBinaryToWriter(message: GetBuildingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetBuildingRequest;
  static deserializeBinaryFromReader(message: GetBuildingRequest, reader: jspb.BinaryReader): GetBuildingRequest;
}

export namespace GetBuildingRequest {
  export type AsObject = {
    id: number,
  }
}

export class UpdateBuildingRequest extends jspb.Message {
  getId(): number;
  setId(value: number): UpdateBuildingRequest;

  getBuilding(): Building | undefined;
  setBuilding(value?: Building): UpdateBuildingRequest;
  hasBuilding(): boolean;
  clearBuilding(): UpdateBuildingRequest;

  getMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateBuildingRequest;
  hasMask(): boolean;
  clearMask(): UpdateBuildingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateBuildingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateBuildingRequest): UpdateBuildingRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateBuildingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateBuildingRequest;
  static deserializeBinaryFromReader(message: UpdateBuildingRequest, reader: jspb.BinaryReader): UpdateBuildingRequest;
}

export namespace UpdateBuildingRequest {
  export type AsObject = {
    id: number,
    building?: Building.AsObject,
    mask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class DeleteBuildingRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteBuildingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteBuildingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteBuildingRequest): DeleteBuildingRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteBuildingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteBuildingRequest;
  static deserializeBinaryFromReader(message: DeleteBuildingRequest, reader: jspb.BinaryReader): DeleteBuildingRequest;
}

export namespace DeleteBuildingRequest {
  export type AsObject = {
    id: number,
  }
}

export class ListBuildingRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListBuildingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListBuildingRequest): ListBuildingRequest.AsObject;
  static serializeBinaryToWriter(message: ListBuildingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListBuildingRequest;
  static deserializeBinaryFromReader(message: ListBuildingRequest, reader: jspb.BinaryReader): ListBuildingRequest;
}

export namespace ListBuildingRequest {
  export type AsObject = {
  }
}

export class ListBuildingResponse extends jspb.Message {
  getBuildingsList(): Array<Building>;
  setBuildingsList(value: Array<Building>): ListBuildingResponse;
  clearBuildingsList(): ListBuildingResponse;
  addBuildings(value?: Building, index?: number): Building;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListBuildingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListBuildingResponse): ListBuildingResponse.AsObject;
  static serializeBinaryToWriter(message: ListBuildingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListBuildingResponse;
  static deserializeBinaryFromReader(message: ListBuildingResponse, reader: jspb.BinaryReader): ListBuildingResponse;
}

export namespace ListBuildingResponse {
  export type AsObject = {
    buildingsList: Array<Building.AsObject>,
  }
}

export class DeleteBuildingResponse extends jspb.Message {
  getStatus(): content_pb.STATUS;
  setStatus(value: content_pb.STATUS): DeleteBuildingResponse;

  getBuilding(): Building | undefined;
  setBuilding(value?: Building): DeleteBuildingResponse;
  hasBuilding(): boolean;
  clearBuilding(): DeleteBuildingResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteBuildingResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteBuildingResponse): DeleteBuildingResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteBuildingResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteBuildingResponse;
  static deserializeBinaryFromReader(message: DeleteBuildingResponse, reader: jspb.BinaryReader): DeleteBuildingResponse;
}

export namespace DeleteBuildingResponse {
  export type AsObject = {
    status: content_pb.STATUS,
    building?: Building.AsObject,
  }
}

