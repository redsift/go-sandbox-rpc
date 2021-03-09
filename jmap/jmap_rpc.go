//
//  jmap_rpc.go
//  JMAP model
//
//  Created by Randal Pinto on 31/03/2015.
//  Copyright (c) 2015 RedSift Limited. All rights reserved.
//
// Package jmap implements an imap to jmap encoder library
//

package jmap

import (
	"fmt"
	"strings"
)

type SourceModel int64

const (
	SourceUnknown SourceModel = iota
	SourcePull
	SourcePush
)

type Meta struct {
	RequestID string      `json:"requestId,omitempty"`
	ToB       int64       `json:"tob,omitempty"`    // time of birth (used to measure latency)
	ToR       int64       `json:"tor,omitempty"`    // time of request (used to measure latency)
	Source    SourceModel `json:"source,omitempty"` // message source
}

// Message is the structure of a JMAP message.
type Message struct {
	ID                 string              `json:"id,omitempty"`
	ThreadID           string              `json:"threadId,omitempty"`
	MailboxIDs         []string            `json:"mailboxIds,omitempty"`
	InReplyToMessageID string              `json:"inReplyToMessageId,omitempty"` // Can be null
	IsUnread           *bool               `json:"isUnread,omitempty"`           // Mutable
	IsFlagged          *bool               `json:"isFlagged,omitempty"`          // Mutable
	IsAnswered         *bool               `json:"isAnswered,omitempty"`         // Mutable
	IsDraft            *bool               `json:"isDraft,omitempty"`            // Mutable by the server
	HasAttachment      *bool               `json:"hasAttachment,omitempty"`
	Headers            map[string]string   `json:"headers,omitempty"`
	From               *Emailer            `json:"from,omitempty"`    // Can be null
	To                 []*Emailer          `json:"to,omitempty"`      // Can be null
	Cc                 []*Emailer          `json:"cc,omitempty"`      // Can be null
	Bcc                []*Emailer          `json:"bcc,omitempty"`     // Can be null
	ReplyTo            *Emailer            `json:"replyTo,omitempty"` // Can be null
	Subject            string              `json:"subject,omitempty"`
	Date               string              `json:"date,omitempty"`
	Size               uint32              `json:"size"`
	Preview            string              `json:"preview,omitempty"`
	TextBody           string              `json:"textBody,omitempty"`         // Can be null
	HTMLBody           string              `json:"htmlBody,omitempty"`         // Can be null
	StrippedHTMLBody   string              `json:"strippedHtmlBody,omitempty"` // Can be null
	Attachments        []*Attachment       `json:"attachments,omitempty"`      // Can be null
	AttachedMessages   map[string]*Message `json:"attachedMessages,omitempty"` // Can be null
	User               string              `json:"user,omitempty"`
	ExtSecReport       *ExtSecReport       `json:"extsecrep,omitempty"` // Can be null
	ToB                int64               `json:"tob,omitempty"`       // time of birth (used to measure latency)
	Meta               Meta                `json:"meta,omitempty"`
}

func (msg *Message) String() string {
	return fmt.Sprintf("%s", map[string]interface{}{
		"id": msg.ID, "subject": msg.Subject, "from": msg.From,
		"to": msg.To, "cc": msg.Cc, "date": msg.Date, "textBody": msg.TextBody, "htmlBody": msg.HTMLBody,
	})
}

func (msg *Message) InMailbox(mailbox string) bool {
	if mailbox == "" || strings.ToLower(mailbox) == "all" {
		return true
	}

	for _, mb := range msg.MailboxIDs {
		if strings.ToLower(mb) == strings.ToLower(mailbox) {
			return true
		}
	}

	return false
}

// Emailer
type Emailer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (e *Emailer) String() string {
	return fmt.Sprintf(`%q <%s>`, e.Name, e.Email)
}

// Attachment
type Attachment struct {
	BlobID   string `json:"blobId"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Size     uint32 `json:"size"`
	IsInline bool   `json:"isInline"`
	Width    uint32 `json:"width,omitempty"`  // Optional
	Height   uint32 `json:"height,omitempty"` // Optional
	Content  []byte `json:"-"`
}

func (a *Attachment) String() string {
	return fmt.Sprintf("[%s] %s", a.BlobID, a.Name)
}
