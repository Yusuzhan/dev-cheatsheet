---
title: Dapr
icon: fa-gears
primary: "#0D2D6B"
lang: yaml
locale: zhs
---

## fa-terminal CLI 基础

```yaml
dapr init
dapr run --app-id myapp --app-port 3000 -- dotnet run
dapr stop --app-id myapp
dapr uninstall
```

```yaml
dapr publish --publish-app-id myapp --pubsub pubsub --topic orders --data '{"orderId":"1"}'
dapr invoke --app-id myapp --method orders --verb POST --data '{"id":"1"}'
dapr list
dapr status -k
```

## fa-arrow-right 服务调用

```yaml
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: myapp-config
spec:
  tracing:
    samplingRate: "1"
```

```go
resp, err := client.InvokeMethod(ctx, "ordersvc", "orders/1", "get")
```

```yaml
GET http://localhost:3500/v1.0/invoke/ordersvc/method/orders/1
```

## fa-database 状态管理

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: statestore
spec:
  type: state.redis
  metadata:
    - name: redisHost
      value: localhost:6379
    - name: redisPassword
      value: ""
```

```yaml
POST http://localhost:3500/v1.0/state/statestore
Content-Type: application/json

[{"key": "order1", "value": {"itemId": "A", "qty": 3}}]
```

```yaml
GET http://localhost:3500/v1.0/state/statestore/order1
DELETE http://localhost:3500/v1.0/state/statestore/order1
```

## fa-envelope 发布订阅

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: pubsub
spec:
  type: pubsub.redis
  metadata:
    - name: redisHost
      value: localhost:6379
    - name: redisPassword
      value: ""
```

```yaml
POST http://localhost:3500/v1.0/publish/pubsub/orders
Content-Type: application/json

{"orderId": "1", "status": "new"}
```

```yaml
POST http://localhost:3500/v1.0/publish/pubsub/orders?metadata.rawPayload=true
```

## fa-plug 绑定

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: queuebinding
spec:
  type: bindings.rabbitmq
  metadata:
    - name: queueName
      value: orders
    - name: host
      value: amqp://localhost:5672
```

```yaml
POST http://localhost:3500/v1.0/bindings/queuebinding
Content-Type: application/json

{ "operation": "create", "data": { "orderId": "1" } }
```

## fa-users Actor 模型

```yaml
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: myapp-config
spec:
  actorService:
    actorIdleTimeout: "30m"
    actorScanInterval: "30s"
    drainOngoingCallTimeout: "60s"
    drainRebalancedActors: true
```

```yaml
POST http://localhost:3500/v1.0/actors/orderactor/1/method/submit
Content-Type: application/json

{"itemId": "A", "qty": 5}
```

```yaml
GET http://localhost:3500/v1.0/actors/orderactor/1/state/status
```

## fa-eye 可观测性

```yaml
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: tracing
spec:
  tracing:
    samplingRate: "1"
    otel:
      endpointAddress: "localhost:4317"
      isSecure: false
      protocol: grpc
  metrics:
    enabled: true
    http:
      increasedCardinality: false
```

## fa-key 密钥存储

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: secretstore
spec:
  type: secretstores.local.file
  metadata:
    - name: secretsFile
      value: ./secrets.json
    - name: nestedSeparator
      value: ":"
```

```yaml
GET http://localhost:3500/v1.0/secrets/secretstore/db-password
GET http://localhost:3500/v1.0/secrets/secretstore/bulk
```

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: aws-secretstore
spec:
  type: secretstores.aws.secretmanager
  metadata:
    - name: region
      value: us-east-1
```

## fa-sliders 配置 API

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: appconfig
spec:
  type: configuration.redis
  metadata:
    - name: redisHost
      value: localhost:6379
    - name: redisPassword
      value: ""
```

```yaml
GET http://localhost:3500/v1.0/configuration/appconfig?key=debugMode
```

## fa-puzzle-piece Dapr 组件

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: mystate
  namespace: default
spec:
  type: state.postgresql
  version: v1
  metadata:
    - name: connectionString
      value: "host=localhost user=postgres password=secret port=5432"
    - name: tableName
      value: state
  scopes:
    - myapp
    - anotherapp
```

## fa-laptop 自托管模式

```yaml
dapr init --runtime-version 1.14.0
dapr init --slim
dapr run --app-id myapp --app-port 3000 --components-path ./components -- dotnet run
```

```yaml
.
├── app.py
└── components/
    ├── statestore.yaml
    ├── pubsub.yaml
    └── queuebinding.yaml
```

## fa-dharmachakra Kubernetes 模式

```yaml
dapr init -k --wait
dapr status -k
```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  template:
    metadata:
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "myapp"
        dapr.io/app-port: "3000"
        dapr.io/config: "tracing"
    spec:
      containers:
        - name: myapp
          image: myapp:latest
```

## fa-lock 分布式锁

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: lock
spec:
  type: lock.redis
  metadata:
    - name: redisHost
      value: localhost:6379
    - name: redisPassword
      value: ""
```

```yaml
POST http://localhost:3500/v1.0-alpha/lock/lock
Content-Type: application/json

{ "resourceId": "order-1", "lockOwner": "node-A", "expiryInSeconds": 30 }
```

```yaml
POST http://localhost:3500/v1.0-alpha/unlock/lock
Content-Type: application/json

{ "resourceId": "order-1", "lockOwner": "node-A" }
```

## fa-stream 工作流

```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: orderworkflow
spec:
  type: workflow
  metadata: []
```

```go
import "github.com/dapr/durabletask-go/task"

wf := task.NewWorkflow()
wf.AddActivity(task.Activity(func(ctx task.ActivityContext) (string, error) {
    var input string
    ctx.GetInput(&input)
    return "processed: " + input, nil
}))
```

```yaml
POST http://localhost:3500/v1.0-alpha/workflows/dapr/orderworkflow/start
Content-Type: application/json

{ "instanceId": "order-1", "input": "{\"orderId\":\"1\"}" }
```

```yaml
GET http://localhost:3500/v1.0-alpha/workflows/dapr/orderworkflow/order-1
POST http://localhost:3500/v1.0-alpha/workflows/dapr/orderworkflow/order-1/terminate
POST http://localhost:3500/v1.0-alpha/workflows/dapr/orderworkflow/order-1/pause
POST http://localhost:3500/v1.0-alpha/workflows/dapr/orderworkflow/order-1/resume
```
