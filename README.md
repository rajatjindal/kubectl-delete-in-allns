# kubectl-remove-in-allns

<span style="color:red; font:24px; ">This has not been tested thoroughly. Not recommended for use in Production and use at your own risk. </span>

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