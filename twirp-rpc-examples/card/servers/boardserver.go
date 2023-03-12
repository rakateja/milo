package servers

import (
	"context"
	"time"

	timestampPb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/rakateja/milo/twirp-rpc-examples/card/domains/board"
	pb "github.com/rakateja/milo/twirp-rpc-examples/card/proto/rpcproto"
	"github.com/twitchtv/twirp"
)

type BoardServer struct {
	boardSvc *board.Service
}

func NewBoardServer(svc *board.Service) pb.BoardService {
	return &BoardServer{boardSvc: svc}
}

func (svc *BoardServer) CreateBoard(ctx context.Context, input *pb.BoardCreateInput) (*pb.Board, error) {
	res, err := svc.boardSvc.Create(ctx, ToBoardInputFromCreateInputPb(input))
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}
	return ToBoardPb(*res), nil
}

func (svc *BoardServer) UpdateBoard(ctx context.Context, input *pb.BoardUpdateInput) (*pb.Board, error) {
	return nil, nil
}

func (svc *BoardServer) GetByID(ctx context.Context, input *pb.GetByIDInput) (*pb.Board, error) {
	res, err := svc.boardSvc.ResolveByID(ctx, input.Id)
	if err != nil {
		return nil, twirp.NewError(twirp.Internal, err.Error())
	}
	return ToBoardPb(*res), nil
}

func ToBoardInputFromCreateInputPb(pbInput *pb.BoardCreateInput) board.Input {
	return board.Input{
		Title: pbInput.Title,
	}
}

func ToBoardInputFromUpdateInputPb(pbInput *pb.BoardUpdateInput) board.Input {
	return board.Input{
		Title: pbInput.Title,
	}
}

func ToBoardPb(entity board.Board) *pb.Board {
	return &pb.Board{
		Id:        entity.ID,
		Title:     entity.Title,
		CreatedAt: ToTimestampPb(entity.CreatedAt),
		UpdatedAt: ToTimestampPb(entity.UpdatedAt),
	}
}

func ToTimestampPb(ts time.Time) *timestampPb.Timestamp {
	return &timestampPb.Timestamp{
		Seconds: ts.Unix(),
		Nanos:   int32(ts.UnixNano()),
	}
}
