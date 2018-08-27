package main

type SellerApi struct {
	Client
}

func (api SellerApi) GetServiceStatus() {
}

/**
 * List Marketplace Participations
 *
 * Returns a list of marketplaces that the seller submitting the request can sell in,
 * and a list of participations that include seller-specific information in that marketplace.
 *
 * @see ListMarketplaceParticipationsRequest
 * @return ListMarketplaceParticipationsResponse
 */
func (api SellerApi) ListMarketplaceParticipations() {
}

/**
 * List Marketplace Participations By Next Token
 *
 * Returns the next page of marketplaces and participations using the NextToken value
 * that was returned by your previous request to either ListMarketplaceParticipations or
 * ListMarketplaceParticipationsByNextToken.
 *
 * @see ListMarketplaceParticipationsByNextTokenRequest
 * @return ListMarketplaceParticipationsByNextTokenResponse
 */
func (api SellerApi) ListMarketplaceParticipationsByNextToken() {
}
