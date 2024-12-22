package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"gihub.com/kengkeng852/SocialWebsiteGo/internal/store"
	"github.com/go-chi/chi/v5"
)

const userContextKey contextKey = "user"

func(app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w,http.StatusOK, user); err != nil {
		app.internalServerError(w,r,err)
	}
}

type FollowUser struct {
	UserID int64 `json:"user_id"`
}

func(app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowerUser := getUserFromContext(r)

	//Todo: revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w,r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return 
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, unfollowerUser.ID, payload.UserID); err != nil {
		switch err{
			
		}
	}

	if err := app.jsonResponse(w,http.StatusNoContent, nil); err != nil {
		app.internalServerError(w,r,err)
	}
}

func(app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowedUser := getUserFromContext(r)

	//Todo: revert back to auth userID from ctx
	var payload FollowUser
	if err := readJSON(w,r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return 
	}

	ctx := r.Context()

	if err := app.store.Followers.UnFollow(ctx, unfollowedUser.ID, payload.UserID); err != nil {
		app.internalServerError(w,r,err)
		return
	}

	if err := app.jsonResponse(w,http.StatusNoContent, nil); err != nil {
		app.internalServerError(w,r,err)
	}
}

func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10,64)
	if err != nil {
		app.badRequestResponse(w,r,err)
		return
	}

	ctx := r.Context()

	user, err := app.store.Users.GetUserByID(ctx, userID)
	if err != nil {
		switch{
		case errors.Is(err, store.ErrNotFound):
			app.badRequestResponse(w,r,err)
			return
		default:
			app.internalServerError(w,r,err)
			return
		}
		
	}

	ctx = context.WithValue(ctx, userContextKey, user)
	next.ServeHTTP(w,r.WithContext(ctx))
	})
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userContextKey).(*store.User)
	return user
}