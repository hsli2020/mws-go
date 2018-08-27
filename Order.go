package main

type OrderApi struct {
	Client
}

/**
 * Get Order
 *
 * This operation takes up to 50 order ids and returns the corresponding orders.
 *
 * @see GetOrderRequest
 * @return GetOrderResponse
 */
func (api OrderApi) GetOrder() {
}

/**
 * Get Service Status
 *
 * Returns the service status of a particular MWS API section. The operation
 * takes no input.
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api OrderApi) GetServiceStatus() {
}

/**
 * List Order Items
 *
 * This operation can be used to list the items of the order indicated by the
 * given order id (only a single Amazon order id is allowed).
 *
 * @see ListOrderItemsRequest
 * @return ListOrderItemsResponse
 */
func (api OrderApi) ListOrderItems() {
}

/**
 * List Order Items By Next Token
 *
 * If ListOrderItems cannot return all the order items in one go, it will
 * provide a nextToken. That nextToken can be used with this operation to
 * retrive the next batch of items for that order.
 *
 * @see ListOrderItemsByNextTokenRequest
 * @return ListOrderItemsByNextTokenResponse
 */
func (api OrderApi) ListOrderItemsByNextToken() {
}

/**
 * List Orders
 *
 * ListOrders can be used to find orders that meet the specified criteria.
 *
 * @see ListOrdersRequest
 * @return ListOrdersResponse
 */
func (api OrderApi) ListOrders() {
}

/**
 * List Orders By Next Token
 *
 * If ListOrders returns a nextToken, thus indicating that there are more orders
 * than returned that matched the given filter criteria, ListOrdersByNextToken
 * can be used to retrieve those other orders using that nextToken.
 *
 * @see ListOrdersByNextTokenRequest
 * @return ListOrdersByNextTokenResponse
 */
func (api OrderApi) ListOrdersByNextToken() {
}
