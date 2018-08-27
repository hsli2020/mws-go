package main

type RecommendationApi struct {
	Client
}

/**
 * Get Last Updated Time For Recommendations
 *
 * Retrieving last updated time for all recommendation categories for the
 * given marketplace and seller. If last updated time for a category is null,
 * it indicates no active recommendations for this seller in the given
 * marketplace for this category.
 *
 * @see GetLastUpdatedTimeForRecommendationsRequest
 * @return GetLastUpdatedTimeForRecommendationsResponse
 */
func (api RecommendationApi) GetLastUpdatedTimeForRecommendations() {
}

/**
 * List Recommendations
 *
 * Retrieving first batch of recommendations.
 *
 * @see ListRecommendationsRequest
 * @return ListRecommendationsResponse
 */
func (api RecommendationApi) ListRecommendations() {
}

/**
 * List Recommendations By Next Token
 *
 * Retrieving recommendation by next token.
 *
 * @see ListRecommendationsByNextTokenRequest
 * @return ListRecommendationsByNextTokenResponse
 */
func (api RecommendationApi) ListRecommendationsByNextToken() {
}

/**
 * Get Service Status
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api RecommendationApi) GetServiceStatus() {
}
