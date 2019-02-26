package cmd

import (
	"commonfunc"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search returns the single search result",
	Long:  `search indirectly calls the search api and returns the single search result`,
	Run: func(cmd *cobra.Command, args []string) {
		titlename, _ := cmd.Flags().GetString("title")
		searchtype, _ := cmd.Flags().GetString("type")
		baseurl := commonfunc.GenerateBaseURL()
		v := baseurl.Query()
		v.Add("t", titlename)
		v.Add("type", searchtype)
		baseurl.RawQuery = v.Encode()

		// fmt.Println(baseurl.String())
		receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
		var result map[string]interface{}

		jsonErr := json.Unmarshal(receivedBytes, &result)
		commonfunc.Validate(jsonErr)

		var moviesearch = commonfunc.ProcessSearchResult(result)
		//fmt.Println(moviesearch)

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)

		defer w.Flush()

		fmt.Fprintf(w, "Title \t%s\n", moviesearch.Title)
		fmt.Fprintf(w, "Type \t%s\n", moviesearch.Movietype)
		fmt.Fprintf(w, "Imdb \t%s\n", moviesearch.Rating.Imdb)
		fmt.Fprintf(w, "Rotten Tomatoes \t%s\n", moviesearch.Rating.Rottentamotoes)
		fmt.Fprintf(w, "Metascore \t%s\n", moviesearch.Rating.Metacritic)
		fmt.Fprintf(w, "Year released \t%s\n", moviesearch.Year)
		fmt.Fprintf(w, "Genre \t%s\n", moviesearch.Genre)
		fmt.Fprintf(w, "Description \t%s\n", moviesearch.Plot)
		fmt.Fprintf(w, "Runtime \t%s\n", moviesearch.Runtime)
		fmt.Fprintf(w, "Actors \t%s\n", moviesearch.Actors)
		fmt.Fprintf(w, "Country \t%s\n", moviesearch.Country)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringP("title", "s", "", "Movie Title to search for")
	searchCmd.PersistentFlags().StringP("type", "t", "", "Type of search (either series or movies)")
}
