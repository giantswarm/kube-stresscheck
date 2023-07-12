# kube-stresscheck

Script to check Kubernetes nodes on stress (CPU/RAM) resistance.

It will run 5 iterations of `stress` command with following parameters:
- CPU forks equal to number of cores.
- Memory forks will be computed to allocate whole memory.

For details what `stress` does under the hood see [manual](https://linux.die.net/man/1/stress).

## Quick start

```
kubectl apply -f https://raw.githubusercontent.com/giantswarm/kube-stresscheck/master/examples/node.yaml
```

Wait at least 30 seconds and check for `CrashLooping` pods and `NotReady` nodes.

```
kubectl get pods -n kube-system
kubectl get nodes
```

```
kubectl delete -f https://raw.githubusercontent.com/giantswarm/kube-stresscheck/master/examples/node.yaml
```

Usually pods like `kube-proxy`, `ingress-nginx-controller`, `calico-node` are crashlooping. If kubelet or docker was affected by stress test then node will become `NotReady`.

## Stress test whole cluster

If you're really brave, then you can start stress check on the whole cluster (including a master node).

```
kubectl apply -f https://raw.githubusercontent.com/giantswarm/kube-stresscheck/master/examples/cluster.yaml
```

Most probably you have just KILLED your cluster. It will probably require manual intervention after this command.

If cluster is still ok, then remove stress check.
```
kubectl delete -f https://raw.githubusercontent.com/giantswarm/kube-stresscheck/master/examples/node.yaml
```
