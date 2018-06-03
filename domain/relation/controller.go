package relation

import (
	"github.com/gin-gonic/gin"
)

// Controller : controller for relation domain
type Controller struct {
}

// PostFriendConnection : As a user, I need an API to create a friend connection between two email addresses
func (ctrl *Controller) PostFriendConnection(c *gin.Context) {
	var relationService Service
	result, msg := relationService.CreateFriendConnection(c)

	if result.Success {
		c.JSON(201, gin.H{"success": true})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}

// PostFriendList : As a user, I need an API to retrieve the friends list for an email address
func (ctrl *Controller) PostFriendList(c *gin.Context) {
	var relationService Service
	result, msg := relationService.GetFriendList(c)

	if result.Success {
		c.JSON(200, gin.H{"success": true, "friends": result.Friends, "count": result.Count})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}

// PostFriendCommonList : As a user, I need an API to retrieve the common friends list between two email addresses
func (ctrl *Controller) PostFriendCommonList(c *gin.Context) {
	var relationService Service
	result, msg := relationService.GetFriendCommonList(c)

	if result.Success {
		c.JSON(200, gin.H{"success": true, "friends": result.Friends, "count": result.Count})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}

// PostFriendSubscribe : As a user, I need an API to subscrive to updates from an email address
func (ctrl *Controller) PostFriendSubscribe(c *gin.Context) {
	var relationService Service
	result, msg := relationService.SubscribeFriend(c)

	if result.Success {
		c.JSON(200, gin.H{"success": true})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}

// PostFriendBlock : As a user, I need an API to block updates from an email address
func (ctrl *Controller) PostFriendBlock(c *gin.Context) {
	var relationService Service
	result, msg := relationService.BlockFriend(c)

	if result.Success {
		c.JSON(200, gin.H{"success": true})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}

// PostFriendUpdates : As a user, I need an API to retrieve all email addresses that can receive updates from an email address
func (ctrl *Controller) PostFriendUpdates(c *gin.Context) {
	var relationService Service
	result, msg := relationService.MentionFriend(c)

	if result.Success {
		c.JSON(200, gin.H{"success": true, "recipients": result.Recipients})
	} else {
		c.JSON(422, gin.H{"error": msg})
	}
}
