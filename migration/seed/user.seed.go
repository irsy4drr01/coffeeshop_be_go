package seed

import (
	"log"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

func SeedUsers(db *sqlx.DB) {
	users := models.Users{
		{
			Username: "Guntur A",
			Email:    "guntura@gmail.com",
			Password: "$2a$10$n8hQ5G112r7qR6OJOl9Y.Obste7czv1CA.3rgfelFB6A2zyvngPO.",
			Role:     "user",
		},
		{
			Username: "Guntur",
			Email:    "guntur@gmail.com",
			Password: "$2a$10$LgLmnpo8YvVhzEh.KjZMTuVnDGZ9SVbusxGZSqBdcGSPS.B0rZ83S",
			Role:     "user",
			Image:    "https://res.cloudinary.com/dye6aug6j/image/upload/v1724188035/user-image-792.png",
		},
		{
			Username: "admin1",
			Email:    "admin1@gmail.com",
			Password: "$2a$10$i/o6V8CMsFfOA69od8SfVOFZxpIMomv6DxbbN9EZzsZtftyqcdE2S",
			Role:     "admin",
		},
		{
			Username: "user2",
			Email:    "user1@gmail.com",
			Password: "$2a$10$1rgEe.p59ssRcHmhqv..tel2nYyAx3XXLxPwRADNvydsMQueu6ybW",
			Role:     "user",
			Image:    "https://res.cloudinary.com/dye6aug6j/image/upload/v1726227724/user-image-670.png",
		},
		{
			Username: "admin2",
			Email:    "admin2@gmail.com",
			Password: "$2a$10$1vftDtwWWqOmEgwKvD8bDeiu11pVAWI1LvLyzUmJ2hPS6eZh4dX7i",
			Role:     "admin",
		},
		{
			Username: "admin3",
			Email:    "admin3@gmail.com",
			Password: "$2a$10$v0eYfGyjKSqZ529gYkluOeRm8s9EsoOg4XKlBKX.Eiy/vF5BL46.e",
			Role:     "admin",
		},
		{
			Username: "user2",
			Email:    "user2@gmail.com",
			Password: "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
			Role:     "user",
			Image:    "https://res.cloudinary.com/dye6aug6j/image/upload/v1727347823/user-image-885.png",
		},
		{
			Username: "user3",
			Email:    "user3@gmail.com",
			Password: "$2a$10$uZiHZo6aNe8KR7QG7843WeGRdDchqsolCelgJg3fzUnFVXlCgdney",
			Role:     "user",
		},
		{
			Username: "user4",
			Email:    "user4@gmail.com",
			Password: "$2a$10$jEFowdFXsXYH9aBapTnmoeFnIGtRqsRAf8DGRDFrxqtZdsLoNgd.2",
			Role:     "user",
		},
		{
			Username: "user5",
			Email:    "user5@gmail.com",
			Password: "$2a$10$Xwdt2b52vtktTDv3wRoqIuIToF1blkSCgQdQlBFPTfxP0.n0aurHq",
			Role:     "user",
		},
		{
			Username: "user6",
			Email:    "user6@gmail.com",
			Password: "$2a$10$eweqwGfA5JJo2G8Pg6OUlOHqTDlbg1ToVNtVcMjl058l9gdD8O1cK",
			Role:     "user",
		},
	}

	query := `INSERT INTO public.users (username, email, password, role, image) 
		VALUES (:username, :email, :password, :role, :image)`

	for _, user := range users {
		_, err := db.NamedExec(query, user)
		if err != nil {
			log.Printf("Failed to seed user: %s. Error: %v", user.Username, err)
		} else {
			log.Printf("Successfully seeded Users: %s", user.Username)
		}
	}
}
