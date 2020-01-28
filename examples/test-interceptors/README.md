# Test Interceptors Example

## Install
[Install `ko`](https://github.com/google/ko) if you have not already.
```bash
ko apply -f examples/test-interceptors
```

## Test
### Port-forward the EventListener service
```bash
kubectl port-forward svc/el-test-interceptors 8080:8080
```

### Curl the EventListener service
```bash
curl -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{"hello": "there"}'
```

### View the created TaskRun resource
```bash
kubectl get taskrun
```

### Clean up created resources
```bash
kubectl delete taskrun -l tekton.dev/trigger=test-interceptors
```
