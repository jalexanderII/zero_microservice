/**
 * @fileoverview gRPC-Web generated client stub for listings
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as apartment_pb from './apartment_pb';
import * as building_pb from './building_pb';
import * as realtor_pb from './realtor_pb';
import * as content_pb from './content_pb';


export class ListingsClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoCreateApartment = new grpcWeb.MethodDescriptor(
    '/listings.Listings/CreateApartment',
    grpcWeb.MethodType.UNARY,
    apartment_pb.CreateApartmentRequest,
    apartment_pb.Apartment,
    (request: apartment_pb.CreateApartmentRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.Apartment.deserializeBinary
  );

  createApartment(
    request: apartment_pb.CreateApartmentRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.Apartment>;

  createApartment(
    request: apartment_pb.CreateApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void): grpcWeb.ClientReadableStream<apartment_pb.Apartment>;

  createApartment(
    request: apartment_pb.CreateApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/CreateApartment',
        request,
        metadata || {},
        this.methodInfoCreateApartment,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/CreateApartment',
    request,
    metadata || {},
    this.methodInfoCreateApartment);
  }

  methodInfoGetApartment = new grpcWeb.MethodDescriptor(
    '/listings.Listings/GetApartment',
    grpcWeb.MethodType.UNARY,
    apartment_pb.GetApartmentRequest,
    apartment_pb.Apartment,
    (request: apartment_pb.GetApartmentRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.Apartment.deserializeBinary
  );

  getApartment(
    request: apartment_pb.GetApartmentRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.Apartment>;

  getApartment(
    request: apartment_pb.GetApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void): grpcWeb.ClientReadableStream<apartment_pb.Apartment>;

  getApartment(
    request: apartment_pb.GetApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/GetApartment',
        request,
        metadata || {},
        this.methodInfoGetApartment,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/GetApartment',
    request,
    metadata || {},
    this.methodInfoGetApartment);
  }

  methodInfoListApartments = new grpcWeb.MethodDescriptor(
    '/listings.Listings/ListApartments',
    grpcWeb.MethodType.UNARY,
    apartment_pb.ListApartmentRequest,
    apartment_pb.ListApartmentResponse,
    (request: apartment_pb.ListApartmentRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.ListApartmentResponse.deserializeBinary
  );

  listApartments(
    request: apartment_pb.ListApartmentRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.ListApartmentResponse>;

  listApartments(
    request: apartment_pb.ListApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.ListApartmentResponse) => void): grpcWeb.ClientReadableStream<apartment_pb.ListApartmentResponse>;

  listApartments(
    request: apartment_pb.ListApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.ListApartmentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/ListApartments',
        request,
        metadata || {},
        this.methodInfoListApartments,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/ListApartments',
    request,
    metadata || {},
    this.methodInfoListApartments);
  }

  methodInfoUpdateApartment = new grpcWeb.MethodDescriptor(
    '/listings.Listings/UpdateApartment',
    grpcWeb.MethodType.UNARY,
    apartment_pb.UpdateApartmentRequest,
    apartment_pb.Apartment,
    (request: apartment_pb.UpdateApartmentRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.Apartment.deserializeBinary
  );

  updateApartment(
    request: apartment_pb.UpdateApartmentRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.Apartment>;

  updateApartment(
    request: apartment_pb.UpdateApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void): grpcWeb.ClientReadableStream<apartment_pb.Apartment>;

  updateApartment(
    request: apartment_pb.UpdateApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.Apartment) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/UpdateApartment',
        request,
        metadata || {},
        this.methodInfoUpdateApartment,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/UpdateApartment',
    request,
    metadata || {},
    this.methodInfoUpdateApartment);
  }

  methodInfoDeleteApartment = new grpcWeb.MethodDescriptor(
    '/listings.Listings/DeleteApartment',
    grpcWeb.MethodType.UNARY,
    apartment_pb.DeleteApartmentRequest,
    apartment_pb.DeleteApartmentResponse,
    (request: apartment_pb.DeleteApartmentRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.DeleteApartmentResponse.deserializeBinary
  );

  deleteApartment(
    request: apartment_pb.DeleteApartmentRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.DeleteApartmentResponse>;

  deleteApartment(
    request: apartment_pb.DeleteApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.DeleteApartmentResponse) => void): grpcWeb.ClientReadableStream<apartment_pb.DeleteApartmentResponse>;

  deleteApartment(
    request: apartment_pb.DeleteApartmentRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.DeleteApartmentResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/DeleteApartment',
        request,
        metadata || {},
        this.methodInfoDeleteApartment,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/DeleteApartment',
    request,
    metadata || {},
    this.methodInfoDeleteApartment);
  }

  methodInfoGetNearbySchools = new grpcWeb.MethodDescriptor(
    '/listings.Listings/GetNearbySchools',
    grpcWeb.MethodType.UNARY,
    apartment_pb.GetNearbySchoolsRequest,
    apartment_pb.GetNearbySchoolsResponse,
    (request: apartment_pb.GetNearbySchoolsRequest) => {
      return request.serializeBinary();
    },
    apartment_pb.GetNearbySchoolsResponse.deserializeBinary
  );

  getNearbySchools(
    request: apartment_pb.GetNearbySchoolsRequest,
    metadata: grpcWeb.Metadata | null): Promise<apartment_pb.GetNearbySchoolsResponse>;

  getNearbySchools(
    request: apartment_pb.GetNearbySchoolsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: apartment_pb.GetNearbySchoolsResponse) => void): grpcWeb.ClientReadableStream<apartment_pb.GetNearbySchoolsResponse>;

  getNearbySchools(
    request: apartment_pb.GetNearbySchoolsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: apartment_pb.GetNearbySchoolsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/GetNearbySchools',
        request,
        metadata || {},
        this.methodInfoGetNearbySchools,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/GetNearbySchools',
    request,
    metadata || {},
    this.methodInfoGetNearbySchools);
  }

  methodInfoCreateBuilding = new grpcWeb.MethodDescriptor(
    '/listings.Listings/CreateBuilding',
    grpcWeb.MethodType.UNARY,
    building_pb.CreateBuildingRequest,
    building_pb.Building,
    (request: building_pb.CreateBuildingRequest) => {
      return request.serializeBinary();
    },
    building_pb.Building.deserializeBinary
  );

  createBuilding(
    request: building_pb.CreateBuildingRequest,
    metadata: grpcWeb.Metadata | null): Promise<building_pb.Building>;

  createBuilding(
    request: building_pb.CreateBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void): grpcWeb.ClientReadableStream<building_pb.Building>;

  createBuilding(
    request: building_pb.CreateBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/CreateBuilding',
        request,
        metadata || {},
        this.methodInfoCreateBuilding,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/CreateBuilding',
    request,
    metadata || {},
    this.methodInfoCreateBuilding);
  }

  methodInfoGetBuilding = new grpcWeb.MethodDescriptor(
    '/listings.Listings/GetBuilding',
    grpcWeb.MethodType.UNARY,
    building_pb.GetBuildingRequest,
    building_pb.Building,
    (request: building_pb.GetBuildingRequest) => {
      return request.serializeBinary();
    },
    building_pb.Building.deserializeBinary
  );

  getBuilding(
    request: building_pb.GetBuildingRequest,
    metadata: grpcWeb.Metadata | null): Promise<building_pb.Building>;

  getBuilding(
    request: building_pb.GetBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void): grpcWeb.ClientReadableStream<building_pb.Building>;

  getBuilding(
    request: building_pb.GetBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/GetBuilding',
        request,
        metadata || {},
        this.methodInfoGetBuilding,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/GetBuilding',
    request,
    metadata || {},
    this.methodInfoGetBuilding);
  }

  methodInfoListBuildings = new grpcWeb.MethodDescriptor(
    '/listings.Listings/ListBuildings',
    grpcWeb.MethodType.UNARY,
    building_pb.ListBuildingRequest,
    building_pb.ListBuildingResponse,
    (request: building_pb.ListBuildingRequest) => {
      return request.serializeBinary();
    },
    building_pb.ListBuildingResponse.deserializeBinary
  );

  listBuildings(
    request: building_pb.ListBuildingRequest,
    metadata: grpcWeb.Metadata | null): Promise<building_pb.ListBuildingResponse>;

  listBuildings(
    request: building_pb.ListBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: building_pb.ListBuildingResponse) => void): grpcWeb.ClientReadableStream<building_pb.ListBuildingResponse>;

  listBuildings(
    request: building_pb.ListBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: building_pb.ListBuildingResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/ListBuildings',
        request,
        metadata || {},
        this.methodInfoListBuildings,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/ListBuildings',
    request,
    metadata || {},
    this.methodInfoListBuildings);
  }

  methodInfoUpdateBuilding = new grpcWeb.MethodDescriptor(
    '/listings.Listings/UpdateBuilding',
    grpcWeb.MethodType.UNARY,
    building_pb.UpdateBuildingRequest,
    building_pb.Building,
    (request: building_pb.UpdateBuildingRequest) => {
      return request.serializeBinary();
    },
    building_pb.Building.deserializeBinary
  );

  updateBuilding(
    request: building_pb.UpdateBuildingRequest,
    metadata: grpcWeb.Metadata | null): Promise<building_pb.Building>;

  updateBuilding(
    request: building_pb.UpdateBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void): grpcWeb.ClientReadableStream<building_pb.Building>;

  updateBuilding(
    request: building_pb.UpdateBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: building_pb.Building) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/UpdateBuilding',
        request,
        metadata || {},
        this.methodInfoUpdateBuilding,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/UpdateBuilding',
    request,
    metadata || {},
    this.methodInfoUpdateBuilding);
  }

  methodInfoDeleteBuilding = new grpcWeb.MethodDescriptor(
    '/listings.Listings/DeleteBuilding',
    grpcWeb.MethodType.UNARY,
    building_pb.DeleteBuildingRequest,
    building_pb.DeleteBuildingResponse,
    (request: building_pb.DeleteBuildingRequest) => {
      return request.serializeBinary();
    },
    building_pb.DeleteBuildingResponse.deserializeBinary
  );

  deleteBuilding(
    request: building_pb.DeleteBuildingRequest,
    metadata: grpcWeb.Metadata | null): Promise<building_pb.DeleteBuildingResponse>;

  deleteBuilding(
    request: building_pb.DeleteBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: building_pb.DeleteBuildingResponse) => void): grpcWeb.ClientReadableStream<building_pb.DeleteBuildingResponse>;

  deleteBuilding(
    request: building_pb.DeleteBuildingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: building_pb.DeleteBuildingResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/DeleteBuilding',
        request,
        metadata || {},
        this.methodInfoDeleteBuilding,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/DeleteBuilding',
    request,
    metadata || {},
    this.methodInfoDeleteBuilding);
  }

  methodInfoCreateRealtor = new grpcWeb.MethodDescriptor(
    '/listings.Listings/CreateRealtor',
    grpcWeb.MethodType.UNARY,
    realtor_pb.CreateRealtorRequest,
    realtor_pb.Realtor,
    (request: realtor_pb.CreateRealtorRequest) => {
      return request.serializeBinary();
    },
    realtor_pb.Realtor.deserializeBinary
  );

  createRealtor(
    request: realtor_pb.CreateRealtorRequest,
    metadata: grpcWeb.Metadata | null): Promise<realtor_pb.Realtor>;

  createRealtor(
    request: realtor_pb.CreateRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void): grpcWeb.ClientReadableStream<realtor_pb.Realtor>;

  createRealtor(
    request: realtor_pb.CreateRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/CreateRealtor',
        request,
        metadata || {},
        this.methodInfoCreateRealtor,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/CreateRealtor',
    request,
    metadata || {},
    this.methodInfoCreateRealtor);
  }

  methodInfoGetRealtor = new grpcWeb.MethodDescriptor(
    '/listings.Listings/GetRealtor',
    grpcWeb.MethodType.UNARY,
    realtor_pb.GetRealtorRequest,
    realtor_pb.Realtor,
    (request: realtor_pb.GetRealtorRequest) => {
      return request.serializeBinary();
    },
    realtor_pb.Realtor.deserializeBinary
  );

  getRealtor(
    request: realtor_pb.GetRealtorRequest,
    metadata: grpcWeb.Metadata | null): Promise<realtor_pb.Realtor>;

  getRealtor(
    request: realtor_pb.GetRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void): grpcWeb.ClientReadableStream<realtor_pb.Realtor>;

  getRealtor(
    request: realtor_pb.GetRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/GetRealtor',
        request,
        metadata || {},
        this.methodInfoGetRealtor,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/GetRealtor',
    request,
    metadata || {},
    this.methodInfoGetRealtor);
  }

  methodInfoListRealtors = new grpcWeb.MethodDescriptor(
    '/listings.Listings/ListRealtors',
    grpcWeb.MethodType.UNARY,
    realtor_pb.ListRealtorRequest,
    realtor_pb.ListRealtorResponse,
    (request: realtor_pb.ListRealtorRequest) => {
      return request.serializeBinary();
    },
    realtor_pb.ListRealtorResponse.deserializeBinary
  );

  listRealtors(
    request: realtor_pb.ListRealtorRequest,
    metadata: grpcWeb.Metadata | null): Promise<realtor_pb.ListRealtorResponse>;

  listRealtors(
    request: realtor_pb.ListRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: realtor_pb.ListRealtorResponse) => void): grpcWeb.ClientReadableStream<realtor_pb.ListRealtorResponse>;

  listRealtors(
    request: realtor_pb.ListRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: realtor_pb.ListRealtorResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/ListRealtors',
        request,
        metadata || {},
        this.methodInfoListRealtors,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/ListRealtors',
    request,
    metadata || {},
    this.methodInfoListRealtors);
  }

  methodInfoUpdateRealtor = new grpcWeb.MethodDescriptor(
    '/listings.Listings/UpdateRealtor',
    grpcWeb.MethodType.UNARY,
    realtor_pb.UpdateRealtorRequest,
    realtor_pb.Realtor,
    (request: realtor_pb.UpdateRealtorRequest) => {
      return request.serializeBinary();
    },
    realtor_pb.Realtor.deserializeBinary
  );

  updateRealtor(
    request: realtor_pb.UpdateRealtorRequest,
    metadata: grpcWeb.Metadata | null): Promise<realtor_pb.Realtor>;

  updateRealtor(
    request: realtor_pb.UpdateRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void): grpcWeb.ClientReadableStream<realtor_pb.Realtor>;

  updateRealtor(
    request: realtor_pb.UpdateRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: realtor_pb.Realtor) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/UpdateRealtor',
        request,
        metadata || {},
        this.methodInfoUpdateRealtor,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/UpdateRealtor',
    request,
    metadata || {},
    this.methodInfoUpdateRealtor);
  }

  methodInfoDeleteRealtor = new grpcWeb.MethodDescriptor(
    '/listings.Listings/DeleteRealtor',
    grpcWeb.MethodType.UNARY,
    realtor_pb.DeleteRealtorRequest,
    realtor_pb.DeleteRealtorResponse,
    (request: realtor_pb.DeleteRealtorRequest) => {
      return request.serializeBinary();
    },
    realtor_pb.DeleteRealtorResponse.deserializeBinary
  );

  deleteRealtor(
    request: realtor_pb.DeleteRealtorRequest,
    metadata: grpcWeb.Metadata | null): Promise<realtor_pb.DeleteRealtorResponse>;

  deleteRealtor(
    request: realtor_pb.DeleteRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: realtor_pb.DeleteRealtorResponse) => void): grpcWeb.ClientReadableStream<realtor_pb.DeleteRealtorResponse>;

  deleteRealtor(
    request: realtor_pb.DeleteRealtorRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: realtor_pb.DeleteRealtorResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/listings.Listings/DeleteRealtor',
        request,
        metadata || {},
        this.methodInfoDeleteRealtor,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/listings.Listings/DeleteRealtor',
    request,
    metadata || {},
    this.methodInfoDeleteRealtor);
  }

}

