+++
title = "KUBEEDGE - 简单了解Kubeedge的部分内容"
author = ["LIYUNFAN"]
description = "最近工作上做了一些调整，涉及到了基于Kubeedge和公司业务的一些定制开发、解决异常工作。所以也对Kubeedge进行了一定程度的了解。"
date = 2020-09-28
tags = ["KUBEEDGE"]
draft = false
+++

最近工作上做了一些调整，调整之后涉及到了Kubeedge相关的内容，公司业务上需要对边端节点进行管控，其中使用到了Kubeedge。
基于业务内容，我也对Kubeedge做了一些简单的了解。

<!--more-->

首先，Kubeedge是什么？

Kubeedge的官网上是这么描述的

> KubeEdge是一个开源系统，用于将容器化应用程序编排功能扩展到Edge的主机。它基于kubernetes构建，并为网络应用程序提供基础架构支持。云和边缘之间的部署和元数据同步。

先简单了解一下在构建边缘计算平台的时候，需要解决的一些问题。

1.  边端性能：边端节点性能普遍不高，无法部署整套kubernetes集群节点所需要的服务
2.  边端网络：云端与边端主机距离远导致的网络不稳定，导致节点与云端通信不顺，甚至会经常发生离线，而kubernetes本身依赖list-watch机制，无法离线运行
3.  边端设备数量：伴随着边端设备的不断增多，云端可能成为性能瓶颈无法满足业务需求

而Kubeedge通过引入cloudcore和edgecore两部分内容解决了前两个问题。

{{< figure src="https://docs.kubeedge.io/en/latest/%5Fimages/kubeedge%5Farch.png" link="https://docs.kubeedge.io/en/latest/%5Fimages/kubeedge%5Farch.png" >}}

对于 **边端设备数量** 的问题则是通过edgecore部分的 **EdgeD** 和 **MetadataManager** 组件实现。
**EdgeD** 作为简化版本的 **kubelet** 为边端提供了足够用的功能。
**MetadataManager** 则是提供了云端数据的本地存储，降低了边端对云端的压力，同时也为离线自治提供了对必要数据存储的功能。


## CloudCore {#cloudcore}

cloudcore主要由3部分构成

1.  **CloudHub**
2.  **EdgeController**
3.  **DeviceController**


### **CloudHub** {#cloudhub}

**CloudHub** 是cloudcore的一个组件，其主要作用是负责边端和云端的通信。
它支持Websocket和Quic两种通信协议，通过配置文件可以同时启用两种通信协议进行通信。

在从 **CloudHub** 发送消息到边端的通信过程中， **CloudHub** 将通信内容封装成为一个类似于Http包的结构（ **Header** 、 **Router** 、 **Content** ）。
通过Websocket协议发送到边端，再由边端对该包进行解析发送到指定的边端组件上进行处理。

而对于从 **CloudHub** 发送消息到云端其他组件的时候，则是通过管道直接发送到其他云端组件处理。


### **EdgeController** {#edgecontroller}

**EdgeController** 是edgecore部分和apiserver的一个中间层。

它主要负责将两个方面内容：

1.  将apiserver中的增、改、删操作通过 **CloudHub** 下发同步给edgecore。
2.  将edgecore部分上报的资源对象的状态信息和对资源对象的监听请求信息发送给apiserver。

除此之外， **EdgeController** 还通过内部的实现对ConfigMap、Pod和Node之间的绑定关系进行了管理。
当有Pod、ConfigMap相关资源的请求时， **EdgeController** 可以获知到该ConfigMap绑定到了哪些Pod上，进而将消息只下发到对应的节点。


### **DeviceController** {#devicecontroller}

由于工作上的业务只涉及到云端到边端而没有涉及到边端到设备，所以并没有了解非常多该方面内容...

该组件大体上是使用了CRD抽象出了设备相关信息，并通过与 **EdgeController** 类似的方式进行了与边端数据通信。


## EdgeCore {#edgecore}

edgecore主要由5部分构成

1.  EdgeHub
2.  MetaManager
3.  DeviceTwin
4.  EdgeD
5.  EventBus


### EdgeHub {#edgehub}

EdgeHub与 **CloudHub** 类似，是edgecore部分请求的集散地，主要负责边端与云端的通信。
它也支持Websocket和Quic两种通信协议，但只能在同一时间内使用一种进行通信。

它将会监听云端发来的数据，并通过解析请求消息内的路由规则，将数据包发送到边端不同的组件上。
同时也会将边端的消息通过Websocket发送给云端，交由云端处理。


### MetaManager {#metamanager}

MetaManager是EdgeD到edgehub消息的处理组件，主要作用是对云端到边端和边端到云端的消息进行处理。
中间可能会将一些信息存储到本地的SQLite中，在一些情况下可以降低云端的压力，同时也减少对云端的依赖，提升离线自治能力。

例如当云端对某个Pod进行更新时，当EdgeHub将信息交到MetaManager后，MetaManager会先通过resourceversion检查资源对象是否有改变。
如果有改变再将该资源对象的最新版本存储到SQLite中，之后交付到EdgeD对资源对象进行具体的操作。


### EdgeD {#edged}

EdgeD是边端对节点上资源对象管理的模块。作用个人感觉类似于kubelet，是真正对容器实施操作的模块（刚刚那些都是动嘴的，这是个跑腿的）。

不过为了适配边端低性能的硬件，Kubeedge对其进行了一定程度删减阉割和修改，在保证功能正常的前提下使其更轻量。


### EventBus、DeviceTwin {#eventbus-devicetwin}

由于工作上的业务只涉及到云端到边端而没有涉及到边端到设备，所以并没有了解非常多这两方面内容...

这两个模块看起来解决了边端设备相关信息同步的问题。包括设备到边端、设备到云端、云端到设备等方向的同步。


## 总结 {#总结}

Kubeedge是一个基于kubernetes的边缘计算平台。它通过cloudcore和edgecore两部分将云和边和端链接了起来。

由于无法保证云端到边端的网络状态，kubeedge选择在边端使用SQLite搭建本地存储。

当边端需要获取的资源对象在本地没有缓存，主动向云端发送消息进行请求后 或是 云端上对资源对象有创建、更新、删除的操作时，云端会主动下发资源对象的最新状态到边端。
边端会先使用SQLite数据库对结果进行存储，在存储成功之后在对真实的容器进行操作处理。

同时边端也会不时的把资源对象的状态内容推送到云端进行同步，保持两侧的数据一致。

这样即使在网络条件不好的情况下，边端也可以较快的获得其他组件相关的信息和状态。同时在离线状态也能有一定的自制能力。

暂时了解的内容就这么多，在之后的工作中有了解到新内容之后再更新...
