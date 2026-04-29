---
title: Envoy / Istio
icon: fa-network-wired
primary: "#466BB0"
lang: yaml
---

## fa-sitemap Envoy Architecture

```yaml
# Envoy core concepts:
# - Listener: accepts inbound connections on a port
# - Filter chain: L4/L7 filters applied to connections (e.g. HTTP, TLS)
# - Route: maps requests to upstream clusters
# - Cluster: group of upstream endpoints with load balancing
# - Endpoint: a single upstream host:port

# xDS API protocols:
# - LDS (Listener Discovery), RDS (Route), CDS (Cluster), EDS (Endpoint)
# - EDS delivers dynamic endpoint updates without full config reload
```

## fa-ear-listen Listeners & Filters

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

## fa-route Routes

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

## fa-circle-nodes Clusters & Endpoints

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

## fa-diagram-project Istio Architecture

```yaml
# Istio components:
# - istiod: control plane (Pilot + Citadel + Galley merged)
# - Envoy proxy: sidecar injected into each pod
# - istio-ingressgateway: entry point for external traffic
# - istio-egressgateway: controlled exit point

# Traffic flow:
# Client → Ingress Gateway → Sidecar (inbound) → App Container
# App Container → Sidecar (outbound) → Sidecar (inbound) → Destination

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

## fa-syringe Sidecar Injection

```yaml
# Automatic injection via namespace label
# kubectl label namespace default istio-injection=enabled

# Manual injection
# istioctl kube-inject -f deployment.yaml | kubectl apply -f -
```

```yaml
# Pod-level override
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

## fa-chart-line Observability

```yaml
# Prometheus annotations for scraping
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
# Kiali dashboard (service mesh visualization)
# istioctl dashboard kiali

# Jaeger tracing (distributed tracing)
# istioctl dashboard jaeger

# Envoy access logs
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

## fa-arrows-split-up-and-left Traffic Management Patterns

```yaml
# Canary release with weighted routing
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
# A/B testing by header
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
# Mirroring (shadow traffic)
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
