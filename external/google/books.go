package googleapi

type Books struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []Book `json:"items"`
}

type Book struct {
	Kind       string     `json:"kind"`
	ID         string     `json:"id"`
	Etag       string     `json:"etag"`
	SelfLink   string     `json:"selfLink"`
	VolumeInfo VolumeInfo `json:"volumeInfo"`
	SaleInfo   SaleInfo   `json:"saleInfo"`
	AccessInfo AccessInfo `json:"accessInfo"`
	SearchInfo SearchInfo `json:"searchInfo"`
}

type IndustryIdentifiers struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type ReadingModes struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type PanelizationSummary struct {
	ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
	ContainsImageBubbles bool `json:"containsImageBubbles"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type VolumeInfo struct {
	Title               string                `json:"title"`
	Authors             []string              `json:"authors"`
	Publisher           string                `json:"publisher"`
	PublishedDate       string                `json:"publishedDate"`
	Description         string                `json:"description"`
	IndustryIdentifiers []IndustryIdentifiers `json:"industryIdentifiers"`
	ReadingModes        ReadingModes          `json:"readingModes"`
	PageCount           int                   `json:"pageCount"`
	PrintType           string                `json:"printType"`
	Categories          []string              `json:"categories"`
	AverageRating       int                   `json:"averageRating"`
	RatingsCount        int                   `json:"ratingsCount"`
	MaturityRating      string                `json:"maturityRating"`
	AllowAnonLogging    bool                  `json:"allowAnonLogging"`
	ContentVersion      string                `json:"contentVersion"`
	PanelizationSummary PanelizationSummary   `json:"panelizationSummary"`
	ImageLinks          ImageLinks            `json:"imageLinks"`
	Language            string                `json:"language"`
	PreviewLink         string                `json:"previewLink"`
	InfoLink            string                `json:"infoLink"`
	CanonicalVolumeLink string                `json:"canonicalVolumeLink"`
}

type SaleInfo struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
}

type Epub struct {
	IsAvailable bool `json:"isAvailable"`
}

type Pdf struct {
	IsAvailable bool `json:"isAvailable"`
}

type AccessInfo struct {
	Country                string `json:"country"`
	Viewability            string `json:"viewability"`
	Embeddable             bool   `json:"embeddable"`
	PublicDomain           bool   `json:"publicDomain"`
	TextToSpeechPermission string `json:"textToSpeechPermission"`
	Epub                   Epub   `json:"epub"`
	Pdf                    Pdf    `json:"pdf"`
	WebReaderLink          string `json:"webReaderLink"`
	AccessViewStatus       string `json:"accessViewStatus"`
	QuoteSharingAllowed    bool   `json:"quoteSharingAllowed"`
}

type SearchInfo struct {
	TextSnippet string `json:"textSnippet"`
}
