package SocialServices

import (
	"errors"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/models/Dtos"
	"magical-crwler/services/Logger"
)

type SocialServices struct {
	repository database.IRepository
	logger     *Logger.Logger
}

func (s SocialServices) Add(bookmark Dtos.BookmarkDto) error {
	return s.repository.CreateBookmark(bookmark)
}
func (s SocialServices) Remove(adid, userid uint) error {
	return s.repository.DeleteBookmark(adid, userid)
}
func (s SocialServices) GetAllByUserId(userid uint) ([]Dtos.BookmarkToShowDto, error) {
	return s.repository.GetBookmarksByUserID(userid)
}
func (s SocialServices) GetPublicsByUserId(userid uint) ([]Dtos.BookmarkToShowDto, error) {
	return s.repository.GetPublicBookmarksByUserID(userid)
}

// ///////////////////////////////////////////////////
func (s SocialServices) GetAccessLevels() []models.AccessLevel {
	return s.repository.GetAllAccessLevels()
}
func (s SocialServices) SetAccess(access Dtos.AccessDto) error {
	return s.repository.AddAccess(access)
}
func (s SocialServices) SetBookmarkForAnotherUser(userid uint, bookmark Dtos.BookmarkDto) error {
	ac := s.repository.GetAccessByIds(bookmark.UserID, userid)
	if ac.AccessLevelID < 4 {
		return errors.New("Not Enough Access")
	}
	return s.repository.CreateBookmark(bookmark)
}
func (s SocialServices) GetAnotherUserBookmarks(userid, toBeWtchedUserId uint) ([]Dtos.BookmarkToShowDto, error) {
	ac := s.repository.GetAccessByIds(toBeWtchedUserId, userid)
	if ac.AccessLevelID == 1 {
		return []Dtos.BookmarkToShowDto{}, errors.New("Not Enough Access")
	} else if ac.AccessLevelID == 2 {
		return s.GetPublicsByUserId(toBeWtchedUserId)
	}
	return s.repository.GetBookmarksByUserID(toBeWtchedUserId)
}
