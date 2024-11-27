___

### API Endpoint

**POST** `/v1/oauth/introspect`

___

### Request Headers

-   **Authorization**: `Basic dGVzdF9jbGllbnRfMTp0ZXN0X3NlY3JldA==`
    -   Description: Base64-encoded client credentials for basic authentication.
-   **Content-Type**: `application/json`
    -   Description: Indicates the request payload is in JSON format.

___

### Request Body

The request body must be in JSON format and contain the following fields:

```
{
  "token": "string",
  "token_type_hint": "string"
}
```

-   **token** (string, required): The OAuth 2.0 token to be introspected.
-   **token\_type\_hint** (string, optional): A hint about the type of the token submitted for introspection (e.g., `access_token` or `refresh_token`).

___

### Response

The response will be in JSON format, providing details about the token's validity and associated information.

#### Successful Response

```
{
  "active": true,
  "scope": "read write",
  "client_id": "test_client_1",
  "username": "user123",
  "token_type": "Bearer",
  "exp": 1640995200,
  "iat": 1640918800,
  "nbf": 1640918800,
  "sub": "user123",
  "aud": "my_api",
  "iss": "http://localhost:4000",
  "jti": "b8e3bc09-cf9a-40c5-a5ad-e5f1240c2f4a"
}

```

-   **active** (boolean): Indicates if the token is currently active.
-   **scope** (string): The scopes associated with the token.
-   **client\_id** (string): The client identifier for the token.
-   **username** (string): The username associated with the token, if applicable.
-   **token\_type** (string): The type of the token (e.g., `Bearer`).
-   **exp** (integer): The token expiration time in Unix epoch format.
-   **iat** (integer): The time at which the token was issued in Unix epoch format.
-   **nbf** (integer): The time before which the token must not be accepted, in Unix epoch format.
-   **sub** (string): The subject of the token, usually a user ID or username.
-   **aud** (string): The intended audience of the token.
-   **iss** (string): The issuer of the token.
-   **jti** (string): A unique identifier for the token.

___

#### Error Response

-   **active** (boolean): Indicates the token is not active or invalid.

___

{
"code": "404000",
"status_code": 404,
"status": "undefined",
"error": "Access token not found"
}