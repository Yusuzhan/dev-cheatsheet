---
title: ArgoCD
icon: fa-rotate
primary: "#EF7B4D"
lang: yaml
---

## fa-circle-info Core Concepts

```yaml
# GitOps: declarative, pull-based continuous delivery
# ArgoCD watches Git repos and syncs Kubernetes manifests

# Key concepts:
# - Application: a group of Kubernetes resources from a Git source
# - AppProject: logical grouping with RBAC and resource restrictions
# - Sync: process of making cluster state match desired state (Git)
# - Health: aggregated status of application resources
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

## fa-layer-group App of Apps Pattern

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

## fa-shield-halved Projects

```yaml
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: my-project
  namespace: argocd
spec:
  description: My team project
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

## fa-arrows-rotate Sync Policies

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

## fa-sliders Sync Options

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

## fa-code-merge Multi-source Apps

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

## fa-gears Config Management

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

# Jsonnet / Plain YAML — just point to path
spec:
  source:
    path: manifests/
    directory:
      recurse: true
      exclude: "docs/*"
```

## fa-bell Notifications

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

## fa-terminal CLI Commands

```bash
argocd login argocd.example.com            # login to ArgoCD
argocd account update-password             # update password

argocd app list                            # list applications
argocd app get my-app                      # app details
argocd app create my-app \
  --repo https://github.com/org/app.git \
  --path k8s \
  --dest-namespace default \
  --dest-server https://kubernetes.default.svc

argocd app sync my-app                     # sync application
argocd app sync my-app --prune             # sync and prune
argocd app diff my-app                     # diff live vs Git state
argocd app history my-app                  # deployment history
argocd app rollback my-app <revision>      # rollback
argocd app set my-app -p image.tag=v2.0   # set parameter

argocd proj list                           # list projects
argocd proj create my-project              # create project

argocd repo add https://github.com/org/app.git --username git --password token
argocd repo list                           # list repos
```

## fa-heart-pulse Health Checks

```yaml
# Custom health check via resource.customizations
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

## fa-link Resource Hooks

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
# Hook types: PreSync, Sync, PostSync, SyncFail, Skip
# Delete policies: HookSucceeded, HookFailed, BeforeHookCreation
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
