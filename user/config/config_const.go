package config

const (
	EBPF_PROG_NONE uint32 = iota
	EBPF_SYS_ENTER
	EBPF_SYS_EXIT
	EBPF_SYS_ALL
	EBPF_UPROBE_ENTER
)

const MAX_COUNT = 1024
const MAX_FILTER_COUNT = 6

const (
	TRACE_COMMON uint32 = iota
	TRACE_ALL
	// TRACE_FILE
	// TRACE_PROCESS
	// TRACE_NET
	// TRACE_SIGNAL
	// TRACE_STAT
)

// const MAX_LOOP_COUNT = 6
// const MAX_BUF_READ_SIZE uint32 = 4096

// const (
// 	REG_ARM64_X0 uint32 = iota
// 	REG_ARM64_X1
// 	REG_ARM64_X2
// 	REG_ARM64_X3
// 	REG_ARM64_X4
// 	REG_ARM64_X5
// 	REG_ARM64_X6
// 	REG_ARM64_X7
// 	REG_ARM64_X8
// 	REG_ARM64_X9
// 	REG_ARM64_X10
// 	REG_ARM64_X11
// 	REG_ARM64_X12
// 	REG_ARM64_X13
// 	REG_ARM64_X14
// 	REG_ARM64_X15
// 	REG_ARM64_X16
// 	REG_ARM64_X17
// 	REG_ARM64_X18
// 	REG_ARM64_X19
// 	REG_ARM64_X20
// 	REG_ARM64_X21
// 	REG_ARM64_X22
// 	REG_ARM64_X23
// 	REG_ARM64_X24
// 	REG_ARM64_X25
// 	REG_ARM64_X26
// 	REG_ARM64_X27
// 	REG_ARM64_X28
// 	REG_ARM64_X29
// 	REG_ARM64_LR
// 	REG_ARM64_SP
// 	REG_ARM64_PC
// 	REG_ARM64_MAX
// 	REG_ARM64_INDEX
// 	REG_ARM64_ABS
// )

// type FlagOp struct {
// 	Name  string
// 	Value int32
// }

// var PermissionFlags []FlagOp = []FlagOp{
// 	{"S_IFMT", int32(00170000)},
// 	{"S_IFSOCK", int32(0140000)},
// 	{"S_IFLNK", int32(0120000)},
// 	{"S_IFREG", int32(0100000)},
// 	{"S_IFBLK", int32(0060000)},
// 	{"S_IFDIR", int32(0040000)},
// 	{"S_IFCHR", int32(0020000)},
// 	{"S_IFIFO", int32(0010000)},
// 	{"S_ISUID", int32(0004000)},
// 	{"S_ISGID", int32(0002000)},
// 	{"S_ISVTX", int32(0001000)},

// 	{"S_IRWXU", int32(00700)},
// 	{"S_IRUSR", int32(00400)},
// 	{"S_IWUSR", int32(00200)},
// 	{"S_IXUSR", int32(00100)},

// 	{"S_IRWXG", int32(00070)},
// 	{"S_IRGRP", int32(00040)},
// 	{"S_IWGRP", int32(00020)},
// 	{"S_IXGRP", int32(00010)},

// 	{"S_IRWXO", int32(00007)},
// 	{"S_IROTH", int32(00004)},
// 	{"S_IWOTH", int32(00002)},
// 	{"S_IXOTH", int32(00001)},
// }

// var ProtFlags []FlagOp = []FlagOp{
// 	{"PROT_READ", int32(0x1)},
// 	{"PROT_WRITE", int32(0x2)},
// 	{"PROT_EXEC", int32(0x4)},
// 	{"PROT_SEM", int32(0x8)},
// 	// {"PROT_NONE", int32(0x0)},
// 	{"PROT_GROWSDOWN", int32(0x01000000)},
// 	{"PROT_GROWSUP", int32(0x02000000)},
// }

// var FileFlags []FlagOp = []FlagOp{
// 	// https://cs.android.com/android/kernel/superproject/+/common-android-mainline:common/include/uapi/asm-generic/fcntl.h
// 	{"O_ACCMODE", int32(00000003)},
// 	// {"O_RDONLY", int32(00000000)},
// 	{"O_WRONLY", int32(00000001)},
// 	{"O_RDWR", int32(00000002)},
// 	{"O_CREAT", int32(00000100)},
// 	{"O_EXCL", int32(00000200)},
// 	{"O_NOCTTY", int32(00000400)},
// 	{"O_TRUNC", int32(00001000)},
// 	{"O_APPEND", int32(00002000)},
// 	{"O_NONBLOCK", int32(00004000)},
// 	{"O_DSYNC", int32(00010000)},
// 	{"FASYNC", int32(00020000)},
// 	// {"O_DIRECT", int32(00040000)},
// 	// {"O_LARGEFILE", int32(00100000)},
// 	// {"O_DIRECTORY", int32(00200000)},
// 	// {"O_NOFOLLOW", int32(00400000)},
// 	{"O_NOATIME", int32(01000000)},
// 	{"O_CLOEXEC", int32(02000000)},
// 	// 注意不同架构的 flag 定义不一样
// 	// https://cs.android.com/android/kernel/superproject/+/common-android-mainline:common/arch/arm64/include/uapi/asm/fcntl.h
// 	{"O_DIRECTORY", int32(00040000)},
// 	{"O_NOFOLLOW", int32(00100000)},
// 	{"O_DIRECT", int32(00200000)},
// 	{"O_LARGEFILE", int32(00400000)},
// }

// var MapFlags []FlagOp = []FlagOp{
// 	{"MAP_SHARED", int32(0x01)},
// 	{"MAP_PRIVATE", int32(0x02)},
// 	{"MAP_SHARED_VALIDATE", int32(0x03)},
// 	{"MAP_TYPE", int32(0x0f)},
// 	{"MAP_FIXED", int32(0x10)},
// 	{"MAP_ANONYMOUS", int32(0x20)},
// 	{"MAP_POPULATE", int32(0x008000)},
// 	{"MAP_NONBLOCK", int32(0x010000)},
// 	{"MAP_STACK", int32(0x020000)},
// 	{"MAP_HUGETLB", int32(0x040000)},
// 	{"MAP_SYNC", int32(0x080000)},
// 	{"MAP_FIXED_NOREPLACE", int32(0x100000)},
// 	{"MAP_UNINITIALIZED", int32(0x4000000)},
// }

// var MreapFlags []FlagOp = []FlagOp{
// 	{"MREMAP_MAYMOVE", 1},
// 	{"MREMAP_FIXED", 2},
// 	{"MREMAP_DONTUNMAP", 4},
// }

// var SocketTypes []FlagOp = []FlagOp{
// 	{"SOCK_STREAM", int32(1)},
// 	{"SOCK_DGRAM", int32(2)},
// 	{"SOCK_RAW", int32(3)},
// 	{"SOCK_RDM", int32(4)},
// 	{"SOCK_SEQPACKET", int32(5)},
// 	{"SOCK_DCCP", int32(6)},
// 	{"SOCK_PACKET", int32(10)},

// 	{"SOCK_CLOEXEC", int32(02000000)},
// 	{"SOCK_NONBLOCK", int32(00004000)},
// }
