package articlemanager

type ArchivesArchivedArticle struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Img            string `json:"img"`
	Abstract       string `json:"abstract"`
	Author         string `json:"author"`
	AddedTimeStamp int64  `json:"addedts"`
}
