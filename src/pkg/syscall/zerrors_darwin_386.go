// mkerrors.sh -f -m32
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// godefs -c gcc -f -m32 -gsyscall -f -m32 _const.c

// MACHINE GENERATED - DO NOT EDIT.

package syscall

// Constants
const (
	AF_APPLETALK                = 0x10
	AF_CCITT                    = 0xa
	AF_CHAOS                    = 0x5
	AF_CNT                      = 0x15
	AF_COIP                     = 0x14
	AF_DATAKIT                  = 0x9
	AF_DECnet                   = 0xc
	AF_DLI                      = 0xd
	AF_E164                     = 0x1c
	AF_ECMA                     = 0x8
	AF_HYLINK                   = 0xf
	AF_IEEE80211                = 0x25
	AF_IMPLINK                  = 0x3
	AF_INET                     = 0x2
	AF_INET6                    = 0x1e
	AF_IPX                      = 0x17
	AF_ISDN                     = 0x1c
	AF_ISO                      = 0x7
	AF_LAT                      = 0xe
	AF_LINK                     = 0x12
	AF_LOCAL                    = 0x1
	AF_MAX                      = 0x26
	AF_NATM                     = 0x1f
	AF_NDRV                     = 0x1b
	AF_NETBIOS                  = 0x21
	AF_NS                       = 0x6
	AF_OSI                      = 0x7
	AF_PPP                      = 0x22
	AF_PUP                      = 0x4
	AF_RESERVED_36              = 0x24
	AF_ROUTE                    = 0x11
	AF_SIP                      = 0x18
	AF_SNA                      = 0xb
	AF_SYSTEM                   = 0x20
	AF_UNIX                     = 0x1
	AF_UNSPEC                   = 0
	BIOCFLUSH                   = 0x20004268
	BIOCGBLEN                   = 0x40044266
	BIOCGDLT                    = 0x4004426a
	BIOCGDLTLIST                = 0xc00c4279
	BIOCGETIF                   = 0x4020426b
	BIOCGHDRCMPLT               = 0x40044274
	BIOCGRSIG                   = 0x40044272
	BIOCGRTIMEOUT               = 0x4008426e
	BIOCGSEESENT                = 0x40044276
	BIOCGSTATS                  = 0x4008426f
	BIOCIMMEDIATE               = 0x80044270
	BIOCPROMISC                 = 0x20004269
	BIOCSBLEN                   = 0xc0044266
	BIOCSDLT                    = 0x80044278
	BIOCSETF                    = 0x80084267
	BIOCSETIF                   = 0x8020426c
	BIOCSHDRCMPLT               = 0x80044275
	BIOCSRSIG                   = 0x80044273
	BIOCSRTIMEOUT               = 0x8008426d
	BIOCSSEESENT                = 0x80044277
	BIOCVERSION                 = 0x40044271
	BPF_A                       = 0x10
	BPF_ABS                     = 0x20
	BPF_ADD                     = 0
	BPF_ALIGNMENT               = 0x4
	BPF_ALU                     = 0x4
	BPF_AND                     = 0x50
	BPF_B                       = 0x10
	BPF_DIV                     = 0x30
	BPF_H                       = 0x8
	BPF_IMM                     = 0
	BPF_IND                     = 0x40
	BPF_JA                      = 0
	BPF_JEQ                     = 0x10
	BPF_JGE                     = 0x30
	BPF_JGT                     = 0x20
	BPF_JMP                     = 0x5
	BPF_JSET                    = 0x40
	BPF_K                       = 0
	BPF_LD                      = 0
	BPF_LDX                     = 0x1
	BPF_LEN                     = 0x80
	BPF_LSH                     = 0x60
	BPF_MAJOR_VERSION           = 0x1
	BPF_MAXBUFSIZE              = 0x80000
	BPF_MAXINSNS                = 0x200
	BPF_MEM                     = 0x60
	BPF_MEMWORDS                = 0x10
	BPF_MINBUFSIZE              = 0x20
	BPF_MINOR_VERSION           = 0x1
	BPF_MISC                    = 0x7
	BPF_MSH                     = 0xa0
	BPF_MUL                     = 0x20
	BPF_NEG                     = 0x80
	BPF_OR                      = 0x40
	BPF_RELEASE                 = 0x30bb6
	BPF_RET                     = 0x6
	BPF_RSH                     = 0x70
	BPF_ST                      = 0x2
	BPF_STX                     = 0x3
	BPF_SUB                     = 0x10
	BPF_TAX                     = 0
	BPF_TXA                     = 0x80
	BPF_W                       = 0
	BPF_X                       = 0x8
	CTL_MAXNAME                 = 0xc
	CTL_NET                     = 0x4
	DLT_APPLE_IP_OVER_IEEE1394  = 0x8a
	DLT_ARCNET                  = 0x7
	DLT_ATM_CLIP                = 0x13
	DLT_ATM_RFC1483             = 0xb
	DLT_AX25                    = 0x3
	DLT_CHAOS                   = 0x5
	DLT_CHDLC                   = 0x68
	DLT_C_HDLC                  = 0x68
	DLT_EN10MB                  = 0x1
	DLT_EN3MB                   = 0x2
	DLT_FDDI                    = 0xa
	DLT_IEEE802                 = 0x6
	DLT_IEEE802_11              = 0x69
	DLT_IEEE802_11_RADIO        = 0x7f
	DLT_IEEE802_11_RADIO_AVS    = 0xa3
	DLT_LINUX_SLL               = 0x71
	DLT_LOOP                    = 0x6c
	DLT_NULL                    = 0
	DLT_PFLOG                   = 0x75
	DLT_PFSYNC                  = 0x12
	DLT_PPP                     = 0x9
	DLT_PPP_BSDOS               = 0x10
	DLT_PPP_SERIAL              = 0x32
	DLT_PRONET                  = 0x4
	DLT_RAW                     = 0xc
	DLT_SLIP                    = 0x8
	DLT_SLIP_BSDOS              = 0xf
	DT_BLK                      = 0x6
	DT_CHR                      = 0x2
	DT_DIR                      = 0x4
	DT_FIFO                     = 0x1
	DT_LNK                      = 0xa
	DT_REG                      = 0x8
	DT_SOCK                     = 0xc
	DT_UNKNOWN                  = 0
	DT_WHT                      = 0xe
	E2BIG                       = 0x7
	EACCES                      = 0xd
	EADDRINUSE                  = 0x30
	EADDRNOTAVAIL               = 0x31
	EAFNOSUPPORT                = 0x2f
	EAGAIN                      = 0x23
	EALREADY                    = 0x25
	EAUTH                       = 0x50
	EBADARCH                    = 0x56
	EBADEXEC                    = 0x55
	EBADF                       = 0x9
	EBADMACHO                   = 0x58
	EBADMSG                     = 0x5e
	EBADRPC                     = 0x48
	EBUSY                       = 0x10
	ECANCELED                   = 0x59
	ECHILD                      = 0xa
	ECONNABORTED                = 0x35
	ECONNREFUSED                = 0x3d
	ECONNRESET                  = 0x36
	EDEADLK                     = 0xb
	EDESTADDRREQ                = 0x27
	EDEVERR                     = 0x53
	EDOM                        = 0x21
	EDQUOT                      = 0x45
	EEXIST                      = 0x11
	EFAULT                      = 0xe
	EFBIG                       = 0x1b
	EFTYPE                      = 0x4f
	EHOSTDOWN                   = 0x40
	EHOSTUNREACH                = 0x41
	EIDRM                       = 0x5a
	EILSEQ                      = 0x5c
	EINPROGRESS                 = 0x24
	EINTR                       = 0x4
	EINVAL                      = 0x16
	EIO                         = 0x5
	EISCONN                     = 0x38
	EISDIR                      = 0x15
	ELAST                       = 0x67
	ELOOP                       = 0x3e
	EMFILE                      = 0x18
	EMLINK                      = 0x1f
	EMSGSIZE                    = 0x28
	EMULTIHOP                   = 0x5f
	ENAMETOOLONG                = 0x3f
	ENEEDAUTH                   = 0x51
	ENETDOWN                    = 0x32
	ENETRESET                   = 0x34
	ENETUNREACH                 = 0x33
	ENFILE                      = 0x17
	ENOATTR                     = 0x5d
	ENOBUFS                     = 0x37
	ENODATA                     = 0x60
	ENODEV                      = 0x13
	ENOENT                      = 0x2
	ENOEXEC                     = 0x8
	ENOLCK                      = 0x4d
	ENOLINK                     = 0x61
	ENOMEM                      = 0xc
	ENOMSG                      = 0x5b
	ENOPOLICY                   = 0x67
	ENOPROTOOPT                 = 0x2a
	ENOSPC                      = 0x1c
	ENOSR                       = 0x62
	ENOSTR                      = 0x63
	ENOSYS                      = 0x4e
	ENOTBLK                     = 0xf
	ENOTCONN                    = 0x39
	ENOTDIR                     = 0x14
	ENOTEMPTY                   = 0x42
	ENOTSOCK                    = 0x26
	ENOTSUP                     = 0x2d
	ENOTTY                      = 0x19
	ENXIO                       = 0x6
	EOPNOTSUPP                  = 0x66
	EOVERFLOW                   = 0x54
	EPERM                       = 0x1
	EPFNOSUPPORT                = 0x2e
	EPIPE                       = 0x20
	EPROCLIM                    = 0x43
	EPROCUNAVAIL                = 0x4c
	EPROGMISMATCH               = 0x4b
	EPROGUNAVAIL                = 0x4a
	EPROTO                      = 0x64
	EPROTONOSUPPORT             = 0x2b
	EPROTOTYPE                  = 0x29
	EPWROFF                     = 0x52
	ERANGE                      = 0x22
	EREMOTE                     = 0x47
	EROFS                       = 0x1e
	ERPCMISMATCH                = 0x49
	ESHLIBVERS                  = 0x57
	ESHUTDOWN                   = 0x3a
	ESOCKTNOSUPPORT             = 0x2c
	ESPIPE                      = 0x1d
	ESRCH                       = 0x3
	ESTALE                      = 0x46
	ETIME                       = 0x65
	ETIMEDOUT                   = 0x3c
	ETOOMANYREFS                = 0x3b
	ETXTBSY                     = 0x1a
	EUSERS                      = 0x44
	EVFILT_AIO                  = -0x3
	EVFILT_FS                   = -0x9
	EVFILT_MACHPORT             = -0x8
	EVFILT_PROC                 = -0x5
	EVFILT_READ                 = -0x1
	EVFILT_SESSION              = -0xb
	EVFILT_SIGNAL               = -0x6
	EVFILT_SYSCOUNT             = 0xb
	EVFILT_THREADMARKER         = 0xb
	EVFILT_TIMER                = -0x7
	EVFILT_USER                 = -0xa
	EVFILT_VNODE                = -0x4
	EVFILT_WRITE                = -0x2
	EV_ADD                      = 0x1
	EV_CLEAR                    = 0x20
	EV_DELETE                   = 0x2
	EV_DISABLE                  = 0x8
	EV_DISPATCH                 = 0x80
	EV_ENABLE                   = 0x4
	EV_EOF                      = 0x8000
	EV_ERROR                    = 0x4000
	EV_FLAG0                    = 0x1000
	EV_FLAG1                    = 0x2000
	EV_ONESHOT                  = 0x10
	EV_OOBAND                   = 0x2000
	EV_POLL                     = 0x1000
	EV_RECEIPT                  = 0x40
	EV_SYSFLAGS                 = 0xf000
	EV_TRIGGER                  = 0x100
	EWOULDBLOCK                 = 0x23
	EXDEV                       = 0x12
	FD_CLOEXEC                  = 0x1
	FD_SETSIZE                  = 0x400
	F_ADDFILESIGS               = 0x3d
	F_ADDSIGS                   = 0x3b
	F_ALLOCATEALL               = 0x4
	F_ALLOCATECONTIG            = 0x2
	F_CHKCLEAN                  = 0x29
	F_DUPFD                     = 0
	F_FREEZE_FS                 = 0x35
	F_FULLFSYNC                 = 0x33
	F_GETFD                     = 0x1
	F_GETFL                     = 0x3
	F_GETLK                     = 0x7
	F_GETOWN                    = 0x5
	F_GETPATH                   = 0x32
	F_GETPROTECTIONCLASS        = 0x3e
	F_GLOBAL_NOCACHE            = 0x37
	F_LOG2PHYS                  = 0x31
	F_MARKDEPENDENCY            = 0x3c
	F_NOCACHE                   = 0x30
	F_PATHPKG_CHECK             = 0x34
	F_PEOFPOSMODE               = 0x3
	F_PREALLOCATE               = 0x2a
	F_RDADVISE                  = 0x2c
	F_RDAHEAD                   = 0x2d
	F_RDLCK                     = 0x1
	F_READBOOTSTRAP             = 0x2e
	F_SETFD                     = 0x2
	F_SETFL                     = 0x4
	F_SETLK                     = 0x8
	F_SETLKW                    = 0x9
	F_SETOWN                    = 0x6
	F_SETPROTECTIONCLASS        = 0x3f
	F_SETSIZE                   = 0x2b
	F_THAW_FS                   = 0x36
	F_UNLCK                     = 0x2
	F_VOLPOSMODE                = 0x4
	F_WRITEBOOTSTRAP            = 0x2f
	F_WRLCK                     = 0x3
	IFF_ALLMULTI                = 0x200
	IFF_ALTPHYS                 = 0x4000
	IFF_BROADCAST               = 0x2
	IFF_DEBUG                   = 0x4
	IFF_LINK0                   = 0x1000
	IFF_LINK1                   = 0x2000
	IFF_LINK2                   = 0x4000
	IFF_LOOPBACK                = 0x8
	IFF_MULTICAST               = 0x8000
	IFF_NOARP                   = 0x80
	IFF_NOTRAILERS              = 0x20
	IFF_OACTIVE                 = 0x400
	IFF_POINTOPOINT             = 0x10
	IFF_PROMISC                 = 0x100
	IFF_RUNNING                 = 0x40
	IFF_SIMPLEX                 = 0x800
	IFF_UP                      = 0x1
	IFNAMSIZ                    = 0x10
	IFT_1822                    = 0x2
	IFT_AAL5                    = 0x31
	IFT_ARCNET                  = 0x23
	IFT_ARCNETPLUS              = 0x24
	IFT_ATM                     = 0x25
	IFT_BRIDGE                  = 0xd1
	IFT_CARP                    = 0xf8
	IFT_CEPT                    = 0x13
	IFT_DS3                     = 0x1e
	IFT_ENC                     = 0xf4
	IFT_EON                     = 0x19
	IFT_ETHER                   = 0x6
	IFT_FAITH                   = 0x38
	IFT_FDDI                    = 0xf
	IFT_FRELAY                  = 0x20
	IFT_FRELAYDCE               = 0x2c
	IFT_GIF                     = 0x37
	IFT_HDH1822                 = 0x3
	IFT_HIPPI                   = 0x2f
	IFT_HSSI                    = 0x2e
	IFT_HY                      = 0xe
	IFT_IEEE1394                = 0x90
	IFT_IEEE8023ADLAG           = 0x88
	IFT_ISDNBASIC               = 0x14
	IFT_ISDNPRIMARY             = 0x15
	IFT_ISO88022LLC             = 0x29
	IFT_ISO88023                = 0x7
	IFT_ISO88024                = 0x8
	IFT_ISO88025                = 0x9
	IFT_ISO88026                = 0xa
	IFT_L2VLAN                  = 0x87
	IFT_LAPB                    = 0x10
	IFT_LOCALTALK               = 0x2a
	IFT_LOOP                    = 0x18
	IFT_MIOX25                  = 0x26
	IFT_MODEM                   = 0x30
	IFT_NSIP                    = 0x1b
	IFT_OTHER                   = 0x1
	IFT_P10                     = 0xc
	IFT_P80                     = 0xd
	IFT_PARA                    = 0x22
	IFT_PDP                     = 0xff
	IFT_PFLOG                   = 0xf5
	IFT_PFSYNC                  = 0xf6
	IFT_PPP                     = 0x17
	IFT_PROPMUX                 = 0x36
	IFT_PROPVIRTUAL             = 0x35
	IFT_PTPSERIAL               = 0x16
	IFT_RS232                   = 0x21
	IFT_SDLC                    = 0x11
	IFT_SIP                     = 0x1f
	IFT_SLIP                    = 0x1c
	IFT_SMDSDXI                 = 0x2b
	IFT_SMDSICIP                = 0x34
	IFT_SONET                   = 0x27
	IFT_SONETPATH               = 0x32
	IFT_SONETVT                 = 0x33
	IFT_STARLAN                 = 0xb
	IFT_STF                     = 0x39
	IFT_T1                      = 0x12
	IFT_ULTRA                   = 0x1d
	IFT_V35                     = 0x2d
	IFT_X25                     = 0x5
	IFT_X25DDN                  = 0x4
	IFT_X25PLE                  = 0x28
	IFT_XETHER                  = 0x1a
	IN_CLASSA_HOST              = 0xffffff
	IN_CLASSA_MAX               = 0x80
	IN_CLASSA_NET               = 0xff000000
	IN_CLASSA_NSHIFT            = 0x18
	IN_CLASSB_HOST              = 0xffff
	IN_CLASSB_MAX               = 0x10000
	IN_CLASSB_NET               = 0xffff0000
	IN_CLASSB_NSHIFT            = 0x10
	IN_CLASSC_HOST              = 0xff
	IN_CLASSC_NET               = 0xffffff00
	IN_CLASSC_NSHIFT            = 0x8
	IN_CLASSD_HOST              = 0xfffffff
	IN_CLASSD_NET               = 0xf0000000
	IN_CLASSD_NSHIFT            = 0x1c
	IN_LINKLOCALNETNUM          = 0xa9fe0000
	IN_LOOPBACKNET              = 0x7f
	IPPROTO_3PC                 = 0x22
	IPPROTO_ADFS                = 0x44
	IPPROTO_AH                  = 0x33
	IPPROTO_AHIP                = 0x3d
	IPPROTO_APES                = 0x63
	IPPROTO_ARGUS               = 0xd
	IPPROTO_AX25                = 0x5d
	IPPROTO_BHA                 = 0x31
	IPPROTO_BLT                 = 0x1e
	IPPROTO_BRSATMON            = 0x4c
	IPPROTO_CFTP                = 0x3e
	IPPROTO_CHAOS               = 0x10
	IPPROTO_CMTP                = 0x26
	IPPROTO_CPHB                = 0x49
	IPPROTO_CPNX                = 0x48
	IPPROTO_DDP                 = 0x25
	IPPROTO_DGP                 = 0x56
	IPPROTO_DIVERT              = 0xfe
	IPPROTO_DONE                = 0x101
	IPPROTO_DSTOPTS             = 0x3c
	IPPROTO_EGP                 = 0x8
	IPPROTO_EMCON               = 0xe
	IPPROTO_ENCAP               = 0x62
	IPPROTO_EON                 = 0x50
	IPPROTO_ESP                 = 0x32
	IPPROTO_ETHERIP             = 0x61
	IPPROTO_FRAGMENT            = 0x2c
	IPPROTO_GGP                 = 0x3
	IPPROTO_GMTP                = 0x64
	IPPROTO_GRE                 = 0x2f
	IPPROTO_HELLO               = 0x3f
	IPPROTO_HMP                 = 0x14
	IPPROTO_HOPOPTS             = 0
	IPPROTO_ICMP                = 0x1
	IPPROTO_ICMPV6              = 0x3a
	IPPROTO_IDP                 = 0x16
	IPPROTO_IDPR                = 0x23
	IPPROTO_IDRP                = 0x2d
	IPPROTO_IGMP                = 0x2
	IPPROTO_IGP                 = 0x55
	IPPROTO_IGRP                = 0x58
	IPPROTO_IL                  = 0x28
	IPPROTO_INLSP               = 0x34
	IPPROTO_INP                 = 0x20
	IPPROTO_IP                  = 0
	IPPROTO_IPCOMP              = 0x6c
	IPPROTO_IPCV                = 0x47
	IPPROTO_IPEIP               = 0x5e
	IPPROTO_IPIP                = 0x4
	IPPROTO_IPPC                = 0x43
	IPPROTO_IPV4                = 0x4
	IPPROTO_IPV6                = 0x29
	IPPROTO_IRTP                = 0x1c
	IPPROTO_KRYPTOLAN           = 0x41
	IPPROTO_LARP                = 0x5b
	IPPROTO_LEAF1               = 0x19
	IPPROTO_LEAF2               = 0x1a
	IPPROTO_MAX                 = 0x100
	IPPROTO_MAXID               = 0x34
	IPPROTO_MEAS                = 0x13
	IPPROTO_MHRP                = 0x30
	IPPROTO_MICP                = 0x5f
	IPPROTO_MTP                 = 0x5c
	IPPROTO_MUX                 = 0x12
	IPPROTO_ND                  = 0x4d
	IPPROTO_NHRP                = 0x36
	IPPROTO_NONE                = 0x3b
	IPPROTO_NSP                 = 0x1f
	IPPROTO_NVPII               = 0xb
	IPPROTO_OSPFIGP             = 0x59
	IPPROTO_PGM                 = 0x71
	IPPROTO_PIGP                = 0x9
	IPPROTO_PIM                 = 0x67
	IPPROTO_PRM                 = 0x15
	IPPROTO_PUP                 = 0xc
	IPPROTO_PVP                 = 0x4b
	IPPROTO_RAW                 = 0xff
	IPPROTO_RCCMON              = 0xa
	IPPROTO_RDP                 = 0x1b
	IPPROTO_ROUTING             = 0x2b
	IPPROTO_RSVP                = 0x2e
	IPPROTO_RVD                 = 0x42
	IPPROTO_SATEXPAK            = 0x40
	IPPROTO_SATMON              = 0x45
	IPPROTO_SCCSP               = 0x60
	IPPROTO_SDRP                = 0x2a
	IPPROTO_SEP                 = 0x21
	IPPROTO_SRPC                = 0x5a
	IPPROTO_ST                  = 0x7
	IPPROTO_SVMTP               = 0x52
	IPPROTO_SWIPE               = 0x35
	IPPROTO_TCF                 = 0x57
	IPPROTO_TCP                 = 0x6
	IPPROTO_TP                  = 0x1d
	IPPROTO_TPXX                = 0x27
	IPPROTO_TRUNK1              = 0x17
	IPPROTO_TRUNK2              = 0x18
	IPPROTO_TTP                 = 0x54
	IPPROTO_UDP                 = 0x11
	IPPROTO_VINES               = 0x53
	IPPROTO_VISA                = 0x46
	IPPROTO_VMTP                = 0x51
	IPPROTO_WBEXPAK             = 0x4f
	IPPROTO_WBMON               = 0x4e
	IPPROTO_WSN                 = 0x4a
	IPPROTO_XNET                = 0xf
	IPPROTO_XTP                 = 0x24
	IPV6_BINDV6ONLY             = 0x1b
	IPV6_CHECKSUM               = 0x1a
	IPV6_DEFAULT_MULTICAST_HOPS = 0x1
	IPV6_DEFAULT_MULTICAST_LOOP = 0x1
	IPV6_DEFHLIM                = 0x40
	IPV6_DSTOPTS                = 0x17
	IPV6_FAITH                  = 0x1d
	IPV6_FLOWINFO_MASK          = 0xffffff0f
	IPV6_FLOWLABEL_MASK         = 0xffff0f00
	IPV6_FRAGTTL                = 0x78
	IPV6_FW_ADD                 = 0x1e
	IPV6_FW_DEL                 = 0x1f
	IPV6_FW_FLUSH               = 0x20
	IPV6_FW_GET                 = 0x22
	IPV6_FW_ZERO                = 0x21
	IPV6_HLIMDEC                = 0x1
	IPV6_HOPLIMIT               = 0x14
	IPV6_HOPOPTS                = 0x16
	IPV6_IPSEC_POLICY           = 0x1c
	IPV6_JOIN_GROUP             = 0xc
	IPV6_LEAVE_GROUP            = 0xd
	IPV6_MAXHLIM                = 0xff
	IPV6_MAXPACKET              = 0xffff
	IPV6_MMTU                   = 0x500
	IPV6_MULTICAST_HOPS         = 0xa
	IPV6_MULTICAST_IF           = 0x9
	IPV6_MULTICAST_LOOP         = 0xb
	IPV6_NEXTHOP                = 0x15
	IPV6_PKTINFO                = 0x13
	IPV6_PKTOPTIONS             = 0x19
	IPV6_PORTRANGE              = 0xe
	IPV6_PORTRANGE_DEFAULT      = 0
	IPV6_PORTRANGE_HIGH         = 0x1
	IPV6_PORTRANGE_LOW          = 0x2
	IPV6_RECVTCLASS             = 0x23
	IPV6_RTHDR                  = 0x18
	IPV6_RTHDR_LOOSE            = 0
	IPV6_RTHDR_STRICT           = 0x1
	IPV6_RTHDR_TYPE_0           = 0
	IPV6_SOCKOPT_RESERVED1      = 0x3
	IPV6_TCLASS                 = 0x24
	IPV6_UNICAST_HOPS           = 0x4
	IPV6_V6ONLY                 = 0x1b
	IPV6_VERSION                = 0x60
	IPV6_VERSION_MASK           = 0xf0
	IP_ADD_MEMBERSHIP           = 0xc
	IP_BOUND_IF                 = 0x19
	IP_DEFAULT_MULTICAST_LOOP   = 0x1
	IP_DEFAULT_MULTICAST_TTL    = 0x1
	IP_DF                       = 0x4000
	IP_DROP_MEMBERSHIP          = 0xd
	IP_DUMMYNET_CONFIGURE       = 0x3c
	IP_DUMMYNET_DEL             = 0x3d
	IP_DUMMYNET_FLUSH           = 0x3e
	IP_DUMMYNET_GET             = 0x40
	IP_FAITH                    = 0x16
	IP_FW_ADD                   = 0x28
	IP_FW_DEL                   = 0x29
	IP_FW_FLUSH                 = 0x2a
	IP_FW_GET                   = 0x2c
	IP_FW_RESETLOG              = 0x2d
	IP_FW_ZERO                  = 0x2b
	IP_HDRINCL                  = 0x2
	IP_IPSEC_POLICY             = 0x15
	IP_MAXPACKET                = 0xffff
	IP_MAX_MEMBERSHIPS          = 0x14
	IP_MF                       = 0x2000
	IP_MSS                      = 0x240
	IP_MULTICAST_IF             = 0x9
	IP_MULTICAST_LOOP           = 0xb
	IP_MULTICAST_TTL            = 0xa
	IP_MULTICAST_VIF            = 0xe
	IP_NAT__XXX                 = 0x37
	IP_OFFMASK                  = 0x1fff
	IP_OLD_FW_ADD               = 0x32
	IP_OLD_FW_DEL               = 0x33
	IP_OLD_FW_FLUSH             = 0x34
	IP_OLD_FW_GET               = 0x36
	IP_OLD_FW_RESETLOG          = 0x38
	IP_OLD_FW_ZERO              = 0x35
	IP_OPTIONS                  = 0x1
	IP_PORTRANGE                = 0x13
	IP_PORTRANGE_DEFAULT        = 0
	IP_PORTRANGE_HIGH           = 0x1
	IP_PORTRANGE_LOW            = 0x2
	IP_RECVDSTADDR              = 0x7
	IP_RECVIF                   = 0x14
	IP_RECVOPTS                 = 0x5
	IP_RECVRETOPTS              = 0x6
	IP_RECVTTL                  = 0x18
	IP_RETOPTS                  = 0x8
	IP_RF                       = 0x8000
	IP_RSVP_OFF                 = 0x10
	IP_RSVP_ON                  = 0xf
	IP_RSVP_VIF_OFF             = 0x12
	IP_RSVP_VIF_ON              = 0x11
	IP_STRIPHDR                 = 0x17
	IP_TOS                      = 0x3
	IP_TRAFFIC_MGT_BACKGROUND   = 0x41
	IP_TTL                      = 0x4
	MADV_CAN_REUSE              = 0x9
	MADV_DONTNEED               = 0x4
	MADV_FREE                   = 0x5
	MADV_FREE_REUSABLE          = 0x7
	MADV_FREE_REUSE             = 0x8
	MADV_NORMAL                 = 0
	MADV_RANDOM                 = 0x1
	MADV_SEQUENTIAL             = 0x2
	MADV_WILLNEED               = 0x3
	MADV_ZERO_WIRED_PAGES       = 0x6
	MAP_ANON                    = 0x1000
	MAP_COPY                    = 0x2
	MAP_FILE                    = 0
	MAP_FIXED                   = 0x10
	MAP_HASSEMAPHORE            = 0x200
	MAP_NOCACHE                 = 0x400
	MAP_NOEXTEND                = 0x100
	MAP_NORESERVE               = 0x40
	MAP_PRIVATE                 = 0x2
	MAP_RENAME                  = 0x20
	MAP_RESERVED0080            = 0x80
	MAP_SHARED                  = 0x1
	MCL_CURRENT                 = 0x1
	MCL_FUTURE                  = 0x2
	MSG_CTRUNC                  = 0x20
	MSG_DONTROUTE               = 0x4
	MSG_DONTWAIT                = 0x80
	MSG_EOF                     = 0x100
	MSG_EOR                     = 0x8
	MSG_FLUSH                   = 0x400
	MSG_HAVEMORE                = 0x2000
	MSG_HOLD                    = 0x800
	MSG_NEEDSA                  = 0x10000
	MSG_OOB                     = 0x1
	MSG_PEEK                    = 0x2
	MSG_RCVMORE                 = 0x4000
	MSG_SEND                    = 0x1000
	MSG_TRUNC                   = 0x10
	MSG_WAITALL                 = 0x40
	MSG_WAITSTREAM              = 0x200
	MS_ASYNC                    = 0x1
	MS_DEACTIVATE               = 0x8
	MS_INVALIDATE               = 0x2
	MS_KILLPAGES                = 0x4
	MS_SYNC                     = 0x10
	NAME_MAX                    = 0xff
	NET_RT_DUMP                 = 0x1
	NET_RT_DUMP2                = 0x7
	NET_RT_FLAGS                = 0x2
	NET_RT_IFLIST               = 0x3
	NET_RT_IFLIST2              = 0x6
	NET_RT_MAXID                = 0x8
	NET_RT_STAT                 = 0x4
	NET_RT_TRASH                = 0x5
	O_ACCMODE                   = 0x3
	O_ALERT                     = 0x20000000
	O_APPEND                    = 0x8
	O_ASYNC                     = 0x40
	O_CREAT                     = 0x200
	O_DIRECTORY                 = 0x100000
	O_DSYNC                     = 0x400000
	O_EVTONLY                   = 0x8000
	O_EXCL                      = 0x800
	O_EXLOCK                    = 0x20
	O_FSYNC                     = 0x80
	O_NDELAY                    = 0x4
	O_NOCTTY                    = 0x20000
	O_NOFOLLOW                  = 0x100
	O_NONBLOCK                  = 0x4
	O_POPUP                     = 0x80000000
	O_RDONLY                    = 0
	O_RDWR                      = 0x2
	O_SHLOCK                    = 0x10
	O_SYMLINK                   = 0x200000
	O_SYNC                      = 0x80
	O_TRUNC                     = 0x400
	O_WRONLY                    = 0x1
	PROT_EXEC                   = 0x4
	PROT_NONE                   = 0
	PROT_READ                   = 0x1
	PROT_WRITE                  = 0x2
	PT_ATTACH                   = 0xa
	PT_ATTACHEXC                = 0xe
	PT_CONTINUE                 = 0x7
	PT_DENY_ATTACH              = 0x1f
	PT_DETACH                   = 0xb
	PT_FIRSTMACH                = 0x20
	PT_FORCEQUOTA               = 0x1e
	PT_KILL                     = 0x8
	PT_READ_D                   = 0x2
	PT_READ_I                   = 0x1
	PT_READ_U                   = 0x3
	PT_SIGEXC                   = 0xc
	PT_STEP                     = 0x9
	PT_THUPDATE                 = 0xd
	PT_TRACE_ME                 = 0
	PT_WRITE_D                  = 0x5
	PT_WRITE_I                  = 0x4
	PT_WRITE_U                  = 0x6
	RTAX_AUTHOR                 = 0x6
	RTAX_BRD                    = 0x7
	RTAX_DST                    = 0
	RTAX_GATEWAY                = 0x1
	RTAX_GENMASK                = 0x3
	RTAX_IFA                    = 0x5
	RTAX_IFP                    = 0x4
	RTAX_MAX                    = 0x8
	RTAX_NETMASK                = 0x2
	RTA_AUTHOR                  = 0x40
	RTA_BRD                     = 0x80
	RTA_DST                     = 0x1
	RTA_GATEWAY                 = 0x2
	RTA_GENMASK                 = 0x8
	RTA_IFA                     = 0x20
	RTA_IFP                     = 0x10
	RTA_NETMASK                 = 0x4
	RTF_BLACKHOLE               = 0x1000
	RTF_BROADCAST               = 0x400000
	RTF_CLONING                 = 0x100
	RTF_CONDEMNED               = 0x2000000
	RTF_DELCLONE                = 0x80
	RTF_DONE                    = 0x40
	RTF_DYNAMIC                 = 0x10
	RTF_GATEWAY                 = 0x2
	RTF_HOST                    = 0x4
	RTF_IFREF                   = 0x4000000
	RTF_IFSCOPE                 = 0x1000000
	RTF_LLINFO                  = 0x400
	RTF_LOCAL                   = 0x200000
	RTF_MODIFIED                = 0x20
	RTF_MULTICAST               = 0x800000
	RTF_PINNED                  = 0x100000
	RTF_PRCLONING               = 0x10000
	RTF_PROTO1                  = 0x8000
	RTF_PROTO2                  = 0x4000
	RTF_PROTO3                  = 0x40000
	RTF_REJECT                  = 0x8
	RTF_STATIC                  = 0x800
	RTF_UP                      = 0x1
	RTF_WASCLONED               = 0x20000
	RTF_XRESOLVE                = 0x200
	RTM_ADD                     = 0x1
	RTM_CHANGE                  = 0x3
	RTM_DELADDR                 = 0xd
	RTM_DELETE                  = 0x2
	RTM_DELMADDR                = 0x10
	RTM_GET                     = 0x4
	RTM_GET2                    = 0x14
	RTM_IFINFO                  = 0xe
	RTM_IFINFO2                 = 0x12
	RTM_LOCK                    = 0x8
	RTM_LOSING                  = 0x5
	RTM_MISS                    = 0x7
	RTM_NEWADDR                 = 0xc
	RTM_NEWMADDR                = 0xf
	RTM_NEWMADDR2               = 0x13
	RTM_OLDADD                  = 0x9
	RTM_OLDDEL                  = 0xa
	RTM_REDIRECT                = 0x6
	RTM_RESOLVE                 = 0xb
	RTM_RTTUNIT                 = 0xf4240
	RTM_VERSION                 = 0x5
	RTV_EXPIRE                  = 0x4
	RTV_HOPCOUNT                = 0x2
	RTV_MTU                     = 0x1
	RTV_RPIPE                   = 0x8
	RTV_RTT                     = 0x40
	RTV_RTTVAR                  = 0x80
	RTV_SPIPE                   = 0x10
	RTV_SSTHRESH                = 0x20
	SCM_CREDS                   = 0x3
	SCM_RIGHTS                  = 0x1
	SCM_TIMESTAMP               = 0x2
	SHUT_RD                     = 0
	SHUT_RDWR                   = 0x2
	SHUT_WR                     = 0x1
	SIGABRT                     = 0x6
	SIGALRM                     = 0xe
	SIGBUS                      = 0xa
	SIGCHLD                     = 0x14
	SIGCONT                     = 0x13
	SIGEMT                      = 0x7
	SIGFPE                      = 0x8
	SIGHUP                      = 0x1
	SIGILL                      = 0x4
	SIGINFO                     = 0x1d
	SIGINT                      = 0x2
	SIGIO                       = 0x17
	SIGIOT                      = 0x6
	SIGKILL                     = 0x9
	SIGPIPE                     = 0xd
	SIGPROF                     = 0x1b
	SIGQUIT                     = 0x3
	SIGSEGV                     = 0xb
	SIGSTOP                     = 0x11
	SIGSYS                      = 0xc
	SIGTERM                     = 0xf
	SIGTRAP                     = 0x5
	SIGTSTP                     = 0x12
	SIGTTIN                     = 0x15
	SIGTTOU                     = 0x16
	SIGURG                      = 0x10
	SIGUSR1                     = 0x1e
	SIGUSR2                     = 0x1f
	SIGVTALRM                   = 0x1a
	SIGWINCH                    = 0x1c
	SIGXCPU                     = 0x18
	SIGXFSZ                     = 0x19
	SIOCADDMULTI                = 0x80206931
	SIOCAIFADDR                 = 0x8040691a
	SIOCALIFADDR                = 0x8118691d
	SIOCARPIPLL                 = 0xc0206928
	SIOCATMARK                  = 0x40047307
	SIOCAUTOADDR                = 0xc0206926
	SIOCAUTONETMASK             = 0x80206927
	SIOCDELMULTI                = 0x80206932
	SIOCDIFADDR                 = 0x80206919
	SIOCDIFPHYADDR              = 0x80206941
	SIOCDLIFADDR                = 0x8118691f
	SIOCGDRVSPEC                = 0xc01c697b
	SIOCGETSGCNT                = 0xc014721c
	SIOCGETVIFCNT               = 0xc014721b
	SIOCGETVLAN                 = 0xc020697f
	SIOCGHIWAT                  = 0x40047301
	SIOCGIFADDR                 = 0xc0206921
	SIOCGIFALTMTU               = 0xc0206948
	SIOCGIFASYNCMAP             = 0xc020697c
	SIOCGIFBOND                 = 0xc0206947
	SIOCGIFBRDADDR              = 0xc0206923
	SIOCGIFCONF                 = 0xc0086924
	SIOCGIFDEVMTU               = 0xc0206944
	SIOCGIFDSTADDR              = 0xc0206922
	SIOCGIFFLAGS                = 0xc0206911
	SIOCGIFGENERIC              = 0xc020693a
	SIOCGIFKPI                  = 0xc0206987
	SIOCGIFMAC                  = 0xc0206982
	SIOCGIFMEDIA                = 0xc0286938
	SIOCGIFMETRIC               = 0xc0206917
	SIOCGIFMTU                  = 0xc0206933
	SIOCGIFNETMASK              = 0xc0206925
	SIOCGIFPDSTADDR             = 0xc0206940
	SIOCGIFPHYS                 = 0xc0206935
	SIOCGIFPSRCADDR             = 0xc020693f
	SIOCGIFSTATUS               = 0xc331693d
	SIOCGIFVLAN                 = 0xc020697f
	SIOCGIFWAKEFLAGS            = 0xc0206988
	SIOCGLIFADDR                = 0xc118691e
	SIOCGLIFPHYADDR             = 0xc1186943
	SIOCGLOWAT                  = 0x40047303
	SIOCGPGRP                   = 0x40047309
	SIOCIFCREATE                = 0xc0206978
	SIOCIFCREATE2               = 0xc020697a
	SIOCIFDESTROY               = 0x80206979
	SIOCRSLVMULTI               = 0xc008693b
	SIOCSDRVSPEC                = 0x801c697b
	SIOCSETVLAN                 = 0x8020697e
	SIOCSHIWAT                  = 0x80047300
	SIOCSIFADDR                 = 0x8020690c
	SIOCSIFALTMTU               = 0x80206945
	SIOCSIFASYNCMAP             = 0x8020697d
	SIOCSIFBOND                 = 0x80206946
	SIOCSIFBRDADDR              = 0x80206913
	SIOCSIFDSTADDR              = 0x8020690e
	SIOCSIFFLAGS                = 0x80206910
	SIOCSIFGENERIC              = 0x80206939
	SIOCSIFKPI                  = 0x80206986
	SIOCSIFLLADDR               = 0x8020693c
	SIOCSIFMAC                  = 0x80206983
	SIOCSIFMEDIA                = 0xc0206937
	SIOCSIFMETRIC               = 0x80206918
	SIOCSIFMTU                  = 0x80206934
	SIOCSIFNETMASK              = 0x80206916
	SIOCSIFPHYADDR              = 0x8040693e
	SIOCSIFPHYS                 = 0x80206936
	SIOCSIFVLAN                 = 0x8020697e
	SIOCSLIFPHYADDR             = 0x81186942
	SIOCSLOWAT                  = 0x80047302
	SIOCSPGRP                   = 0x80047308
	SOCK_DGRAM                  = 0x2
	SOCK_MAXADDRLEN             = 0xff
	SOCK_RAW                    = 0x3
	SOCK_RDM                    = 0x4
	SOCK_SEQPACKET              = 0x5
	SOCK_STREAM                 = 0x1
	SOL_SOCKET                  = 0xffff
	SOMAXCONN                   = 0x80
	SO_ACCEPTCONN               = 0x2
	SO_BROADCAST                = 0x20
	SO_DEBUG                    = 0x1
	SO_DONTROUTE                = 0x10
	SO_DONTTRUNC                = 0x2000
	SO_ERROR                    = 0x1007
	SO_KEEPALIVE                = 0x8
	SO_LABEL                    = 0x1010
	SO_LINGER                   = 0x80
	SO_LINGER_SEC               = 0x1080
	SO_NKE                      = 0x1021
	SO_NOADDRERR                = 0x1023
	SO_NOSIGPIPE                = 0x1022
	SO_NOTIFYCONFLICT           = 0x1026
	SO_NP_EXTENSIONS            = 0x1083
	SO_NREAD                    = 0x1020
	SO_NWRITE                   = 0x1024
	SO_OOBINLINE                = 0x100
	SO_PEERLABEL                = 0x1011
	SO_RANDOMPORT               = 0x1082
	SO_RCVBUF                   = 0x1002
	SO_RCVLOWAT                 = 0x1004
	SO_RCVTIMEO                 = 0x1006
	SO_RESTRICTIONS             = 0x1081
	SO_RESTRICT_DENYIN          = 0x1
	SO_RESTRICT_DENYOUT         = 0x2
	SO_RESTRICT_DENYSET         = 0x80000000
	SO_REUSEADDR                = 0x4
	SO_REUSEPORT                = 0x200
	SO_REUSESHAREUID            = 0x1025
	SO_SNDBUF                   = 0x1001
	SO_SNDLOWAT                 = 0x1003
	SO_SNDTIMEO                 = 0x1005
	SO_TIMESTAMP                = 0x400
	SO_TYPE                     = 0x1008
	SO_UPCALLCLOSEWAIT          = 0x1027
	SO_USELOOPBACK              = 0x40
	SO_WANTMORE                 = 0x4000
	SO_WANTOOBFLAG              = 0x8000
	S_IEXEC                     = 0x40
	S_IFBLK                     = 0x6000
	S_IFCHR                     = 0x2000
	S_IFDIR                     = 0x4000
	S_IFIFO                     = 0x1000
	S_IFLNK                     = 0xa000
	S_IFMT                      = 0xf000
	S_IFREG                     = 0x8000
	S_IFSOCK                    = 0xc000
	S_IFWHT                     = 0xe000
	S_IREAD                     = 0x100
	S_IRGRP                     = 0x20
	S_IROTH                     = 0x4
	S_IRUSR                     = 0x100
	S_IRWXG                     = 0x38
	S_IRWXO                     = 0x7
	S_IRWXU                     = 0x1c0
	S_ISGID                     = 0x400
	S_ISTXT                     = 0x200
	S_ISUID                     = 0x800
	S_ISVTX                     = 0x200
	S_IWGRP                     = 0x10
	S_IWOTH                     = 0x2
	S_IWRITE                    = 0x80
	S_IWUSR                     = 0x80
	S_IXGRP                     = 0x8
	S_IXOTH                     = 0x1
	S_IXUSR                     = 0x40
	TCP_CONNECTIONTIMEOUT       = 0x20
	TCP_KEEPALIVE               = 0x10
	TCP_MAXBURST                = 0x4
	TCP_MAXHLEN                 = 0x3c
	TCP_MAXOLEN                 = 0x28
	TCP_MAXSEG                  = 0x2
	TCP_MAXWIN                  = 0xffff
	TCP_MAX_SACK                = 0x3
	TCP_MAX_WINSHIFT            = 0xe
	TCP_MINMSS                  = 0xd8
	TCP_MINMSSOVERLOAD          = 0x3e8
	TCP_MSS                     = 0x200
	TCP_NODELAY                 = 0x1
	TCP_NOOPT                   = 0x8
	TCP_NOPUSH                  = 0x4
	WCONTINUED                  = 0x10
	WCOREFLAG                   = 0x80
	WEXITED                     = 0x4
	WNOHANG                     = 0x1
	WNOWAIT                     = 0x20
	WORDSIZE                    = 0x20
	WSTOPPED                    = 0x8
	WUNTRACED                   = 0x2
)

// Types


// Error table
var errors = [...]string{
	1:   "operation not permitted",
	2:   "no such file or directory",
	3:   "no such process",
	4:   "interrupted system call",
	5:   "input/output error",
	6:   "device not configured",
	7:   "argument list too long",
	8:   "exec format error",
	9:   "bad file descriptor",
	10:  "no child processes",
	11:  "resource deadlock avoided",
	12:  "cannot allocate memory",
	13:  "permission denied",
	14:  "bad address",
	15:  "block device required",
	16:  "resource busy",
	17:  "file exists",
	18:  "cross-device link",
	19:  "operation not supported by device",
	20:  "not a directory",
	21:  "is a directory",
	22:  "invalid argument",
	23:  "too many open files in system",
	24:  "too many open files",
	25:  "inappropriate ioctl for device",
	26:  "text file busy",
	27:  "file too large",
	28:  "no space left on device",
	29:  "illegal seek",
	30:  "read-only file system",
	31:  "too many links",
	32:  "broken pipe",
	33:  "numerical argument out of domain",
	34:  "result too large",
	35:  "resource temporarily unavailable",
	36:  "operation now in progress",
	37:  "operation already in progress",
	38:  "socket operation on non-socket",
	39:  "destination address required",
	40:  "message too long",
	41:  "protocol wrong type for socket",
	42:  "protocol not available",
	43:  "protocol not supported",
	44:  "socket type not supported",
	45:  "operation not supported",
	46:  "protocol family not supported",
	47:  "address family not supported by protocol family",
	48:  "address already in use",
	49:  "can't assign requested address",
	50:  "network is down",
	51:  "network is unreachable",
	52:  "network dropped connection on reset",
	53:  "software caused connection abort",
	54:  "connection reset by peer",
	55:  "no buffer space available",
	56:  "socket is already connected",
	57:  "socket is not connected",
	58:  "can't send after socket shutdown",
	59:  "too many references: can't splice",
	60:  "operation timed out",
	61:  "connection refused",
	62:  "too many levels of symbolic links",
	63:  "file name too long",
	64:  "host is down",
	65:  "no route to host",
	66:  "directory not empty",
	67:  "too many processes",
	68:  "too many users",
	69:  "disc quota exceeded",
	70:  "stale NFS file handle",
	71:  "too many levels of remote in path",
	72:  "RPC struct is bad",
	73:  "RPC version wrong",
	74:  "RPC prog. not avail",
	75:  "program version wrong",
	76:  "bad procedure for program",
	77:  "no locks available",
	78:  "function not implemented",
	79:  "inappropriate file type or format",
	80:  "authentication error",
	81:  "need authenticator",
	82:  "device power is off",
	83:  "device error",
	84:  "value too large to be stored in data type",
	85:  "bad executable (or shared library)",
	86:  "bad CPU type in executable",
	87:  "shared library version mismatch",
	88:  "malformed Mach-o file",
	89:  "operation canceled",
	90:  "identifier removed",
	91:  "no message of desired type",
	92:  "illegal byte sequence",
	93:  "attribute not found",
	94:  "bad message",
	95:  "EMULTIHOP (Reserved)",
	96:  "no message available on STREAM",
	97:  "ENOLINK (Reserved)",
	98:  "no STREAM resources",
	99:  "not a STREAM",
	100: "protocol error",
	101: "STREAM ioctl timeout",
	102: "operation not supported on socket",
	103: "policy not found",
}
