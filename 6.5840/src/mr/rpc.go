package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

const (
	MapFinished = iota
	ReduceFinished
	Idle
)

type RegisterArgs struct {
	WorkerId int
}

type RegisterReply struct {
	WorkerId int
}

type HeartRequest struct {
	WorkerId int
}

type HeartReply struct {
}

type TaskRequest struct {
	WorkerState int
	WorkerId    int
	FileName    string
	ReduceId    int
	MapId       int // add MapId
}

type TaskResponse struct {
	TaskType string
	FileName string
	ReduceId int
	MapId    int
	MapCount int
	NReduce  int
	AllDone  bool
	Data     []byte // add Data
}

// Add your RPC definitions here.
type SendFileArgs struct {
	MapId    int
	ReduceId int
	Data     []byte
}

type SendFileReply struct {
	Success bool
}

type FetchReduceInputArgs struct {
	ReduceId int
	MapId    int
}

type FetchReduceInputReply struct {
	Data []byte
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
