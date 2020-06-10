package models

type Link struct {
	ShortUrl string `db:"short_url" json:"shortUrl" binding:"required"`
	LongUrl  string `db:"long_url" json:"longUrl" binding:"required"`
}

type DeleteLink struct {
	ShortUrl string `db:"short_url"`
}

// type CreateShortener struct {
// 	ShortUrl string `json:"shortUrl" binding:"required"`
// 	LongUrl  string `json:"longUrl" binding:"required"`
// }
