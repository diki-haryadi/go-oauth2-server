# API Feature Common
1. oauth2/revoke
2. /oauth2/userinfo
3. Register Endpoint (/register):
```
POST /register
{
"email": "user@example.com",
"password": "password123",
"name": "John Doe",
"additional_fields": {}
}

Response:
{
"user_id": "12345",
"message": "Registration successful"
}
```

3. Password Reset Flow:
```
// Request reset
POST /forgot-password
{
"email": "user@example.com"
}

// Reset with token
POST /reset-password
{
"token": "reset_token",
"new_password": "newpass123"
}
```
4. Profile Management:
```
// Get profile
GET /profile
Authorization: Bearer <access_token>

// Update profile
PUT /profile
Authorization: Bearer <access_token>
{
  "name": "Updated Name",
  "additional_fields": {}
}
```