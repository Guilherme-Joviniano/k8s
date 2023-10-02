# K8S é baseado em states, configurando o state de cada objeto

## Kubernetes Master

Kube Api-Server
Kube Controller Manager
Kube Scheduler

## Outros Nodes

Kubelet
Kubeproxy

# Cluster

Conjunto de Máquinas (Node)
Cada node possui uma quantidade de vCPU e Memória

# Pods

dentro das nodes, é a unidade que contêm os containers
representando a menor instancia de processamento

# Deployment

Configuração dos Pods

ReplicaSets
B = backend => 3 réplicas
F = frontend => 2 réplicas

Realocação dos pods automaticamente e dinamismo em relação aos recursos disponíveis

## Hierarquia de Funcionamento

Cluster[Nodes[Service[Deployment --> ReplicaSet --> Pod ]]]

Camada de service funciona
como um load balancer

## Configuração

Variáveis de ambiente no k8s 