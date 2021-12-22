import * as jspb from 'google-protobuf'



export class UploadContentRequest extends jspb.Message {
  getInfo(): ContentInfo | undefined;
  setInfo(value?: ContentInfo): UploadContentRequest;
  hasInfo(): boolean;
  clearInfo(): UploadContentRequest;

  getChunkData(): Uint8Array | string;
  getChunkData_asU8(): Uint8Array;
  getChunkData_asB64(): string;
  setChunkData(value: Uint8Array | string): UploadContentRequest;

  getDataCase(): UploadContentRequest.DataCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadContentRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UploadContentRequest): UploadContentRequest.AsObject;
  static serializeBinaryToWriter(message: UploadContentRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadContentRequest;
  static deserializeBinaryFromReader(message: UploadContentRequest, reader: jspb.BinaryReader): UploadContentRequest;
}

export namespace UploadContentRequest {
  export type AsObject = {
    info?: ContentInfo.AsObject,
    chunkData: Uint8Array | string,
  }

  export enum DataCase { 
    DATA_NOT_SET = 0,
    INFO = 1,
    CHUNK_DATA = 2,
  }
}

export class ContentInfo extends jspb.Message {
  getSourceId(): number;
  setSourceId(value: number): ContentInfo;

  getContentType(): ContentInfo.ContentType;
  setContentType(value: ContentInfo.ContentType): ContentInfo;

  getContentSource(): ContentInfo.ContentSource;
  setContentSource(value: ContentInfo.ContentSource): ContentInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ContentInfo.AsObject;
  static toObject(includeInstance: boolean, msg: ContentInfo): ContentInfo.AsObject;
  static serializeBinaryToWriter(message: ContentInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ContentInfo;
  static deserializeBinaryFromReader(message: ContentInfo, reader: jspb.BinaryReader): ContentInfo;
}

export namespace ContentInfo {
  export type AsObject = {
    sourceId: number,
    contentType: ContentInfo.ContentType,
    contentSource: ContentInfo.ContentSource,
  }

  export enum ContentType { 
    TYPE_UNKNOWN = 0,
    IMAGE = 1,
    VIDEO = 2,
  }

  export enum ContentSource { 
    SOURCE_UNKNOWN = 0,
    APARTMENT = 1,
    BUILDING = 2,
  }
}

export class UploadContentResponse extends jspb.Message {
  getContentId(): number;
  setContentId(value: number): UploadContentResponse;

  getSize(): number;
  setSize(value: number): UploadContentResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UploadContentResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UploadContentResponse): UploadContentResponse.AsObject;
  static serializeBinaryToWriter(message: UploadContentResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UploadContentResponse;
  static deserializeBinaryFromReader(message: UploadContentResponse, reader: jspb.BinaryReader): UploadContentResponse;
}

export namespace UploadContentResponse {
  export type AsObject = {
    contentId: number,
    size: number,
  }
}

export enum STATUS { 
  STATUS_UNKNOWN = 0,
  SUCCESS = 1,
  FAIL = 2,
}
