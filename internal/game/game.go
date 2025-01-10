package game

import (
	"dragon_battle/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateName() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindind:"required"`
			Name     string `json:"name" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.UpdateName(request.PlayerID, request.Name)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Update name successfull!!"})
	}
}

func Crea_Clan() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindding:"required"`
			N_clan   string `json:"name_clan" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.Crea_Clan(request.PlayerID, request.N_clan)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Create Clan successfull!!"})
	}

}

func UpdateAvatar() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			PlayerID int    `json:"player_id" bindind:"required"`
			Avatar   string `json:"avatar" bindding:"required"`
		}
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.UpdateAvatar(request.PlayerID, request.Avatar)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		g.JSON(http.StatusOK, gin.H{"message": "Update avatar successfull!!"})
	}
}

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

func Crate_Dragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Name    string `json:"name_dragon" binding:"required"`
			Point   int    `json:"point" binding:"required"`
			Hp      int    `json:"hp" binding:"required"`
			Mp      int    `json:"mp" binding:"required"`
			Dame    int    `json:"dame" binding:"required"`
			Defense int    `json:"defense" binding:"required"`
			Speed   int    `json:"speed" binding:"required"`
			Rarity  string `json:"rarity" binding:"required"`  // Common,epic,legendary,immo
			Element string `json:"element" binding:"required"` // Fire,Natural,Terra,Ice,Metal,Dark,Light,Chaos
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Crea_Dragon(request.Name, request.Point, request.Hp, request.Mp, request.Dame, request.Defense, request.Speed, request.Rarity, request.Element)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Add Dragon successfully"})
	}
}

func Delete_Dragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		var request struct {
			Name string `json:"name_dr" binding:"required"`
		}

		// Kiểm tra dữ liệu đầu vào
		if err := g.ShouldBindJSON(&request); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gọi hàm DeleteFriend để xóa bạn bè
		err := db.Delete_Dragon(request.Name)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Add Dragon successfully"})
	}
}

func GetAllDragon() func(g *gin.Context) {
	return func(g *gin.Context) {
		// Gọi hàm DeleteFriend để xóa bạn bè
		dragons, err := db.GetAllDragon()
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Get Add Dragon", "details": err.Error()})
			return
		}

		// Trả về thông báo thành công
		g.JSON(http.StatusOK, gin.H{"message": "Get All Dragons successfully", "dragons": dragons})
	}
}
