package utils

import (
	"fmt"

	"github.com/omise/omise-go"
)

type OmiseHelper struct{}

func (s *OmiseHelper) GetOmiseEventId(scope omise.SearchScope, status omise.ChargeStatus) string {
	return fmt.Sprintf("%s.%s", scope, status)
}
