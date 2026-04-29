---
title: etcd
icon: fa-key
primary: "#419EDA"
lang: bash
---

## fa-terminal CLI Basics (etcdctl)

```bash
etcdctl put mykey "hello world"
etcdctl get mykey
etcdctl del mykey

ETCDCTL_API=3 etcdctl --endpoints=http://127.0.0.1:2379 put foo bar
etcdctl --cacert=/etc/etcd/ca.crt --cert=/etc/etcd/client.crt --key=/etc/etcd/client.key get foo

etcdctl endpoint health
etcdctl endpoint status --write-out=table
etcdctl endpoint hashkv --write-out=table
```

## fa-pen-to-square Put / Get / Delete

```bash
etcdctl put /config/db/host "db.example.com"
etcdctl put /config/db/port "5432"
etcdctl get /config/db/host

etcdctl get /config/db/ --prefix
etcdctl get /config/ --prefix --keys-only
etcdctl get /config/db/host --print-value-only

etcdctl del /config/db/host
etcdctl del /config/db/ --prefix
etcdctl del /config/ --prefix --prev-kv

etcdctl get "" --prefix --limit=10
etcdctl get /config/ --prefix --sort-by=KEY --order=DESCEND
```

## fa-eye Watch

```bash
etcdctl watch /config/db/host
etcdctl watch /config/ --prefix

etcdctl watch /config/ --prev-kv

etcdctl watch -i <<EOF
watch /config/db/host
watch /config/db/port
EOF

etcdctl get /config/db/host -w=json
```

## fa-clock Lease

```bash
LEASE_ID=$(etcdctl lease grant 300 | grep -oP 'ID \K\S+')
etcdctl put /session/key "value" --lease=$LEASE_ID
etcdctl get /session/key

etcdctl lease list
etcdctl lease timetolive $LEASE_ID
etcdctl lease timetolive $LEASE_ID --keys

etcdctl lease revoke $LEASE_ID
```

## fa-heart-pulse Keep Alive

```bash
LEASE_ID=$(etcdctl lease grant 60 | grep -oP 'ID \K\S+')
etcdctl put /service/node1 "10.0.0.1:8080" --lease=$LEASE_ID

etcdctl lease keep-alive $LEASE_ID &

etcdctl lease timetolive $LEASE_ID
```

## fa-code-branch Transactions

```bash
etcdctl txn <<EOF
compares:
value("/config/ready") = "true"
success:
put /config/db/host "db.example.com"
failure:
put /config/error "not ready"
EOF

etcdctl txn -i <<EOF
compare:
mod("key1") > "0"
success:
put key1 "updated"
get key1
EOF
```

## fa-compress Compaction

```bash
etcdctl get mykey -w=json

REVISION=$(etcdctl get mykey -w=json | python3 -c "import sys,json; print(json.load(sys.stdin)['header']['revision'])")
etcdctl compact $REVISION

etcdctl get mykey --rev=$((REVISION-1))
etcdctl get "" --prefix --write-out=json | python3 -c "
import sys,json
d=json.load(sys.stdin)
for kv in d.get('kvs',[]):
    print(f\"{kv['key'].decode()}={kv['value'].decode()}\")
"
```

## fa-broom Defragment

```bash
etcdctl defrag
etcdctl defrag --cluster

etcdctl endpoint status --write-out=table

du -sh /var/lib/etcd/member/
etcdctl endpoint status -w table | grep DB_SIZE
```

## fa-camera Snapshot & Restore

```bash
etcdctl snapshot save /backup/etcd-$(date +%Y%m%d).db
etcdctl snapshot status /backup/etcd-snapshot.db --write-out=table

etcdctl snapshot restore /backup/etcd-snapshot.db \
  --name etcd-node1 \
  --initial-cluster etcd-node1=http://10.0.0.1:2380 \
  --initial-advertise-peer-urls http://10.0.0.1:2380 \
  --data-dir /var/lib/etcd/new-data
```

## fa-server Cluster Membership

```bash
etcdctl member list --write-out=table
etcdctl member add etcd-node4 --peer-urls=http://10.0.0.4:2380
etcdctl member remove abc123

etcdctl member update abc123 --peer-urls=http://10.0.0.1:2381

etcdctl endpoint health --cluster --write-out=table
etcdctl endpoint status --cluster --write-out=table
```

## fa-user-shield Auth & RBAC

```bash
etcdctl user add root
etcdctl auth enable

etcdctl user add appuser --interactive=false --new-user-password=secret123
etcdctl user list
etcdctl user get appuser

etcdctl role add approle
etcdctl role grant-permission approle readwrite --prefix=true /app/
etcdctl role get approle
etcdctl user grant-role appuser approle

etcdctl auth disable
etcdctl --user=root:secret get /app/config
```

## fa-hourglass TTL Keys

```bash
etcdctl put /cache/user:123 "json_data" --lease=$(etcdctl lease grant 60 | grep -oP 'with TTL \K\S+')

LEASE_ID=$(etcdctl lease grant 120 | awk '/ID/{print $2}')
etcdctl put /lock/leader "node1" --lease=$LEASE_ID
etcdctl lease timetolive $LEASE_ID

etcdctl get /lock/leader -w=json
```

## fa-lock Distributed Lock

```bash
etcdctl lock /locks/migrate --ttl=60 -- /usr/local/bin/migrate.sh

etcdctl lock /locks/deploy --ttl=300 -- /opt/scripts/deploy.sh

LEASE_ID=$(etcdctl lease grant 30 | awk '/ID/{print $2}')
etcdctl put /locks/resource "node1-owner" --lease=$LEASE_ID
etcdctl get /locks/resource --prefix
etcdctl lease revoke $LEASE_ID
```

## fa-crown Leader Election

```bash
etcdctl elect /elections/leader "candidate-1" --ttl=60

etcdctl elect /elections/primary node-a --ttl=30 &
ELECT_PID=$!

etcdctl get /elections/leader --prefix
kill $ELECT_PID
```
