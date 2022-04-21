## Instalation
    - wget -q -O - https://raw.githubusercontent.com/rancher/k3d/main/install.sh | bash 
    OR
    - install kubeclt : https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/


## Build image
    - docker build -t mrpsousa/hello-go 

## step by step
- build app image(s) / docker build -t go-server . /  docker tag go-server mrpsousa/go-server / docker push
- k3d cluster create
- kubectl apply -f k8s/pod.yaml
- kubectl get pod

## Commands kubernets
    - "up" pod - kubectl apply -f k8s/pod.yaml
    - list pod - kubectl get pod
    - kubectl port-forward pod/goserver 8000:80
    - kubectl delete pod NAME

## debug
    - kubectl get pods --field-selector=status.phase=Pending
    - kubectl describe pod
    - kubectl get nodes
    - kubectl describe node <node-name>
    - kubectl describe node <nodename> | grep Taints
    - kubectl taint node fullcycle-control-plane node-role.kubernetes.io/master:NoSchedule-


