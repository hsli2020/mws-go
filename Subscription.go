package mws

type Subscription struct {
	Client
}

/**
 * Create Subscription
 * Create a new subscription.
 *
 * @param mixed $request array of parameters for CreateSubscription request or CreateSubscription object itself
 * @see CreateSubscriptionInput
 * @return CreateSubscriptionResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) CreateSubscription() {
}

/**
 * Delete Subscription
 * Delete a subscription.
 *
 * @param mixed $request array of parameters for DeleteSubscription request or DeleteSubscription object itself
 * @see DeleteSubscriptionInput
 * @return DeleteSubscriptionResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) DeleteSubscription() {
}

/**
 * Deregister Destination
 * Delete a destination.
 *
 * @param mixed $request array of parameters for DeregisterDestination request or DeregisterDestination object itself
 * @see DeregisterDestinationInput
 * @return DeregisterDestinationResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) DeregisterDestination() {
}

/**
 * Get Subscription
 * Retrieve subscription information.
 *
 * @param mixed $request array of parameters for GetSubscription request or GetSubscription object itself
 * @see GetSubscriptionInput
 * @return GetSubscriptionResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) GetSubscription() {
}

/**
 * List Registered Destinations
 * List all the destinations for the specified seller created by the developer.
 *
 * @param mixed $request array of parameters for ListRegisteredDestinations request or ListRegisteredDestinations object itself
 * @see ListRegisteredDestinationsInput
 * @return ListRegisteredDestinationsResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) ListRegisteredDestinations() {
}

/**
 * List Subscriptions
 * List all the subscriptions for the specified seller created by the current developer.
 *
 * @param mixed $request array of parameters for ListSubscriptions request or ListSubscriptions object itself
 * @see ListSubscriptionsInput
 * @return ListSubscriptionsResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) ListSubscriptions() {
}

/**
 * Register Destination
 * Create a new destination.
 *
 * @param mixed $request array of parameters for RegisterDestination request or RegisterDestination object itself
 * @see RegisterDestinationInput
 * @return RegisterDestinationResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) RegisterDestination() {
}

/**
 * Send Test Notification To Destination
 * Send a test Notification to the specified destination.
 *
 * @param mixed $request array of parameters for SendTestNotificationToDestination request or SendTestNotificationToDestination object itself
 * @see SendTestNotificationToDestinationInput
 * @return SendTestNotificationToDestinationResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) SendTestNotificationToDestination() {
}

/**
 * Update Subscription
 * Update a subscription.
 *
 * @param mixed $request array of parameters for UpdateSubscription request or UpdateSubscription object itself
 * @see UpdateSubscriptionInput
 * @return UpdateSubscriptionResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) UpdateSubscription() {
}

/**
 * Get Service Status
 *
 *
 * @param mixed $request array of parameters for GetServiceStatus request or GetServiceStatus object itself
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 *
 * @throws SubscriptionsException
 */
func (api Subscription) GetServiceStatus() {
}
