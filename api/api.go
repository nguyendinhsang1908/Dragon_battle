package api

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInfoPlayer_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			PlayerID int `json:"player_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pl, err := db.GetPOnePlayer(request.PlayerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "Player Information retrieved successfully", "Information": pl})

	}

}

func Message_API() func(c *gin.Context) {
	return func(c *gin.Context) {
		var request struct {
			ID1  int `json:"player_id_from" binding:"required"`
			ID2  int `json:"player_id_to" binding:"required"`
			Text int `json:"messanges" binding:"required"`
		}
		// Kiểm tra dữ liệu đầu vào
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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
