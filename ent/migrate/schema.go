// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "url", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "feed_title", Type: field.TypeString},
		{Name: "feed_description", Type: field.TypeString},
		{Name: "feed_link", Type: field.TypeString},
		{Name: "feed_feed_link", Type: field.TypeString},
		{Name: "feed_updated", Type: field.TypeString},
		{Name: "feed_published", Type: field.TypeString},
		{Name: "feed_author", Type: field.TypeString},
		{Name: "feed_authors", Type: field.TypeString},
		{Name: "feed_language", Type: field.TypeString},
		{Name: "feed_image", Type: field.TypeString},
		{Name: "feed_copyright", Type: field.TypeString},
		{Name: "feed_generator", Type: field.TypeString},
		{Name: "feed_categories", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:       "feeds",
		Columns:    FeedsColumns,
		PrimaryKey: []*schema.Column{FeedsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "feed_url_username_password",
				Unique:  true,
				Columns: []*schema.Column{FeedsColumns[1], FeedsColumns[2], FeedsColumns[3]},
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "item_title", Type: field.TypeString},
		{Name: "item_description", Type: field.TypeString},
		{Name: "item_content", Type: field.TypeString},
		{Name: "item_link", Type: field.TypeString},
		{Name: "item_updated", Type: field.TypeString},
		{Name: "item_published", Type: field.TypeString},
		{Name: "item_author", Type: field.TypeString},
		{Name: "item_authors", Type: field.TypeString},
		{Name: "item_guid", Type: field.TypeString},
		{Name: "item_image", Type: field.TypeString},
		{Name: "item_categories", Type: field.TypeString},
		{Name: "item_enclosures", Type: field.TypeString},
		{Name: "crawler_title", Type: field.TypeString},
		{Name: "crawler_author", Type: field.TypeString},
		{Name: "crawler_excerpt", Type: field.TypeString},
		{Name: "crawler_site_name", Type: field.TypeString},
		{Name: "crawler_image", Type: field.TypeString},
		{Name: "crawler_content_html", Type: field.TypeString},
		{Name: "crawler_content_text", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "feed_items", Type: field.TypeUUID, Nullable: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_feeds_items",
				Columns:    []*schema.Column{ItemsColumns[22]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReadsColumns holds the columns for the "reads" table.
	ReadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "item_id", Type: field.TypeUUID},
	}
	// ReadsTable holds the schema information for the "reads" table.
	ReadsTable = &schema.Table{
		Name:       "reads",
		Columns:    ReadsColumns,
		PrimaryKey: []*schema.Column{ReadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reads_users_user",
				Columns:    []*schema.Column{ReadsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "reads_items_item",
				Columns:    []*schema.Column{ReadsColumns[3]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "read_user_id_item_id",
				Unique:  true,
				Columns: []*schema.Column{ReadsColumns[2], ReadsColumns[3]},
			},
		},
	}
	// SubscriptionsColumns holds the columns for the "subscriptions" table.
	SubscriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "group", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "feed_id", Type: field.TypeUUID},
	}
	// SubscriptionsTable holds the schema information for the "subscriptions" table.
	SubscriptionsTable = &schema.Table{
		Name:       "subscriptions",
		Columns:    SubscriptionsColumns,
		PrimaryKey: []*schema.Column{SubscriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subscriptions_users_user",
				Columns:    []*schema.Column{SubscriptionsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "subscriptions_feeds_feed",
				Columns:    []*schema.Column{SubscriptionsColumns[5]},
				RefColumns: []*schema.Column{FeedsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subscription_user_id_feed_id",
				Unique:  true,
				Columns: []*schema.Column{SubscriptionsColumns[4], SubscriptionsColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FeedsTable,
		ItemsTable,
		ReadsTable,
		SubscriptionsTable,
		UsersTable,
	}
)

func init() {
	ItemsTable.ForeignKeys[0].RefTable = FeedsTable
	ReadsTable.ForeignKeys[0].RefTable = UsersTable
	ReadsTable.ForeignKeys[1].RefTable = ItemsTable
	SubscriptionsTable.ForeignKeys[0].RefTable = UsersTable
	SubscriptionsTable.ForeignKeys[1].RefTable = FeedsTable
}
