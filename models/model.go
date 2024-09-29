package models

type GetAccessToken struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          struct {
		AccessToken  string `json:"accessToken"`
		ExpiredTime  int    `json:"expiredTime"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
	Success bool `json:"success"`
}

type UserInfo struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          struct {
		UserID          string      `json:"userId"`
		SocialType      string      `json:"socialType"`
		RiskPassed      bool        `json:"riskPassed"`
		Qualified       bool        `json:"qualified"`
		Participated    bool        `json:"participated"`
		Bound           bool        `json:"bound"`
		BinanceUserInfo interface{} `json:"binanceUserInfo"`
		MetaInfo        struct {
			TotalGrade                  int `json:"totalGrade"`
			ReferralTotalGrade          int `json:"referralTotalGrade"`
			TotalAttempts               int `json:"totalAttempts"`
			ConsumedAttempts            int `json:"consumedAttempts"`
			AttemptRefreshCountDownTime int `json:"attemptRefreshCountDownTime"`
		} `json:"metaInfo"`
	} `json:"data"`
	Success bool `json:"success"`
}

type VideoTkn struct {
	Code          string `json:"code"`
	Success       bool   `json:"success"`
	Message       any    `json:"message"`
	MessageDetail any    `json:"messageDetail"`
	Dfp           string `json:"dfp"`
	Dt            string `json:"dt"`
}

type CompleteGameBody struct {
	ResourceID string `json:"resourceId"`
	Payload    string `json:"payload"`
	Log        int    `json:"log"`
}
type StandardResp struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          interface{} `json:"data"`
	Success       bool        `json:"success"`
}

type CompleteTask struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          struct {
		UserID          int64       `json:"userId"`
		ResourceID      int         `json:"resourceId"`
		Type            string      `json:"type"`
		SubType         interface{} `json:"subType"`
		InviteeTaskType interface{} `json:"inviteeTaskType"`
		Status          string      `json:"status"`
		CompletedCount  int         `json:"completedCount"`
		TotalCount      int         `json:"totalCount"`
		RewardList      interface{} `json:"rewardList"`
		StartTime       int64       `json:"startTime"`
		Code            string      `json:"code"`
		ErrorCode       interface{} `json:"errorCode"`
		UpdatedTime     int64       `json:"updatedTime"`
	} `json:"data"`
	Success bool `json:"success"`
}
type LicenseChecker struct {
	Message   string `json:"message"`
	Whitelist struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		IPAddress any    `json:"ipAddress"`
	} `json:"whitelist"`
}
type StartGame struct {
	Code          string      `json:"code"`
	Message       string      `json:"message"`
	MessageDetail interface{} `json:"messageDetail"`
	Data          struct {
		GameTag           string      `json:"gameTag"`
		SpecialRewardID   interface{} `json:"specialRewardId"`
		CryptoMinerConfig struct {
			GameDuration        interface{} `json:"gameDuration"`
			HookSwipeSpeed      int         `json:"hookSwipeSpeed"`
			FinalHookSwipeSpeed int         `json:"finalHookSwipeSpeed"`
			ItemSettingList     []struct {
				Type            string `json:"type"`
				Speed           int    `json:"speed"`
				Size            int    `json:"size"`
				Quantity        int    `json:"quantity"`
				RewardValueList []int  `json:"rewardValueList"`
			} `json:"itemSettingList"`
		} `json:"cryptoMinerConfig"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GameKey struct {
	Encrypted string `json:"encrypted"`
	Point     int    `json:"point"`
}

type TaskList struct {
	Code          string `json:"code"`
	Message       any    `json:"message"`
	MessageDetail any    `json:"messageDetail"`
	Data          struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Data      []struct {
			ResourceID   int    `json:"resourceId"`
			ResourceCode string `json:"resourceCode"`
			Participated bool   `json:"participated"`
			TaskList     struct {
				PageIndex int `json:"pageIndex"`
				PageSize  int `json:"pageSize"`
				Data      []struct {
					UserID          int64  `json:"userId"`
					ResourceID      int    `json:"resourceId"`
					Type            string `json:"type"`
					SubType         any    `json:"subType"`
					InviteeTaskType any    `json:"inviteeTaskType"`
					Status          string `json:"status"`
					CompletedCount  int    `json:"completedCount"`
					TotalCount      int    `json:"totalCount"`
					RewardList      []struct {
						ID             int    `json:"id"`
						Type           string `json:"type"`
						Code           string `json:"code"`
						Asset          any    `json:"asset"`
						Amount         int    `json:"amount"`
						RemainingCount int    `json:"remainingCount"`
						TotalCount     int    `json:"totalCount"`
						RewardSubType  any    `json:"rewardSubType"`
						TaskSubType    any    `json:"taskSubType"`
						UserRewardList []struct {
							ID             string `json:"id"`
							ResourceID     int    `json:"resourceId"`
							RewardID       int    `json:"rewardId"`
							RemainingCount int    `json:"remainingCount"`
							TotalCount     int    `json:"totalCount"`
							Status         string `json:"status"`
							RewardAmount   any    `json:"rewardAmount"`
							RewardAsset    any    `json:"rewardAsset"`
							Mine           bool   `json:"mine"`
							CreatedTime    int64  `json:"createdTime"`
							UpdatedTime    int64  `json:"updatedTime"`
						} `json:"userRewardList"`
					} `json:"rewardList"`
					StartTime   int64  `json:"startTime"`
					Code        string `json:"code"`
					ErrorCode   any    `json:"errorCode"`
					UpdatedTime int64  `json:"updatedTime"`
				} `json:"data"`
				Total int `json:"total"`
			} `json:"taskList"`
			FinalRewardList []any `json:"finalRewardList"`
		} `json:"data"`
		Total int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}
