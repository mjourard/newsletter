package articlemanager

type ArchivesScannedArchivedArticle struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Img            string `json:"img"`
	Abstract       string `json:"abstract"`
	Author         string `json:"author"`
	AddedTimeStamp int    `json:"addedts"`
}
