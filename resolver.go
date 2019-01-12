package cms

import (
	"context"
	"database/sql"
	"time"

	"github.com/deslee/cms/data"
)

type Resolver struct{
	sql.DB
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
func (r *assetResolver) CreatedAt(ctx context.Context, obj *data.Asset) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}
func (r *assetResolver) LastUpdatedAt(ctx context.Context, obj *data.Asset) (string, error) {
	return obj.LastUpdatedAt.Format(time.RFC3339), nil
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Items(ctx context.Context, obj *data.Group) ([]data.Item, error) {
	return obj.Items(ctx, r.DB)
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Groups(ctx context.Context, obj *data.Item) ([]data.Group, error) {
	return obj.Groups(ctx, r.DB)
}
func (r *itemResolver) CreatedAt(ctx context.Context, obj *data.Item) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}
func (r *itemResolver) LastUpdatedAt(ctx context.Context, obj *data.Item) (string, error) {
	return obj.LastUpdatedAt.Format(time.RFC3339), nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUserToSite(ctx context.Context, userId string, siteId string) (data.GenericResult, error) {
	return data.AddUserToSite(ctx, r.DB, userId, siteId)
}
func (r *mutationResolver) DeleteAsset(ctx context.Context, assetId string) (data.GenericResult, error) {
	return data.DeleteAsset(ctx, r.DB, assetId)
}
func (r *mutationResolver) DeleteItem(ctx context.Context, itemId string) (data.GenericResult, error) {
	return data.DeleteItem(ctx, r.DB, itemId)
}
func (r *mutationResolver) DeleteSite(ctx context.Context, siteId string) (data.GenericResult, error) {
	return data.DeleteSite(ctx, r.DB, siteId)
}
func (r *mutationResolver) Login(ctx context.Context, login data.LoginInput) (data.LoginResult, error) {
	return data.Login(ctx, r.DB, login)
}
func (r *mutationResolver) Register(ctx context.Context, registration data.RegisterInput) (data.UserResult, error) {
	return data.Register(ctx, r.DB, registration)
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user data.UserInput) (data.UserResult, error) {
	return data.UpdateUser(ctx, r.DB, user)
}
func (r *mutationResolver) UpsertItem(ctx context.Context, item data.ItemInput, siteId string) (data.ItemResult, error) {
	return data.UpsertItem(ctx, r.DB, item, siteId)
}
func (r *mutationResolver) UpsertSite(ctx context.Context, site data.SiteInput) (data.SiteResult, error) {
	return data.UpsertSite(ctx, r.DB, site)
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
func (r *siteResolver) CreatedAt(ctx context.Context, obj *data.Site) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}
func (r *siteResolver) LastUpdatedAt(ctx context.Context, obj *data.Site) (string, error) {
	return obj.LastUpdatedAt.Format(time.RFC3339), nil
}

type userResolver struct{ *Resolver }

func (r *userResolver) Sites(ctx context.Context, obj *data.User) ([]data.Site, error) {
	return obj.Sites(ctx, r.DB)
}
