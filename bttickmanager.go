
package malgova

type btTickManager struct {
	observerAlgoIDs []string
}

func (s *btTickManager) addObserver(algoID string) {