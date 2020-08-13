// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type Kategori struct {
	ID           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

type Membership struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type MembershipDetail struct {
	MembershipID string `json:"membershipId"`
	UserID       string `json:"userId"`
	Date         string `json:"date"`
	Bill         int    `json:"bill"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewCategory struct {
	CategoryName string `json:"categoryName"`
}

type NewComment struct {
	VideoID string `json:"videoId"`
	Comment string `json:"comment"`
}

type NewMembership struct {
	Type string `json:"type"`
}

type NewMembershipDetail struct {
	MembershipID *string `json:"membershipId"`
	Bill         int     `json:"bill"`
}

type NewPlaylist struct {
	Title   string `json:"title"`
	Privacy bool   `json:"privacy"`
}

type NewPlaylistDetail struct {
	PlaylistID string `json:"playlistId"`
	VideoID    string `json:"videoID"`
}

type NewPlaylistSub struct {
	PlaylistID string `json:"playlistID"`
	UserID     string `json:"userID"`
}

type NewReply struct {
	CommentID string `json:"commentId"`
	UserID    string `json:"userId"`
	Reply     string `json:"reply"`
}

type NewSubscription struct {
	UserID       string `json:"userID"`
	SubscriberID string `json:"subscriberID"`
}

type NewUser struct {
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   *string `json:"password"`
	ProfilePic string  `json:"profilePic"`
}

type NewVideo struct {
	Link        string  `json:"link"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumbnail   string  `json:"thumbnail"`
	Category    int     `json:"category"`
	Label       bool    `json:"label"`
	Privacy     bool    `json:"privacy"`
	Location    string  `json:"location"`
	DatePublish *string `json:"datePublish"`
}

type UpdatedPlaylist struct {
	Title   *string `json:"title"`
	Privacy *bool   `json:"privacy"`
	Sort    *string `json:"sort"`
}
