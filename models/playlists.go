package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DataCriacao     primitive.DateTime `bson:"data_criacao" json:"dataCriacao,omitempty"`
	Sincronizado    string             `bson:"sincronizado" json:"sincronizado,omitempty"`
	Youtube         *YoutubeInfo       `bson:"youtube" json:"youtube,omitempty"`
	DataAtualizacao primitive.DateTime `bson:"data_atualizacao" json:"dataAtualizacao,omitempty"`
	SlugUrl         string             `bson:"slug_url" json:"slugUrl,omitempty"`
	Videos          []Video            `bson:"videos" json:"videos,omitempty"`
}

type YoutubeInfo struct {
	ID           string             `bson:"id" json:"id,omitempty"`
	DataCriacao  primitive.DateTime `bson:"data_criacao" json:"dataCriacao,omitempty"`
	IdCanal      string             `bson:"id_canal" json:"idCanal,omitempty"`
	TituloCanal  string             `bson:"titulo_canal" json:"tituloCanal,omitempty"`
	Titulo       string             `bson:"titulo" json:"titulo,omitempty"`
	Descricao    string             `bson:"descricao" json:"descricao,omitempty"`
	Status       string             `bson:"status" json:"status,omitempty"`
	NumeroVideos uint16             `bson:"numero_videos" json:"numeroVideos,omitempty"`
	Imagens      *Imagem            `bson:"imagens" json:"imagens,omitempty"`
}

type Imagem struct {
	Padrao   *ImagemInfo `bson:"padrao" json:"padrao,omitempty"`
	Media    *ImagemInfo `bson:"media" json:"media,omitempty"`
	Alta     *ImagemInfo `bson:"alta" json:"alta,omitempty"`
	Superior *ImagemInfo `bson:"superior" json:"superior,omitempty"`
	Maxima   *ImagemInfo `bson:"maxima" json:"maxima,omitempty"`
}

type ImagemInfo struct {
	Url     string `bson:"url" json:"url,omitempty"`
	Largura uint16 `bson:"largura" json:"largura,omitempty"`
	Altura  uint16 `bson:"altura" json:"altura,omitempty"`
}

type Video struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IdVideo        string             `bson:"id_video" json:"idVideo,omitempty"`
	Posicao        uint16             `bson:"posicao" json:"posicao,omitempty"`
	DataAdicionado primitive.DateTime `bson:"data_adicionado" json:"dataAdicionado,omitempty"`
}
