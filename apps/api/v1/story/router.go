package story

import (
	"github.com/BooBooStory/v1/creators"
	"github.com/BooBooStory/v1/readers/progresses"
	"github.com/BooBooStory/v1/story/bookmarks"
	"github.com/BooBooStory/v1/story/quizs"
	rarewords "github.com/BooBooStory/v1/story/rareWords"
	"github.com/BooBooStory/v1/story/stories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func StoryRouter(rg *gin.RouterGroup, db *gorm.DB) {
	// all story router
	rg_story := rg.Group("/story")
	stories.StoriesRouter(rg_story, db)
	rarewords.RareWordsRouter(rg_story, db)
	bookmarks.BookmarkRouter(rg_story, db)
	quizs.QuizsRouter(rg_story, db)

	// creator router
	rg_creators := rg.Group("/creators")
	creators.CreatorRouter(rg_creators, db)

	// reader router
	rg_readers := rg.Group("/readers")
	progresses.ProgressReadRouter(rg_readers, db)
}