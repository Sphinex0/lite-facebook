package ratelimiter

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"sync"
	"time"

	"social-network/internal/models"
	utils "social-network/pkg"
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

func GetUserFromContext(ctx context.Context) (models.UserInfo, bool) {
	user, ok := ctx.Value(utils.UserIDKey).(models.UserInfo)
	return user, ok
}

func (rl *RateLimiter) RateMiddleware(next http.Handler, maxTokens int, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := strings.Split(r.RemoteAddr, ":")[0] + r.URL.Path
		rl.Mu.Lock()
		if _, ok := rl.Users[user]; !ok {
			rl.Users[user] = NewBucketToken(maxTokens, duration)
		}
		bucket := rl.Users[user]
		rl.Mu.Unlock()

		if !bucket.Allow() {
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

var CreateArticleLimiter = NewRateLimiter()
