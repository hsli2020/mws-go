package mws

/**
 * Cancel Shipment
 * Cancels an existing shipment.  This will only succeed if the cancellation window has not passed and if the shipment
 *     has not been cancelled already.
 *
 * @param mixed request array of parameters for CancelShipment request or CancelShipment object itself
 * @see CancelShipmentRequest
 * @return CancelShipmentResponse
 *
 * @throws MerchantFulfillmentException
 */
func CancelShipment(request) {
}

/**
 * Create Shipment
 * Creates a shipment for the shipping information specified.  Purchases and returns a label for the specified
 *     shipping service or shipping service offering.
 *
 * @param mixed request array of parameters for CreateShipment request or CreateShipment object itself
 * @see CreateShipmentRequest
 * @return CreateShipmentResponse
 *
 * @throws MerchantFulfillmentException
 */
func CreateShipment(request) {
}

/**
 * Get Eligible Shipping Services
 * Gets a list of eligible shipping services for the shipment information specified.  The ShippingServiceId or
 *     ShippingServiceOfferingId can be used in CreateShipment to specify the shipping service or the specific offer
 *     respectively.  A list of carriers that are temporarily unavailable is also returned.
 *
 * @param mixed request array of parameters for GetEligibleShippingServices request or GetEligibleShippingServices object itself
 * @see GetEligibleShippingServicesRequest
 * @return GetEligibleShippingServicesResponse
 *
 * @throws MerchantFulfillmentException
 */
func GetEligibleShippingServices(request) {
}

/**
 * Get Shipment
 * Gets an existing shipment, including the label status, label content, shipping information.
 *
 * @param mixed request array of parameters for GetShipment request or GetShipment object itself
 * @see GetShipmentRequest
 * @return GetShipmentResponse
 *
 * @throws MerchantFulfillmentException
 */
func GetShipment(request) {
}

/**
 * Get Service Status
 *
 *
 * @param mixed request array of parameters for GetServiceStatus request or GetServiceStatus object itself
 * @see GetServiceStatusRequest
 * @return GetServiceStatusResponse
 *
 * @throws MerchantFulfillmentException
 */
func GetServiceStatus(request) {
}
