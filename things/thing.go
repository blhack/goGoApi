package things

type Attribute struct {
	AttributeName		string `json:"attributeName"`
	AttributeValue		string `json:"attributeValue"`
}

type Thing struct {
	Id           int    	 `json:"id"`
	Title        string 	 `json:"title"`
	Url          string 	 `json:"url"`
	Text         string 	 `json:"text"`
	Username     string 	 `json:"username"`
	CreationDate string 	 `json:"creationDate"`
	Uuid         string 	 `json:"uuid"`
	FileUuid     string 	 `json:"fileUuid"`
	FilePath     string 	 `json:"filePath"`
	Attributes	 []Attribute `json:"attributes"`
}
