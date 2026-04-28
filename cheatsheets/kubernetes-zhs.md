---
title: Kubernetes
icon: fa-dharmachakra
primary: "#326CE5"
lang: bash
locale: zhs
---

## fa-circle-info 集群信息

```bash
kubectl version                          # 显示客户端和服务端版本
kubectl cluster-info                     # 显示集群端点 URL
kubectl get nodes                        # 列出集群节点
kubectl describe node node-1             # 节点详细信息
kubectl top nodes                        # 节点资源使用（需要 metrics-server）
kubectl get namespaces                   # 列出命名空间
kubectl config get-contexts              # 列出 kubeconfig 上下文
kubectl config use-context my-cluster    # 切换上下文
```

## fa-cube Pod 管理

```bash
kubectl get pods                         # 列出当前命名空间的 Pod
kubectl get pods -A                      # 列出所有命名空间的 Pod
kubectl get pods -o wide                 # 显示 Pod IP 和所在节点
kubectl describe pod my-pod              # Pod 详细信息和事件
kubectl logs my-pod                      # 查看 Pod 日志（stdout）
kubectl logs my-pod -c sidecar           # 查看指定容器日志
kubectl logs my-pod -f                   # 实时跟踪日志
kubectl logs my-pod --since=1h           # 最近 1 小时的日志
kubectl exec -it my-pod -- bash          # 进入 Pod 交互式终端
kubectl port-forward my-pod 8080:80      # 本地 8080 转发到 Pod 80 端口
```

## fa-tags 标签与选择器

```bash
kubectl get pods --show-labels           # 显示 Pod 标签
kubectl label pod my-pod env=prod        # 添加标签
kubectl label pod my-pod env=staging --overwrite  # 更新标签
kubectl label pod my-pod env-            # 删除标签
kubectl get pods -l app=nginx            # 按标签过滤
kubectl get pods -l 'env in (prod,staging)'       # 标签集合查询
kubectl get pods -l 'app=nginx,env!=dev'           # 多选择器组合
```

## fa-arrow-down-wide-short Deployment

```bash
kubectl create deployment nginx --image=nginx:latest    # 创建 Deployment
kubectl get deployments                    # 列出 Deployment
kubectl describe deployment nginx          # Deployment 详情
kubectl scale deployment nginx --replicas=5   # 扩缩副本数
kubectl rollout status deployment nginx     # 查看滚动更新状态
kubectl rollout history deployment nginx    # 查看更新历史
kubectl rollout undo deployment nginx       # 回滚到上一版本
kubectl rollout undo deployment nginx --to-revision=2  # 回滚到指定版本
kubectl set image deployment/nginx nginx=nginx:1.25  # 更新容器镜像
```

## fa-network-wired Service 与网络

```bash
kubectl get services                      # 列出 Service
kubectl expose deployment nginx --port=80 --target-port=80 --type=NodePort  # 创建 NodePort 服务
kubectl expose deployment nginx --port=80 --type=LoadBalancer   # 创建 LoadBalancer 服务
kubectl describe svc nginx                # Service 详情
kubectl get endpoints nginx               # Service 后端端点

kubectl get ingress                       # 列出 Ingress
kubectl describe ingress my-ingress       # Ingress 详情

kubectl get networkpolicies               # 列出网络策略
```

## fa-hard-drive ConfigMap 与 Secret

```bash
kubectl create configmap my-config --from-literal=key1=value1       # 从键值对创建
kubectl create configmap my-config --from-file=config.yaml          # 从文件创建
kubectl create configmap my-config --from-env-file=.env             # 从 env 文件创建
kubectl get configmaps                    # 列出 ConfigMap
kubectl describe configmap my-config      # 查看 ConfigMap 数据

kubectl create secret generic my-secret --from-literal=password=s3cret  # 创建通用 Secret
kubectl create secret tls tls-secret --cert=tls.crt --key=tls.key  # 创建 TLS Secret
kubectl get secrets                       # 列出 Secret
kubectl describe secret my-secret         # 查看元数据（不显示值）
kubectl get secret my-secret -o jsonpath='{.data.password}' | base64 -d  # 解码值
```

## fa-box-archive 存储与数据卷

```bash
kubectl get pv                            # 列出持久卷
kubectl get pvc                           # 列出持久卷声明
kubectl describe pvc my-pvc               # PVC 详情

# 存储类
kubectl get storageclasses                # 列出可用存储类
kubectl describe storageclass standard    # 存储类详情
```

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce                      # RWO 单节点读写, ROX 多节点只读, RWX 多节点读写
  resources:
    requests:
      storage: 10Gi
  storageClassName: standard
```

## fa-cronjob Job 与 CronJob

```bash
kubectl get jobs                          # 列出 Job
kubectl describe job my-job               # Job 详情
kubectl delete job my-job                 # 删除 Job
kubectl logs job/my-job                   # Job Pod 日志

kubectl get cronjobs                      # 列出 CronJob
kubectl describe cronjob my-cron          # CronJob 详情
kubectl create job --from=cronjob/my-cron manual-run  # 手动触发一次
```

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: backup
spec:
  schedule: "0 2 * * *"                  # cron 格式：每天凌晨 2 点
  jobTemplate:
    spec:
      backoffLimit: 3
      template:
        spec:
          containers:
            - name: backup
              image: backup-tool:latest
              command: ["/bin/sh", "-c", "backup.sh"]
          restartPolicy: OnFailure
```

## fa-wrench 故障排查

```bash
kubectl get events --sort-by='.lastTimestamp'   # 按时间排序的集群事件
kubectl get pods --field-selector=status.phase=Failed  # 查找失败的 Pod
kubectl describe pod my-pod | grep -A5 "Events"  # Pod 事件
kubectl logs my-pod --previous             # 上一次崩溃容器的日志
kubectl debug my-pod -it --image=busybox  # 临时调试容器

# 资源使用
kubectl top pods                          # Pod CPU/内存使用
kubectl top pods -l app=nginx             # 按标签过滤

# 集群问题
kubectl get componentstatuses             # 组件健康状态（1.19+ 已弃用）
kubectl get cs                            # 简写
```

## fa-file-code 资源管理

```bash
kubectl apply -f manifest.yaml            # 从文件创建/更新资源
kubectl apply -f ./manifests/             # 应用目录下所有文件
kubectl apply -k ./kustomization/         # 应用 kustomization
kubectl delete -f manifest.yaml           # 从文件删除资源
kubectl get all                           # 列出命名空间中常见资源

kubectl edit deployment nginx             # 在编辑器中修改（打开 vim）
kubectl patch deployment nginx -p '{"spec":{"replicas":3}}'  # 补丁更新

# 输出格式
kubectl get pods -o yaml                  # YAML 输出
kubectl get pods -o json                  # JSON 输出
kubectl get pods -o name                  # 仅资源名称
kubectl get pods -o custom-columns=NAME:.metadata.name,STATUS:.status.phase
```

## fa-lightbulb 实用别名与技巧

```bash
# 常用别名（添加到 ~/.bashrc）
alias k='kubectl'
alias kgp='kubectl get pods'
alias kgs='kubectl get svc'
alias kgd='kubectl get deployments'
alias kl='kubectl logs'
alias ke='kubectl exec -it'
alias kd='kubectl describe'
alias kaf='kubectl apply -f'
alias kdf='kubectl delete -f'

# 快速创建资源
kubectl run test --image=busybox --rm -it --restart=Never -- sh  # 临时调试 Pod
kubectl run nginx --image=nginx --dry-run=client -o yaml > deploy.yaml  # 生成清单文件

# 命名空间简写
kubectl get pods -n kube-system           # 指定命名空间
kubectl -n monitoring get pods            # namespace 参数位置灵活
```
