package main

import (
	"fmt"
	"net/http"

	// gin: HTTP web framework
	// import は githubのurlを指定する（キモすぎるんだけど，慣れるものなんかな？）
	"github.com/gin-gonic/gin"
)

// gofmt コマンドで整形できる
type album struct {
	// json:"hoge" で serialize する時の json の key を指定できる
	// 指定しなかったら，フィールド名がそのまま json の key になる
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 出力フォーマット
func (a album) String() string {
	return fmt.Sprintf("album{ID: %s, Title: %s, Artist: %s, Price: %f}", a.ID, a.Title, a.Artist, a.Price)
}

// album struct 型 の slice
// Python でいう list が Go の slice ?
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

// endpoint: /albums
// Goは一文字変数が好きらしい（慣れないな...）
func getAlbums(c *gin.Context) {
	// docs:
	// IndentedJSON serializes the given struct as pretty JSON (indented + endlines)
	// into the response body.
	c.IndentedJSON(http.StatusOK, albums)
}

// endpoint: /albums/:id
func getAlbumByID(c *gin.Context) {
	// gin.Contextって便利だね．
	// docs:
	// Param returns the value of the URL param.
	// It is a shortcut for c.Params.ByName(key)
	id := c.Param("id")

	// for ... range で Python の enumerate になる
	// _i みたいにしても unused var って言われるのは気に食わない
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// docs:
	// H is a shortcut for map[string]any
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// endpoint: /albums
func postAlbums(c *gin.Context) {
	// varで宣言して初期化しない場合，デフォルト値が入る
	var newAlbum album
	{
		// デフォルト値の確認をしてみる
		// 普通に%vで出力すると，{ 0} が出力されてよく分からないので， %+v を使う
		// fmt.Printf("uninitialized newAlbum: %+v\n", newAlbum)

		// album 構造体に String() メソッドを実装しても良い
		fmt.Printf("uninitialized newAlbum: %v\n", newAlbum)

		// encoding/json パッケージの MarshalIndent 関数を使うこともできる
		// j, _ := json.MarshalIndent(newAlbum, "", "  ")
		// fmt.Printf("uninitialized newAlbum: %s\n", j)
	}

	// if x := hoge; {} とすると， xは{}のスコープ内に閉じられる
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// albums.append(newAlbum) って書きたくなる
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	// docs:
	// Default returns an Engine instance
	// with the Logger and Recovery middleware already attached.
	router := gin.Default()
	router.GET("/", healthCheck)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
