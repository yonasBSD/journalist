package feeds

import (
	"context"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/mrusme/journalist/ent"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/rss"

	"go.uber.org/zap"
)

type FeedUpdateResponse struct {
	Success bool           `json:"success"`
	Feed    *FeedUpdateModel `json:"feed"`
	Message string         `json:"message"`
}

// Update godoc
// @Summary      Update a feed
// @Description  Update feed by ID
// @Tags         feeds
// @Accept       json
// @Produce      json
// @Param        id   path      string true "Feed ID"
// @Success      200  {object}  FeedUpdateResponse
// @Failure      400  {object}  FeedUpdateResponse
// @Failure      404  {object}  FeedUpdateResponse
// @Failure      500  {object}  FeedUpdateResponse
// @Router       /feeds/{id}/update [get]
// @security     BasicAuth
func (h *handler) Update(ctx *fiber.Ctx) error {
	var err error

	param_id := ctx.Params("id")
	id, err := uuid.Parse(param_id)
	if err != nil {
		h.logger.Debug(
			"Could not parse user ID",
			zap.Error(err),
		)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(FeedUpdateResponse{
				Success: false,
				Feed:    nil,
				Message: err.Error(),
			})
	}

	dbFeed, err := h.entClient.Feed.
		Query().
		Where(
			feed.ID(id),
		).
		Only(context.Background())
	if err != nil {
		h.logger.Debug(
			"Could not query feed",
			zap.String("feedID", param_id),
			zap.Error(err),
		)
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(FeedUpdateResponse{
				Success: false,
				Feed:    nil,
				Message: err.Error(),
			})
	}

	showFeed := FeedUpdateModel{
		ID:    dbFeed.ID.String(),
		Name:  dbFeed.FeedTitle,
		URL:   dbFeed.FeedFeedLink,
		Group: "*",
	}

	h.logger.Debug(
		"Refresh feed starting, refreshing all feeds",
	)

	var feedIds []uuid.UUID = make([]uuid.UUID, 1)
	feedIds[0] = dbFeed.ID

	errs := h.Refresh(feedIds)

	if len(errs) > 0 {
		h.logger.Error(
			"Refresh feed completed with errors",
			zap.Errors("errors", errs),
		)
	} else {
		h.logger.Debug(
			"Refresh feed completed",
		)
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(FeedUpdateResponse{
			Success: true,
			Feed:    &showFeed,
			Message: "",
		})
}

func (h *handler) Refresh(feedIds []uuid.UUID) []error {
	var errs []error

	dbFeeds, err := h.entClient.Feed.
		Query().
		Where(
			feed.IDIn(feedIds...),
		).
		WithItems(func(query *ent.ItemQuery) {
			query.
				Select(item.FieldItemGUID).
				Where(item.CrawlerContentHTMLNEQ(""))
		}).
		All(context.Background())
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	for _, dbFeed := range dbFeeds {
		var exceptItemGUIDs []string
		for _, exceptItem := range dbFeed.Edges.Items {
			exceptItemGUIDs = append(exceptItemGUIDs, exceptItem.ItemGUID)
		}

		rc, errr := rss.NewClient(
			dbFeed.URL,
			dbFeed.Username,
			dbFeed.Password,
			true,
			exceptItemGUIDs,
			h.logger,
		)
		if len(errr) > 0 {
			errs = append(errs, errr...)
			continue
		}

		dbFeedTmp := h.entClient.Feed.
			Create()
		rc.SetFeed(
			dbFeed.URL,
			dbFeed.Username,
			dbFeed.Password,
			dbFeedTmp,
		)
		feedID, err := dbFeedTmp.
			OnConflictColumns("url", "username", "password").
			UpdateNewValues().
			ID(context.Background())
		if err != nil {
			errs = append(errs, err)
		}

		dbItems := make([]*ent.ItemCreate, len(rc.Feed.Items))
		for i := 0; i < len(rc.Feed.Items); i++ {
			dbItems[i] = h.entClient.Item.
				Create()
			dbItems[i] = rc.SetItem(
				feedID,
				i,
				dbItems[i],
			)
		}
		err = h.entClient.Item.
			CreateBulk(dbItems...).
			OnConflictColumns("item_guid").
			UpdateNewValues().
			Exec(context.Background())
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
