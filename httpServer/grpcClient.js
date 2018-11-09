const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const GRPC_PORT = 8383;
const SERVICE_PROTO_PATH = __dirname + '/../protobuff/service.proto';

console.log(`Loading proto from ${SERVICE_PROTO_PATH}`);

const packageDefinition = protoLoader.loadSync(
    SERVICE_PROTO_PATH,
    {keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const service = grpc.loadPackageDefinition(packageDefinition).service;
const client = new service.OCRService(`localhost:${GRPC_PORT}`, grpc.credentials.createInsecure());

module.exports = client;