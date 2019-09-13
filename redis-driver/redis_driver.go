package redis_driver

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"time"
)

type Error struct {
	errno int
	str   string
}

const (
	redisBlock = 0x1
)

type RedisDriver struct {
	context *redisContext
	auth    string
}

func (r *RedisDriver) Connect(ip string, port int) {
	r.context = newRedisContext()
	r.context.flags |= redisBlock

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
	}

}

func (r *RedisDriver) Command(format string, args ...string) int {
	return 0
}

func (r *RedisDriver) FormatCommand(format string, args []string) (formatted string, err error) {
	fLen := len(format)
	if fLen == 1 {
		if format[0] == '%' {
			// TODO: throw error
		} else {
			formatted = format
		}
	} else {
	}
	buf := bytes.Buffer{}
	c := format[i]
	i := 0
	for i < fLen {
		if c != '%' || i == fLen-1 {
			if c == ' ' {

			} else {

			}
		} else {
			switch c {
			case 's':
			case 'b':
			case '%':
			default:

			}
		}
	}
	return

}

type redisContext struct {
	err    Error
	fd     int
	flags  int
	outBuf string
	reader *redisReader
}

func newRedisContext() *redisContext {
	return &redisContext{
		err:    Error{},
		fd:     0,
		flags:  0,
		outBuf: "",
		reader: newRedisReader(),
	}
}

// redis reply types
const (
	redisReplyString = iota
	redisReplyArray
	redisReplyInteger
	redisReplyNil
	redisReplyStatus
	redisReplyError
)

type RedisReply struct {
	typ      int
	integer  int64
	str      string
	elements []*RedisReply
}

type redisReadTask struct {
	typ         int
	idx         int
	obj         interface{}
	parent      *redisReadTask
	privateData interface{}
}

const (
	redisReaderMaxBuf = 1024 * 16
)

type redisReader struct {
	err Error

	buf              string
	pos, len, maxBuf int

	rsTask []redisReadTask
	rIdx   int
	reply  interface{}

	privateDate interface{}
}

func newRedisReader() *redisReader {
	return &redisReader{
		err:         Error{},
		buf:         "",
		pos:         0,
		len:         0,
		maxBuf:      redisReaderMaxBuf,
		rsTask:      nil,
		rIdx:        -1,
		reply:       nil,
		privateDate: nil,
	}
}

func (r *redisReader) readBytes(nBytes int) []byte {
	return nil
}

func (r *redisReader) readLine(size int) []byte {
	return nil
}

func (r *redisReader) moveToNextTask() {

}

func (r *redisReader) processLineItem() int {
	return 0
}

func (r *redisReader) processBulkItem() int {
	return 0
}

func (r *redisReader) processMultiBulkItem() int {
	return 0
}

func (r *redisReader) processItem() int {
	return 0
}

func (r *redisReader) feed(buf string) int {
	return 0
}

func (r *redisReader) getReply(reply []interface{}) int {
	return 0
}

func createReplyObject(typ int) *RedisReply {
	return &RedisReply{typ: typ}
}

func createStringObject(task *redisReadTask, str string) *RedisReply {
	if task.typ != redisReplyError &&
		task.typ != redisReplyStatus &&
		task.typ != redisReplyString {
		panic(errors.New("invalid argument"))
	}

	r := createReplyObject(task.typ)
	r.str = str
	if task.parent != nil {
		parent := task.parent.obj.(*RedisReply)
		if parent.typ != redisReplyArray {
			panic(errors.New("invalid"))
		}
		parent.elements[task.idx] = r
	}
	return r
}

func createArrayObject(task *redisReadTask, capacity int) *RedisReply {
	r := createReplyObject(redisReplyArray)
	if capacity > 0 {
		r.elements = make([]*RedisReply, capacity)
	}
	if task.parent != nil {
		parent := task.parent.obj.(*RedisReply)
		if parent.typ != redisReplyArray {
			panic(errors.New("invalid"))
		}
		parent.elements[task.idx] = r
	}
	return r
}

func createIntegerObject(task *redisReadTask, value int64) *RedisReply {
	r := createReplyObject(redisReplyInteger)
	r.integer = value
	if task.parent != nil {
		parent := task.parent.obj.(*RedisReply)
		if parent.typ != redisReplyArray {
			panic(errors.New("invalid"))
		}
		parent.elements[task.idx] = r
	}
	return r
}

func createNilObject(task *redisReadTask) *RedisReply {
	r := createReplyObject(redisReplyNil)
	if task.parent != nil {
		parent := task.parent.obj.(*RedisReply)
		if parent.typ != redisReplyArray {
			panic(errors.New("invalid"))
		}
		parent.elements[task.idx] = r
	}
	return r
}

func redisConnect(ip string, port int) {

}

func redisConnectWithTimeout(ip string, port int, timeout time.Duration) {

}
