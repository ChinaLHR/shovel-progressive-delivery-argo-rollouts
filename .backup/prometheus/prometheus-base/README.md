## prometheus-configmap-job

- kubernetes-kubelet: 从集群各节点kubelet组件中获取节点kubelet的基本运行状态的监控指标。（通过Kubernetes的api-server提供的代理API访问各个节点中kubelet的metrics服务）。
- kubernetes-cadvisor: 从集群各节点kubelet内置的cAdvisor中获取，节点中运行的容器的监控指标。（通过api-server提供的代理地址访问kubelet的/metrics/cadvisor地址）。
- kubernetes-pods: 从部署到各个节点的Node Exporter中采集主机资源相关的运行资源。
- kubernetes-apiservers: 获取API Server组件的访问地址，并从中获取Kubernetes集群相关的运行监控指标。