package winapi

var SE_DEBUG_PRIVILEGE = 20

var STANDARD_RIGHTS_REQUIRED = 0xf0000
var SYNCHRONIZE              = 0x00100000

var TOKEN_QUERY = 8
var TIC_TOKEN_USER = 1

type SID struct {}

type SID_AND_ATTRIBUTES struct{
    Sid        *SID
    Attributes uint32
}

type TOKEN_USER struct {
    User SID_AND_ATTRIBUTES
}
