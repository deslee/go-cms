package cms

import (
	"context"

	"github.com/deslee/cms/data"
)

type Resolver struct{}

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
	panic("not implemented")
}
func (r *assetResolver) CreatedAt(ctx context.Context, obj *data.Asset) (string, error) {
	panic("not implemented")
}
func (r *assetResolver) LastUpdatedAt(ctx context.Context, obj *data.Asset) (string, error) {
	panic("not implemented")
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Items(ctx context.Context, obj *data.Group) ([]data.Item, error) {
	panic("not implemented")
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Groups(ctx context.Context, obj *data.Item) ([]data.Group, error) {
	panic("not implemented")
}
func (r *itemResolver) CreatedAt(ctx context.Context, obj *data.Item) (string, error) {
	panic("not implemented")
}
func (r *itemResolver) LastUpdatedAt(ctx context.Context, obj *data.Item) (string, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUserToSite(ctx context.Context, userId string, siteId string) (data.GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAsset(ctx context.Context, assetId string) (data.GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteItem(ctx context.Context, itemId string) (data.GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteSite(ctx context.Context, siteId string) (data.GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, login data.LoginInput) (data.LoginResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) Register(ctx context.Context, registration data.RegisterInput) (data.UserResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user data.UserInput) (data.UserResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertItem(ctx context.Context, item data.ItemInput, siteId string) (data.ItemResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertSite(ctx context.Context, site data.SiteInput) (data.SiteResult, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Asset(ctx context.Context, assetId string) (*data.Asset, error) {
	panic("not implemented")
}
func (r *queryResolver) Items(ctx context.Context, siteId string) ([]data.Item, error) {
	panic("not implemented")
}
func (r *queryResolver) Item(ctx context.Context, itemId string) (*data.Item, error) {
	panic("not implemented")
}
func (r *queryResolver) Me(ctx context.Context) (*data.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Site(ctx context.Context, siteId string) (*data.Site, error) {
	panic("not implemented")
}
func (r *queryResolver) Sites(ctx context.Context) ([]data.Site, error) {
	panic("not implemented")
}

type siteResolver struct{ *Resolver }

func (r *siteResolver) Assets(ctx context.Context, obj *data.Site) ([]data.Asset, error) {
	panic("not implemented")
}
func (r *siteResolver) Groups(ctx context.Context, obj *data.Site) ([]data.Group, error) {
	panic("not implemented")
}
func (r *siteResolver) Items(ctx context.Context, obj *data.Site) ([]data.Item, error) {
	panic("not implemented")
}
func (r *siteResolver) CreatedAt(ctx context.Context, obj *data.Site) (string, error) {
	panic("not implemented")
}
func (r *siteResolver) LastUpdatedAt(ctx context.Context, obj *data.Site) (string, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Sites(ctx context.Context, obj *data.User) ([]data.Site, error) {
	panic("not implemented")
}
