# ops-meta   

管理ops-updater的组件   
接收ops-updater汇报上来的agent real state，返回最新的agent desired state

## 设计理念

- 对于一个公司而言，agent并不多，也就有个监控agent、部署agent、naming agent，所以ops-meta直接采用配置文件而不是数据库之类的大型存储来存放agent信息
- 公司级别agent升级慢一点没关系，比如一晚上升级完问题都不大，所以ops-updater与ops-meta的通信周期默认是5min，比较长。如果做成长连接，周期调小，是否就可以不光用来部署agent，也可以部署一些业务程序？不要这么做！部署其他业务组件是部署agent的责任，ops-updater做的事情少才不容易出错。ops-updater推荐在装机的时候直接安装好，功能少基本不升级。
- 配置文件中针对各个agent有个default配置，有个others配置，这个others配置是为了解决小流量问题，对于某些前缀的机器可以采用与default不同的配置，也就间接解决了小流量测试问题
- ops-updater会汇报自己管理的各个agent的状态、版本号，这个信息直接存放在ops-meta模块的内存中，因为数据量真没多少，100w机器，3个agent……

## 使用方法

- 1. 把要升级的agent打好tarball，交给http server
- 2. agent命名规范是`<agent-name>-<version>.tar.gz`，md5生成方式和命名：`md5sum <agent-name>-<version>.tar.gz > <agent-name>-<version>.tar.gz.md5`，比如：falcon-agent，全名：falcon-agent-1.0.0.tar.gz
- 3. 修改ops-meta的配置文件，agent太重要了，最好有个admin专门来审核、上线

agent tarball最终下载地址是：`{$tarball}/${arch}/{$name}-{$version}.tar.gz`，为啥不在tarball这里配置成全路径呢？为了规范！就是这么横！

