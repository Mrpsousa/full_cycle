## comandos
    - docker exec -it consul-01 sh
    - consul agent -dev
    - consul members -(infos) / curl localhost:8500/v1/catalog/nodes

## 3 agentes
    - dentro do container:
        mkdir /etc/consul.d 
        mkdir /var/lib/consul
    - ifconfig - pegar eth0 - inet addr: IPAQUI
    -> consul agent -server -bootstrap-expect=3 -node=consulserver01 -bind=172.19.0.1 -data-dir=/var/lib/consul -config-dir=/etc/consul.d

    - Para Cliente:
        -   ifconfig
        -   mkdir /etc/consul.d
        -   mkdir /var/lib/consul
        -   consul agent -bind=172.18.0.5 -data-dir=/var/lib/consul -config-dir=/etc/consul.d