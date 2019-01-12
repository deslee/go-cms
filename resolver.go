package cms

import (
	"context"
	"database/sql"
	"github.com/deslee/cms/data"
	"log"
	"runtime/debug"
)

type Resolver struct {
	*sql.DB
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

func (r *assetResolver) Items(ctx context.Context, obj *data.Asset) ([]data.Item, error) {
	return obj.Items(ctx, r.DB)
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Items(ctx context.Context, obj *data.Group) ([]data.Item, error) {
	return obj.Items(ctx, r.DB)
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Groups(ctx context.Context, obj *data.Item) ([]data.Group, error) {
	return obj.Groups(ctx, r.DB)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUserToSite(ctx context.Context, userId string, siteId string) (res data.GenericResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.GenericError()
		}
	}()
	res, err = data.AddUserToSite(ctx, r.DB, userId, siteId)
	return res, err
}
func (r *mutationResolver) DeleteAsset(ctx context.Context, assetId string) (res data.GenericResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.GenericError()
		}
	}()
	res, err = data.DeleteAsset(ctx, r.DB, assetId)
	return res, err
}
func (r *mutationResolver) DeleteItem(ctx context.Context, itemId string) (res data.GenericResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.GenericError()
		}
	}()
	res, err = data.DeleteItem(ctx, r.DB, itemId)
	return res, err
}
func (r *mutationResolver) DeleteSite(ctx context.Context, siteId string) (res data.GenericResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.GenericError()
		}
	}()
	res, err = data.DeleteSite(ctx, r.DB, siteId)
	return res, err
}
func (r *mutationResolver) Login(ctx context.Context, login data.LoginInput) (res data.LoginResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.LoginResult{GenericResult: data.GenericError()}
		}
	}()
	res, err = data.Login(ctx, r.DB, login)
	return res, err
}
func (r *mutationResolver) Register(ctx context.Context, registration data.RegisterInput) (res data.UserResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.UserResult{GenericResult: data.GenericError()}
		}
	}()
	res, err = data.Register(ctx, r.DB, registration)
	return res, err
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user data.UserInput) (res data.UserResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.UserResult{GenericResult: data.GenericError()}
		}
	}()
	res, err = data.UpdateUser(ctx, r.DB, user)
	return res, err
}
func (r *mutationResolver) UpsertItem(ctx context.Context, item data.ItemInput, siteId string) (res data.ItemResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.ItemResult{GenericResult: data.GenericError()}
		}
	}()
	res, err = data.UpsertItem(ctx, r.DB, item, siteId)
	return res, err
}
func (r *mutationResolver) UpsertSite(ctx context.Context, site data.SiteInput) (res data.SiteResult, err error) {
	defer func(){
		if r := recover(); r != nil {
			log.Printf("%s %s", r, debug.Stack())
			res = data.SiteResult{GenericResult: data.GenericError()}
		}
	}()
	res, err = data.UpsertSite(ctx, r.DB, site)
	return res, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Asset(ctx context.Context, assetId string) (*data.Asset, error) {
	return data.GetAsset(ctx, r.DB, assetId)
}
func (r *queryResolver) Items(ctx context.Context, siteId string) ([]data.Item, error) {
	return data.GetItems(ctx, r.DB, siteId)
}
func (r *queryResolver) Item(ctx context.Context, itemId string) (*data.Item, error) {
	return data.GetItem(ctx, r.DB, itemId)
}
func (r *queryResolver) Me(ctx context.Context) (*data.User, error) {
	return data.UserFromContext(ctx, r.DB)
}
func (r *queryResolver) Site(ctx context.Context, siteId string) (*data.Site, error) {
	return data.GetSite(ctx, r.DB, siteId)
}
func (r *queryResolver) Sites(ctx context.Context) ([]data.Site, error) {
	return data.GetSites(ctx, r.DB)
}

type siteResolver struct{ *Resolver }

func (r *siteResolver) Assets(ctx context.Context, obj *data.Site) ([]data.Asset, error) {
	return obj.Assets(ctx, r.DB)
}
func (r *siteResolver) Groups(ctx context.Context, obj *data.Site) ([]data.Group, error) {
	return obj.Groups(ctx, r.DB)
}
func (r *siteResolver) Items(ctx context.Context, obj *data.Site) ([]data.Item, error) {
	return obj.Items(ctx, r.DB)
}

type userResolver struct{ *Resolver }

func (r *userResolver) Sites(ctx context.Context, obj *data.User) ([]data.Site, error) {
	return obj.Sites(ctx, r.DB)
}
