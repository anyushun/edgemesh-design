## 用户体验
随着微服务框架的流行和边缘节点计算和网络能力的增强，把部分服务部署到边缘处理设备产生的数据轻松解决时延的问题。但由于安全等原因，大多数服务仍然运行在云上。**云**跟**边**、**边**跟**边**之间的协同是必须要解决的问题，edgemesh正是为此而生。用户通过**服务名称**来访问通过kubeedge部署的服务，甚至脱离于kubeedge的自维护服务，而不需要关心服务部署到哪里

![user view](docs/images/user%20view.jpg)

## 节点内协同

![edge](docs/images/inedge.jpg)

## 边边协同

![edge2edge](docs/images/edge2edge.jpg)

## 边云协同

![edge2cloud](docs/images/edge2cloud.jpg)

## edgemesh内部实现

![edgemesh](docs/images/edgemesh.jpg)