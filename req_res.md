OAuth 2.0 is a protocol for authorization that allows users to grant limited access to resources on a server without sharing their credentials. Below are examples of requests and responses for various OAuth 2.0 flows.

---

## **1. Authorization Code Flow**
Used by web apps for server-side exchange.

### **Step 1: Authorization Request**
The user is redirected to the authorization server with a URL like this:

```
GET /authorize?response_type=code
              &client_id=CLIENT_ID
              &redirect_uri=REDIRECT_URI
              &scope=SCOPE
              &state=STATE
```

### Example:
```
GET https://authorization-server.com/oauth2/authorize?response_type=code
    &client_id=12345
    &redirect_uri=https://yourapp.com/callback
    &scope=read write
    &state=abc123
```

### **Step 2: Authorization Response**
The server redirects back to the client app with a code:

```
https://yourapp.com/callback?code=AUTH_CODE&state=abc123
```

---

### **Step 3: Token Request**
Exchange the authorization code for an access token.

```
POST /token
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=AUTH_CODE
&redirect_uri=REDIRECT_URI
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET
```

### Example:
```http
POST https://authorization-server.com/oauth2/token
Content-Type: application/x-www-form-urlencoded

grant_type=authorization_code
&code=abcd1234
&redirect_uri=https://yourapp.com/callback
&client_id=12345
&client_secret=xyzsecret
```

### **Step 4: Token Response**
The server responds with an access token.

```json
{
  "access_token": "ACCESS_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "REFRESH_TOKEN",
  "scope": "read write"
}
```

---

## **2. Client Credentials Flow**
Used for server-to-server communication.

### **Request**
```
POST /token
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET
&scope=SCOPE
```

### Example:
```http
POST https://authorization-server.com/oauth2/token
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
&client_id=12345
&client_secret=xyzsecret
&scope=read
```

### **Response**
```json
{
  "access_token": "ACCESS_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600,
  "scope": "read"
}
```

---

## **3. Password Grant Flow**
Used for trusted applications to authenticate users directly.

### **Request**
```
POST /token
Host: authorization-server.com
Content-Type: application/x-www-form-urlencoded

grant_type=password
&username=USERNAME
&password=PASSWORD
&client_id=CLIENT_ID
&client_secret=CLIENT_SECRET
&scope=SCOPE
```

### Example:
```http
POST https://authorization-server.com/oauth2/token
Content-Type: application/x-www-form-urlencoded

grant_type=password
&username=johndoe
&password=1234password
&client_id=12345
&client_secret=xyzsecret
&scope=read write
```

### **Response**
```json
{
  "access_token": "ACCESS_TOKEN",
  "token_type": "Bearer",
  "expires_in": 3600,
  "scope": "read write"
}
```

---

## **4. Implicit Flow**
Used by single-page applications (SPA).

### **Request**
```
GET /authorize?response_type=token
              &client_id=CLIENT_ID
              &redirect_uri=REDIRECT_URI
              &scope=SCOPE
              &state=STATE
```

### Example:
```
GET https://authorization-server.com/oauth2/authorize?response_type=token
    &client_id=12345
    &redirect_uri=https://yourapp.com/callback
    &scope=read write
    &state=abc123
```

### **Response**
The server redirects back with the token in the URL fragment.

```
https://yourapp.com/callback#access_token=ACCESS_TOKEN
&token_type=Bearer
&expires_in=3600
&state=abc123
```

---

These examples should give you a good starting point for understanding how OAuth 2.0 works in various scenarios.