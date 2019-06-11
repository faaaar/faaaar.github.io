---
title: Golang - 内存分配
date: 2019-06-05
updated: 2019-06-11
categories:
    - Go
tags:
    - 内存分配
    - 垃圾回收
---

# 内存管理

查看的源码版本为1.12。

由于知识不足，理解很肤浅。在有了更深的理解之后持续更新。

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

- next、prev
  指针域，指向前一个和后一个mspan，说明mspan是一个双向链表。
- npages
  span中存储的页的数量, mspan 的大小为 page 大小的整数倍。
- nelems
  span中可以存储元素的数量
- freeindex
  是一个大小在0到nelems的数字，用来标记下一次分配空间时开始扫描的位置，可以简单的理解为分配到了第几个区块。如果freeindex==nelem,说明当前span已经没有空闲空间了。
- spanclass
  0~_NumSizeClasses之间的一个值。是当前mspan的类型。不同类型的mspan所分割的区块的大小也不一样。该值与class_to_size中存储的值相关。
- elemsize
  元素大小。根据sizeclass计算得出。在分配大对象的时候，该值与npages相关，值为`pagesize*npages`。

### mcache

[mcache(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcache.go#L19 "mcache(点击查看结构定义)")

在Go的GPM模型中可以知道，P是一个虚拟资源，同一时间只能有一个G来访问P，所以P中的资源不需要加锁。为了在分配对象的时候有更好的性能，则每个P都会有一个mspan的缓存mcache

在mspan的结构定义中我们可以看到，alloc存储着指向mspan的指针的长度为[numSpanClasses](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L614 "numSpanClasses")的数组。而numSpanClasses的长度与[_NumSizeClasses](https://github.com/golang/go/blob/go1.12.5/src/runtime/sizeclasses.go#L79 "_NumSizeClasses")有关。计算可得出numSpanClasses的值为134。这134个值分别对应着67个类型的scan和noscan的值。如下图

![mcache-alloc](/images/go/gc/mcache-alloc.png "mcache-alloc")

其中scan和noscan的区别在于，如果对象包含了指针，分配对象时会使用scan的span，如果对象不包含指针, 分配对象时会使用noscan的span。
而把span分为scan和noscan的意义在于，GC扫描对象的时候对于noscan的span可以不去查看bitmap区域来标记子对象，这样可以大幅提升标记的效率。

### mcentral
[mcentral(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcentral.go#L20 "mcentral(点击查看结构定义)") 将内存按照mspan进行划分管理。访问的时候需要加锁，因为会有多个P对它进行访问。

结构体分析

- lock
  mcentral是带锁的，因为会有多个P来访问它，应保证并发安全。
- spanclass
  与mspan处的spanclass一致，所以我们可以推测mcentral也与mcache的alloc类似，有`numSpanClasses`个
- nonempty
  类型为[mSpanList](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L310 "mSpanList")的一个列表。其中包含了指向列表头和尾的指针。这个列表中存储了空闲的mspan的指针，当mcache中的mspan不够用的时候，可以向mcentral申请更多的空间。
- empty
  类型与nonempty相同，然而是存储了已经被申请走使用中(或者被mcache缓存起来)的mspan的指针。

### mheap
[mheap(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L31 "mheap(点击查看结构定义)")

mheap在系统初始化的时候([mallocinit](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L359 "mallocinit"))进行初始化，是一个全局变量。

结构体分析

- lock
  全局变量，应当保证并发安全
- allspans
  存储了创建的所有的mspans，并且是手动管理的，可以重新分配，并伴随着堆的增长而不断更新存储位置。
- central
  如前面所说，central也存在`numSpanClasses`个
- arenaHints
  是一组地址列表，会在系统初始化的时候进行初始化并尝试添加多的arenaHints。在初始化的时候填充了一组通用的地址，然后会根据真正的arena的边界进行扩展增长。
- arenas
  是一个arena的map，它的定义是`[1<<arenaL1Bits]*[1<<arenaL2Bits]*heapArena`，指向了heap虚拟地址中包含的所有的arena的元数据。`arenaL1Bits`的值与平台有关。关于`heapArena`的结构在下面再聊。

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

系统在启动的时候，会调用[mallocinit](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L359 "mallocinit")函数。但是在进一步了解这个函数之前需要了解一些其他的内容。

### Go的虚拟内存与arena

先来一个简单的Go程序

```go
func main() {
    for {}
}
```

linux用户可以通过下面的命令来进行查看应用程序的实际内存和虚拟内存的分配、使用情况

```shell
# linux
ps ux | grep xxx
cat /prox/PID/smaps | grep -B 1 -s "Size"

# macos
/usr/bin/vmmap PID | grep VM_ALLOCATE
```

![vm_allocate](/images/go/gc/vm_allocate.png "vm_allocate")

从统计信息中可以看到，程序的虚拟内存是由几个内存块组成的，大小分别为接近于512K、2M、32M、64M。这是因为Go的虚拟内存是由一组arena组成。在程序初始化的时候，会映射一个arena，所以是64MB(根据[heapArenaBytes](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L226 "heapArenaBytes")处的注释获得)。而存储arena用的arenas的L2长度为4M，即存储4M个指针，每个指针8字节，所以在我当前的电脑上需要32MB空间。

** 所以arena到底是什么? **

Go的虚拟内存是由一组arena组成的，这一组arena组成我们所说的堆(mheap.arenas)。初始情况下mheap会映射一个arena，也就是64MB(不同系统的arena大小不同)。所以我们的程序的当前内存会按照需要小增量进行映射，并且每次以一个arena为单位。在旧的Go版本中，Go程序是采用预先保留连续的虚拟地址的方案，在64位的系统上，会预先保留512G的虚拟内存空间，但是不可增长，而当前版本中，虚拟内存的地址长度被设置为了48位，理论上可以支持2^48字节的内存使用。

既然如此，那么是时候一探`arenas`中`heapArena`的内容了

[heapArena(点击查看结构定义)](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L209 "点我查看heapArena(点击查看结构定义)")

- bitmap [heapArenaBitmapBytes]byte
  用于表示arena区域中哪些地址保存了对象, 并且对象中哪些地址包含了指针。bitmap的长度为`heapArenaBytes / (sys.PtrSize * 8 / 2)`根据`heapArenaBytes`的定义可以得出在我的计算机上一个heapArena中bitmap的长度为2M。
  每个bitmap长度为1byte，对应了arena区域中4个指针大小的内存，即2bit对应一个指针大小内存，这两个bit分别表示是否应该继续扫描和是否包含指针(用于GC)。其结构如下

  ![bitmap-arena-ptr](/images/go/gc/bitmap-arena-ptr.png "bitmap-arena-ptr")


- spans [pagesPerArena]*mspan

  `pagesPerArena = heapArenaBytes / pageSize`
  用于表示arena区域中的某一个page属于哪个span(mspan)，将page与span进行对应。其结构如下s

  ![spans-arena-page](/images/go/gc/spans-arena-page.png "spans-arena-page")


- pageInUse [pagesPerArena / 8]uint8
  ```go
  // pageInUse is a bitmap that indicates which spans are in
  // state mSpanInUse. This bitmap is indexed by page number,
  // but only the bit corresponding to the first page in each
  // span is used.
  //
  // Writes are protected by mheap_.lock.
  ```
- pageMarks [pagesPerArena / 8]uint8
  ```go
  // pageMarks is a bitmap that indicates which spans have any
  // marked objects on them. Like pageInUse, only the bit
  // corresponding to the first page in each span is used.
  //
  // Writes are done atomically during marking. Reads are
  // non-atomic and lock-free since they only occur during
  // sweeping (and hence never race with writes).
  //
  // This is used to quickly find whole spans that can be freed.
  //
  // TODO(austin): It would be nice if this was uint64 for
  // faster scanning, but we don't have 64-bit atomic bit
  // operations.
  ```

### mheap初始化

[点我查看mheap.init()](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L777 "点我查看mheap.init()")

mheap结构在系统调用[mallocinit](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L359 "mallocinit")函数时执行。主要做了以下事情

- 初始化各种结构的分配器
- 初始化heap.central中各个类型的mcentral

### mcache初始化

也是对当前G的M的初始化，需要先获取到当前的G

```go
_g_ := getg()
_g_.m.mcache = allocmcache()
```

[allocmacache](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcache.go#L19 "allocmcache")函数主要通过`mheap.cachealloc`从堆上划分一块内存，这期间mheap是加锁状态。之后初始化不同类型的span。


### 虚拟地址申请

[虚拟地址申请代码](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L397 "虚拟地址申请代码")

程序在启动的时候，会根据当前系统信息构建一块虚拟内存。

在64为操作系统上，程序根据不同的`GOARCH`和`GOOS`初始化`mehap.arenaHints`，一般填充的是一组通用的地址，然后会根据真正的arena的边界进行扩展增长。

## 内存分配

堆上内存分配调用了runtime包的[newobject](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L1067 "newobject")函数。

### 主要分配流程

根据要分配的对象的大小使用不同的分配策略

1. objectsize>32768，则使用heap直接进行分配
2. objectsize<16&&noscan，则使用tiny分配器进行分配
3. 其他情况根据对象具体大小使用不同的sizeclass来进行分配

### 大对象分配

对于大于32K的对象来说，分配直接跳过了mcache和mcentral两个部分，直接通过调用[largeAlloc](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L1039 "largeAlloc")方法，先对齐然后计算所需要的整数页，然后使用mheap.alloc进行分配。

### 小对象分配

使用[tiny allocator](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L863 "tiny allocator")进行小对象分配。tiny实际上只是一个指针，用来记录当前mspan(sizeclass=2&&noscan的元素)中内存的起始位置，通过tinyoffset计算当前对象应该分配在哪里。

在分配过程中，先将要分配的对象的大小与tinyoffset进行地址对齐。如果当前块中还有剩余的空间而且足够，直接分配，并记录新的tinyoffset的值。如果不够则重新申请一个新的内存块，在简单初始化这块内存之后将对象分配进去，并记录tinyoffset和新的内存块的起始位置。

### 普通对象分配
[对象在16b和32K之间时的分配逻辑。](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L928 "对象在16b和32K之间时的分配逻辑。")

1. 通过对象的大小以及是否含有指针(noscan)计算出应当分配的sizeclass，然后通过调用[nextFreeFast](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L749 "nextFreeFast")尝试进行分配(从mcache.allocc中分配)。

2. 如果分配失败了，则调用[nextFree](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L776 "nextFree")方法尝试从mcentral或者mheap中申请更多内存使用。
   - 在nextFree方法中调用[refill](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcache.go#L119 "refill")向mcentral中申请内存。通过调用[cacheSpan](https://github.com/golang/go/blob/go1.12.5/src/runtime/malloc.go#L135 "cacheSpan")申请新的span。
   - 如果在这个过程中，mcentral空了，则会调用[mcentral.grow](https://github.com/golang/go/blob/go1.12.5/src/runtime/mcentral.go#L251 "mcentral.grow")方法向堆申请空间。
   - 如果堆空了则会调用[alloc_m](https://github.com/golang/go/blob/go1.12.5/src/runtime/mheap.go#L961 "alloc_m")方法向操作系统申请内存空间。

## 总结

之所以写这个内容，最开始是因为需要进行组内分享，想把自己看的内容给整理一下，然后好整理成PPT。然而在不知不觉中自己已经接触、了解到了远超分享本身需要准备的内容，这对我个人的成长是巨大的。

回到标题本身来讲，虽然硬着头皮、强行写完了，然而很多内容是不理解的，可能需要其他知识点的辅助吧。Go的内存分配是基于`tcmalloc`的，之后会先去了解这块内容，进一步去理解整个内存分配的流程。如果能解开自己的一些疑惑当然是最好的，到时候也会再来一篇文章，或者更新本文来记录学习到的知识点。

End.
