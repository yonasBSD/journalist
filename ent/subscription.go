// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/user"
)

// Subscription is the model entity for the Subscription schema.
type Subscription struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// FeedID holds the value of the "feed_id" field.
	FeedID uuid.UUID `json:"feed_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscriptionQuery when eager-loading is set.
	Edges SubscriptionEdges `json:"edges"`
}

// SubscriptionEdges holds the relations/edges for other nodes in the graph.
type SubscriptionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Feed holds the value of the feed edge.
	Feed *Feed `json:"feed,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubscriptionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// FeedOrErr returns the Feed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubscriptionEdges) FeedOrErr() (*Feed, error) {
	if e.loadedTypes[1] {
		if e.Feed == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: feed.Label}
		}
		return e.Feed, nil
	}
	return nil, &NotLoadedError{edge: "feed"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscription) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscription.FieldName, subscription.FieldGroup:
			values[i] = new(sql.NullString)
		case subscription.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case subscription.FieldID, subscription.FieldUserID, subscription.FieldFeedID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subscription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscription fields.
func (s *Subscription) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscription.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subscription.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				s.UserID = *value
			}
		case subscription.FieldFeedID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field feed_id", values[i])
			} else if value != nil {
				s.FeedID = *value
			}
		case subscription.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subscription.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				s.Group = value.String
			}
		case subscription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Subscription entity.
func (s *Subscription) QueryUser() *UserQuery {
	return NewSubscriptionClient(s.config).QueryUser(s)
}

// QueryFeed queries the "feed" edge of the Subscription entity.
func (s *Subscription) QueryFeed() *FeedQuery {
	return NewSubscriptionClient(s.config).QueryFeed(s)
}

// Update returns a builder for updating this Subscription.
// Note that you need to call Subscription.Unwrap() before calling this method if this Subscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscription) Update() *SubscriptionUpdateOne {
	return NewSubscriptionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscription) Unwrap() *Subscription {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subscription is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscription) String() string {
	var builder strings.Builder
	builder.WriteString("Subscription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("feed_id=")
	builder.WriteString(fmt.Sprintf("%v", s.FeedID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(s.Group)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Subscriptions is a parsable slice of Subscription.
type Subscriptions []*Subscription
