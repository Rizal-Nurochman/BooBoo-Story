package story

import (
	"github.com/BooBooStory/v1/creators"
	"github.com/BooBooStory/v1/story/bookmarks"
	rarewords "github.com/BooBooStory/v1/story/rareWords"
	"github.com/BooBooStory/v1/story/stories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func StoryRouter(rg *gin.RouterGroup, db *gorm.DB) {
	// all story router
	rg = rg.Group("/story")
	stories.StoriesRouter(rg, db)
	rarewords.RareWordsRouter(rg, db)
	bookmarks.BookmarkRouter(rg, db)

	// creator router
	rg = rg.Group("/creators")
	creators.CreatorRouter(rg, db)
}