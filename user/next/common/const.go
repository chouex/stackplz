package common

const MAX_IOV_COUNT = 6
const MAX_LOOP_COUNT = 32
const MAX_OP_COUNT = 192
const MAX_BUF_READ_SIZE = 4096

const (
	REG_ARM64_X0 uint32 = iota
	REG_ARM64_X1
	REG_ARM64_X2
	REG_ARM64_X3
	REG_ARM64_X4
	REG_ARM64_X5
	REG_ARM64_X6
	REG_ARM64_X7
	REG_ARM64_X8
	REG_ARM64_X9
	REG_ARM64_X10
	REG_ARM64_X11
	REG_ARM64_X12
	REG_ARM64_X13
	REG_ARM64_X14
	REG_ARM64_X15
	REG_ARM64_X16
	REG_ARM64_X17
	REG_ARM64_X18
	REG_ARM64_X19
	REG_ARM64_X20
	REG_ARM64_X21
	REG_ARM64_X22
	REG_ARM64_X23
	REG_ARM64_X24
	REG_ARM64_X25
	REG_ARM64_X26
	REG_ARM64_X27
	REG_ARM64_X28
	REG_ARM64_X29
	REG_ARM64_LR
	REG_ARM64_SP
	REG_ARM64_PC
	REG_ARM64_MAX
	REG_ARM64_INDEX
	REG_ARM64_ABS
)

const (
	CONST_ARGTYPE_START uint32 = iota
	POINTER
	INT
	UINT
	INT8
	INT16
	INT32
	INT64
	UINT8
	UINT16
	UINT32
	UINT64
	NUM
	STRING
	STRUCT
	ARRAY
	BUFFER
	IOVEC
	MSGHDR
	SOCKLEN_T
	SIZE_T
	SSIZE_T
	SOCKADDR
	TIMESPEC
	STAT
	POLLFD
	SIGACTION
	SIGINFO
	STACK_T
	LINUX_DIRENT64
	STRING_ARRAY
	ITTMERSPEC
	RUSAGE
	UTSNAME
	TIMEVAL
	TIMEZONE
	SYSINFO
	STATFS
	EPOLLEVENT
	INT_ARRAY_1
	INT_ARRAY_2
	SIGINFO_V2
	UINT_ARRAY_1
	SIGSET
	INT_PTR
	UINT_PTR
	BUFFER_X2
	IOVEC_X2
	INT_FCNTL_FLAGS
	INT_STATX_FLAGS
	INT_UNLINK_FLAGS
	INT_SOCKET_FLAGS
	INT_FILE_FLAGS
	INT16_PERM_FLAGS
	CONST_ARGTYPE_END
)
