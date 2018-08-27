package main

type ProductApi struct {
	Client
}

/**
 * Get Competitive Pricing For ASIN
 *
 * Gets competitive pricing and related information for a product identified by
 * the MarketplaceId and ASIN.
 *
 * @see GetCompetitivePricingForASINRequest
 * @return GetCompetitivePricingForASINResponse
 */
func (api ProductApi) GetCompetitivePricingForASIN() {
}

/**
 * Get Competitive Pricing For SKU
 *
 * Gets competitive pricing and related information for a product identified by
 * the SellerId and SKU.
 *
 * @see GetCompetitivePricingForSKURequest
 * @return GetCompetitivePricingForSKUResponse
 */
func (api ProductApi) GetCompetitivePricingForSKU() {
}

/**
 * Get Lowest Offer Listings For ASIN
 *
 * Gets some of the lowest prices based on the product identified by the given
 * MarketplaceId and ASIN.
 *
 * @see GetLowestOfferListingsForASINRequest
 * @return GetLowestOfferListingsForASINResponse
 */
func (api ProductApi) GetLowestOfferListingsForASIN() {
}

/**
 * Get Lowest Offer Listings For SKU
 *
 * Gets some of the lowest prices based on the product identified by the given
 * SellerId and SKU.
 *
 * @see GetLowestOfferListingsForSKURequest
 * @return GetLowestOfferListingsForSKUResponse
 */
func (api ProductApi) GetLowestOfferListingsForSKU() {
}

/**
 * Get Lowest Priced Offers For ASIN
 *
 * Retrieves the lowest priced offers based on the product identified by the given
 * ASIN.
 *
 * @see GetLowestPricedOffersForASINRequest
 * @return GetLowestPricedOffersForASINResponse
 */
func (api ProductApi) GetLowestPricedOffersForASIN() {
}

/**
 * Get Lowest Priced Offers For SKU
 *
 * Retrieves the lowest priced offers based on the product identified by the given
 * SellerId and SKU.
 *
 * @see GetLowestPricedOffersForSKURequest
 * @return GetLowestPricedOffersForSKUResponse
 */
func (api ProductApi) GetLowestPricedOffersForSKU() {
}

/**
 * Get Matching Product
 *
 * GetMatchingProduct will return the details (attributes) for the
 * given ASIN.
 *
 * @see GetMatchingProductRequest
 * @return GetMatchingProductResponse
 */
func (api ProductApi) GetMatchingProduct() {
}

/**
 * Get Matching Product For Id
 *
 * GetMatchingProduct will return the details (attributes) for the
 * given Identifier list. Identifer type can be one of [SKU|ASIN|UPC|EAN|ISBN|GTIN|JAN]
 *
 * @see GetMatchingProductForIdRequest
 * @return GetMatchingProductForIdResponse
 */
func (api ProductApi) GetMatchingProductForId() {
}

/**
 * Get My Fees Estimate
 *
 * Retrieves the fees estimate for the products identified by the given
 * ASIN/SKU list.
 *
 * @see GetMyFeesEstimateRequest
 * @return GetMyFeesEstimateResponse
 */
func (api ProductApi) GetMyFeesEstimate() {
}

/**
 * Get My Price For ASIN
 *
 * @see GetMyPriceForASINRequest
 * @return GetMyPriceForASINResponse
 */
func (api ProductApi) GetMyPriceForASIN() {
}

/**
 * Get My Price For SKU
 *
 * @see GetMyPriceForSKURequest
 * @return GetMyPriceForSKUResponse
 */
func (api ProductApi) GetMyPriceForSKU() {
}

/**
 * Get Product Categories For ASIN
 *
 * Gets categories information for a product identified by
 * the MarketplaceId and ASIN.
 *
 * @see GetProductCategoriesForASINRequest
 * @return GetProductCategoriesForASINResponse
 */
func (api ProductApi) GetProductCategoriesForASIN() {
}

/**
 * Get Product Categories For SKU
 *
 * Gets categories information for a product identified by
 * the SellerId and SKU.
 *
 * @see GetProductCategoriesForSKURequest
 * @return GetProductCategoriesForSKUResponse
 */
func (api ProductApi) GetProductCategoriesForSKU() {
}

/**
 * Get Service Status
 *
 * Returns the service status of a particular MWS API section. The operation
 * takes no input.
 *
 * All API sections within the API are required to implement this operation.
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api ProductApi) GetServiceStatus() {
}

/**
 * List Matching Products
 *
 * ListMatchingProducts can be used to find products that match the given criteria.
 *
 * @see ListMatchingProductsRequest
 * @return ListMatchingProductsResponse
 */
func (api ProductApi) ListMatchingProducts() {
}
