package ratelimiter

import (
	"context"
	"database/sql"
	"net/http"
	"social-network/internal/models"
	utils "social-network/pkg"
	"strconv"
	"sync"
	"time"
)

type BucketToken struct {
	Tokens     int
	MaxTokens  int
	RefillTime time.Duration
	LastRefill time.Time
	Mu         sync.Mutex
}

func NewBucketToken(maxTokens int, refillTime time.Duration) *BucketToken {
	return &BucketToken{
		Tokens:     maxTokens,
		MaxTokens:  maxTokens,
		RefillTime: refillTime,
		LastRefill: time.Now(),
	}
}

func (bt *BucketToken) Allow() bool {
	now := time.Now()
	elapsed := now.Sub(bt.LastRefill)
	tokensToAdd := int(elapsed / bt.RefillTime)

	if tokensToAdd > 0 {
		bt.Tokens += tokensToAdd
		if bt.Tokens > bt.MaxTokens {
			bt.Tokens = bt.MaxTokens
		}
		bt.LastRefill = now
	}
	if bt.Tokens > 0 {
		bt.Tokens--
		return true
	}

	return false
}

type RateLimiter struct {
	Users map[string]*BucketToken
	Mu    sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		Users: make(map[string]*BucketToken),
	}
}

type ContextKey string

const UserIDKey ContextKey = "user"

func GetUserFromContext(ctx context.Context) (models.UserInfo, bool) {
	user, ok := ctx.Value(UserIDKey).(models.UserInfo)
	return user, ok
}

func (rl *RateLimiter) RateMiddleware(next http.Handler, maxTokens int, duration time.Duration, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := GetUserFromContext(r.Context()); if !ok {
			utils.WriteJson(w, http.StatusUnauthorized, "User not Found")
			return
		}

		rl.Mu.Lock()
		defer rl.Mu.Unlock()
		id := strconv.Itoa(user.ID)
		if _, ok := rl.Users[id]; !ok {
			rl.Users[id] = NewBucketToken(maxTokens, duration)
		}

		if !rl.Users[id].Allow() {
			http.Error(w, "Too many request", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *RateLimiter) RemoveSleepUsers() {
	for key, rateLimiter := range rl.Users {
		now := time.Now()
		elapsed := now.Sub(rateLimiter.LastRefill)

		if elapsed > (120 * time.Minute) {
			delete(rl.Users, key)
		}
	}
}

func (rl *RateLimiter) GetUserID(userUID string, db *sql.DB) (int, error) {
	var userID int
	err := db.QueryRow("SELECT id FROM user WHERE uid = ?", userUID).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

var (
	CreateArticleLimiter = NewRateLimiter()
)
