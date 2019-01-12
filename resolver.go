package cms

import (
	"context"

	"github.com/deslee/cms/database/models"
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

func (r *assetResolver) FileName(ctx context.Context, obj *models.Asset) (string, error) {
	panic("not implemented")
}
func (r *assetResolver) Extension(ctx context.Context, obj *models.Asset) (string, error) {
	panic("not implemented")
}
func (r *assetResolver) Items(ctx context.Context, obj *models.Asset) ([]*models.Item, error) {
	panic("not implemented")
}

type groupResolver struct{ *Resolver }

func (r *groupResolver) Items(ctx context.Context, obj *models.Group) ([]*models.Item, error) {
	panic("not implemented")
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Groups(ctx context.Context, obj *models.Item) ([]*models.Group, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUserToSite(ctx context.Context, userId string, siteId string) (GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAsset(ctx context.Context, assetId string) (GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteItem(ctx context.Context, itemId string) (GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteSite(ctx context.Context, siteId string) (GenericResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, login LoginInput) (LoginResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) Register(ctx context.Context, registration RegisterInput) (UserResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, user UserInput) (UserResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertItem(ctx context.Context, item ItemInput, siteId string) (ItemResult, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpsertSite(ctx context.Context, site SiteInput) (SiteResult, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Asset(ctx context.Context, assetId string) (*models.Asset, error) {
	panic("not implemented")
}
func (r *queryResolver) Items(ctx context.Context, siteId string) ([]*models.Item, error) {
	panic("not implemented")
}
func (r *queryResolver) Item(ctx context.Context, itemId string) (*models.Item, error) {
	panic("not implemented")
}
func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Site(ctx context.Context, siteId string) (*models.Site, error) {
	panic("not implemented")
}
func (r *queryResolver) Sites(ctx context.Context) ([]*models.Site, error) {
	panic("not implemented")
}

type siteResolver struct{ *Resolver }

func (r *siteResolver) Assets(ctx context.Context, obj *models.Site) ([]*models.Asset, error) {
	panic("not implemented")
}
func (r *siteResolver) Groups(ctx context.Context, obj *models.Site) ([]*models.Group, error) {
	panic("not implemented")
}
func (r *siteResolver) Items(ctx context.Context, obj *models.Site) ([]*models.Item, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Sites(ctx context.Context, obj *models.User) ([]*models.Site, error) {
	panic("not implemented")
}
