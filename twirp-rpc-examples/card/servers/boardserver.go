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

func (svc *BoardServer) AddMember(context.Context, *pb.AddMemberInput) (*pb.Board, error) {
	return nil, nil
}

func (svc *BoardServer) AddLabel(context.Context, *pb.AddLabelInput) (*pb.Board, error) {
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
		Title:   pbInput.Title,
		Members: ToBoardMemberInputFromPb(pbInput.Members),
		Lists:   ToBoardListInputFromPb(pbInput.Lists),
		Labels:  ToBoardLabelInputFromPb(pbInput.Labels),
	}
}

func ToBoardMemberInputFromPb(ls []*pb.AddMemberInput) []board.MemberInput {
	res := make([]board.MemberInput, 0)
	for _, inputPb := range ls {
		res = append(res, board.MemberInput{
			UserID: inputPb.UserId,
		})
	}
	return res
}

func ToBoardLabelInputFromPb(ls []*pb.AddLabelInput) []board.LabelInput {
	res := make([]board.LabelInput, 0)
	for _, inputPb := range ls {
		res = append(res, board.LabelInput{
			Title: inputPb.Name,
			Color: inputPb.Color,
		})
	}
	return res
}

func ToBoardListInputFromPb(ls []*pb.AddListInput) []board.ListInput {
	res := make([]board.ListInput, 0)
	for _, inputPb := range ls {
		res = append(res, board.ListInput{
			Title:    inputPb.Name,
			Position: int(inputPb.Position),
		})
	}
	return res
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
		Members:   ToBoardMembersPb(entity.Members),
		Lists:     ToBoardListsPb(entity.Lists),
		Labels:    ToBoardLabelsPb(entity.Labels),
		CreatedAt: ToTimestampPb(entity.CreatedAt),
		UpdatedAt: ToTimestampPb(entity.UpdatedAt),
	}
}

func ToTimestampPb(ts time.Time) *timestampPb.Timestamp {
	return &timestampPb.Timestamp{
		Seconds: ts.Unix(),
		Nanos:   int32(ts.Nanosecond()),
	}
}

func ToBoardMembersPb(ls []board.BoardMember) []*pb.BoardMember {
	res := make([]*pb.BoardMember, 0)
	for _, entity := range ls {
		res = append(res, &pb.BoardMember{
			Id:        entity.ID,
			BoardId:   entity.BoardID,
			UserId:    entity.UserID,
			CreatedAt: ToTimestampPb(entity.CreatedAt),
		})
	}
	return res
}

func ToBoardListsPb(ls []board.BoardList) []*pb.BoardList {
	res := make([]*pb.BoardList, 0)
	for _, entity := range ls {
		res = append(res, &pb.BoardList{
			Id:        entity.ID,
			BoardId:   entity.BoardID,
			PublicId:  entity.PublicID,
			Title:     entity.Title,
			Position:  int32(entity.Position),
			CreatedAt: ToTimestampPb(entity.CreatedAt),
			UpdatedAt: ToTimestampPb(entity.UpdatedAt),
		})
	}
	return res
}

func ToBoardLabelsPb(ls []board.Label) []*pb.BoardLabel {
	res := make([]*pb.BoardLabel, 0)
	for _, entity := range ls {
		res = append(res, &pb.BoardLabel{
			Id:        entity.ID,
			BoardId:   entity.BoardID,
			Slug:      entity.Slug,
			Title:     entity.Title,
			Color:     entity.Color,
			CreatedAt: ToTimestampPb(entity.CreatedAt),
			UpdatedAt: ToTimestampPb(entity.CreatedAt),
		})
	}
	return res
}
