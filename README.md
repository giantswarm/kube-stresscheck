# kube-stresscheck

Script to check Kubernetes nodes on stress (CPU/RAM) resistance.

It will run 5 iterations of `stress` command with following parameters:
- CPU forks equal to number of cores.
- Memory forks will be computed to allocate whole memory.

For details what `stress` does under the hood see [manual](https://linux.die.net/man/1/stress).

## Quick start
### Install
```
helm repo add kube-stresscheck https://giantswarm.github.io/kube-stresscheck
helm install kube-stresscheck kube-stresscheck/kube-stresscheck -n kube-system
```
### Delete
```
helm delete kube-stresscheck -n kube-system
```

## How it works
Wait at least 30 seconds and check for `CrashLooping` pods and `NotReady` nodes.

```
kubectl get pods -n kube-system
kubectl get nodes
```

Usually pods like `kube-proxy`, `nginx-ingress-controller`, `calico-node` are crashlooping. If kubelet or docker was affected by stress test then node will become `NotReady`.

## Advanced usage
### Stress test whole cluster

If you're really brave, then you can start stress check on the whole cluster (including master nodes).

```
helm install kube-stresscheck kube-stresscheck/kube-stresscheck -n kube-system --set kind=Daemonset
```

Most probably you have just KILLED your cluster. It will probably require manual intervention after this command.

If cluster is still ok, then remove stress check.
