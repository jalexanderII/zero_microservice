import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as content_pb from './content_pb';


export class Apartment extends jspb.Message {
  getId(): number;
  setId(value: number): Apartment;

  getName(): string;
  setName(value: string): Apartment;

  getFullAddress(): string;
  setFullAddress(value: string): Apartment;

  getStreet(): string;
  setStreet(value: string): Apartment;

  getCity(): string;
  setCity(value: string): Apartment;

  getState(): string;
  setState(value: string): Apartment;

  getZipCode(): number;
  setZipCode(value: number): Apartment;

  getNeighborhood(): string;
  setNeighborhood(value: string): Apartment;

  getUnit(): string;
  setUnit(value: string): Apartment;

  getLat(): number;
  setLat(value: number): Apartment;

  getLng(): number;
  setLng(value: number): Apartment;

  getRent(): number;
  setRent(value: number): Apartment;

  getSqft(): number;
  setSqft(value: number): Apartment;

  getBeds(): number;
  setBeds(value: number): Apartment;

  getBaths(): number;
  setBaths(value: number): Apartment;

  getAvailableOn(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setAvailableOn(value?: google_protobuf_timestamp_pb.Timestamp): Apartment;
  hasAvailableOn(): boolean;
  clearAvailableOn(): Apartment;

  getDaysOnMarket(): number;
  setDaysOnMarket(value: number): Apartment;

  getDescription(): string;
  setDescription(value: string): Apartment;

  getAmenitiesList(): Array<string>;
  setAmenitiesList(value: Array<string>): Apartment;
  clearAmenitiesList(): Apartment;
  addAmenities(value: string, index?: number): Apartment;

  getUploadIdsList(): Array<string>;
  setUploadIdsList(value: Array<string>): Apartment;
  clearUploadIdsList(): Apartment;
  addUploadIds(value: string, index?: number): Apartment;

  getBuildingRef(): number;
  setBuildingRef(value: number): Apartment;

  getRealtorRef(): number;
  setRealtorRef(value: number): Apartment;

  getIsArchived(): boolean;
  setIsArchived(value: boolean): Apartment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Apartment.AsObject;
  static toObject(includeInstance: boolean, msg: Apartment): Apartment.AsObject;
  static serializeBinaryToWriter(message: Apartment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Apartment;
  static deserializeBinaryFromReader(message: Apartment, reader: jspb.BinaryReader): Apartment;
}

export namespace Apartment {
  export type AsObject = {
    id: number,
    name: string,
    fullAddress: string,
    street: string,
    city: string,
    state: string,
    zipCode: number,
    neighborhood: string,
    unit: string,
    lat: number,
    lng: number,
    rent: number,
    sqft: number,
    beds: number,
    baths: number,
    availableOn?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    daysOnMarket: number,
    description: string,
    amenitiesList: Array<string>,
    uploadIdsList: Array<string>,
    buildingRef: number,
    realtorRef: number,
    isArchived: boolean,
  }
}

export class CreateApartmentRequest extends jspb.Message {
  getApartment(): Apartment | undefined;
  setApartment(value?: Apartment): CreateApartmentRequest;
  hasApartment(): boolean;
  clearApartment(): CreateApartmentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateApartmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateApartmentRequest): CreateApartmentRequest.AsObject;
  static serializeBinaryToWriter(message: CreateApartmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateApartmentRequest;
  static deserializeBinaryFromReader(message: CreateApartmentRequest, reader: jspb.BinaryReader): CreateApartmentRequest;
}

export namespace CreateApartmentRequest {
  export type AsObject = {
    apartment?: Apartment.AsObject,
  }
}

export class GetApartmentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): GetApartmentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetApartmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetApartmentRequest): GetApartmentRequest.AsObject;
  static serializeBinaryToWriter(message: GetApartmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetApartmentRequest;
  static deserializeBinaryFromReader(message: GetApartmentRequest, reader: jspb.BinaryReader): GetApartmentRequest;
}

export namespace GetApartmentRequest {
  export type AsObject = {
    id: number,
  }
}

export class UpdateApartmentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): UpdateApartmentRequest;

  getApartment(): Apartment | undefined;
  setApartment(value?: Apartment): UpdateApartmentRequest;
  hasApartment(): boolean;
  clearApartment(): UpdateApartmentRequest;

  getMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateApartmentRequest;
  hasMask(): boolean;
  clearMask(): UpdateApartmentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateApartmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateApartmentRequest): UpdateApartmentRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateApartmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateApartmentRequest;
  static deserializeBinaryFromReader(message: UpdateApartmentRequest, reader: jspb.BinaryReader): UpdateApartmentRequest;
}

export namespace UpdateApartmentRequest {
  export type AsObject = {
    id: number,
    apartment?: Apartment.AsObject,
    mask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class DeleteApartmentRequest extends jspb.Message {
  getId(): number;
  setId(value: number): DeleteApartmentRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteApartmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteApartmentRequest): DeleteApartmentRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteApartmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteApartmentRequest;
  static deserializeBinaryFromReader(message: DeleteApartmentRequest, reader: jspb.BinaryReader): DeleteApartmentRequest;
}

export namespace DeleteApartmentRequest {
  export type AsObject = {
    id: number,
  }
}

export class ListApartmentRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListApartmentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListApartmentRequest): ListApartmentRequest.AsObject;
  static serializeBinaryToWriter(message: ListApartmentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListApartmentRequest;
  static deserializeBinaryFromReader(message: ListApartmentRequest, reader: jspb.BinaryReader): ListApartmentRequest;
}

export namespace ListApartmentRequest {
  export type AsObject = {
  }
}

export class ListApartmentResponse extends jspb.Message {
  getApartmentsList(): Array<Apartment>;
  setApartmentsList(value: Array<Apartment>): ListApartmentResponse;
  clearApartmentsList(): ListApartmentResponse;
  addApartments(value?: Apartment, index?: number): Apartment;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListApartmentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListApartmentResponse): ListApartmentResponse.AsObject;
  static serializeBinaryToWriter(message: ListApartmentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListApartmentResponse;
  static deserializeBinaryFromReader(message: ListApartmentResponse, reader: jspb.BinaryReader): ListApartmentResponse;
}

export namespace ListApartmentResponse {
  export type AsObject = {
    apartmentsList: Array<Apartment.AsObject>,
  }
}

export class DeleteApartmentResponse extends jspb.Message {
  getStatus(): content_pb.STATUS;
  setStatus(value: content_pb.STATUS): DeleteApartmentResponse;

  getApartment(): Apartment | undefined;
  setApartment(value?: Apartment): DeleteApartmentResponse;
  hasApartment(): boolean;
  clearApartment(): DeleteApartmentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteApartmentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteApartmentResponse): DeleteApartmentResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteApartmentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteApartmentResponse;
  static deserializeBinaryFromReader(message: DeleteApartmentResponse, reader: jspb.BinaryReader): DeleteApartmentResponse;
}

export namespace DeleteApartmentResponse {
  export type AsObject = {
    status: content_pb.STATUS,
    apartment?: Apartment.AsObject,
  }
}

export class GetNearbySchoolsRequest extends jspb.Message {
  getLat(): number;
  setLat(value: number): GetNearbySchoolsRequest;

  getLng(): number;
  setLng(value: number): GetNearbySchoolsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNearbySchoolsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetNearbySchoolsRequest): GetNearbySchoolsRequest.AsObject;
  static serializeBinaryToWriter(message: GetNearbySchoolsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNearbySchoolsRequest;
  static deserializeBinaryFromReader(message: GetNearbySchoolsRequest, reader: jspb.BinaryReader): GetNearbySchoolsRequest;
}

export namespace GetNearbySchoolsRequest {
  export type AsObject = {
    lat: number,
    lng: number,
  }
}

export class GetNearbySchoolsResponse extends jspb.Message {
  getResultsList(): Array<PlacesResult>;
  setResultsList(value: Array<PlacesResult>): GetNearbySchoolsResponse;
  clearResultsList(): GetNearbySchoolsResponse;
  addResults(value?: PlacesResult, index?: number): PlacesResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetNearbySchoolsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetNearbySchoolsResponse): GetNearbySchoolsResponse.AsObject;
  static serializeBinaryToWriter(message: GetNearbySchoolsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetNearbySchoolsResponse;
  static deserializeBinaryFromReader(message: GetNearbySchoolsResponse, reader: jspb.BinaryReader): GetNearbySchoolsResponse;
}

export namespace GetNearbySchoolsResponse {
  export type AsObject = {
    resultsList: Array<PlacesResult.AsObject>,
  }
}

export class PlacesResult extends jspb.Message {
  getFormattedAddress(): string;
  setFormattedAddress(value: string): PlacesResult;

  getGeometry(): Coordinates | undefined;
  setGeometry(value?: Coordinates): PlacesResult;
  hasGeometry(): boolean;
  clearGeometry(): PlacesResult;

  getName(): string;
  setName(value: string): PlacesResult;

  getTypesList(): Array<string>;
  setTypesList(value: Array<string>): PlacesResult;
  clearTypesList(): PlacesResult;
  addTypes(value: string, index?: number): PlacesResult;

  getPermanentlyClosed(): boolean;
  setPermanentlyClosed(value: boolean): PlacesResult;

  getBusinessStatus(): string;
  setBusinessStatus(value: string): PlacesResult;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PlacesResult.AsObject;
  static toObject(includeInstance: boolean, msg: PlacesResult): PlacesResult.AsObject;
  static serializeBinaryToWriter(message: PlacesResult, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PlacesResult;
  static deserializeBinaryFromReader(message: PlacesResult, reader: jspb.BinaryReader): PlacesResult;
}

export namespace PlacesResult {
  export type AsObject = {
    formattedAddress: string,
    geometry?: Coordinates.AsObject,
    name: string,
    typesList: Array<string>,
    permanentlyClosed: boolean,
    businessStatus: string,
  }
}

export class Coordinates extends jspb.Message {
  getLat(): number;
  setLat(value: number): Coordinates;

  getLng(): number;
  setLng(value: number): Coordinates;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Coordinates.AsObject;
  static toObject(includeInstance: boolean, msg: Coordinates): Coordinates.AsObject;
  static serializeBinaryToWriter(message: Coordinates, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Coordinates;
  static deserializeBinaryFromReader(message: Coordinates, reader: jspb.BinaryReader): Coordinates;
}

export namespace Coordinates {
  export type AsObject = {
    lat: number,
    lng: number,
  }
}

