package income

//========================================================================
//                           Initial Service
//========================================================================

// NewService New Service
func NewService(repo Repository) Service {
	svc := service{repo}
	return &svc
}

type service struct {
	repo Repository
}

//========================================================================
//                           Helper Service
//========================================================================
// Choose "Not" To Create ServiceRepo For Use Outsite Service
// Example (s *service) DemoList { return s.repo.DemoList}
