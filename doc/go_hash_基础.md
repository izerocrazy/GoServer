**Hash**
	
go 语言中有一个Hash Packget，它提供 Hash Hash32 Hash64 三种Interface

三者 Interface的接口如下：

	type Hash interface {
		// oh…它有了io.Write的所有的接口
		// 那么可以这样用 hash.Write(btyeSliece)
		// io.Copy(hash, file)
		io.Write
	
		// 将b添加入其中，得到 前面的输入+b 新产生的哈希的结果
		Sum(b []byte) []byte
	
		// 重置hash，也就是转为 零输入
		Reset()
	
		// Sum 返回slice的长度
		Size() int
	
		// 返回hash底层表的块大小，Write能够接受任意大小的输入，但如果是这个块的整数倍，效率会得到提升
		BlockSize() int
	}

	type Hash32 interface {
		Hash
		Sum32() uint32
	}

	type Hash64 interface {
		Hash
		Sum64() uint64
	}

CRC32，CRC是一个Hash算法，32/64位循环冗余校验（常用于检查数据的完整性）；类似的还有MD5，信息摘要算法第五版，SHA1，安全哈希算法。简单看这个就ok

	// 多项式参数，hash算法最重要的一个指标在于碰撞（另一个是时间），IEEE据说是这里面最差的一个，但它亦是使用最多的
	const (
		IEEE = Oxxxxxx
		Castagnoli = 
		Koopman = 
	)

	var IEEETable = makeTable(IEEE)		// 用上述参数，make一个Table，用在算法之中，注意到makeTable是一个不导出函数

	type Table [256]uint32

	// 这里却是一个导出的函数，功效应该是一样的
	func MakeTable(poly uint32) *Table

	// 以下是两个快捷函数，我们希望的是，给一坨数据，然后得到结果——它们正能做到
	func Checksum(data[]byte, tab *Table) uint32
	func ChecksumIEEE(data []byte)	// 直接使用IEEETable

	// 生成一个hash实例
	func New(tab *Table) hash.Hash32
	func NewIEEE() hash.Hash32

	// 类似于上面Hash中的 Sum()
	func Update(crc uint32, tab *Table, p []byte) uint32

具体示例，介绍两种基本用法，省去自己建Table，因为我对CRC32的原理也不太懂

	str := "hello ketty"
	key := []byte(str)

	// 用法一
	v := crc32.ChecksumIEEE([]byte(key))

	// 用法二
	h := crc32.NewIEEE()
	h.Write(key)
	v2 := h.Sum32()
