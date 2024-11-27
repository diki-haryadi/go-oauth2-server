sequenceDiagram
participant C as Client/Frontend
participant AS as Authorization Server
participant RS as Resource Server

    %% Client Credentials Flow
    rect rgb(200, 220, 240)
    Note over C,RS: Client Credentials Flow
    C->>AS: POST /oauth/tokens
    Note right of C: grant_type=client_credentials<br/>Basic Auth: client_id:client_secret<br/>scope=read_write
    AS->>C: Returns access_token
    C->>RS: API request with access_token
    RS->>AS: POST /oauth/introspect<br/>Validate token
    AS->>RS: Token validation response
    RS->>C: Protected resource response
    end

    %% Password Flow
    rect rgb(220, 240, 200)
    Note over C,RS: Password Flow
    C->>AS: POST /oauth/tokens
    Note right of C: grant_type=password<br/>Basic Auth: client_id:client_secret<br/>username & password<br/>scope=read_write
    AS->>C: Returns access_token + refresh_token
    C->>RS: API request with access_token
    RS->>AS: POST /oauth/introspect<br/>Validate token
    AS->>RS: Token validation response
    RS->>C: Protected resource response
    end

    %% Authorization Code Flow
    rect rgb(240, 220, 200)
    Note over C,RS: Authorization Code Flow
    C->>AS: GET /oauth/authorize
    Note right of C: response_type=code<br/>client_id<br/>redirect_uri<br/>scope
    AS->>C: Redirect with auth code
    C->>AS: POST /oauth/tokens
    Note right of C: grant_type=authorization_code<br/>code<br/>redirect_uri<br/>Basic Auth: client_id:client_secret
    AS->>C: Returns access_token + refresh_token
    C->>RS: API request with access_token
    RS->>AS: POST /oauth/introspect<br/>Validate token
    AS->>RS: Token validation response
    RS->>C: Protected resource response
    end

    %% Refresh Token Flow
    rect rgb(240, 200, 220)
    Note over C,RS: Refresh Token Flow
    C->>AS: POST /oauth/tokens
    Note right of C: grant_type=refresh_token<br/>refresh_token<br/>Basic Auth: client_id:client_secret
    AS->>C: Returns new access_token + refresh_token
    C->>RS: API request with new access_token
    RS->>AS: POST /oauth/introspect<br/>Validate token
    AS->>RS: Token validation response
    RS->>C: Protected resource response
    end

sequenceDiagram
participant Admin
participant Client
participant User
participant AS as Authorization Server

    %% Client Management
    rect rgb(200, 220, 240)
    Note over Admin,AS: Client Management APIs
    Admin->>AS: POST /api/v1/oauth/clients
    Note right of Admin: Register new client<br/>name, redirect_uris, grant_types
    AS->>Admin: Returns client_id & client_secret

    Admin->>AS: GET /api/v1/oauth/clients
    Note right of Admin: List all registered clients

    Admin->>AS: PUT /api/v1/oauth/clients/{client_id}
    Note right of Admin: Update client details

    Admin->>AS: DELETE /api/v1/oauth/clients/{client_id}
    Note right of Admin: Remove client registration
    end

    %% User Management
    rect rgb(220, 240, 200)
    Note over Admin,AS: User Management APIs
    Admin->>AS: POST /api/v1/users
    Note right of Admin: Register new user<br/>username, password, roles

    Admin->>AS: GET /api/v1/users
    Note right of Admin: List all users

    User->>AS: PUT /api/v1/users/profile
    Note right of User: Update user profile

    Admin->>AS: DELETE /api/v1/users/{user_id}
    Note right of Admin: Deactivate user
    end

    %% Scope Management
    rect rgb(240, 220, 200)
    Note over Admin,AS: Scope Management APIs
    Admin->>AS: POST /api/v1/oauth/scopes
    Note right of Admin: Create new scope<br/>name, description

    Admin->>AS: GET /api/v1/oauth/scopes
    Note right of Admin: List all available scopes

    Admin->>AS: PUT /api/v1/oauth/scopes/{scope_id}
    Note right of Admin: Update scope details
    end

    %% Token Management
    rect rgb(240, 200, 220)
    Note over Admin,AS: Token Management APIs
    Admin->>AS: GET /api/v1/oauth/tokens
    Note right of Admin: List active tokens

    Admin->>AS: DELETE /api/v1/oauth/tokens/{token_id}
    Note right of Admin: Revoke specific token

    Client->>AS: POST /api/v1/oauth/tokens/revoke
    Note right of Client: Revoke token by value
    end