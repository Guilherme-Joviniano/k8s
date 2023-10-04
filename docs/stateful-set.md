# **Stateless** **&** **Stateful**

Pod ==> [database]<->[MySQL Master]
Pod ==> [database]<->[MySQL Slave 1]
Pod ==> [database]<->[MySQL Slave 2]

## Como posso saber quem é o master e os slaves nos pods?

O tempo de carregamento, sempre o master sobre primeiro e as ordem dos slaves sobem de maneira sincronizada.
Porém as replicas sobem de maneira automática e não seguem o processo síncrono.

## Mecanismo para processo ordenados de subidas, para que a HPA tenha ordem

Precisamos regular a escala e o "down" dos pods, de maneira ordenada no Stateful!

Para isso utilizamos o StatefulSet. Quase parecido com o deployment porém com as ordens de subidas e os nomes sendo registrados de maneira ordenada.

## Camada de Service Dentro do StatefulSet

utilizado para todos os pods
[MySQL SERVICE]

Pod ==> [database]<->[MySQL Master]
Pod ==> [database]<->[MySQL Slave 1]
Pod ==> [database]<->[MySQL Slave 2]

## Headless service, criando um service para cada POD!
DNS --> headless service --> POD 1 
    --> headless service --> POD 2 