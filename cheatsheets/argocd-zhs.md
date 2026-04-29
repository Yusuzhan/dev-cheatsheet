---
title: ArgoCD
icon: fa-rotate
primary: "#EF7B4D"
lang: yaml
locale: zhs
---

## fa-circle-info 核心概念

```yaml
# GitOps：声明式、基于拉取的持续交付
# ArgoCD 监控 Git 仓库并同步 Kubernetes 清单

# 关键概念：
# - Application：来自 Git 源的一组 Kubernetes 资源
# - AppProject：具有 RBAC 和资源限制的逻辑分组
# - Sync：使集群状态与期望状态（Git）一致的过程
# - Health：应用资源的聚合健康状态
```

## fa-file-code Application CRD

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/org/my-manifests.git
    targetRevision: main
    path: overlays/production
  destination:
    server: https://kubernetes.default.svc
    namespace: my-app
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
```

## fa-layer-group App of Apps 模式

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: apps
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/org/gitops.git
    targetRevision: main
    path: apps
  destination:
    server: https://kubernetes.default.svc
    namespace: argocd
```

```yaml
# apps/app1.yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: app1
  namespace: argocd
spec:
  source:
    repoURL: https://github.com/org/app1.git
    targetRevision: main
    path: k8s
  destination:
    server: https://kubernetes.default.svc
    namespace: app1
```

## fa-shield-halved 项目

```yaml
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: my-project
  namespace: argocd
spec:
  description: 我的团队项目
  sourceRepos:
    - "https://github.com/my-org/*"
  destinations:
    - namespace: "my-app-*"
      server: https://kubernetes.default.svc
  clusterResourceWhitelist:
    - group: ""
      kind: Namespace
  namespaceResourceBlacklist:
    - group: ""
      kind: ResourceQuota
  roles:
    - name: admin
      policies:
        - "p, proj:my-project:admin, applications, *, my-project/*, allow"
      groups:
        - my-org:admins
```

## fa-arrows-rotate 同步策略

```yaml
spec:
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
      allowEmpty: false
    retry:
      limit: 3
      backoff:
        duration: 5s
        factor: 2
        maxDuration: 3m
    syncOptions:
      - CreateNamespace=true
      - PrunePropagationPolicy=foreground
      - PruneLast=true
      - ServerSideApply=true
      - RespectIgnoreDifferences=true
```

## fa-sliders 同步选项

```yaml
spec:
  syncPolicy:
    syncOptions:
      - CreateNamespace=true
      - PrunePropagationPolicy=foreground
      - PruneLast=true
      - ServerSideApply=true
      - RespectIgnoreDifferences=true
      - ApplyOutOfSyncOnly=true

  ignoreDifferences:
    - group: apps
      kind: Deployment
      jsonPointers:
        - /spec/replicas
    - group: ""
      kind: Secret
      jqPathExpressions:
        - .data.password
```

## fa-code-merge 多源应用

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: multi-source-app
  namespace: argocd
spec:
  project: default
  sources:
    - repoURL: https://github.com/org/helm-charts.git
      targetRevision: main
      path: charts/my-app
      helm:
        valueFiles:
          - $values/values/production.yaml
    - repoURL: https://github.com/org/values.git
      targetRevision: main
      ref: values
    - repoURL: https://github.com/org/kustomize.git
      targetRevision: main
      path: plugins
      kustomize:
        namePrefix: prod-
  destination:
    server: https://kubernetes.default.svc
    namespace: my-app
```

## fa-gears 配置管理

```yaml
# Kustomize
spec:
  source:
    kustomize:
      namePrefix: prod-
      nameSuffix: "-v1"
      images:
        - my-app=my-registry/my-app:v2.0.0
      commonLabels:
        env: production
      namespace: my-app

# Helm
spec:
  source:
    helm:
      valueFiles:
        - values.yaml
        - values-prod.yaml
      parameters:
        - name: image.tag
          value: "2.0.0"
      releaseName: my-app
      skipCrds: true

# Jsonnet / 纯 YAML — 只需指定路径
spec:
  source:
    path: manifests/
    directory:
      recurse: true
      exclude: "docs/*"
```

## fa-bell 通知

```yaml
apiVersion: argoproj.io/v1alpha1
kind: NotificationsConfiguration
metadata:
  name: default-notifications
  namespace: argocd
spec:
  triggers:
    - name: on-sync-failed
      condition: app.status.operationState.phase in ["Failed", "Error"]
      template: slack-template
    - name: on-health-degraded
      condition: app.status.health.status == "Degraded"
      template: slack-template
  templates:
    - name: slack-template
      slack:
        attachments: |
          [{
            "title": "{{.app.metadata.name}}",
            "title_link": "{{.context.argocdUrl}}/applications/{{.app.metadata.name}}",
            "color": "#f44336",
            "fields": [{
              "title": "Sync Status",
              "value": "{{.app.status.sync.status}}",
              "short": true
            }]
          }]
```

## fa-user-shield RBAC

```yaml
# argocd-cm ConfigMap
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-rbac-cm
  namespace: argocd
data:
  policy.default: role:readonly
  policy.csv: |
    p, role:admin, applications, *, */*, allow
    p, role:admin, clusters, *, *, allow
    p, role:dev, applications, get, my-project/*, allow
    p, role:dev, applications, sync, my-project/dev-* , allow
    p, role:dev, repositories, *, *, allow
    g, my-org:devs, role:dev
    g, my-org:admins, role:admin
```

## fa-key SSO / OIDC

```yaml
# argocd-cm ConfigMap
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-cm
  namespace: argocd
data:
  url: https://argocd.example.com
  oidc.config: |
    name: Okta
    issuer: https://dev-xxx.okta.com/oauth2/default
    clientID: xxxxx
    clientSecret: $oidc.okta.clientSecret
    requestedScopes:
      - openid
      - profile
      - email
      - groups
    requestedIDTokenClaims:
      groups:
        essential: true
```

## fa-terminal CLI 命令

```bash
argocd login argocd.example.com            # 登录 ArgoCD
argocd account update-password             # 更新密码

argocd app list                            # 列出应用
argocd app get my-app                      # 应用详情
argocd app create my-app \
  --repo https://github.com/org/app.git \
  --path k8s \
  --dest-namespace default \
  --dest-server https://kubernetes.default.svc

argocd app sync my-app                     # 同步应用
argocd app sync my-app --prune             # 同步并清理
argocd app diff my-app                     # 对比实际与 Git 状态
argocd app history my-app                  # 部署历史
argocd app rollback my-app <revision>      # 回滚
argocd app set my-app -p image.tag=v2.0   # 设置参数

argocd proj list                           # 列出项目
argocd proj create my-project              # 创建项目

argocd repo add https://github.com/org/app.git --username git --password token
argocd repo list                           # 列出仓库
```

## fa-heart-pulse 健康检查

```yaml
# 通过 resource.customizations 自定义健康检查
apiVersion: v1
kind: ConfigMap
metadata:
  name: argocd-cm
  namespace: argocd
data:
  resource.customizations: |
    cert-manager.io/Certificate:
      health.lua: |
        hs = {}
        if obj.status ~= nil then
          if obj.status.conditions ~= nil then
            for i, condition in ipairs(obj.status.conditions) do
              if condition.type == "Ready" and condition.status == "False" then
                hs.status = "Degraded"
                hs.message = condition.message
                return hs
              end
            end
          end
        end
        hs.status = "Progressing"
        return hs
```

## fa-link Resource Hook

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: db-migrate
  annotations:
    argocd.argoproj.io/hook: PreSync
    argocd.argoproj.io/hook-delete-policy: HookSucceeded
spec:
  template:
    spec:
      containers:
        - name: migrate
          image: my-app:latest
          command: ["./migrate", "up"]
      restartPolicy: Never
```

```yaml
# Hook 类型：PreSync, Sync, PostSync, SyncFail, Skip
# 删除策略：HookSucceeded, HookFailed, BeforeHookCreation
apiVersion: batch/v1
kind: Job
metadata:
  name: smoke-test
  annotations:
    argocd.argoproj.io/hook: PostSync
    argocd.argoproj.io/hook-delete-policy: BeforeHookCreation
spec:
  template:
    spec:
      containers:
        - name: test
          image: curlimages/curl
          command: ["curl", "-f", "http://my-app:80/health"]
      restartPolicy: Never
```
