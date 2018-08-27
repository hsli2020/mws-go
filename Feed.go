package main

type Feed struct {
	Client
}

/**
 * Submit Feed
 *
 * Uploads a file for processing together with the necessary
 * metadata to process the file, such as which type of feed it is.
 * PurgeAndReplace if true means that your existing e.g. inventory is
 * wiped out and replace with the contents of this feed - use with
 * caution (the default is false).
 *
 * @see SubmitFeedRequest
 * @return SubmitFeedResponse
 */
func (api Feed) SubmitFeed() {
}

/**
 * Get Feed Submission List By Next Token
 *
 * retrieve the next batch of list items and if there are more items to retrieve
 *
 * @see GetFeedSubmissionListByNextTokenRequest
 * @return GetFeedSubmissionListByNextTokenResponse
 */
func (api Feed) GetFeedSubmissionListByNextToken() {
}

/**
 * Cancel Feed Submissions
 *
 * cancels feed submissions - by default all of the submissions of the
 * last 30 days that have not started processing
 *
 * @see CancelFeedSubmissionsRequest
 * @return CancelFeedSubmissionsResponse
 */
func (api Feed) CancelFeedSubmissions() {
}

/**
 * Get Feed Submission Count
 *
 * returns the number of feeds matching all of the specified criteria
 *
 * @see GetFeedSubmissionCountRequest
 * @return GetFeedSubmissionCountResponse
 */
func (api Feed) GetFeedSubmissionCount() {
}

/**
 * Get Feed Submission Result
 *
 * retrieves the feed processing report
 *
 * @see GetFeedSubmissionResultRequest
 * @return GetFeedSubmissionResultResponse
 */
func (api Feed) GetFeedSubmissionResult() {
}

/**
 * Get Feed Submission List
 *
 * returns a list of feed submission identifiers and their associated metadata
 *
 * @see GetFeedSubmissionListRequest
 * @return GetFeedSubmissionListResponse
 */
func (api Feed) GetFeedSubmissionList() {
}
