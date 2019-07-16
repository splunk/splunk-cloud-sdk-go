/*
* Splunk Search Service
*
 */

package search

import "time"

// WaitForJob polls the job until it's completed or errors out
func (s *Service) WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error) {
	for {
		job, err := s.GetJob(jobID)
		if err != nil {
			return nil, err
		}
		// wait for terminal state
		switch *job.Status {
		case SearchStatusDone, SearchStatusFailed, SearchStatusCanceled:
			return *job.Status, nil
		default:
			time.Sleep(pollInterval)
		}
	}
}
