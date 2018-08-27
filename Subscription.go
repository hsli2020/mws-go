package main

type SubscriptionApi struct {
	Client
}

/**
 * Create Subscription
 *
 * Create a new subscription.
 *
 * @see CreateSubscriptionInput
 * @return CreateSubscriptionResponse
 */
func (api SubscriptionApi) CreateSubscription() {
}

/**
 * Delete Subscription
 *
 * Delete a subscription.
 *
 * @see DeleteSubscriptionInput
 * @return DeleteSubscriptionResponse
 */
func (api SubscriptionApi) DeleteSubscription() {
}

/**
 * Deregister Destination
 *
 * Delete a destination.
 *
 * @see DeregisterDestinationInput
 * @return DeregisterDestinationResponse
 */
func (api SubscriptionApi) DeregisterDestination() {
}

/**
 * Get Subscription
 *
 * Retrieve subscription information.
 *
 * @see GetSubscriptionInput
 * @return GetSubscriptionResponse
 */
func (api SubscriptionApi) GetSubscription() {
}

/**
 * List Registered Destinations
 *
 * List all the destinations for the specified seller created by the developer.
 *
 * @see ListRegisteredDestinationsInput
 * @return ListRegisteredDestinationsResponse
 */
func (api SubscriptionApi) ListRegisteredDestinations() {
}

/**
 * List Subscriptions
 *
 * List all the subscriptions for the specified seller created by the current developer.
 *
 * @see ListSubscriptionsInput
 * @return ListSubscriptionsResponse
 */
func (api SubscriptionApi) ListSubscriptions() {
}

/**
 * Register Destination
 *
 * Create a new destination.
 *
 * @see RegisterDestinationInput
 * @return RegisterDestinationResponse
 */
func (api SubscriptionApi) RegisterDestination() {
}

/**
 * Send Test Notification To Destination
 *
 * Send a test Notification to the specified destination.
 *
 * @see SendTestNotificationToDestinationInput
 * @return SendTestNotificationToDestinationResponse
 */
func (api SubscriptionApi) SendTestNotificationToDestination() {
}

/**
 * Update Subscription
 *
 * Update a subscription.
 *
 * @see UpdateSubscriptionInput
 * @return UpdateSubscriptionResponse
 */
func (api SubscriptionApi) UpdateSubscription() {
}

/**
 * Get Service Status
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api SubscriptionApi) GetServiceStatus() {
}
