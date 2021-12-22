import * as jspb from 'google-protobuf'

import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as content_pb from './content_pb';


export class Realtor extends jspb.Message {
  getId(): number;
  setId(value: number): Realtor;

  getName(): string;
  setName(value: string): Realtor;

  getEmail(): string;
  setEmail(value: string): Realtor;

  getPhoneNumber(): string;
  setPhoneNumber(value: string): Realtor;

  getCompany(): string;
  setCompany(value: string): Realtor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Realtor.AsObject;
  static toObject(includeInstance: boolean, msg: Realtor): Realtor.AsObject;
  static serializeBinaryToWriter(message: Realtor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Realtor;
  static deserializeBinaryFromReader(message: Realtor, reader: jspb.BinaryReader): Realtor;
}

export namespace Realtor {
  export type AsObject = {
    id: number,
    name: string,
    email: string,
    phoneNumber: string,
    company: string,
  }
}

export class CreateRealtorRequest extends jspb.Message {
  getRealtor(): Realtor | undefined;
  setRealtor(value?: Realtor): CreateRealtorRequest;
  hasRealtor(): boolean;
  clearRealtor(): CreateRealtorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRealtorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRealtorRequest): CreateRealtorRequest.AsObject;
  static serializeBinaryToWriter(message: CreateRealtorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRealtorRequest;
  static deserializeBinaryFromReader(message: CreateRealtorRequest, reader: jspb.BinaryReader): CreateRealtorRequest;
}

export namespace CreateRealtorRequest {
  export type AsObject = {
    realtor?: Realtor.AsObject,
  }
}

export class GetRealtorRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetRealtorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetRealtorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetRealtorRequest): GetRealtorRequest.AsObject;
  static serializeBinaryToWriter(message: GetRealtorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetRealtorRequest;
  static deserializeBinaryFromReader(message: GetRealtorRequest, reader: jspb.BinaryReader): GetRealtorRequest;
}

export namespace GetRealtorRequest {
  export type AsObject = {
    id: number,
  }
}

export class UpdateRealtorRequest extends jspb.Message {
  getId(): number;
  setId(value: number): UpdateRealtorRequest;

  getRealtor(): Realtor | undefined;
  setRealtor(value?: Realtor): UpdateRealtorRequest;
  hasRealtor(): boolean;
  clearRealtor(): UpdateRealtorRequest;

  getMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateRealtorRequest;
  hasMask(): boolean;
  clearMask(): UpdateRealtorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateRealtorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateRealtorRequest): UpdateRealtorRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateRealtorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateRealtorRequest;
  static deserializeBinaryFromReader(message: UpdateRealtorRequest, reader: jspb.BinaryReader): UpdateRealtorRequest;
}

export namespace UpdateRealtorRequest {
  export type AsObject = {
    id: number,
    realtor?: Realtor.AsObject,
    mask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class DeleteRealtorRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteRealtorRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRealtorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRealtorRequest): DeleteRealtorRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteRealtorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRealtorRequest;
  static deserializeBinaryFromReader(message: DeleteRealtorRequest, reader: jspb.BinaryReader): DeleteRealtorRequest;
}

export namespace DeleteRealtorRequest {
  export type AsObject = {
    id: number,
  }
}

export class ListRealtorRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRealtorRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListRealtorRequest): ListRealtorRequest.AsObject;
  static serializeBinaryToWriter(message: ListRealtorRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRealtorRequest;
  static deserializeBinaryFromReader(message: ListRealtorRequest, reader: jspb.BinaryReader): ListRealtorRequest;
}

export namespace ListRealtorRequest {
  export type AsObject = {
  }
}

export class ListRealtorResponse extends jspb.Message {
  getRealtorsList(): Array<Realtor>;
  setRealtorsList(value: Array<Realtor>): ListRealtorResponse;
  clearRealtorsList(): ListRealtorResponse;
  addRealtors(value?: Realtor, index?: number): Realtor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRealtorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListRealtorResponse): ListRealtorResponse.AsObject;
  static serializeBinaryToWriter(message: ListRealtorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRealtorResponse;
  static deserializeBinaryFromReader(message: ListRealtorResponse, reader: jspb.BinaryReader): ListRealtorResponse;
}

export namespace ListRealtorResponse {
  export type AsObject = {
    realtorsList: Array<Realtor.AsObject>,
  }
}

export class DeleteRealtorResponse extends jspb.Message {
  getStatus(): content_pb.STATUS;
  setStatus(value: content_pb.STATUS): DeleteRealtorResponse;

  getRealtor(): Realtor | undefined;
  setRealtor(value?: Realtor): DeleteRealtorResponse;
  hasRealtor(): boolean;
  clearRealtor(): DeleteRealtorResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteRealtorResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteRealtorResponse): DeleteRealtorResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteRealtorResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteRealtorResponse;
  static deserializeBinaryFromReader(message: DeleteRealtorResponse, reader: jspb.BinaryReader): DeleteRealtorResponse;
}

export namespace DeleteRealtorResponse {
  export type AsObject = {
    status: content_pb.STATUS,
    realtor?: Realtor.AsObject,
  }
}

