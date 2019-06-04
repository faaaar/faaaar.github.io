---
title: Golang - 内存管理
date: 2019-06-05
updated: 2019-06-05
categories:
    - Go
tags:
    - 内存分配
    - 垃圾回收
---

# 内存管理

查看的源码版本为1.12

## 基础结构

在`/src/runtime/malloc.go`文件中我们可以看到，内存分配器的主要数据结构如下

- fixalloc a free-list allocator for fixed-size off-heap objects, used to manage storage used by the allocator.
- mheap: the malloc heap, managed at page (8192-byte) granularity.
- mspan: a run of pages managed by the mheap.
- mcentral: collects all spans of a given size class.
- mcache: a per-P cache of mspans with free space.
- mstats: allocation statistics.

### mspan

[mspan(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L316 "mspan(点击查看结构定义)")

在go的内存管理中mspan是一种基本单位，是一个用于分配对象的区块。它存储了一系列同样大小的元素。下图为mspan的结构

![mspan-struct](/images/go/gc/mspan-struct.png "mspan-struct")

其中比较重要的点

- next、prev 指针域，指向前一个和后一个mspan，说明mspan是一个双向链表。
- span中存储的页的数量
- npages: mspan 的大小为 page 大小的整数倍。
- nelems span中可以存储元素的数量
- freeindex 是一个大小在0到nelems的数字，用来标记下一次分配空间时开始扫描的位置，可以简单的理解为分配到了第几个区块。如果freeindex==nelem,说明当前span已经没有空闲空间了。
- spanclass 0~_NumSizeClasses之间的一个值。是当前mspan的类型。不同类型的mspan所分割的区块的大小也不一样。该值与class_to_size中存储的值相关。
- elemsize 元素大小。根据sizeclass计算得出。在分配大对象的时候，该值与npages相关，值为pagesize*npages。

### mcache

[mcache(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcache.go#L19 "mcache(点击查看结构定义)")

在Go的GPM模型中可以知道，P是一个虚拟资源，同一时间只能有一个G来访问P，所以P中的资源不需要加锁。为了在分配对象的时候有更好的性能，则每个P都会有一个mspan的缓存mcache

在mspan的结构定义中我们可以看到，alloc存储着指向mspan的指针的长度为[numSpanClasses](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L614 "numSpanClasses")的数组。而numSpanClasses的长度与[_NumSizeClasses](https://github.com/golang/go/blob/go1.12.5/src/runtime/sizeclasses.go#L79 "_NumSizeClasses")有关。计算可得出numSpanClasses的值为134。这134个值分别对应着67个类型的scan和noscan的值。如下图

![mcache-alloc](/images/go/gc/mcache-alloc.png "mcache-alloc")

其中scan和noscan的区别在于，如果对象包含了指针，分配对象时会使用scan的span，如果对象不包含指针, 分配对象时会使用noscan的span。
而把span分为scan和noscan的意义在于，GC扫描对象的时候对于noscan的span可以不去查看bitmap区域来标记子对象，这样可以大幅提升标记的效率。

### mcentral
[mcentral(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcentral.go#L20 "mcentral(点击查看结构定义)")

结构体分析

- lock mcentral是带锁的，因为会有多个P来访问它，应保证并发安全。
- spanclass 与mspan处的spanclass一致，所以我们可以推测mcentral也与mcache的alloc类似，有`numSpanClasses`个
- nonempty 类型为[mSpanList](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L310 "mSpanList")的一个列表。其中包含了指向列表头和尾的指针。这个列表中存储了空闲的mspan的指针，当mcache中的mspan不够用的时候，可以向mcentral申请更多的空间。
- empty 类型与nonempty相同，然而是存储了已经被申请走使用中(或者被mcache缓存起来)的mspan的指针。

mcentral只要结合了nonempty和empty两个字段中的数据，就可以追踪到当前mcentral所分配出去的所有对象。

### mheap
[mheap(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L31 "mheap(点击查看结构定义)")

mheap在系统初始化的时候([mallocinit](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L359 "mallocinit"))进行初始化，是一个全局变量。

结构体分析

- lock 全局变量，应当保证并发安全
- allspans 存储了创建的所有的mspans，并且是手动管理的，可以重新分配，并伴随着堆的增长而不断更新存储位置。
- central 如前面所说，central也存在`numSpanClasses`个
- arenaHints 是一组地址列表，会在系统初始化的时候进行初始化并尝试添加多的arenaHints。在初始化的时候填充了一组通用的地址，然后会根据真正的arena的边界进行扩展增长。
- arenas 是一个arena的map，可以访问到heap中的所有heapArena的值。其中包含了每个arena中的bitmap、spans等数据。具体的arena、bitmap和spans的内容后面再说。

### fixalloc

[fixalloc(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mfixalloc.go#L27 "fixalloc(点击查看结构定义)")

fixalloc是go的一个内存分配器的结构定义。在mheap的结构体中会看到`spanalloc`、`cachealloc`、`arenaHintAlloc`等很多内存分配器的。它主要实现了三个方法`init` `alloc` `free`。

- init
  初始化分配器
- alloc
  进行内存分配。先检查fixalloc的list，如果不为空，则直接分配里面的内存块来用。如果chunk中有足够的空间来用，则使用chunk中的空间，如果没有足够的空间，则向操作系统申请一个新的chunk代替老的chunk。
- free
  释放内存。然而并不会还给操作系统，而是放到list中等待分配。

## 内存初始化

### arena、bitmap、spans

### 虚拟地址申请

### mheap初始化

### mcache初始化

## 内存分配

### 主要分配流程

### 普通对象分配

### 大对象分配

### 小对象分配

## 内存释放

### mcache

### mcentral

### mheap
