package cmd

import (
	"MovieApp_Go/commonfunc"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list returns the list of all the movies",
	Long:  `list calls the list api indirectly and returns the list of all the titles of the movies`,
	Run: func(cmd *cobra.Command, args []string) {
		titlename, _ := cmd.Flags().GetString("title")
		searchtype, _ := cmd.Flags().GetString("type")
		baseurl := commonfunc.GenerateBaseURL()
		v := baseurl.Query()

		v.Add("s", titlename)
		v.Add("type", searchtype)
		baseurl.RawQuery = v.Encode()

		receivedBytes := commonfunc.SendAndReceiveRequest(baseurl)
		var result map[string]interface{}

		jsonErr := json.Unmarshal(receivedBytes, &result)
		commonfunc.Validate(jsonErr)

		var movielist = commonfunc.ProcessListResults(result)
		//fmt.Println(movielist)
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		defer w.Flush()
		for _, movie := range movielist {
			if movie.Title != "" {
				fmt.Fprintf(w, "Title \t%s\n", movie.Title)
				fmt.Fprintf(w, "Type \t%s\n", movie.Type)
				fmt.Fprintf(w, "Year \t%s\n", movie.Year)
				fmt.Fprintf(w, "\n")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringP("title", "s", "", "Movie Title to search for")
	listCmd.PersistentFlags().StringP("type", "t", "", "Type of search (either series or movies)")
}
