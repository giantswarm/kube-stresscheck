# kube-stresscheck

Script to check Kubernetes nodes on stress (CPU/RAM) resistance.

It will run 5 iterations of `stress` command with following parameters:
- CPU forks equal to number of cores.
- Memory forks will be computed to allocate whole memory.

For details what `stress` does under the hood see [manual](https://linux.die.net/man/1/stress).
