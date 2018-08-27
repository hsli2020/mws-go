package main

type FinanceApi struct {
	Client
}

/**
 * List Financial Event Groups
 *
 * ListFinancialEventGroups can be used to find financial event groups that meet filter criteria.
 *
 * @see ListFinancialEventGroupsRequest
 * @return ListFinancialEventGroupsResponse
 */
func (api FinanceApi) ListFinancialEventGroups() {
}

/**
 * List Financial Event Groups By Next Token
 *
 * If ListFinancialEventGroups returns a nextToken, thus indicating that there are more groups
 * than returned that matched the given filter criteria, ListFinancialEventGroupsByNextToken
 * can be used to retrieve those groups using that nextToken.
 *
 * @see ListFinancialEventGroupsByNextTokenRequest
 * @return ListFinancialEventGroupsByNextTokenResponse
 */
func (api FinanceApi) ListFinancialEventGroupsByNextToken() {
}

/**
 * List Financial Events
 *
 * ListFinancialEvents can be used to find financial events that meet the specified criteria.
 *
 * @see ListFinancialEventsRequest
 * @return ListFinancialEventsResponse
 */
func (api FinanceApi) ListFinancialEvents() {
}

/**
 * List Financial Events By Next Token
 *
 * If ListFinancialEvents returns a nextToken, thus indicating that there are more financial events
 * than returned that matched the given filter criteria, ListFinancialEventsByNextToken
 * can be used to retrieve those events using that nextToken.
 *
 * @see ListFinancialEventsByNextTokenRequest
 * @return ListFinancialEventsByNextTokenResponse
 */
func (api FinanceApi) ListFinancialEventsByNextToken() {
}

/**
 * Get Service Status
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api FinanceApi) GetServiceStatus() {
}
