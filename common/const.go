package common

import "unsafe"

const (
	MaxTaosSqlLen   = 1048576
	DefaultUser     = "root"
	DefaultPassword = "taosdata"
)

const (
	PrecisionMilliSecond = 0
	PrecisionMicroSecond = 1
	PrecisionNanoSecond  = 2
)

const (
	TSDB_OPTION_LOCALE = iota
	TSDB_OPTION_CHARSET
	TSDB_OPTION_TIMEZONE
	TSDB_OPTION_CONFIGDIR
	TSDB_OPTION_SHELL_ACTIVITY_TIMER
	TSDB_MAX_OPTIONS
)

const (
	TSDB_DATA_TYPE_NULL       = 0  // 1 bytes
	TSDB_DATA_TYPE_BOOL       = 1  // 1 bytes
	TSDB_DATA_TYPE_TINYINT    = 2  // 1 byte
	TSDB_DATA_TYPE_SMALLINT   = 3  // 2 bytes
	TSDB_DATA_TYPE_INT        = 4  // 4 bytes
	TSDB_DATA_TYPE_BIGINT     = 5  // 8 bytes
	TSDB_DATA_TYPE_FLOAT      = 6  // 4 bytes
	TSDB_DATA_TYPE_DOUBLE     = 7  // 8 bytes
	TSDB_DATA_TYPE_BINARY     = 8  // string
	TSDB_DATA_TYPE_TIMESTAMP  = 9  // 8 bytes
	TSDB_DATA_TYPE_NCHAR      = 10 // unicode string
	TSDB_DATA_TYPE_UTINYINT   = 11 // 1 byte
	TSDB_DATA_TYPE_USMALLINT  = 12 // 2 bytes
	TSDB_DATA_TYPE_UINT       = 13 // 4 bytes
	TSDB_DATA_TYPE_UBIGINT    = 14 // 8 bytes
	TSDB_DATA_TYPE_JSON       = 15
	TSDB_DATA_TYPE_VARBINARY  = 16
	TSDB_DATA_TYPE_DECIMAL    = 17
	TSDB_DATA_TYPE_BLOB       = 18
	TSDB_DATA_TYPE_MEDIUMBLOB = 19
	TSDB_DATA_TYPE_MAX        = 20
)

const (
	TSDB_DATA_TYPE_NULL_Str      = "NULL"
	TSDB_DATA_TYPE_BOOL_Str      = "BOOL"
	TSDB_DATA_TYPE_TINYINT_Str   = "TINYINT"
	TSDB_DATA_TYPE_SMALLINT_Str  = "SMALLINT"
	TSDB_DATA_TYPE_INT_Str       = "INT"
	TSDB_DATA_TYPE_BIGINT_Str    = "BIGINT"
	TSDB_DATA_TYPE_FLOAT_Str     = "FLOAT"
	TSDB_DATA_TYPE_DOUBLE_Str    = "DOUBLE"
	TSDB_DATA_TYPE_BINARY_Str    = "VARCHAR"
	TSDB_DATA_TYPE_TIMESTAMP_Str = "TIMESTAMP"
	TSDB_DATA_TYPE_NCHAR_Str     = "NCHAR"
	TSDB_DATA_TYPE_UTINYINT_Str  = "TINYINT UNSIGNED"
	TSDB_DATA_TYPE_USMALLINT_Str = "SMALLINT UNSIGNED"
	TSDB_DATA_TYPE_UINT_Str      = "INT UNSIGNED"
	TSDB_DATA_TYPE_UBIGINT_Str   = "BIGINT UNSIGNED"
	TSDB_DATA_TYPE_JSON_Str      = "JSON"
)

var TypeNameMap = map[int]string{
	TSDB_DATA_TYPE_NULL:      TSDB_DATA_TYPE_NULL_Str,
	TSDB_DATA_TYPE_BOOL:      TSDB_DATA_TYPE_BOOL_Str,
	TSDB_DATA_TYPE_TINYINT:   TSDB_DATA_TYPE_TINYINT_Str,
	TSDB_DATA_TYPE_SMALLINT:  TSDB_DATA_TYPE_SMALLINT_Str,
	TSDB_DATA_TYPE_INT:       TSDB_DATA_TYPE_INT_Str,
	TSDB_DATA_TYPE_BIGINT:    TSDB_DATA_TYPE_BIGINT_Str,
	TSDB_DATA_TYPE_FLOAT:     TSDB_DATA_TYPE_FLOAT_Str,
	TSDB_DATA_TYPE_DOUBLE:    TSDB_DATA_TYPE_DOUBLE_Str,
	TSDB_DATA_TYPE_BINARY:    TSDB_DATA_TYPE_BINARY_Str,
	TSDB_DATA_TYPE_TIMESTAMP: TSDB_DATA_TYPE_TIMESTAMP_Str,
	TSDB_DATA_TYPE_NCHAR:     TSDB_DATA_TYPE_NCHAR_Str,
	TSDB_DATA_TYPE_UTINYINT:  TSDB_DATA_TYPE_UTINYINT_Str,
	TSDB_DATA_TYPE_USMALLINT: TSDB_DATA_TYPE_USMALLINT_Str,
	TSDB_DATA_TYPE_UINT:      TSDB_DATA_TYPE_UINT_Str,
	TSDB_DATA_TYPE_UBIGINT:   TSDB_DATA_TYPE_UBIGINT_Str,
	TSDB_DATA_TYPE_JSON:      TSDB_DATA_TYPE_JSON_Str,
}

var NameTypeMap = map[string]int{
	TSDB_DATA_TYPE_NULL_Str:      TSDB_DATA_TYPE_NULL,
	TSDB_DATA_TYPE_BOOL_Str:      TSDB_DATA_TYPE_BOOL,
	TSDB_DATA_TYPE_TINYINT_Str:   TSDB_DATA_TYPE_TINYINT,
	TSDB_DATA_TYPE_SMALLINT_Str:  TSDB_DATA_TYPE_SMALLINT,
	TSDB_DATA_TYPE_INT_Str:       TSDB_DATA_TYPE_INT,
	TSDB_DATA_TYPE_BIGINT_Str:    TSDB_DATA_TYPE_BIGINT,
	TSDB_DATA_TYPE_FLOAT_Str:     TSDB_DATA_TYPE_FLOAT,
	TSDB_DATA_TYPE_DOUBLE_Str:    TSDB_DATA_TYPE_DOUBLE,
	TSDB_DATA_TYPE_BINARY_Str:    TSDB_DATA_TYPE_BINARY,
	TSDB_DATA_TYPE_TIMESTAMP_Str: TSDB_DATA_TYPE_TIMESTAMP,
	TSDB_DATA_TYPE_NCHAR_Str:     TSDB_DATA_TYPE_NCHAR,
	TSDB_DATA_TYPE_UTINYINT_Str:  TSDB_DATA_TYPE_UTINYINT,
	TSDB_DATA_TYPE_USMALLINT_Str: TSDB_DATA_TYPE_USMALLINT,
	TSDB_DATA_TYPE_UINT_Str:      TSDB_DATA_TYPE_UINT,
	TSDB_DATA_TYPE_UBIGINT_Str:   TSDB_DATA_TYPE_UBIGINT,
	TSDB_DATA_TYPE_JSON_Str:      TSDB_DATA_TYPE_JSON,
}

const (
	TMQ_RES_INVALID    = -1
	TMQ_RES_DATA       = 1
	TMQ_RES_TABLE_META = 2
	TMQ_RES_METADATA   = 3
)

var TypeLengthMap = map[int]int{
	TSDB_DATA_TYPE_NULL:      1,
	TSDB_DATA_TYPE_BOOL:      1,
	TSDB_DATA_TYPE_TINYINT:   1,
	TSDB_DATA_TYPE_SMALLINT:  2,
	TSDB_DATA_TYPE_INT:       4,
	TSDB_DATA_TYPE_BIGINT:    8,
	TSDB_DATA_TYPE_FLOAT:     4,
	TSDB_DATA_TYPE_DOUBLE:    8,
	TSDB_DATA_TYPE_TIMESTAMP: 8,
	TSDB_DATA_TYPE_UTINYINT:  1,
	TSDB_DATA_TYPE_USMALLINT: 2,
	TSDB_DATA_TYPE_UINT:      4,
	TSDB_DATA_TYPE_UBIGINT:   8,
}

const (
	Int8Size    = unsafe.Sizeof(int8(0))
	Int16Size   = unsafe.Sizeof(int16(0))
	Int32Size   = unsafe.Sizeof(int32(0))
	Int64Size   = unsafe.Sizeof(int64(0))
	UInt8Size   = unsafe.Sizeof(uint8(0))
	UInt16Size  = unsafe.Sizeof(uint16(0))
	UInt32Size  = unsafe.Sizeof(uint32(0))
	UInt64Size  = unsafe.Sizeof(uint64(0))
	Float32Size = unsafe.Sizeof(float32(0))
	Float64Size = unsafe.Sizeof(float64(0))
)
