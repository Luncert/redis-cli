package redis_driver

import "errors"

type redisContext struct {
	err    int
	errStr string
	fd     int
	flags  int
	outBuf string
	reader *redisReader
}

type RedisDriver struct {
	context *redisContext
	auth    string
}

func (r *RedisDriver) Connect(ip string, port int) {
	r.initContext()
	r.context.flags |= redisBlock
}

func (r *RedisDriver) initContext() {
	r.context = &redisContext{
		err:    0,
		errStr: "",
		fd:     0,
		flags:  0,
		outBuf: "",
		reader: nil,
	}
}

func (r *RedisDriver) Auth(auth string) {

}

func (r *RedisDriver) command(cmd string) {

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
