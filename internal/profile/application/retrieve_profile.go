package application

import (
	api "hc/api/profile"
	"hc/internal/profile/domain"
)

type InfoRetriever struct {
	ProfileStore domain.Store
}

func (i InfoRetriever) Retrieve(accountId int) (api.Info, error) {
	profile, err := i.ProfileStore.FindByAccountID(accountId)
	if err != nil {
		return api.Info{}, err
	}

	return api.Info{
		Motto:  profile.Motto,
		Figure: profile.Look,
		Sex:    profile.Sex,
	}, nil
}
