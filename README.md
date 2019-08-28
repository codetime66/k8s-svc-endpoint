# k8s-svc-endpoint

go mod init github.com/codetime66/k8s-svc-endpoint

./bin/k8s-svc-endpoint --help
Usage of ./bin/k8s-svc-endpoint:
  -kubeconfig string
    	(optional) absolute path to the kubeconfig file (default "/home/carlosfe/.kube/config")


./bin/k8s-svc-endpoint --kubeconfig ~/projects/kubeland/hzubernetes/.kube/conf --service-name api/v1/namespaces/credenciamento/endpoints/credenciamento-validacao-telefone

---------
#metrics:
kubectl --kubeconfig ~/projects/kubeland/hzubernetes/.kube/conf -n infra exec -it mytool-fdf57d5bb-hzb7p -- curl -v -k https://kubernetes.default.svc:443/apis/metrics.k8s.io/v1beta1/pods

#k8s api/v1
kubectl --kubeconfig ~/projects/kubeland/hzubernetes/.kube/conf -n infra exec mytool-648f5658b8-4m5b4 -- curl -v -k https://kubernetes.default.svc:443/api/v1/

kubectl --kubeconfig ~/projects/kubeland/hzubernetes/.kube/conf -n infra exec mytool-648f5658b8-4m5b4 -- curl -v -k https://kubernetes.default.svc:443/api/v1/endpoints

kubectl --kubeconfig ~/projects/kubeland/hzubernetes/.kube/conf -n infra exec mytool-648f5658b8-4m5b4 -- curl -v -k https://kubernetes.default.svc:443/api/v1/namespaces/credenciamento/endpoints/gsurf-svc
