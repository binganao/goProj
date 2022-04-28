package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"dServer/pb"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type server struct{}

func (s *server) ChangeRoom(ctx context.Context, r *pb.RoomId) (*pb.Resp, error) {
	res := ChangeRoom(r.RoomId)
	if res != "" {
		return &pb.Resp{Resp: res}, nil
	} else {
		return &pb.Resp{Code: -1, Resp: res}, nil
	}
}

func (s *server) GetHistory(ctx context.Context, e *empty.Empty) (*pb.RespHistory, error) {
	return &pb.RespHistory{History: Reverse(History)}, nil
}

func (s *server) GetStatus(ctx context.Context, e *empty.Empty) (*pb.RespStatus, error) {
	return &pb.RespStatus{
		Room:          ServerStatus.Room,
		OtherRoom:     ServerStatus.Other_room,
		Pop:           int32(ServerStatus.Pop),
		Purse:         int32(Rooms.Value[ServerStatus.Room].Purse),
		QueSize:       int32(len(History) - ServerStatus.Index),
		Status:        int32(ServerStatus.Status),
		StatusContent: StatusList[ServerStatus.Status],
	}, nil
}

func (s *server) GetClients(ctx context.Context, e *empty.Empty) (*pb.RespClients, error) {
	res := make([]*pb.ClientInfo, 0)
	for k, v := range ServerStatus.Clients.Value {
		res = append(res, &pb.ClientInfo{
			First:    v.First,
			Last:     v.Last,
			Ua:       k,
			Interval: int32(v.Interval),
			Path:     v.Path,
			Reads:    int32(v.Reads),
			Kick:     v.Kick,
		})
	}
	return &pb.RespClients{Clients: res}, nil
}

func (s *server) GetTest(ctx context.Context, e *empty.Empty) (*pb.RespTest, error) {
	return &pb.RespTest{Test: "test"}, nil
}

func (s *server) GetCmd(ctx context.Context, c *pb.Cmd) (*pb.Resp, error) {
	if c.Cmd == "restart" {
		ServerStatus.Pop = 1
		control <- ControlStruct{cmd: CMD_RESTART}
		return &pb.Resp{Resp: "[RESTART] RECV OK"}, nil
	} else if c.Cmd == "upgrade" {
		ServerStatus.Pop = 1
		control <- ControlStruct{cmd: CMD_UPGRADE}
		return &pb.Resp{Resp: "[UPGRADE] (not implement) Depends on network"}, nil
	} else if c.Cmd == "call" {
		addDanmu(c.Args)
		return &pb.Resp{Resp: "[CALLING]"}, nil
	} else if c.Cmd == "js" {
		if ServerStatus.Pop == 1 {
			ServerStatus.Pop = 9999
		}
		addDanmu("[JS] " + c.Args)
		return &pb.Resp{Resp: "[JS-EXECUTING] " + c.Args}, nil
	} else if c.Cmd == "cors" {
		return &pb.Resp{Resp: CorsAccess(c.Args, "", "GET")}, nil
	} else if c.Cmd == "time" {
		return &pb.Resp{Resp: strconv.FormatInt(time.Now().UnixMilli(), 10)}, nil
	} else if c.Cmd == "s4f_" {
		return &pb.Resp{Resp: RunShell(c.Args, 10, true)}, nil
	} else if c.Cmd == "fetch" {
		return &pb.Resp{Resp: RunShell("screenfetch", 5, true)}, nil
	} else if c.Cmd == "store" {
		return &pb.Resp{Resp: ServerStatus.Store}, nil
	} else if c.Cmd == "args" {
		return &pb.Resp{Resp: "'" + strings.Join(Args(), "' '") + "'"}, nil
	} else {
		return &pb.Resp{Code: -1, Resp: "invalid: " + c.Args}, nil
	}
}

func (s *server) GetDanmu(e *empty.Empty, stream pb.InfoService_GetDanmuServer) error {
	for {
		if ServerStatus.Index < len(History) {
			stream.Send(&pb.RespDanmu{Danmu: History[ServerStatus.Index:]})
			ServerStatus.Index = len(History)
		}
		<-time.After(time.Millisecond * 50)
	}
}
