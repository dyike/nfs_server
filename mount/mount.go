package mount

import (
	"bytes"
	"encoding/binary"
)

// RFC 1057

const (
	PROGRAM    uint32 = 100005
	VERSION    uint32 = 3
	MNTPATHLEN uint32 = 1024 // Maximum bytes in a path name
	MNTNAMLEN  uint32 = 255  // Maximum bytes in a name
	FHSIZE3    uint32 = 64   // Maximum bytes in v3 file handle
)

type mountstat3 uint32

const (
	MNT3_OK             mountstat3 = 0     // no error
	MNT3ERR_PERM        mountstat3 = 1     // not owner
	MNT3ERR_NOENT       mountstat3 = 2     // no such file or directory
	MNT3ERR_IO          mountstat3 = 5     // io error
	MNT3ERR_ACCES       mountstat3 = 13    // permission denied
	MNT3ERR_NOTDIR      mountstat3 = 20    // not a directory
	MNT3ERR_INVAL       mountstat3 = 22    // invalid argument
	MNT3ERR_NAMETOOLONG mountstat3 = 63    // filename too long
	MNT3ERR_NOTSUPP     mountstat3 = 10004 // opeation not supported
	MNT3ERR_SERVERFAULT mountstat3 = 10006 // a failure on the server
)

func SerializeMountStat3(ms mountstat3) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(ms))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeMountStat3(data []byte) (mountstat3, error) {
	var ms uint32
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.BigEndian, &ms)
	if err != nil {
		return 0, err
	}
	return mountstat3(ms), nil
}
