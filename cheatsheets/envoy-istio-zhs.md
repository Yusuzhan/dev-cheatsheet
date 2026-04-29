---
title: Envoy / Istio
icon: fa-network-wired
primary: "#466BB0"
lang: yaml
locale: zhs
---

## fa-sitemap Envoy 架构

```yaml
# Envoy 核心概念：
# - Listener：在端口上接收入站连接
# - Filter chain：应用于连接的 L4/L7 过滤器（如 HTTP、TLS）
# - Route：将请求映射到上游集群
# - Cluster：具有负载均衡策略的上游端点组
# - Endpoint：单个上游 host:port

# xDS API 协议：
# - LDS（监听器发现）、RDS（路由）、CDS（集群）、EDS（端点）
# - EDS 提供动态端点更新，无需完整配置重载
```

## fa-ear-listen Listener 与 Filter

```yaml
static_resources:
  listeners:
    - name: http_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 10000
      filter_chains:
        - filters:
            - name: envoy.filters.network.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                stat_prefix: ingress_http
                codec_type: AUTO
                route_config:
                  name: local_route
                  virtual_hosts:
                    - name: backend
                      domains: ["*"]
                      routes:
                        - match:
                            prefix: "/"
                          route:
                            cluster: service_a
                http_filters:
                  - name: envoy.filters.http.router
```

## fa-route 路由

```yaml
route_config:
  virtual_hosts:
    - name: all
      domains: ["*"]
      routes:
        - match:
            prefix: "/api/v1"
          route:
            cluster: api_v1
            timeout: 30s
            retry_policy:
              retry_on: 5xx
              num_retries: 3
              per_try_timeout: 5s
        - match:
            prefix: "/api/v2"
          route:
            cluster: api_v2
            weighted_clusters:
              clusters:
                - name: api_v2_canary
                  weight: 10
                - name: api_v2_stable
                  weight: 90
        - match:
            prefix: "/"
            headers:
              - name: "x-canary"
                string_match:
                  exact: "true"
          route:
            cluster: canary
        - match:
            prefix: "/"
          route:
            cluster: stable
```

## fa-circle-nodes Cluster 与 Endpoint

```yaml
static_resources:
  clusters:
    - name: service_a
      type: STRICT_DNS
      lb_policy: ROUND_ROBIN
      connect_timeout: 5s
      health_checks:
        - timeout: 2s
          interval: 10s
          unhealthy_threshold: 3
          healthy_threshold: 2
          http_health_check:
            path: /healthz
            host: service-a
      load_assignment:
        cluster_name: service_a
        endpoints:
          - lb_endpoints:
              - endpoint:
                  address:
                    socket_address:
                      address: service-a
                      port_value: 8080
              - endpoint:
                  address:
                    socket_address:
                      address: service-a-v2
                      port_value: 8080

    - name: service_b
      type: EDS
      eds_cluster_config:
        eds_config:
          api_config_source:
            api_type: GRPC
            grpc_services:
              - envoy_grpc:
                  cluster_name: xds_cluster
```

## fa-diagram-project Istio 架构

```yaml
# Istio 组件：
# - istiod：控制平面（Pilot + Citadel + Galley 合并）
# - Envoy proxy：注入到每个 Pod 的 sidecar
# - istio-ingressgateway：外部流量入口
# - istio-egressgateway：受控出口

# 流量路径：
# 客户端 → Ingress Gateway → Sidecar（入站）→ 应用容器
# 应用容器 → Sidecar（出站）→ Sidecar（入站）→ 目标

# istioctl install --set profile=demo
# kubectl label namespace default istio-injection=enabled
```

## fa-shuffle VirtualService

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: my-app
spec:
  hosts:
    - my-app
    - my-app.example.com
  gateways:
    - mesh
    - my-gateway
  http:
    - match:
        - headers:
            x-canary:
              exact: "true"
      route:
        - destination:
            host: my-app
            subset: canary
          weight: 100
    - route:
        - destination:
            host: my-app
            subset: stable
          weight: 90
        - destination:
            host: my-app
            subset: canary
          weight: 10
      timeout: 30s
      retries:
        attempts: 3
        perTryTimeout: 5s
        retryOn: 5xx,reset
      fault:
        abort:
          percentage:
            value: 0.1
          httpStatus: 500
```

## fa-sliders DestinationRule

```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: my-app
spec:
  host: my-app
  trafficPolicy:
    connectionPool:
      tcp:
        maxConnections: 100
      http:
        h2UpgradePolicy: DEFAULT
        http1MaxPendingRequests: 100
        http2MaxRequests: 100
    outlierDetection:
      consecutive5xxErrors: 3
      interval: 30s
      baseEjectionTime: 30s
      maxEjectionPercent: 50
  subsets:
    - name: stable
      labels:
        version: v1
      trafficPolicy:
        connectionPool:
          http:
            http1MaxPendingRequests: 200
    - name: canary
      labels:
        version: v2
    - name: v3
      labels:
        version: v3
```

## fa-door-open Gateway

```yaml
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: my-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
    - port:
        number: 80
        name: http
        protocol: HTTP
      hosts:
        - my-app.example.com
      tls:
        httpsRedirect: true
    - port:
        number: 443
        name: https
        protocol: HTTPS
      tls:
        mode: SIMPLE
        credentialName: my-app-tls
      hosts:
        - my-app.example.com
    - port:
        number: 15443
        name: tls-passthrough
        protocol: TLS
      tls:
        mode: PASSTHROUGH
      hosts:
        - "*.example.com"
```

## fa-lock PeerAuthentication

```yaml
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: istio-system
spec:
  mtls:
    mode: STRICT
```

```yaml
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: my-app
  namespace: production
spec:
  selector:
    matchLabels:
      app: my-app
  mtls:
    mode: STRICT
  portLevelMtls:
    8080:
      mode: PERMISSIVE
```

## fa-user-shield AuthorizationPolicy

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: allow-frontend
  namespace: production
spec:
  selector:
    matchLabels:
      app: backend
  action: ALLOW
  rules:
    - from:
        - source:
            principals: ["cluster.local/ns/production/sa/frontend"]
            namespaces: ["production"]
      to:
        - operation:
            methods: ["GET", "POST"]
            paths: ["/api/*"]
      when:
        - key: request.headers[x-token]
          notValues: [""]
```

```yaml
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: deny-external
  namespace: production
spec:
  action: DENY
  rules:
    - from:
        - source:
            notNamespaces: ["production", "istio-system"]
```

## fa-globe ServiceEntry

```yaml
apiVersion: networking.istio.io/v1beta1
kind: ServiceEntry
metadata:
  name: external-api
spec:
  hosts:
    - api.external-service.com
  location: MESH_EXTERNAL
  ports:
    - number: 443
      name: https
      protocol: TLS
  resolution: DNS
  dnsLookupFamily: V4_ONLY
```

```yaml
apiVersion: networking.istio.io/v1beta1
kind: ServiceEntry
metadata:
  name: vm-service
spec:
  hosts:
    - vm-service.internal
  location: MESH_INTERNAL
  ports:
    - number: 8080
      name: http
      protocol: HTTP
  resolution: STATIC
  endpoints:
    - address: 10.0.0.100
      ports:
        http: 8080
      labels:
        app: legacy-app
```

## fa-syringe Sidecar 注入

```yaml
# 通过命名空间标签自动注入
# kubectl label namespace default istio-injection=enabled

# 手动注入
# istioctl kube-inject -f deployment.yaml | kubectl apply -f -
```

```yaml
# Pod 级别覆盖
apiVersion: v1
kind: Pod
metadata:
  name: my-app
  annotations:
    sidecar.istio.io/inject: "true"
    sidecar.istio.io/proxyCPU: "100m"
    sidecar.istio.io/proxyMemory: "128Mi"
    traffic.sidecar.istio.io/includeInboundPorts: "8080,9090"
    traffic.sidecar.istio.io/excludeOutboundPorts: "3306"
    traffic.sidecar.istio.io/excludeOutboundIPRanges: "10.0.0.0/8"
spec:
  containers:
    - name: app
      image: my-app:latest
```

## fa-chart-line 可观测性

```yaml
# Prometheus 抓取注解
apiVersion: v1
kind: Pod
metadata:
  name: my-app
  annotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "15090"
    prometheus.io/path: "/stats/prometheus"
```

```yaml
# Kiali 仪表盘（服务网格可视化）
# istioctl dashboard kiali

# Jaeger 链路追踪（分布式追踪）
# istioctl dashboard jaeger

# Envoy 访问日志
apiVersion: networking.istio.io/v1beta1
kind: EnvoyFilter
metadata:
  name: access-log
  namespace: istio-system
spec:
  workloadSelector:
    labels:
      istio: ingressgateway
  configPatches:
    - applyTo: NETWORK_FILTER
      match:
        context: GATEWAY
      patch:
        operation: MERGE
        value:
          typed_config:
            "@type": "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager"
            access_log:
              - name: envoy.access_loggers.file
                typed_config:
                  "@type": "type.googleapis.com/envoy.extensions.access_loggers.file.v3.FileAccessLog"
                  path: /dev/stdout
                  log_format:
                    text_format_source:
                      inline_string: "[%START_TIME%] \"%REQ(:METHOD)% %REQ(X-ENVOY-ORIGINAL-PATH?:PATH)%\" %RESPONSE_CODE% %RESPONSE_FLAGS%\n"
```

## fa-arrows-split-up-and-left 流量管理模式

```yaml
# 金丝雀发布（加权路由）
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: canary
spec:
  hosts: [my-app]
  http:
    - route:
        - destination:
            host: my-app
            subset: v1
          weight: 95
        - destination:
            host: my-app
            subset: v2
          weight: 5
```

```yaml
# 基于请求头的 A/B 测试
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: ab-test
spec:
  hosts: [my-app]
  http:
    - match:
        - headers:
            cookie:
              regex: ".*user_group=beta.*"
      route:
        - destination:
            host: my-app
            subset: v2
    - route:
        - destination:
            host: my-app
            subset: v1
```

```yaml
# 流量镜像（影子流量）
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: mirror
spec:
  hosts: [my-app]
  http:
    - route:
        - destination:
            host: my-app
            subset: v1
          weight: 100
      mirror:
        host: my-app
        subset: v2
      mirrorPercentage:
        value: 100
```
