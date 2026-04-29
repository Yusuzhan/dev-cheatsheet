---
title: OpenAPI / Swagger
icon: fa-file-code
primary: "#6BA539"
lang: yaml
---

## fa-sitemap OpenAPI Structure

```yaml
openapi: "3.1.0"
info:
  title: My API
  version: "1.0.0"
servers:
  - url: https://api.example.com/v1
paths: {}
components: {}
```

## fa-circle-info Info & Servers

```yaml
info:
  title: Pet Store API
  description: A sample pet store API
  version: "2.0.0"
  contact:
    name: API Support
    email: support@example.com
  license:
    name: MIT
    url: https://opensource.org/licenses/MIT
servers:
  - url: https://api.example.com/{version}
    variables:
      version:
        default: v1
        enum: [v1, v2]
```

## fa-route Paths & Operations

```yaml
paths:
  /pets:
    get:
      operationId: listPets
      summary: List all pets
      tags: [pets]
      parameters:
        - name: limit
          in: query
          schema:
            type: integer
            default: 20
      responses:
        "200":
          description: A list of pets
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Pet"
    post:
      operationId: createPet
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/NewPet"
      responses:
        "201":
          description: Created
```

## fa-sliders Parameters

```yaml
parameters:
  - name: petId
    in: path
    required: true
    schema:
      type: string
      pattern: "^[a-zA-Z0-9]+$"
  - name: status
    in: query
    schema:
      type: string
      enum: [available, sold]
  - name: X-Request-ID
    in: header
    schema:
      type: string
      format: uuid
  - name: cookie
    in: cookie
    schema:
      type: string
```

## fa-pen-to-square Request Body

```yaml
requestBody:
  required: true
  description: Pet to add
  content:
    application/json:
      schema:
        $ref: "#/components/schemas/NewPet"
    application/xml:
      schema:
        $ref: "#/components/schemas/NewPet"
    multipart/form-data:
      schema:
        type: object
        properties:
          name:
            type: string
          photo:
            type: string
            format: binary
```

## fa-reply Responses

```yaml
responses:
  "200":
    description: Success
    content:
      application/json:
        schema:
          $ref: "#/components/schemas/Pet"
        examples:
          dog:
            summary: A dog example
            value:
              id: "1"
              name: Buddy
              status: available
  "404":
    description: Not found
  default:
    description: Error
    content:
      application/json:
        schema:
          $ref: "#/components/schemas/Error"
```

## fa-cubes Schemas / Models

```yaml
components:
  schemas:
    Pet:
      type: object
      required: [id, name]
      properties:
        id:
          type: string
          format: uuid
        name:
          type: string
          maxLength: 100
        tag:
          type: string
        status:
          type: string
          enum: [available, pending, sold]
      allOf:
        - $ref: "#/components/schemas/NewPet"
        - type: object
          properties:
            id:
              type: string
    NewPet:
      type: object
      required: [name]
      properties:
        name:
          type: string
        tag:
          type: string
    Error:
      type: object
      required: [code, message]
      properties:
        code:
          type: integer
        message:
          type: string
```

## fa-shield-halved Security Schemes

```yaml
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT
    apiKey:
      type: apiKey
      in: header
      name: X-API-Key
    oauth2:
      type: oauth2
      flows:
        authorizationCode:
          authorizationUrl: https://example.com/oauth/authorize
          tokenUrl: https://example.com/oauth/token
          scopes:
            read:pets: Read pets
            write:pets: Write pets
    mutualTLS:
      type: mutualTLS
```

## fa-key Authentication

```yaml
security:
  - bearerAuth: []
  - apiKey: []

paths:
  /public:
    get:
      security: []
      responses:
        "200":
          description: Public endpoint
  /admin:
    get:
      security:
        - oauth2: [admin]
      responses:
        "200":
          description: Admin only
```

## fa-tags Tags & Grouping

```yaml
tags:
  - name: pets
    description: Pet operations
  - name: users
    description: User operations

paths:
  /pets:
    get:
      tags: [pets]
      summary: List pets
  /users:
    get:
      tags: [users]
      summary: List users
```

## fa-link $ref & Components

```yaml
components:
  schemas:
    Address:
      type: object
      properties:
        street:
          type: string
        city:
          type: string
  parameters:
    limitParam:
      name: limit
      in: query
      schema:
        type: integer
        default: 20
  responses:
    NotFound:
      description: Resource not found
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Error"
  examples:
    petExample:
      summary: A sample pet
      value:
        id: "1"
        name: Buddy

paths:
  /items:
    get:
      parameters:
        - $ref: "#/components/parameters/limitParam"
      responses:
        "404":
          $ref: "#/components/responses/NotFound"
```

## fa-lightbulb Examples

```yaml
components:
  examples:
    UserInput:
      summary: Sample user
      value:
        name: Alice
        email: alice@example.com
    ErrorResp:
      summary: Error response
      value:
        code: 400
        message: Bad request

paths:
  /users:
    post:
      requestBody:
        content:
          application/json:
            examples:
              valid:
                $ref: "#/components/examples/UserInput"
      responses:
        "400":
          content:
            application/json:
              examples:
                error:
                  $ref: "#/components/examples/ErrorResp"
```

## fa-gears Code Generation

```bash
openapi-generator-cli generate \
  -i openapi.yaml \
  -g python-fastapi \
  -o ./api

openapi-generator-cli generate \
  -i openapi.yaml \
  -g typescript-axios \
  -o ./client

openapi-generator-cli generate \
  -i openapi.yaml \
  -g go-server \
  -o ./server

openapi-generator-cli generate \
  -i openapi.yaml \
  -g java-spring \
  -o ./spring-app
```

## fa-display Swagger UI / Redoc

```yaml
swagger-ui:
  image: swaggerapi/swagger-ui
  ports:
    - "8080:8080"
  environment:
    SWAGGER_JSON: /openapi.yaml
  volumes:
    - ./openapi.yaml:/openapi.yaml

redoc:
  image: redocly/redoc
  ports:
    - "8081:80"
  environment:
    SPEC_URL: /openapi.yaml
  volumes:
    - ./openapi.yaml:/usr/share/nginx/html/openapi.yaml
```

```html
<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css">
<script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
<div id="swagger-ui"></div>
<script>
SwaggerUIBundle({ url: "openapi.yaml", dom_id: "#swagger-ui" });
</script>
```
