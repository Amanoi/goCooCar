import * as $protobuf from "protobufjs";

// Common aliases
const $Reader = $protobuf.Reader, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const auth = $root.auth = (() => {

    /**
     * Namespace auth.
     * @exports auth
     * @namespace
     */
    const auth = {};

    auth.v1 = (function() {

        /**
         * Namespace v1.
         * @memberof auth
         * @namespace
         */
        const v1 = {};

        v1.LoginResquest = (function() {

            /**
             * Properties of a LoginResquest.
             * @memberof auth.v1
             * @interface ILoginResquest
             * @property {string|null} [code] LoginResquest code
             */

            /**
             * Constructs a new LoginResquest.
             * @memberof auth.v1
             * @classdesc Represents a LoginResquest.
             * @implements ILoginResquest
             * @constructor
             * @param {auth.v1.ILoginResquest=} [properties] Properties to set
             */
            function LoginResquest(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * LoginResquest code.
             * @member {string} code
             * @memberof auth.v1.LoginResquest
             * @instance
             */
            LoginResquest.prototype.code = "";

            /**
             * Creates a new LoginResquest instance using the specified properties.
             * @function create
             * @memberof auth.v1.LoginResquest
             * @static
             * @param {auth.v1.ILoginResquest=} [properties] Properties to set
             * @returns {auth.v1.LoginResquest} LoginResquest instance
             */
            LoginResquest.create = function create(properties) {
                return new LoginResquest(properties);
            };

            /**
             * Decodes a LoginResquest message from the specified reader or buffer.
             * @function decode
             * @memberof auth.v1.LoginResquest
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {auth.v1.LoginResquest} LoginResquest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LoginResquest.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.auth.v1.LoginResquest();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.code = reader.string();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Creates a LoginResquest message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof auth.v1.LoginResquest
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {auth.v1.LoginResquest} LoginResquest
             */
            LoginResquest.fromObject = function fromObject(object) {
                if (object instanceof $root.auth.v1.LoginResquest)
                    return object;
                let message = new $root.auth.v1.LoginResquest();
                if (object.code != null)
                    message.code = String(object.code);
                return message;
            };

            /**
             * Creates a plain object from a LoginResquest message. Also converts values to other types if specified.
             * @function toObject
             * @memberof auth.v1.LoginResquest
             * @static
             * @param {auth.v1.LoginResquest} message LoginResquest
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            LoginResquest.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults)
                    object.code = "";
                if (message.code != null && message.hasOwnProperty("code"))
                    object.code = message.code;
                return object;
            };

            /**
             * Converts this LoginResquest to JSON.
             * @function toJSON
             * @memberof auth.v1.LoginResquest
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            LoginResquest.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return LoginResquest;
        })();

        v1.LoginResponse = (function() {

            /**
             * Properties of a LoginResponse.
             * @memberof auth.v1
             * @interface ILoginResponse
             * @property {string|null} [accessToken] LoginResponse accessToken
             * @property {number|null} [expiresIn] LoginResponse expiresIn
             */

            /**
             * Constructs a new LoginResponse.
             * @memberof auth.v1
             * @classdesc Represents a LoginResponse.
             * @implements ILoginResponse
             * @constructor
             * @param {auth.v1.ILoginResponse=} [properties] Properties to set
             */
            function LoginResponse(properties) {
                if (properties)
                    for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * LoginResponse accessToken.
             * @member {string} accessToken
             * @memberof auth.v1.LoginResponse
             * @instance
             */
            LoginResponse.prototype.accessToken = "";

            /**
             * LoginResponse expiresIn.
             * @member {number} expiresIn
             * @memberof auth.v1.LoginResponse
             * @instance
             */
            LoginResponse.prototype.expiresIn = 0;

            /**
             * Creates a new LoginResponse instance using the specified properties.
             * @function create
             * @memberof auth.v1.LoginResponse
             * @static
             * @param {auth.v1.ILoginResponse=} [properties] Properties to set
             * @returns {auth.v1.LoginResponse} LoginResponse instance
             */
            LoginResponse.create = function create(properties) {
                return new LoginResponse(properties);
            };

            /**
             * Decodes a LoginResponse message from the specified reader or buffer.
             * @function decode
             * @memberof auth.v1.LoginResponse
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {auth.v1.LoginResponse} LoginResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            LoginResponse.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                let end = length === undefined ? reader.len : reader.pos + length, message = new $root.auth.v1.LoginResponse();
                while (reader.pos < end) {
                    let tag = reader.uint32();
                    switch (tag >>> 3) {
                    case 1:
                        message.accessToken = reader.string();
                        break;
                    case 2:
                        message.expiresIn = reader.int32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                    }
                }
                return message;
            };

            /**
             * Creates a LoginResponse message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof auth.v1.LoginResponse
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {auth.v1.LoginResponse} LoginResponse
             */
            LoginResponse.fromObject = function fromObject(object) {
                if (object instanceof $root.auth.v1.LoginResponse)
                    return object;
                let message = new $root.auth.v1.LoginResponse();
                if (object.accessToken != null)
                    message.accessToken = String(object.accessToken);
                if (object.expiresIn != null)
                    message.expiresIn = object.expiresIn | 0;
                return message;
            };

            /**
             * Creates a plain object from a LoginResponse message. Also converts values to other types if specified.
             * @function toObject
             * @memberof auth.v1.LoginResponse
             * @static
             * @param {auth.v1.LoginResponse} message LoginResponse
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            LoginResponse.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                let object = {};
                if (options.defaults) {
                    object.accessToken = "";
                    object.expiresIn = 0;
                }
                if (message.accessToken != null && message.hasOwnProperty("accessToken"))
                    object.accessToken = message.accessToken;
                if (message.expiresIn != null && message.hasOwnProperty("expiresIn"))
                    object.expiresIn = message.expiresIn;
                return object;
            };

            /**
             * Converts this LoginResponse to JSON.
             * @function toJSON
             * @memberof auth.v1.LoginResponse
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            LoginResponse.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return LoginResponse;
        })();

        v1.AuthService = (function() {

            /**
             * Constructs a new AuthService service.
             * @memberof auth.v1
             * @classdesc Represents an AuthService
             * @extends $protobuf.rpc.Service
             * @constructor
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             */
            function AuthService(rpcImpl, requestDelimited, responseDelimited) {
                $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
            }

            (AuthService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = AuthService;

            /**
             * Creates new AuthService service using the specified rpc implementation.
             * @function create
             * @memberof auth.v1.AuthService
             * @static
             * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
             * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
             * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
             * @returns {AuthService} RPC service. Useful where requests and/or responses are streamed.
             */
            AuthService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                return new this(rpcImpl, requestDelimited, responseDelimited);
            };

            /**
             * Callback as used by {@link auth.v1.AuthService#login}.
             * @memberof auth.v1.AuthService
             * @typedef LoginCallback
             * @type {function}
             * @param {Error|null} error Error, if any
             * @param {auth.v1.LoginResponse} [response] LoginResponse
             */

            /**
             * Calls Login.
             * @function login
             * @memberof auth.v1.AuthService
             * @instance
             * @param {auth.v1.ILoginResquest} request LoginResquest message or plain object
             * @param {auth.v1.AuthService.LoginCallback} callback Node-style callback called with the error, if any, and LoginResponse
             * @returns {undefined}
             * @variation 1
             */
            Object.defineProperty(AuthService.prototype.login = function login(request, callback) {
                return this.rpcCall(login, $root.auth.v1.LoginResquest, $root.auth.v1.LoginResponse, request, callback);
            }, "name", { value: "Login" });

            /**
             * Calls Login.
             * @function login
             * @memberof auth.v1.AuthService
             * @instance
             * @param {auth.v1.ILoginResquest} request LoginResquest message or plain object
             * @returns {Promise<auth.v1.LoginResponse>} Promise
             * @variation 2
             */

            return AuthService;
        })();

        return v1;
    })();

    return auth;
})();