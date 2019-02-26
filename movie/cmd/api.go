package cmd

import (
	"MovieApp_Go/commonfunc"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts up the movie api",
	Long: `Initiates and starts up the movie api with the following endpoints:
	/search/title 
	/search/movie/title, 
	/search/series/title, 
	/list/title`,
	Run: func(cmd *cobra.Command, args []string) {
		hostportStr, _ := cmd.Flags().GetString("hostport")
		router := gin.Default()

		router.GET("/search/:title", search)
		router.GET("/search/:title/movies", searchMovies)
		router.GET("/search/:title/series", searchSeries)
		router.GET("/list/:searchitem", list)
		router.Run(":" + hostportStr)
	},
}

func searchMovies(c *gin.Context) {
	titlename := c.Param("title")
	baseurl := commonfunc.GenerateBaseURL()
	v := baseurl.Query()
	v.Add("type", "movie")
	v.Add("t", titlename)

	baseurl.RawQuery = v.Encode()
	receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
	var result map[string]interface{}

	jsonErr := json.Unmarshal(receivedBytes, &result)
	commonfunc.Validate(jsonErr)

	var moviedb = commonfunc.ProcessSearchResult(result)
	c.JSON(http.StatusOK, moviedb)
}

func searchSeries(c *gin.Context) {
	titlename := c.Param("title")
	baseurl := commonfunc.GenerateBaseURL()
	v := baseurl.Query()
	v.Add("type", "series")
	v.Add("t", titlename)

	baseurl.RawQuery = v.Encode()
	receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
	var result map[string]interface{}

	jsonErr := json.Unmarshal(receivedBytes, &result)
	commonfunc.Validate(jsonErr)

	var moviedb = commonfunc.ProcessSearchResult(result)
	c.JSON(http.StatusOK, moviedb)
}

func search(c *gin.Context) {

	titlename := c.Param("title")

	baseurl := commonfunc.GenerateBaseURL()
	v := baseurl.Query()
	v.Add("t", titlename)

	baseurl.RawQuery = v.Encode()
	receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
	var result map[string]interface{}

	jsonErr := json.Unmarshal(receivedBytes, &result)
	commonfunc.Validate(jsonErr)

	var moviesearch = commonfunc.ProcessSearchResult(result)
	c.JSON(http.StatusOK, moviesearch)
}

func list(c *gin.Context) {
	searchitem := c.Param("searchitem")

	baseurl := commonfunc.GenerateBaseURL()
	v := baseurl.Query()

	v.Add("s", searchitem)
	baseurl.RawQuery = v.Encode()

	receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
	var result map[string]interface{}

	jsonErr := json.Unmarshal(receivedBytes, &result)
	commonfunc.Validate(jsonErr)

	var movielist = commonfunc.ProcessListResults(result)
	c.JSON(http.StatusOK, movielist)
}

func init() {
	rootCmd.AddCommand(apiCmd)

	apiCmd.PersistentFlags().StringP("hostport", "p", "3000", "Port to run the Api application on")
}
