---
title: Helm
icon: fa-anchor
primary: "#0F1689"
lang: yaml
---

## fa-folder-tree Chart Structure

```
my-chart/
├── Chart.yaml
├── values.yaml
├── values.schema.json
├── .helmignore
├── templates/
│   ├── _helpers.tpl
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   ├── configmap.yaml
│   ├── secret.yaml
│   ├── hpa.yaml
│   ├── NOTES.txt
│   └── tests/
│       └── test-connection.yaml
├── charts/            # chart dependencies
└── templates/
    └── ...
```

## fa-file-lines Chart.yaml

```yaml
apiVersion: v2
name: my-app
description: A Helm chart for my application
type: application
version: 1.0.0
appVersion: "3.1.0"
kubeVersion: ">=1.24.0"
dependencies:
  - name: postgresql
    version: "~> 14.0"
    repository: "https://charts.bitnami.com/bitnami"
    condition: postgresql.enabled
  - name: redis
    version: "~> 18.0"
    repository: "https://charts.bitnami.com/bitnami"
    alias: cache
maintainers:
  - name: dev-team
    email: team@example.com
```

## fa-list-check values.yaml

```yaml
replicaCount: 2

image:
  repository: myapp
  pullPolicy: IfNotPresent
  tag: "1.0.0"

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  create: true
  annotations: {}
  name: ""

service:
  type: ClusterIP
  port: 80

ingress:
  enabled: false
  className: nginx
  hosts:
    - host: myapp.local
      paths:
        - path: /
          pathType: Prefix

resources:
  limits:
    cpu: 200m
    memory: 256Mi
  requests:
    cpu: 100m
    memory: 128Mi

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 10
  targetCPUUtilizationPercentage: 80

postgresql:
  enabled: true
  auth:
    database: myapp
```

## fa-file-code Templates

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "my-chart.fullname" . }}
  labels:
    {{- include "my-chart.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      {{- include "my-chart.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "my-chart.selectorLabels" . | nindent 8 }}
    spec:
      containers:
        - name: {{ .Chart.Name }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          ports:
            - containerPort: 80
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
```

## fa-cube Built-in Objects

```yaml
{{ .Chart.Name }}
{{ .Chart.Version }}
{{ .Chart.AppVersion }}

{{ .Release.Name }}
{{ .Release.Namespace }}
{{ .Release.IsUpgrade }}
{{ .Release.IsInstall }}
{{ .Release.Revision }}

{{ .Values.replicaCount }}
{{ .Values.image.repository }}

{{ .Files.Get "config.ini" }}
{{ .Files.Glob "certs/**" }}
{{ .Capabilities.APIVersions.Has "apps/v1" }}
{{ .Capabilities.KubeVersion.Major }}
```

## fa-wand-magic-sparkles Functions & Pipelines

```yaml
{{ .Values.image.tag | default "latest" }}
{{ .Values.name | upper }}
{{ .Values.name | lower }}
{{ .Values.name | quote }}
{{ .Values.port | toString }}

{{ .Values.hosts | join ", " }}
{{ "hello-world" | replace "-" "_" }}
{{ "  hello  " | trim }}
{{ .Values.path | base }}

{{ .Values.items | toJson }}
{{ .Values.config | b64enc }}
{{ .Values.data | sha256sum }}

{{ .Values.annotations | toYaml | nindent 4 }}
{{ include "my-chart.labels" . | indent 4 }}

{{ lookup "v1" "ConfigMap" "default" "my-config" }}
{{ randAlphaNum 16 }}
{{ now | date "2006-01-02" }}
```

## fa-code-branch Flow Control

```yaml
{{- if .Values.ingress.enabled }}
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{ include "my-chart.fullname" . }}
  {{- with .Values.ingress.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
spec:
  rules:
    {{- range .Values.ingress.hosts }}
    - host: {{ .host }}
      http:
        paths:
          {{- range .paths }}
          - path: {{ .path }}
            pathType: {{ .pathType }}
            backend:
              service:
                name: {{ include "my-chart.fullname" $ }}
                port:
                  number: {{ $.Values.service.port }}
          {{- end }}
    {{- end }}
{{- end }}
```

## fa-puzzle-piece Named Templates

```yaml
{{- define "my-chart.labels" -}}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/version: {{ .Chart.AppVersion }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "my-chart.selectorLabels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "my-chart.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
```

## fa-link Hooks

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: {{ include "my-chart.fullname" . }}-pre-install
  annotations:
    "helm.sh/hook": pre-install
    "helm.sh/hook-weight": "-5"
    "helm.sh/hook-delete-policy": hook-succeeded
spec:
  template:
    spec:
      containers:
        - name: setup
          image: busybox
          command: ["sh", "-c", "echo 'Running pre-install hook'"]
      restartPolicy: Never
```

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: {{ include "my-chart.fullname" . }}-test
  annotations:
    "helm.sh/hook": test
spec:
  template:
    spec:
      containers:
        - name: test
          image: busybox
          command: ["wget", "-qO-", "http://{{ include \"my-chart.fullname\" . }}:{{ .Values.service.port }}/health"]
      restartPolicy: Never
```

## fa-sitemap Subcharts

```yaml
# parent chart: Chart.yaml
dependencies:
  - name: postgresql
    version: "~> 14.0"
    repository: "https://charts.bitnami.com/bitnami"
    condition: postgresql.enabled
    alias: db
  - name: redis
    version: "~> 18.0"
    repository: "https://charts.bitnami.com/bitnami"
```

```yaml
# parent chart: values.yaml — override subchart values
postgresql:
  enabled: true
  auth:
    database: myapp
    password: secret
  primary:
    persistence:
      size: 20Gi

redis:
  architecture: standalone
  auth:
    enabled: false
```

## fa-terminal CLI Commands

```bash
helm install my-app ./my-chart                # install chart
helm install my-app ./my-chart -n staging     # install into namespace
helm install my-app ./my-chart -f custom.yaml # with custom values
helm install my-app ./my-chart --set image.tag=2.0.0

helm upgrade my-app ./my-chart                # upgrade release
helm upgrade my-app ./my-chart --install      # install or upgrade
helm upgrade my-app ./my-chart --values prod.yaml --set replicaCount=5

helm rollback my-app 1                        # rollback to revision 1
helm rollback my-app 1 -n staging

helm list                                     # list releases
helm list -A                                  # all namespaces
helm status my-app                            # release status
helm history my-app                           # release history
helm show values ./my-chart                   # show default values
helm show all ./my-chart                      # show all chart info

helm uninstall my-app                         # uninstall release
helm uninstall my-app --keep-history          # keep release history
```

## fa-warehouse Repositories

```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add elastic https://helm.elastic.co
helm repo list                               # list repos
helm repo update                             # update repo index
helm repo remove bitnami                     # remove repo

helm search repo nginx                       # search in added repos
helm search hub nginx                        # search Artifact Hub
helm search repo nginx --versions            # show all versions

helm pull bitnami/nginx                      # download chart
helm pull bitnami/nginx --version=15.0.0     # specific version
helm pull bitnami/nginx --untar              # download and extract

helm dependency update ./my-chart            # update chart deps
helm dependency build ./my-chart             # build from lock file
helm dependency list ./my-chart              # list dependencies
```

## fa-vial Chart Testing

```bash
helm lint ./my-chart                         # lint chart
helm lint ./my-chart --strict                # treat warnings as errors

helm template ./my-chart                     # render templates locally
helm template ./my-chart --debug             # debug template rendering
helm template ./my-chart -f dev.yaml         # with custom values
helm template ./my-chart --show-only templates/deployment.yaml

helm test my-app                             # run chart tests
helm test my-app -n staging

helm plugin install https://github.com/helm-unittest/helm-unittest
helm unittest ./my-chart
```

## fa-shield-halved Values Schema

```json
{
  "$schema": "https://json-schema.org/draft-07/schema#",
  "type": "object",
  "required": ["image"],
  "properties": {
    "replicaCount": {
      "type": "integer",
      "minimum": 1,
      "maximum": 100,
      "default": 1
    },
    "image": {
      "type": "object",
      "required": ["repository"],
      "properties": {
        "repository": {
          "type": "string"
        },
        "tag": {
          "type": "string"
        },
        "pullPolicy": {
          "type": "string",
          "enum": ["Always", "IfNotPresent", "Never"]
        }
      }
    },
    "service": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "enum": ["ClusterIP", "NodePort", "LoadBalancer"]
        },
        "port": {
          "type": "integer",
          "minimum": 1,
          "maximum": 65535
        }
      }
    }
  }
}
```
