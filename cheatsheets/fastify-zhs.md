---
title: Fastify
icon: fa-bolt
primary: "#000000"
lang: javascript
locale: zhs
---

## fa-rocket 初始化与基本服务

```javascript
import Fastify from 'fastify'

const app = Fastify({ logger: true })

app.get('/', async (request, reply) => {
  return { hello: 'world' }
})

app.listen({ port: 3000, host: '0.0.0.0' })
```

## fa-route 路由

```javascript
app.get('/users', getUsers)
app.post('/users', createUser)
app.put('/users/:id', updateUser)
app.delete('/users/:id', deleteUser)
app.patch('/users/:id', patchUser)
app.all('/health', healthHandler)

app.register(async (instance) => {
  instance.get('/inner', innerHandler)
}, { prefix: '/api' })
```

## fa-map-pin 路径参数与查询

```javascript
app.get('/users/:id', async (request) => {
  const { id } = request.params
  return { id }
})

app.get('/search', async (request) => {
  const { q, page, limit } = request.query
  return { q, page: page || 1, limit: limit || 10 }
})

app.get('/files/*', async (request) => {
  const { '*': wildcard } = request.params
  return { path: wildcard }
})
```

## fa-filter 请求体与验证

```javascript
app.post('/users', {
  handler: async (request, reply) => {
    return request.body
  },
  schema: {
    body: {
      type: 'object',
      required: ['name', 'email'],
      properties: {
        name: { type: 'string', minLength: 2 },
        email: { type: 'string', format: 'email' },
        age: { type: 'integer', minimum: 0 }
      }
    }
  }
}, async (request, reply) => {})
```

## fa-check JSON Schema 验证

```javascript
const schema = {
  params: {
    type: 'object',
    properties: { id: { type: 'integer' } }
  },
  querystring: {
    type: 'object',
    properties: {
      page: { type: 'integer', default: 1 },
      limit: { type: 'integer', default: 20 }
    }
  },
  headers: {
    type: 'object',
    required: ['authorization'],
    properties: {
      authorization: { type: 'string' }
    }
  },
  response: {
    200: {
      type: 'object',
      properties: {
        id: { type: 'integer' },
        name: { type: 'string' }
      }
    }
  }
}
```

## fa-link 钩子

```javascript
app.addHook('onRequest', async (request, reply) => {
  request.startTime = Date.now()
})

app.addHook('preHandler', async (request, reply) => {
  if (!request.headers.authorization) {
    reply.code(401).send({ error: 'Unauthorized' })
  }
})

app.addHook('onResponse', async (request, reply) => {
  const elapsed = Date.now() - request.startTime
  app.log.info({ elapsed })
})

app.addHook('onSend', async (request, reply, payload) => {
  return payload
})

app.addHook('onError', async (request, reply, error) => {
  app.log.error(error)
})
```

## fa-puzzle-piece 插件

```javascript
import fp from 'fastify-plugin'

const dbPlugin = fp(async (app, options) => {
  const db = await connectDB(options.url)
  app.decorate('db', db)
  app.addHook('onClose', async () => db.close())
})

app.register(dbPlugin, { url: 'mongodb://localhost:27017' })

app.register(import('@fastify/swagger'), {
  swagger: { title: 'API', version: '1.0.0' }
})

app.register(import('@fastify/cors'), { origin: true })
```

## fa-shield 中间件

```javascript
app.register(import('@fastify/rate-limit'), {
  max: 100,
  timeWindow: '1 minute'
})

app.register(import('@fastify/helmet'))

app.register(import('@fastify/compress'))

app.register(import('@fastify/csrf-protection'))

app.addHook('onRequest', async (request, reply) => {
  reply.header('X-Custom', 'value')
})
```

## fa-triangle-exclamation 错误处理

```javascript
app.setErrorHandler((error, request, reply) => {
  if (error.validation) {
    reply.code(400).send({ error: error.message, details: error.validation })
    return
  }
  if (error.statusCode === 404) {
    reply.code(404).send({ error: 'Not found' })
    return
  }
  reply.code(error.statusCode || 500).send({ error: 'Internal error' })
})

app.setNotFoundHandler((request, reply) => {
  reply.code(404).send({ error: 'Route not found' })
})

throw app.httpErrors.unauthorized()
throw app.httpErrors.notFound()
throw app.httpErrors.badRequest('invalid input')
```

## fa-lock JWT 认证

```javascript
import jwt from '@fastify/jwt'

app.register(jwt, { secret: 'supersecret' })

app.post('/login', async (request, reply) => {
  const token = app.jwt.sign({ id: 1, role: 'admin' })
  return { token }
})

app.get('/protected', {
  preHandler: [app.authenticate]
}, async (request) => {
  return request.user
})

app.decorate('authenticate', async (request, reply) => {
  try {
    await request.jwtVerify()
  } catch (err) {
    reply.code(401).send({ error: 'Invalid token' })
  }
})
```

## fa-upload 文件上传

```javascript
import multipart from '@fastify/multipart'

app.register(multipart, { limits: { fileSize: 10 * 1024 * 1024 } })

app.post('/upload', async (request, reply) => {
  const data = await request.file()
  const buffer = await data.toBuffer()
  await fs.writeFile(`./uploads/${data.filename}`, buffer)
  return { filename: data.filename, mimetype: data.mimetype }
})

app.post('/multi-upload', async (request, reply) => {
  const files = request.files()
  const results = []
  for await (const data of files) {
    results.push(data.filename)
  }
  return { files: results }
})
```

## fa-folder-open 静态文件

```javascript
import fastifyStatic from '@fastify/static'
import path from 'path'

app.register(fastifyStatic, {
  root: path.join(import.meta.dirname, 'public'),
  prefix: '/assets/',
  index: ['index.html']
})

app.get('/download', async (request, reply) => {
  return reply.download('./files/report.pdf', 'report.pdf')
})
```

## fa-scroll 日志

```javascript
const app = Fastify({
  logger: {
    level: 'info',
    transport: {
      target: 'pino-pretty',
      options: { colorize: true }
    }
  }
})

app.log.info('服务已启动')
app.log.warn({ event: 'slow_query' }, '查询耗时 5s')
app.log.error({ err }, '请求失败')

app.addHook('onResponse', async (request, reply) => {
  request.log.info({ statusCode: reply.statusCode }, '已发送响应')
})
```

## fa-vial 测试

```javascript
import { test } from 'node:test'
import assert from 'node:assert'
import build from './app.js'

test('GET /', async (t) => {
  const app = build()

  t.after(() => app.close())

  const response = await app.inject({
    method: 'GET',
    url: '/'
  })

  assert.equal(response.statusCode, 200)
  assert.deepEqual(response.json(), { hello: 'world' })
})

test('POST /users', async (t) => {
  const app = build()
  t.after(() => app.close())

  const response = await app.inject({
    method: 'POST',
    url: '/users',
    payload: { name: 'Alice', email: 'alice@test.com' }
  })

  assert.equal(response.statusCode, 200)
})
```

## fa-wand-magic-sparkles 装饰器

```javascript
app.decorate('db', null)
app.decorate('getUser', function (id) {
  return this.db.findUser(id)
})

app.decorateRequest('user', null)
app.addHook('preHandler', async (request) => {
  if (request.headers.authorization) {
    request.user = await verifyToken(request.headers.authorization)
  }
})

app.decorateReply('success', function (data) {
  return this.code(200).send({ success: true, data })
})

if (!app.hasDecorator('db')) {
  app.decorate('db', createConnection())
}
```
