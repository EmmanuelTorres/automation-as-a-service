package app

func (m *MicroserviceServer) getUserIdFromToken(token string) (int64, error) {
	userID, err := m.tokenManager.Parse(token)
	if err != nil {
		return 0, err
	}

	return *userID, nil
}
