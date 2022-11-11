package constants

const (
	PathPrefix                 = "/api/v1"
	PathOAuth2Connect          = "/oauth2/connect"
	PathOAuth2Complete         = "/oauth2/complete"
	PathDownloadUpdateSet      = "/download"
	PathCreateSubscription     = "/subscriptions"
	PathGetAllSubscriptions    = PathCreateSubscription
	PathDeleteSubscription     = PathCreateSubscription + "/{subscription_id:" + ServiceNowSysIDRegex + "}"
	PathEditSubscription       = PathDeleteSubscription
	PathGetUserChannelsForTeam = "/channels/{team_id:[A-Za-z0-9]+}"
	PathSearchRecords          = "/records/{record_type}"
	PathGetSingleRecord        = "/records/{record_type}/{record_id:" + ServiceNowSysIDRegex + "}"
	PathProcessNotification    = "/notification"
	PathGetConnected           = "/connected"
	PathGetConfig              = "/config"
	PathShareRecord            = "/share/{channel_id:[A-Za-z0-9]+}"
	PathCommentsForRecord      = "/comments/{record_type}/{record_id:" + ServiceNowSysIDRegex + "}"
	PathOpenCommentModal       = "/comment-modal"
	PathGetStatesForRecordType = "/states/{record_type}"
	PathUpdateStateOfRecord    = "/states/{record_type}/{record_id:" + ServiceNowSysIDRegex + "}"
	PathOpenStateModal         = "/state-modal"
	PathGetUsers               = "/users"

	// ServiceNow API paths
	PathActivateSubscriptions    = "api/now/table/" + ServiceNowForMattermostNotificationsAppID + "_servicenow_for_mattermost_notifications_auth"
	PathSubscriptionCRUD         = "api/now/table/" + ServiceNowForMattermostNotificationsAppID + "_servicenow_for_mattermost_subscriptions"
	PathGetRecordsFromServiceNow = "api/now/table/{tableName}"
	PathGetStatesFromServiceNow  = "api/" + ServiceNowForMattermostNotificationsAppID + "/getstates/{record_type}"
	PathGetUserFromServiceNow    = "/api/now/table/sys_user"
)
