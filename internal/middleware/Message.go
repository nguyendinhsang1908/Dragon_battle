package middleware

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Post Message Friend Api
func PostMessage_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			ID1     int    `json:"player_id_from" binding:"required"`
			ID2     int    `json:"player_id_to" binding:"required"`
			Clan_ID int    `json:"clan_id,omitempty"`
			Team_id int    `json:"team_id,omitempty"`
			Text    string `json:"messanges" binding:"required"`
			Type    string `json:"type" binding:"required"`
		}
		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.PostMessage(request.ID1, request.ID2, request.Clan_ID, request.Team_id, request.Text, request.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Post success!"})
		}

	}
}

// Get Messages API
func GetMessagesAPI() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			ID1     int    `json:"player_id_from" binding:"required"`
			ID2     int    `json:"player_id_to" binding:"required"`
			Clan_ID int    `json:"clan_id,omitempty"`
			Team_id int    `json:"team_id,omitempty"`
			Type    string `json:"type" binding:"required"`
		}
		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		Messages, err := db.GetMessages(request.ID1, request.ID2, request.Clan_ID, request.Team_id, request.Type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"erro": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "getMessages retrieved successfully", "Mesanges": Messages})

	}
}

func AddFriend_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Nhận dữ liệu JSON từ yêu cầu
		var friendRequest struct {
			PlayerID int `json:"player_id" binding:"required"`
			FriendID int `json:"friend_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&friendRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm AddFriend để thêm bạn bè
		err := db.AddFriend(friendRequest.PlayerID, friendRequest.FriendID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Friend added successfully"})
	}

}

func GetFriend_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Nhận dữ liệu JSON từ yêu cầu
		var friendRequest struct {
			PlayerID int `json:"player_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&friendRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm GetFriends để lấy danh sách bạn bè
		friends, err := db.GetFriends(friendRequest.PlayerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get friends", "details": err.Error()})
			return
		}

		// Trả về danh sách bạn bè nếu không có lỗi
		c.JSON(http.StatusOK, gin.H{"message": "Friends retrieved successfully", "friends": friends})
	}
}

func DeleteFriendAPI() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Nhận dữ liệu JSON từ yêu cầu
		var request struct {
			PlayerID int `json:"player_id" binding:"required"`
			FriendID int `json:"friend_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.DeleteFriend(request.PlayerID, request.FriendID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete friend", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		c.JSON(http.StatusOK, gin.H{"message": "Friend deleted successfully"})
	}
}

func Send_Fr_Mess_Request() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			SenderID   int `json:"sender_id" binding:"required"`
			ReceiverID int `json:"receiver_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Send_Fr_Mess_Request(request.SenderID, request.ReceiverID)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send friend request", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Send friend request successfully"})
	}
}

func Accept_Fr_Request() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			RequestID int `json:"request_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Accept_Fr_Request(request.RequestID)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept friend", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Accept friend successfully"})
	}
}

func Reject_Fr_Request() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			RequestID int `json:"request_id" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Reject_Fr_Request(request.RequestID)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject friend", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Reject friend successfully"})
	}
}
