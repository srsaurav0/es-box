package controllers

import (
	"encoding/json"
	"es-box/dao"
	"io"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

// SearchController handles Elasticsearch queries
type SearchController struct {
	web.Controller
}

// @router / [GET]
func (c *SearchController) Get() {
	// Renders the search page
	c.TplName = "index.tpl"
}

// @router /search [POST]
func (c *SearchController) SearchResults() {
	// Read the raw request body
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": http.StatusText(http.StatusInternalServerError), "message": "Error reading request body"}
		c.ServeJSON()
		return
	}

	// Declare a map to hold the parsed JSON data
	var requestBody map[string]interface{}

	// Unmarshal the body into the requestBody map
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": http.StatusText(http.StatusBadRequest), "message": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	// Retrieve the 'query' value from the parsed JSON
	searchQuery, ok := requestBody["query"].(string)
	if !ok || searchQuery == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": http.StatusText(http.StatusBadRequest), "message": "query parameter is required"}
		c.ServeJSON()
		return
	}

	// Construct the Elasticsearch query using `match_phrase_prefix`
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []interface{}{
					// Exact match with highest boost
					map[string]interface{}{
						"term": map[string]interface{}{
							"products.product_name.keyword": map[string]interface{}{
								"value": searchQuery,
								"boost": 3,
							},
						},
					},
					// Match with analyzer for handling hyphens
					map[string]interface{}{
						"match": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"query":    searchQuery,
								"analyzer": "standard",
								"boost":    2,
							},
						},
					},
					// Fuzzy matching with higher edit distance and lower boost
					map[string]interface{}{
						"match": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"query":                searchQuery,
								"fuzziness":            2,  // Allows up to 2 character changes
								"prefix_length":        1,  // First character must match
								"max_expansions":       50, // Allow more variations
								"boost":                1.5,
								"fuzzy_transpositions": true, // Allow character swaps (e.g., "jakcet" -> "jacket")
							},
						},
					},
					// Handle hyphen variations
					map[string]interface{}{
						"match": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"query": strings.ReplaceAll(strings.ReplaceAll(searchQuery, "-", ""), " ", ""), // Remove hyphens and spaces
								"boost": 1.5,
							},
						},
					},
					// Prefix matching for autocomplete
					map[string]interface{}{
						"prefix": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"value": searchQuery,
								"boost": 1,
							},
						},
					},
				},
				"minimum_should_match": 1,
			},
		},
		"size": 20,
	}

	// Execute the search query
	esClient := dao.Client // Ensure your Elasticsearch client is initialized properly
	result, err := esClient.ExecuteSearch(query)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": http.StatusText(http.StatusInternalServerError), "message": "Error querying Elasticsearch"}
		c.ServeJSON()
		return
	}

	// Parse and send back results as JSON
	suggestions := parseSearchResults(result, searchQuery)
	c.Data["json"] = suggestions
	c.ServeJSON()
}

// Utility function to handle parsing of search results
func parseSearchResults(result map[string]interface{}, searchQuery string) []map[string]interface{} {
	var suggestions []map[string]interface{}

	hitsInterface, ok := result["hits"].(map[string]interface{})
	if !ok {
		return suggestions
	}

	hitsArray, ok := hitsInterface["hits"].([]interface{})
	if !ok {
		return suggestions
	}

	for _, hit := range hitsArray {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		products, ok := source["products"].([]interface{})
		if !ok {
			continue
		}

		for _, product := range products {
			if productMap, ok := product.(map[string]interface{}); ok {
				productName, ok1 := productMap["product_name"].(string)
				price, ok2 := productMap["price"].(float64)

				if ok1 && ok2 && strings.Contains(strings.ToLower(productName), strings.ToLower(searchQuery)) {
					suggestions = append(suggestions, map[string]interface{}{
						"name":  productName,
						"price": price,
					})
				}
			}
		}
	}

	if len(suggestions) > 20 {
		return suggestions[:20]
	}
	return suggestions
}
