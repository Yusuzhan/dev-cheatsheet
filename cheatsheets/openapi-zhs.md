---
title: OpenAPI / Swagger
icon: fa-file-code
primary: "#6BA539"
lang: yaml
locale: zhs
---

## fa-sitemap OpenAPI 结构

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

## fa-circle-info 信息与服务器

```yaml
info:
  title: Pet Store API
  description: 示例宠物商店 API
  version: "2.0.0"
  contact:
    name: API 支持
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

## fa-route 路径与操作

```yaml
paths:
  /pets:
    get:
      operationId: listPets
      summary: 列出所有宠物
      tags: [pets]
      parameters:
        - name: limit
          in: query
          schema:
            type: integer
            default: 20
      responses:
        "200":
          description: 宠物列表
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
          description: 创建成功
```

## fa-sliders 参数

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

## fa-pen-to-square 请求体

```yaml
requestBody:
  required: true
  description: 要添加的宠物
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

## fa-reply 响应

```yaml
responses:
  "200":
    description: 成功
    content:
      application/json:
        schema:
          $ref: "#/components/schemas/Pet"
        examples:
          dog:
            summary: 狗的示例
            value:
              id: "1"
              name: Buddy
              status: available
  "404":
    description: 未找到
  default:
    description: 错误
    content:
      application/json:
        schema:
          $ref: "#/components/schemas/Error"
```

## fa-cubes 模式 / 模型

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

## fa-shield-halved 安全方案

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
            read:pets: 读取宠物
            write:pets: 写入宠物
    mutualTLS:
      type: mutualTLS
```

## fa-key 认证

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
          description: 公开端点
  /admin:
    get:
      security:
        - oauth2: [admin]
      responses:
        "200":
          description: 仅管理员
```

## fa-tags 标签与分组

```yaml
tags:
  - name: pets
    description: 宠物操作
  - name: users
    description: 用户操作

paths:
  /pets:
    get:
      tags: [pets]
      summary: 列出宠物
  /users:
    get:
      tags: [users]
      summary: 列出用户
```

## fa-link $ref 与组件

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
      description: 资源未找到
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/Error"
  examples:
    petExample:
      summary: 示例宠物
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

## fa-lightbulb 示例

```yaml
components:
  examples:
    UserInput:
      summary: 示例用户
      value:
        name: Alice
        email: alice@example.com
    ErrorResp:
      summary: 错误响应
      value:
        code: 400
        message: 请求无效

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

## fa-gears 代码生成

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
