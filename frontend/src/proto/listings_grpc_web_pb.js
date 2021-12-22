/**
 * @fileoverview gRPC-Web generated client stub for listings
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var apartment_pb = require('./apartment_pb.js')

var building_pb = require('./building_pb.js')

var realtor_pb = require('./realtor_pb.js')

var content_pb = require('./content_pb.js')
const proto = {};
proto.listings = require('./listings_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.listings.ListingsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.listings.ListingsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.CreateApartmentRequest,
 *   !proto.listings.Apartment>}
 */
const methodDescriptor_Listings_CreateApartment = new grpc.web.MethodDescriptor(
  '/listings.Listings/CreateApartment',
  grpc.web.MethodType.UNARY,
  apartment_pb.CreateApartmentRequest,
  apartment_pb.Apartment,
  /**
   * @param {!proto.listings.CreateApartmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.Apartment.deserializeBinary
);


/**
 * @param {!proto.listings.CreateApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Apartment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Apartment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.createApartment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/CreateApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateApartment,
      callback);
};


/**
 * @param {!proto.listings.CreateApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Apartment>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.createApartment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/CreateApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateApartment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.GetApartmentRequest,
 *   !proto.listings.Apartment>}
 */
const methodDescriptor_Listings_GetApartment = new grpc.web.MethodDescriptor(
  '/listings.Listings/GetApartment',
  grpc.web.MethodType.UNARY,
  apartment_pb.GetApartmentRequest,
  apartment_pb.Apartment,
  /**
   * @param {!proto.listings.GetApartmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.Apartment.deserializeBinary
);


/**
 * @param {!proto.listings.GetApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Apartment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Apartment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.getApartment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/GetApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_GetApartment,
      callback);
};


/**
 * @param {!proto.listings.GetApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Apartment>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.getApartment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/GetApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_GetApartment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.ListApartmentRequest,
 *   !proto.listings.ListApartmentResponse>}
 */
const methodDescriptor_Listings_ListApartments = new grpc.web.MethodDescriptor(
  '/listings.Listings/ListApartments',
  grpc.web.MethodType.UNARY,
  apartment_pb.ListApartmentRequest,
  apartment_pb.ListApartmentResponse,
  /**
   * @param {!proto.listings.ListApartmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.ListApartmentResponse.deserializeBinary
);


/**
 * @param {!proto.listings.ListApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.ListApartmentResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.ListApartmentResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.listApartments =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/ListApartments',
      request,
      metadata || {},
      methodDescriptor_Listings_ListApartments,
      callback);
};


/**
 * @param {!proto.listings.ListApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.ListApartmentResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.listApartments =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/ListApartments',
      request,
      metadata || {},
      methodDescriptor_Listings_ListApartments);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.UpdateApartmentRequest,
 *   !proto.listings.Apartment>}
 */
const methodDescriptor_Listings_UpdateApartment = new grpc.web.MethodDescriptor(
  '/listings.Listings/UpdateApartment',
  grpc.web.MethodType.UNARY,
  apartment_pb.UpdateApartmentRequest,
  apartment_pb.Apartment,
  /**
   * @param {!proto.listings.UpdateApartmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.Apartment.deserializeBinary
);


/**
 * @param {!proto.listings.UpdateApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Apartment)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Apartment>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.updateApartment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/UpdateApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateApartment,
      callback);
};


/**
 * @param {!proto.listings.UpdateApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Apartment>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.updateApartment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/UpdateApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateApartment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.DeleteApartmentRequest,
 *   !proto.listings.DeleteApartmentResponse>}
 */
const methodDescriptor_Listings_DeleteApartment = new grpc.web.MethodDescriptor(
  '/listings.Listings/DeleteApartment',
  grpc.web.MethodType.UNARY,
  apartment_pb.DeleteApartmentRequest,
  apartment_pb.DeleteApartmentResponse,
  /**
   * @param {!proto.listings.DeleteApartmentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.DeleteApartmentResponse.deserializeBinary
);


/**
 * @param {!proto.listings.DeleteApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.DeleteApartmentResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.DeleteApartmentResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.deleteApartment =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/DeleteApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteApartment,
      callback);
};


/**
 * @param {!proto.listings.DeleteApartmentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.DeleteApartmentResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.deleteApartment =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/DeleteApartment',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteApartment);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.GetNearbySchoolsRequest,
 *   !proto.listings.GetNearbySchoolsResponse>}
 */
const methodDescriptor_Listings_GetNearbySchools = new grpc.web.MethodDescriptor(
  '/listings.Listings/GetNearbySchools',
  grpc.web.MethodType.UNARY,
  apartment_pb.GetNearbySchoolsRequest,
  apartment_pb.GetNearbySchoolsResponse,
  /**
   * @param {!proto.listings.GetNearbySchoolsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  apartment_pb.GetNearbySchoolsResponse.deserializeBinary
);


/**
 * @param {!proto.listings.GetNearbySchoolsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.GetNearbySchoolsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.GetNearbySchoolsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.getNearbySchools =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/GetNearbySchools',
      request,
      metadata || {},
      methodDescriptor_Listings_GetNearbySchools,
      callback);
};


/**
 * @param {!proto.listings.GetNearbySchoolsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.GetNearbySchoolsResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.getNearbySchools =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/GetNearbySchools',
      request,
      metadata || {},
      methodDescriptor_Listings_GetNearbySchools);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.CreateBuildingRequest,
 *   !proto.listings.Building>}
 */
const methodDescriptor_Listings_CreateBuilding = new grpc.web.MethodDescriptor(
  '/listings.Listings/CreateBuilding',
  grpc.web.MethodType.UNARY,
  building_pb.CreateBuildingRequest,
  building_pb.Building,
  /**
   * @param {!proto.listings.CreateBuildingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  building_pb.Building.deserializeBinary
);


/**
 * @param {!proto.listings.CreateBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Building)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Building>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.createBuilding =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/CreateBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateBuilding,
      callback);
};


/**
 * @param {!proto.listings.CreateBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Building>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.createBuilding =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/CreateBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateBuilding);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.GetBuildingRequest,
 *   !proto.listings.Building>}
 */
const methodDescriptor_Listings_GetBuilding = new grpc.web.MethodDescriptor(
  '/listings.Listings/GetBuilding',
  grpc.web.MethodType.UNARY,
  building_pb.GetBuildingRequest,
  building_pb.Building,
  /**
   * @param {!proto.listings.GetBuildingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  building_pb.Building.deserializeBinary
);


/**
 * @param {!proto.listings.GetBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Building)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Building>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.getBuilding =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/GetBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_GetBuilding,
      callback);
};


/**
 * @param {!proto.listings.GetBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Building>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.getBuilding =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/GetBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_GetBuilding);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.ListBuildingRequest,
 *   !proto.listings.ListBuildingResponse>}
 */
const methodDescriptor_Listings_ListBuildings = new grpc.web.MethodDescriptor(
  '/listings.Listings/ListBuildings',
  grpc.web.MethodType.UNARY,
  building_pb.ListBuildingRequest,
  building_pb.ListBuildingResponse,
  /**
   * @param {!proto.listings.ListBuildingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  building_pb.ListBuildingResponse.deserializeBinary
);


/**
 * @param {!proto.listings.ListBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.ListBuildingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.ListBuildingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.listBuildings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/ListBuildings',
      request,
      metadata || {},
      methodDescriptor_Listings_ListBuildings,
      callback);
};


/**
 * @param {!proto.listings.ListBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.ListBuildingResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.listBuildings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/ListBuildings',
      request,
      metadata || {},
      methodDescriptor_Listings_ListBuildings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.UpdateBuildingRequest,
 *   !proto.listings.Building>}
 */
const methodDescriptor_Listings_UpdateBuilding = new grpc.web.MethodDescriptor(
  '/listings.Listings/UpdateBuilding',
  grpc.web.MethodType.UNARY,
  building_pb.UpdateBuildingRequest,
  building_pb.Building,
  /**
   * @param {!proto.listings.UpdateBuildingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  building_pb.Building.deserializeBinary
);


/**
 * @param {!proto.listings.UpdateBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Building)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Building>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.updateBuilding =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/UpdateBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateBuilding,
      callback);
};


/**
 * @param {!proto.listings.UpdateBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Building>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.updateBuilding =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/UpdateBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateBuilding);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.DeleteBuildingRequest,
 *   !proto.listings.DeleteBuildingResponse>}
 */
const methodDescriptor_Listings_DeleteBuilding = new grpc.web.MethodDescriptor(
  '/listings.Listings/DeleteBuilding',
  grpc.web.MethodType.UNARY,
  building_pb.DeleteBuildingRequest,
  building_pb.DeleteBuildingResponse,
  /**
   * @param {!proto.listings.DeleteBuildingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  building_pb.DeleteBuildingResponse.deserializeBinary
);


/**
 * @param {!proto.listings.DeleteBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.DeleteBuildingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.DeleteBuildingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.deleteBuilding =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/DeleteBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteBuilding,
      callback);
};


/**
 * @param {!proto.listings.DeleteBuildingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.DeleteBuildingResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.deleteBuilding =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/DeleteBuilding',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteBuilding);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.CreateRealtorRequest,
 *   !proto.listings.Realtor>}
 */
const methodDescriptor_Listings_CreateRealtor = new grpc.web.MethodDescriptor(
  '/listings.Listings/CreateRealtor',
  grpc.web.MethodType.UNARY,
  realtor_pb.CreateRealtorRequest,
  realtor_pb.Realtor,
  /**
   * @param {!proto.listings.CreateRealtorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  realtor_pb.Realtor.deserializeBinary
);


/**
 * @param {!proto.listings.CreateRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Realtor)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Realtor>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.createRealtor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/CreateRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateRealtor,
      callback);
};


/**
 * @param {!proto.listings.CreateRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Realtor>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.createRealtor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/CreateRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_CreateRealtor);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.GetRealtorRequest,
 *   !proto.listings.Realtor>}
 */
const methodDescriptor_Listings_GetRealtor = new grpc.web.MethodDescriptor(
  '/listings.Listings/GetRealtor',
  grpc.web.MethodType.UNARY,
  realtor_pb.GetRealtorRequest,
  realtor_pb.Realtor,
  /**
   * @param {!proto.listings.GetRealtorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  realtor_pb.Realtor.deserializeBinary
);


/**
 * @param {!proto.listings.GetRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Realtor)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Realtor>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.getRealtor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/GetRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_GetRealtor,
      callback);
};


/**
 * @param {!proto.listings.GetRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Realtor>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.getRealtor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/GetRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_GetRealtor);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.ListRealtorRequest,
 *   !proto.listings.ListRealtorResponse>}
 */
const methodDescriptor_Listings_ListRealtors = new grpc.web.MethodDescriptor(
  '/listings.Listings/ListRealtors',
  grpc.web.MethodType.UNARY,
  realtor_pb.ListRealtorRequest,
  realtor_pb.ListRealtorResponse,
  /**
   * @param {!proto.listings.ListRealtorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  realtor_pb.ListRealtorResponse.deserializeBinary
);


/**
 * @param {!proto.listings.ListRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.ListRealtorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.ListRealtorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.listRealtors =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/ListRealtors',
      request,
      metadata || {},
      methodDescriptor_Listings_ListRealtors,
      callback);
};


/**
 * @param {!proto.listings.ListRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.ListRealtorResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.listRealtors =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/ListRealtors',
      request,
      metadata || {},
      methodDescriptor_Listings_ListRealtors);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.UpdateRealtorRequest,
 *   !proto.listings.Realtor>}
 */
const methodDescriptor_Listings_UpdateRealtor = new grpc.web.MethodDescriptor(
  '/listings.Listings/UpdateRealtor',
  grpc.web.MethodType.UNARY,
  realtor_pb.UpdateRealtorRequest,
  realtor_pb.Realtor,
  /**
   * @param {!proto.listings.UpdateRealtorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  realtor_pb.Realtor.deserializeBinary
);


/**
 * @param {!proto.listings.UpdateRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.Realtor)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.Realtor>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.updateRealtor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/UpdateRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateRealtor,
      callback);
};


/**
 * @param {!proto.listings.UpdateRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.Realtor>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.updateRealtor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/UpdateRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_UpdateRealtor);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.listings.DeleteRealtorRequest,
 *   !proto.listings.DeleteRealtorResponse>}
 */
const methodDescriptor_Listings_DeleteRealtor = new grpc.web.MethodDescriptor(
  '/listings.Listings/DeleteRealtor',
  grpc.web.MethodType.UNARY,
  realtor_pb.DeleteRealtorRequest,
  realtor_pb.DeleteRealtorResponse,
  /**
   * @param {!proto.listings.DeleteRealtorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  realtor_pb.DeleteRealtorResponse.deserializeBinary
);


/**
 * @param {!proto.listings.DeleteRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.listings.DeleteRealtorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.listings.DeleteRealtorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.listings.ListingsClient.prototype.deleteRealtor =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/listings.Listings/DeleteRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteRealtor,
      callback);
};


/**
 * @param {!proto.listings.DeleteRealtorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.listings.DeleteRealtorResponse>}
 *     Promise that resolves to the response
 */
proto.listings.ListingsPromiseClient.prototype.deleteRealtor =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/listings.Listings/DeleteRealtor',
      request,
      metadata || {},
      methodDescriptor_Listings_DeleteRealtor);
};


module.exports = proto.listings;

