package main

type MerchantFulfillmentApi struct {
	Client
}

/**
 * Cancel Shipment
 *
 * Cancels an existing shipment. This will only succeed if the cancellation 
 * window has not passed and if the shipment has not been cancelled already.
 *
 * @see CancelShipmentRequest
 * @return CancelShipmentResponse
 */
func (api MerchantFulfillmentApi) CancelShipment() {
}

/**
 * Create Shipment
 *
 * Creates a shipment for the shipping information specified. Purchases and
 * returns a label for the specified shipping service or shipping service offering.
 *
 * @see CreateShipmentRequest
 * @return CreateShipmentResponse
 */
func (api MerchantFulfillmentApi) CreateShipment() {
}

/**
 * Get Eligible Shipping Services
 *
 * Gets a list of eligible shipping services for the shipment information specified.
 * The ShippingServiceId or ShippingServiceOfferingId can be used in CreateShipment 
 * to specify the shipping service or the specific offer respectively. A list of 
 * carriers that are temporarily unavailable is also returned.
 *
 * @see GetEligibleShippingServicesRequest
 * @return GetEligibleShippingServicesResponse
 */
func (api MerchantFulfillmentApi) GetEligibleShippingServices() {
}

/**
 * Get Shipment
 *
 * Gets an existing shipment, including the label status, label content, shipping information.
 *
 * @see GetShipmentRequest
 * @return GetShipmentResponse
 */
func (api MerchantFulfillmentApi) GetShipment() {
}

/**
 * Get Service Status
 *
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 */
func (api MerchantFulfillmentApi) GetServiceStatus() {
}
