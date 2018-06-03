package relation

import (
	"regexp"

	dbFactory "github.com/anovafawzi/socialmedia/db"
	models "github.com/anovafawzi/socialmedia/models"
	"github.com/gin-gonic/gin"
)

// Service : service for relation domain
type Service struct {
}

// CreateFriendConnection : create friend connection
func (service *Service) CreateFriendConnection(c *gin.Context) (models.Result, string) {
	db := dbFactory.InitDb()
	defer db.Close()

	// get parameter input
	var connections models.Connections
	c.Bind(&connections)

	// count email
	ec := len(connections.Friends)

	if ec > 1 {
		// check if friend already connected
		var relations []models.Relations
		db.Where("(email1 = ? AND email2 = ?) OR (email1 = ? AND email2 = ?)", connections.Friends[0], connections.Friends[1], connections.Friends[1], connections.Friends[0]).Find(&relations)
		if len(relations) > 0 {
			for _, v := range relations {
				// check if already friends or not
				if v.Friend {
					return models.Result{Success: false}, "they are already friends"
				}
				// check if blocked
				if v.Block {
					return models.Result{Success: false}, "one of them have blocked the other"
				}
			}

			// if not error, update the relations into friends (if exists)
			for _, v := range relations {
				// update each data
				v.Friend = true
				db.Save(&v)
			}
		} else {
			// add connection here if no relationship exists before
			// friend added will default have reversed data
			addfriends1 := models.Relations{Email1: connections.Friends[0], Email2: connections.Friends[1], Friend: true, Subscribe: false, Block: false}
			addfriends2 := models.Relations{Email1: connections.Friends[1], Email2: connections.Friends[0], Friend: true, Subscribe: false, Block: false}
			db.Create(&addfriends1)
			db.Create(&addfriends2)
		}

		return models.Result{Success: true}, ""
	} else {
		return models.Result{Success: false}, "emails are empty or numbers of email not sufficient to make a friend connections"
	}

}

// GetFriendList : get friend list based on email address
func (service *Service) GetFriendList(c *gin.Context) (models.Result, string) {
	// Connection to the database
	db := dbFactory.InitDb()
	// Close connection database
	defer db.Close()

	// get parameter value
	var user models.Users
	c.Bind(&user)

	// get friend list
	var relations []models.Relations
	db.Where("email1 = ?", user.Email).Find(&relations)
	if len(relations) > 0 {
		// convert to array to meet requirements
		emailOnly := convertRelationsResultToArray(relations)

		// return result
		return models.Result{Success: true, Friends: emailOnly, Count: len(relations)}, ""
	} else {
		// return result
		return models.Result{Success: false}, "this user dont have friends"
	}
}

// GetFriendCommonList : get common friends list between two email addresses
func (service *Service) GetFriendCommonList(c *gin.Context) (models.Result, string) {
	db := dbFactory.InitDb()
	defer db.Close()

	// get parameter input
	var connections models.Connections
	c.Bind(&connections)

	// count email
	ec := len(connections.Friends)

	if ec > 1 {
		// check the same friend
		var relations []models.Relations
		db.Where("email1 = ? OR email1 = ?", connections.Friends[0], connections.Friends[1]).Group("email2").Having("count(*) > 1").Find(&relations)
		if len(relations) > 0 {
			// convert to array to meet requirements
			emailOnly := convertRelationsResultToArray(relations)

			// return result
			return models.Result{Success: true, Friends: emailOnly, Count: len(relations)}, ""
		} else {
			return models.Result{Success: false}, "they don't have the same friends"
		}

	} else {
		return models.Result{Success: false}, "you must input two email addresses"
	}
}

// SubscribeFriend : subscribe to a friend
func (service *Service) SubscribeFriend(c *gin.Context) (models.Result, string) {
	db := dbFactory.InitDb()
	defer db.Close()

	// get parameter input
	var ru models.RelatedUsers
	c.Bind(&ru)

	var relation models.Relations
	db.Where("email1 = ? AND email2 = ? AND friend = ?", ru.Requestor, ru.Target, true).First(&relation)

	if relation.ID != 0 {
		// update subscribe status
		relation.Subscribe = true
		db.Save(&relation)

		// return result
		return models.Result{Success: true}, ""
	} else {
		// return error
		return models.Result{Success: false}, "these user are not friends"
	}
}

// BlockFriend : block updates from an email address
func (service *Service) BlockFriend(c *gin.Context) (models.Result, string) {
	db := dbFactory.InitDb()
	defer db.Close()

	// get parameter input
	var ru models.RelatedUsers
	c.Bind(&ru)

	var relation models.Relations
	db.Where("email1 = ? AND email2 = ?", ru.Requestor, ru.Target).First(&relation)

	if relation.ID != 0 {
		// update current relation with block status
		relation.Block = true
		db.Save(&relation)

		// return result
		return models.Result{Success: true}, ""
	} else {
		// create new relation but not as friend
		// friend added will default have reversed data
		addrel1 := models.Relations{Email1: ru.Requestor, Email2: ru.Target, Friend: false, Subscribe: false, Block: true}
		addrel2 := models.Relations{Email1: ru.Target, Email2: ru.Requestor, Friend: false, Subscribe: false, Block: false}
		db.Create(&addrel1)
		db.Create(&addrel2)

		// return result
		return models.Result{Success: true}, ""
	}

}

// MentionFriend : retrieve all email addresses that can receive updates from an email address
func (service *Service) MentionFriend(c *gin.Context) (models.Result, string) {
	db := dbFactory.InitDb()
	defer db.Close()

	// get parameter input
	var mention models.Mention
	c.Bind(&mention)

	if mention.Sender == "" {
		// return error
		return models.Result{Success: false}, "must input email"
	}

	if mention.Text == "" {
		// return error
		return models.Result{Success: false}, "text cannot empty"
	}

	// regex expressions to find email
	r := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	f := r.FindAllString(mention.Text, -1)
	var emails []string

	// check if user exists
	var relations []models.Relations
	db.Where("email2 = ?", mention.Sender).Find(&relations)
	if len(relations) == 0 {
		// return error
		return models.Result{Success: false}, "no user found"
	}

	for _, v := range relations {
		if (v.Friend || v.Subscribe) && !v.Block {
			emails = append(emails, v.Email1)
		}
	}

	// get from mentioned
	for _, v := range f {
		emails = append(emails, v)
	}

	// return result
	return models.Result{Success: true, Recipients: emails}, ""
}

func convertRelationsResultToArray(relations []models.Relations) []string {
	var s []string
	for _, v := range relations {
		s = append(s, v.Email2)
	}
	return s
}
