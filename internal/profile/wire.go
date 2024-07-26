package profile

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"hc/internal/profile/application"
	"hc/internal/profile/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(ProvideProfileStore, ProvideCreateProfile)

var (
	profileStoreOnce sync.Once
	profileStore     *store.Profile

	createProfileOnce sync.Once
	createProfile     *application.CreateProfile
)

func ProvideProfileStore(db *sqlx.DB) *store.Profile {
	profileStoreOnce.Do(func() {
		profileStore = &store.Profile{DB: db}
	})

	return profileStore
}

func ProvideCreateProfile(profileStore *store.Profile) *application.CreateProfile {
	createProfileOnce.Do(func() {
		createProfile = &application.CreateProfile{Store: profileStore}
	})

	return createProfile
}
