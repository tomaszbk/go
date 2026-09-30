package syscall

var NewProcThreadAttributeList = newProcThreadAttributeList
var UpdateProcThreadAttribute = updateProcThreadAttribute
var DeleteProcThreadAttributeList = deleteProcThreadAttributeList

const PROC_THREAD_ATTRIBUTE_HANDLE_LIST = _PROC_THREAD_ATTRIBUTE_HANDLE_LIST

var EncodeWTF16 = encodeWTF16
var DecodeWTF16 = decodeWTF16

var EnvSorted = envSorted

var RemoveExtendedPrefix = removeExtendedPrefix
