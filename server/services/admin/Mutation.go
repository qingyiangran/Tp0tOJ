package admin

import (
	"context"
	"github.com/kataras/go-sessions/v3"
	"server/services/types"
)

type MutationResolver struct {
}

func (r *MutationResolver) BulletinPub(input types.BulletinPubInput, ctx context.Context) (*types.BulletinPubResult, error) {
	session := ctx.Value("session").(*sessions.Session)
}
func (r *MutationResolver) UserInfoUpdate(input types.ChallengeMutateInput, ctx context.Context) (*types.UserInfoUpdateResult, error) {

}

func (r *MutationResolver) ChallengeMutate(input types.ChallengeMutateInput, ctx context.Context) (*types.ChallengeMutateResult, error) {

}
func (r *MutationResolver) ChallengeRemove(id string) (*types.ChallengeRemoveResult, error) {

}
func (r *MutationResolver) WarmUp() (bool, error) {

}
