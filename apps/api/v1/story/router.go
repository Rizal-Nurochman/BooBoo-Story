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
	rg = rg.Group("/story")
	stories.StoriesRouter(rg, db)
	rarewords.RareWordsRouter(rg, db)
	bookmarks.BookmarkRouter(rg, db)
}

func CreatorRouter(rg *gin.RouterGroup, db *gorm.DB) {
	rg = rg.Group("/creators")
	creators.CreatorRouter(rg, db)
}

func ReaderRouter(rg *gin.RouterGroup, db *gorm.DB) {
	rg = rg.Group("/readers")
	progresses.ProgressReadRouter(rg, db)
}

func QuizRouter(rg *gin.RouterGroup, db *gorm.DB) {
	rg = rg.Group("/story")
	quizs.QuizsRouter(rg, db)
}