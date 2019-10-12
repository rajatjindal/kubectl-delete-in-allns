# kubectl-remove-in-allns

## Warnings

```diff
- This is not recommented for use in Production
- There are no unit tests covering this (yet). please use it at your own risk.
```

This kubectl plugin can delete a given resource from all the namespaces

sample usage

```bash
$ ./kubectl-remove-in-allns configmap my-test-config
INFO[0000] configmap "my-test-config" deleted from namespace "default" 
INFO[0000] configmap "my-test-config" not found in namespace "kube-node-lease" 
INFO[0000] configmap "my-test-config" not found in namespace "kube-public" 
INFO[0000] configmap "my-test-config" deleted from namespace "kube-system" 
INFO[0000] configmap "my-test-config" deleted from namespace "test" 
```

supports following resources right now (PR welcome for supporting other resource types):
- configmaps
- secrets
- ingress
- deployments
