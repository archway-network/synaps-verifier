package synaps

type SessionInfo struct {
	SessionId string `json:"session_id"`
	Sandbox   bool   `json:"sandbox"`
	Status    string `json:"status"`
	Alias     string `json:"alias"`
}

type SessionDetails struct {
	SessionInfo
	Steps map[string]struct {
		Verification struct {
			State string `json:"state"`
		} `json:"verification"`

		Type string `json:"type"`
	} `json:"steps"`
}

/*-------------------*/

func (s *SessionDetails) IsVerified() bool {
	for i := range s.Steps {
		if s.Steps[i].Verification.State != VERIFICATION_STATE_VALIDATED {
			return false
		}
	}

	// If we have no Steps to verify, then we return false
	return len(s.Steps) > 0
}
