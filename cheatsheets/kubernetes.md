---
title: Kubernetes
icon: fa-dharmachakra
primary: "#326CE5"
lang: bash
---

## fa-circle-info Cluster Info

```bash
kubectl version                          # show client & server version
kubectl cluster-info                     # show cluster endpoint URLs
kubectl get nodes                        # list cluster nodes
kubectl describe node node-1             # detailed node info
kubectl top nodes                        # node resource usage (needs metrics-server)
kubectl get namespaces                   # list namespaces
kubectl config get-contexts              # list kubeconfig contexts
kubectl config use-context my-cluster    # switch context
```

## fa-cube Pods

```bash
kubectl get pods                         # list pods in current namespace
kubectl get pods -A                      # list pods across all namespaces
kubectl get pods -o wide                 # show pod IP and node
kubectl describe pod my-pod              # detailed pod info & events
kubectl logs my-pod                      # pod logs (stdout)
kubectl logs my-pod -c sidecar           # logs from specific container
kubectl logs my-pod -f                   # stream logs (follow)
kubectl logs my-pod --since=1h           # logs from last hour
kubectl exec -it my-pod -- bash          # interactive shell in pod
kubectl port-forward my-pod 8080:80      # forward local 8080 to pod port 80
```

## fa-tags Labels & Selectors

```bash
kubectl get pods --show-labels           # show pod labels
kubectl label pod my-pod env=prod        # add label
kubectl label pod my-pod env=staging --overwrite  # update label
kubectl label pod my-pod env-            # remove label
kubectl get pods -l app=nginx            # filter by label
kubectl get pods -l 'env in (prod,staging)'       # label set query
kubectl get pods -l 'app=nginx,env!=dev'           # multiple selectors
```

## fa-arrow-down-wide-short Deployments

```bash
kubectl create deployment nginx --image=nginx:latest    # create deployment
kubectl get deployments                    # list deployments
kubectl describe deployment nginx          # deployment details
kubectl scale deployment nginx --replicas=5   # scale replicas
kubectl rollout status deployment nginx     # watch rollout progress
kubectl rollout history deployment nginx    # rollout revision history
kubectl rollout undo deployment nginx       # rollback to previous revision
kubectl rollout undo deployment nginx --to-revision=2  # rollback to specific revision
kubectl set image deployment/nginx nginx=nginx:1.25  # update container image
```

## fa-network-wired Services & Networking

```bash
kubectl get services                      # list services
kubectl expose deployment nginx --port=80 --target-port=80 --type=NodePort  # create service
kubectl expose deployment nginx --port=80 --type=LoadBalancer   # load balancer service
kubectl describe svc nginx                # service details
kubectl get endpoints nginx               # service endpoints (backend pods)

kubectl get ingress                       # list ingress resources
kubectl describe ingress my-ingress       # ingress details

kubectl get networkpolicies               # list network policies
```

## fa-hard-drive ConfigMaps & Secrets

```bash
kubectl create configmap my-config --from-literal=key1=value1       # from literal
kubectl create configmap my-config --from-file=config.yaml          # from file
kubectl create configmap my-config --from-env-file=.env             # from env file
kubectl get configmaps                    # list configmaps
kubectl describe configmap my-config      # view configmap data

kubectl create secret generic my-secret --from-literal=password=s3cret  # generic secret
kubectl create secret tls tls-secret --cert=tls.crt --key=tls.key  # TLS secret
kubectl get secrets                       # list secrets
kubectl describe secret my-secret         # view metadata (not values)
kubectl get secret my-secret -o jsonpath='{.data.password}' | base64 -d  # decode value
```

## fa-box-archive Volumes & Storage

```bash
kubectl get pv                            # list persistent volumes
kubectl get pvc                           # list persistent volume claims
kubectl describe pvc my-pvc               # PVC details

# storage classes
kubectl get storageclasses                # list available storage classes
kubectl describe storageclass standard    # storage class details
```

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce                      # RWO, ReadOnlyMany (ROX), ReadWriteMany (RWX)
  resources:
    requests:
      storage: 10Gi
  storageClassName: standard
```

## fa-cronjob Jobs & CronJobs

```bash
kubectl get jobs                          # list jobs
kubectl describe job my-job               # job details
kubectl delete job my-job                 # delete job
kubectl logs job/my-job                   # job pod logs

kubectl get cronjobs                      # list cronjobs
kubectl describe cronjob my-cron          # cronjob details
kubectl create job --from=cronjob/my-cron manual-run  # trigger manually
```

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: backup
spec:
  schedule: "0 2 * * *"                  # cron format: daily at 2 AM
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

## fa-wrench Troubleshooting

```bash
kubectl get events --sort-by='.lastTimestamp'   # cluster events sorted by time
kubectl get pods --field-selector=status.phase=Failed  # find failed pods
kubectl describe pod my-pod | grep -A5 "Events"  # pod events
kubectl logs my-pod --previous             # logs from previous crashed container
kubectl debug my-pod -it --image=busybox  # ephemeral debug container

# resource usage
kubectl top pods                          # pod CPU/memory usage
kubectl top pods -l app=nginx             # filtered by label

# cluster issues
kubectl get componentstatuses             # component health (deprecated in 1.19+)
kubectl get cs                            # shorthand
```

## fa-file-code Resource Management

```bash
kubectl apply -f manifest.yaml            # create/update from file
kubectl apply -f ./manifests/             # apply all files in directory
kubectl apply -k ./kustomization/         # apply kustomization
kubectl delete -f manifest.yaml           # delete resources from file
kubectl get all                           # list common resources in namespace

kubectl edit deployment nginx             # edit in editor (opens vim)
kubectl patch deployment nginx -p '{"spec":{"replicas":3}}'  # patch resource

# output formatting
kubectl get pods -o yaml                  # YAML output
kubectl get pods -o json                  # JSON output
kubectl get pods -o name                  # resource names only
kubectl get pods -o custom-columns=NAME:.metadata.name,STATUS:.status.phase
```

## fa-lightbulb Useful Aliases & Tips

```bash
# common aliases (add to ~/.bashrc)
alias k='kubectl'
alias kgp='kubectl get pods'
alias kgs='kubectl get svc'
alias kgd='kubectl get deployments'
alias kl='kubectl logs'
alias ke='kubectl exec -it'
alias kd='kubectl describe'
alias kaf='kubectl apply -f'
alias kdf='kubectl delete -f'

# quick resource creation
kubectl run test --image=busybox --rm -it --restart=Never -- sh  # temporary debug pod
kubectl run nginx --image=nginx --dry-run=client -o yaml > deploy.yaml  # generate manifest

# namespace shortcut
kubectl get pods -n kube-system           # specify namespace
kubectl -n monitoring get pods            # namespace flag anywhere
```
