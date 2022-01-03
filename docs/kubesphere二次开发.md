## Kubesphere二次开发
>  编译ks-apiserver
```shell
make ks-apiserver
```
> 运行ks-apiserver
```shell
cat > /etc/kubesphere/kubesphere.yaml << EOF
authentication:
  authenticateRateLimiterMaxTries: 10
  authenticateRateLimiterDuration: 10m0s
  loginHistoryRetentionPeriod: 168h
  maximumClockSkew: 10s
  multipleLogin: True
  kubectlImage: kubesphere/kubectl:v1.21.0
  jwtSecret: "yElyXC6H9fcp6UbpY9EJcgDOBvjGEt6n"
authorization:
  mode: "AlwaysAllow"

network:
  ippoolType: none
  
monitoring:
  endpoint: http://prometheus-operated.kubesphere-monitoring-system.svc:9090
  enableGPUMonitoring: false
EOF
```
> 启动
```shell
./ks-apiserver --kubeconfig ~/.kube/config
```
***[调高日志等级]***
```shell
./ks-apiserver --kubeconfig ~/.kube/config --v 8
```
> 接口测试验证
- 查询当前开发的组件
```shell
curl -v http://127.0.0.1:9090/kapis/config.kubesphere.io/v1alpha2/configs/configz
```
- 查看当前所有deployment名称
**[非认证查询]**
```shell
curl -v http://127.0.0.1:9090/kapis/resources.kubesphere.io/v1alpha3/deployments
```
**[认证查询]**
```shell
curl -v -u "admin:P@88w0rd" http://127.0.0.1:9090/kapis/resources.kubesphere.io/v1alpha3/deployments
```
- 使用telepresence连同ks-console调试
```shell
sudo telepresence --namespace kubesphere-system --swap-deployment  ks-apiserver --expose 9090
```
> 查看telepresence状态
```shell
sudo telepresence status
```
> 退出telepresence调试模式
```shell
sudo telepresence quit
```
