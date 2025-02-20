package notification

import (
	"context"
	"fmt"
	"strings"

	"github.com/helixml/helix/api/pkg/auth"
	"github.com/helixml/helix/api/pkg/config"
	"github.com/helixml/helix/api/pkg/types"
	"github.com/rs/zerolog/log"
)

type Provider string

const (
	ProviderEmail Provider = "email"
)

type Event int

const (
	EventFinetuningStarted  Event = 1
	EventFinetuningComplete Event = 2
)

func (e Event) String() string {
	switch e {
	case EventFinetuningStarted:
		return "finetuning_started"
	case EventFinetuningComplete:
		return "finetuning_complete"
	default:
		return "unknown_event"
	}
}

type Notification struct {
	Event   Event
	Session *types.Session

	// Populated by the provider
	Email     string
	FirstName string
}

type Notifier interface {
	Notify(ctx context.Context, n *Notification) error
}

type NotificationsProvider struct {
	user_retriever auth.UserRetriever

	email *Email
}

func New(cfg *config.Notifications, user_retriever auth.UserRetriever) (Notifier, error) {
	email, err := NewEmail(cfg)
	if err != nil {
		return nil, err
	}

	return &NotificationsProvider{
		user_retriever: user_retriever,
		email:          email,
	}, nil
}

func (n *NotificationsProvider) Notify(ctx context.Context, notification *Notification) error {
	user, err := n.user_retriever.GetUserByID(ctx, notification.Session.Owner)
	if err != nil {
		return fmt.Errorf("failed to get user '%s' details: %w", notification.Session.Owner, err)
	}

	log.Debug().
		Str("email", user.Email).Str("notification", notification.Event.String()).Msg("sending notification")

	notification.Email = user.Email
	notification.FirstName = strings.Split(user.FullName, " ")[0]

	if n.email.Enabled() {
		err := n.email.Notify(ctx, notification)
		if err != nil {
			return err
		}
	}

	return nil
}
