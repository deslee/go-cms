package cms

import (
	"context"
	. "github.com/deslee/cms/data"
	. "github.com/deslee/cms/model"
	"github.com/jmoiron/sqlx"
	"log"
	"runtime/debug"
)

type Resolver struct {
	*sqlx.DB
}

func (r *Resolver) Asset() AssetResolver {
	return &assetResolver{r}
}
func (r *Resolver) Group() GroupResolver {
	return &groupResolver{r}
}
func (r *Resolver) Item() ItemResolver {
	return &itemResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Site() SiteResolver {
	return &siteResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type assetResolver struct{ *Resolver }

func (r *assetResolver) Items(ctx context.Context, obj *Asset) ([]Item, error) {
	return ItemsFromAsset(ctx, r.DB, *obj)
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Items(ctx context.Context, obj *Group) ([]Item, error) {
	return ItemsFromGroup(ctx, r.DB, *obj)
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Groups(ctx context.Context, obj *Item) ([]Group, error) {
	return GroupsFromItem(ctx, r.DB, *obj)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUserToSite(ctx context.Context, userId string, siteId string) (res GenericResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = GenericError()
		}
	}()
	res, err = AddUserToSite(ctx, r.DB, userId, siteId)
	return res, err
}
func (r *mutationResolver) DeleteAsset(ctx context.Context, assetId string) (res GenericResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = GenericError()
		}
	}()
	res, err = DeleteAsset(ctx, r.DB, assetId)
	return res, err
}
func (r *mutationResolver) DeleteItem(ctx context.Context, itemId string) (res GenericResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = GenericError()
		}
	}()
	res, err = DeleteItem(ctx, r.DB, itemId)
	return res, err
}
func (r *mutationResolver) DeleteSite(ctx context.Context, siteId string) (res GenericResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = GenericError()
		}
	}()
	res, err = DeleteSite(ctx, r.DB, siteId)
	return res, err
}
func (r *mutationResolver) Login(ctx context.Context, login LoginInput) (res LoginResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = LoginResult{GenericResult: GenericError()}
		}
	}()
	res, err = Login(ctx, r.DB, login)
	return res, err
}
func (r *mutationResolver) Register(ctx context.Context, registration RegisterInput) (res UserResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = UserResult{GenericResult: GenericError()}
		}
	}()
	res, err = Register(ctx, r.DB, registration)
	return res, err
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user UserInput) (res UserResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = UserResult{GenericResult: GenericError()}
		}
	}()
	res, err = UpdateUser(ctx, r.DB, user)
	return res, err
}
func (r *mutationResolver) UpsertItem(ctx context.Context, item ItemInput, siteId string) (res ItemResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = ItemResult{GenericResult: GenericError()}
		}
	}()
	res, err = UpsertItem(ctx, r.DB, item, siteId)
	return res, err
}
func (r *mutationResolver) UpsertSite(ctx context.Context, site SiteInput) (res SiteResult, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = SiteResult{GenericResult: GenericError()}
		}
	}()
	res, err = UpsertSite(ctx, r.DB, site)
	return res, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Asset(ctx context.Context, assetId string) (*Asset, error) {
	return GetAsset(ctx, r.DB, assetId)
}
func (r *queryResolver) Items(ctx context.Context, siteId string) ([]Item, error) {
	return GetItems(ctx, r.DB, siteId)
}
func (r *queryResolver) Item(ctx context.Context, itemId string) (*Item, error) {
	return GetItem(ctx, r.DB, itemId)
}
func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	return UserFromContext(ctx, r.DB)
}
func (r *queryResolver) Site(ctx context.Context, siteId string) (*Site, error) {
	return GetSite(ctx, r.DB, siteId)
}
func (r *queryResolver) Sites(ctx context.Context) ([]Site, error) {
	return GetSites(ctx, r.DB)
}

type siteResolver struct{ *Resolver }

func (r *siteResolver) Assets(ctx context.Context, obj *Site) ([]Asset, error) {
	return AssetsFromSite(ctx, r.DB, *obj)
}
func (r *siteResolver) Groups(ctx context.Context, obj *Site) ([]Group, error) {
	return GroupsFromSite(ctx, r.DB, *obj)
}
func (r *siteResolver) Items(ctx context.Context, obj *Site) ([]Item, error) {
	return ItemsFromSite(ctx, r.DB, *obj)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Sites(ctx context.Context, obj *User) ([]Site, error) {
	return SitesFromUser(ctx, r.DB, *obj)
}
