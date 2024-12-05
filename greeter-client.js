// run with: deno run -A greeter-client.js
import * as grpc from "npm:@grpc/grpc-js";
import { promisify } from "node:util";
import * as protoLoader from "npm:@grpc/proto-loader";
import * as path from "node:path";

const PROTO_PATH = "./contract/greeter.proto";
const PORT = "8080";
const ADDR = "127.0.0.1";

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const accessProto = grpc.loadPackageDefinition(packageDefinition).mypackage;
const target = `${ADDR}:${PORT}`;
const structureAccessClient = new accessProto.Greeter(
  target,
  grpc.credentials.createInsecure(),
);

structureAccessClient.SayHello({ name: "World" }, function (err, resp) {
  if (err) {
    console.error("Error", err);
  }

  console.log(resp);
});
