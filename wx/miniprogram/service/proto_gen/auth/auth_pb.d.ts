import * as $protobuf from "protobufjs";
/** Namespace auth. */
export namespace auth {

    /** Namespace v1. */
    namespace v1 {

        /** Properties of a LoginResquest. */
        interface ILoginResquest {

            /** LoginResquest code */
            code?: (string|null);
        }

        /** Represents a LoginResquest. */
        class LoginResquest implements ILoginResquest {

            /**
             * Constructs a new LoginResquest.
             * @param [properties] Properties to set
             */
            constructor(properties?: auth.v1.ILoginResquest);

            /** LoginResquest code. */
            public code: string;

            /**
             * Creates a new LoginResquest instance using the specified properties.
             * @param [properties] Properties to set
             * @returns LoginResquest instance
             */
            public static create(properties?: auth.v1.ILoginResquest): auth.v1.LoginResquest;

            /**
             * Decodes a LoginResquest message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns LoginResquest
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.v1.LoginResquest;

            /**
             * Creates a LoginResquest message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns LoginResquest
             */
            public static fromObject(object: { [k: string]: any }): auth.v1.LoginResquest;

            /**
             * Creates a plain object from a LoginResquest message. Also converts values to other types if specified.
             * @param message LoginResquest
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: auth.v1.LoginResquest, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this LoginResquest to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Properties of a LoginResponse. */
        interface ILoginResponse {

            /** LoginResponse accessToken */
            accessToken?: (string|null);

            /** LoginResponse expiresIn */
            expiresIn?: (number|null);
        }

        /** Represents a LoginResponse. */
        class LoginResponse implements ILoginResponse {

            /**
             * Constructs a new LoginResponse.
             * @param [properties] Properties to set
             */
            constructor(properties?: auth.v1.ILoginResponse);

            /** LoginResponse accessToken. */
            public accessToken: string;

            /** LoginResponse expiresIn. */
            public expiresIn: number;

            /**
             * Creates a new LoginResponse instance using the specified properties.
             * @param [properties] Properties to set
             * @returns LoginResponse instance
             */
            public static create(properties?: auth.v1.ILoginResponse): auth.v1.LoginResponse;

            /**
             * Decodes a LoginResponse message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns LoginResponse
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): auth.v1.LoginResponse;

            /**
             * Creates a LoginResponse message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns LoginResponse
             */
            public static fromObject(object: { [k: string]: any }): auth.v1.LoginResponse;

            /**
             * Creates a plain object from a LoginResponse message. Also converts values to other types if specified.
             * @param message LoginResponse
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: auth.v1.LoginResponse, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this LoginResponse to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }

        /** Represents an AuthService */
        class AuthService extends $protobuf.rpc.Service {

            /**
             * Constructs a new AuthService service.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             */
            constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);

            /**
             * Creates new AuthService service using the specified rpc implementation.
             * @param rpcImpl RPC implementation
             * @param [requestDelimited=false] Whether requests are length-delimited
             * @param [responseDelimited=false] Whether responses are length-delimited
             * @returns RPC service. Useful where requests and/or responses are streamed.
             */
            public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): AuthService;

            /**
             * Calls Login.
             * @param request LoginResquest message or plain object
             * @param callback Node-style callback called with the error, if any, and LoginResponse
             */
            public login(request: auth.v1.ILoginResquest, callback: auth.v1.AuthService.LoginCallback): void;

            /**
             * Calls Login.
             * @param request LoginResquest message or plain object
             * @returns Promise
             */
            public login(request: auth.v1.ILoginResquest): Promise<auth.v1.LoginResponse>;
        }

        namespace AuthService {

            /**
             * Callback as used by {@link auth.v1.AuthService#login}.
             * @param error Error, if any
             * @param [response] LoginResponse
             */
            type LoginCallback = (error: (Error|null), response?: auth.v1.LoginResponse) => void;
        }
    }
}
