package validations

type CreatorCreate struct {
	Name         string `json:"name" binding:"required"`
	Bio          string `json:"bio" binding:"required"`
	ProfileImage string `json:"profile_image"`
	BannerImage  string `json:"banner_image"`
	Occupation   string `json:"occupation"`
	Location     string `json:"location"`
	WebsiteURL   string `json:"website_url"`
	IgUrl        string `json:"ig_url"`
	TiktokUrl    string `json:"tiktok_url"`
	YoutubeURL   string `json:"youtube_url"`
	LinkedInURL  string `json:"linkedin_url"`
	Portfolio    string `json:"portfolio"`
}


type CreatorUpdate struct {
	Name         string  `json:"name" binding:"required"`
	Bio          string  `json:"bio" binding:"required"`
	ProfileImage string  `json:"profile_image,omitempty"`
	BannerImage  string  `json:"banner_image,omitempty"`
	Occupation   string  `json:"occupation,omitempty"`
	Location     string  `json:"location,omitempty"`
	WebsiteURL   string  `json:"website_url,omitempty"`
	IgUrl        string  `json:"ig_url,omitempty"`
	TiktokUrl    string  `json:"tiktok_url,omitempty"`
	YoutubeURL   string  `json:"youtube_url,omitempty"`
	LinkedInURL  string  `json:"linkedin_url,omitempty"`
	Portfolio    string  `json:"portfolio,omitempty"`
}