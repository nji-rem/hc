package profile

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"hc/internal/profile/application"
	"hc/internal/profile/infrastructure/store"
	"sync"
)

var Set = wire.NewSet(ProvideProfileStore, ProvideCreateProfile, ProvideRetrieveProfile, ProvideUpdateProfile)

var (
	profileStoreOnce sync.Once
	profileStore     *store.Profile

	createProfileOnce sync.Once
	createProfile     *application.CreateProfile

	retrieveProfileOnce sync.Once
	retrieveProfile     *application.InfoRetriever

	updateProfileOnce sync.Once
	updateProfile     *application.UpdateProfile
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

func ProvideRetrieveProfile(profileStore *store.Profile) *application.InfoRetriever {
	retrieveProfileOnce.Do(func() {
		retrieveProfile = &application.InfoRetriever{ProfileStore: profileStore}
	})

	return retrieveProfile
}

func ProvideUpdateProfile(profileStore *store.Profile) *application.UpdateProfile {
	updateProfileOnce.Do(func() {
		updateProfile = &application.UpdateProfile{Store: profileStore}
	})

	return updateProfile
}
