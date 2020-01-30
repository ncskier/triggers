# Failed Binding Example

```bash
kubectl apply -f examples/failed-binding/failed-binding.yaml
```

```bash
kubectl port-forward svc/el-failed-binding-eventlistener 8080:8080
```

```bash
curl -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{"param0": "0"}'
curl -X POST http://localhost:8080 -H 'Content-Type: application/json' -d '{"param1": "1"}'
```
