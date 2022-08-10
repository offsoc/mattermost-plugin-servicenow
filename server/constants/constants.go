package constants

import "errors"

const (
	// Bot related constants
	BotUserName    = "servicenow"
	BotDisplayName = "ServiceNow"
	BotDescription = "A bot account created by the ServiceNow plugin."

	HeaderMattermostUserID = "Mattermost-User-ID"
	CommandTrigger         = "servicenow"

	ConnectSuccessMessage = "#### Welcome to the Mattermost ServiceNow Plugin!\n" +
		"You've connected your Mattermost account `%s` to ServiceNow. Read about the features of this plugin below:\n\n" +
		"##### Slash Commands\n"

	ServiceNowForMattermostNotificationsAppID = "x_830655_mm_std"
	ServiceNowSysIDRegex                      = "[0-9a-f]{32}"
	SysQueryParam                             = "sysparm_query"
	SysQueryParamLimit                        = "sysparm_limit"
	SysQueryParamOffset                       = "sysparm_offset"
	SysQueryParamFields                       = "sysparm_fields"
	SysQueryParamDisplayValue                 = "sysparm_display_value"

	UpdateSetNotUploadedMessage = "it looks like the notifications have not been configured in ServiceNow by uploading and committing the update set."
	UpdateSetVersion            = "v1.0"
	UpdateSetFilename           = "servicenow_for_mattermost_notifications_" + UpdateSetVersion + ".xml"

	SubscriptionTypeRecord              = "record"
	SubscriptionRecordTypeProblem       = "problem"
	SubscriptionRecordTypeIncident      = "incident"
	SubscriptionRecordTypeChangeRequest = "change_request"
	SubscriptionEventPriority           = "priority"
	SubscriptionEventState              = "state"
	SubscriptionEventCommented          = "commented"
	SubscriptionEventAssignedTo         = "assigned_to"
	SubscriptionEventAssignmentGroup    = "assignment_group"

	// Used for storing the token in the request context to pass from one middleware to another
	// #nosec G101 -- This is a false positive. The below line is not a hardcoded credential
	ContextTokenKey ServiceNowOAuthToken = "ServiceNow-Oauth-Token"

	DefaultPage             = 0
	DefaultPerPage          = 20
	MaxPerPage              = 100
	QueryParamPage          = "page"
	QueryParamPerPage       = "per_page"
	QueryParamChannelID     = "channel_id"
	QueryParamUserID        = "user_id"
	QueryParamSearchTerm    = "search"
	PathParamSubscriptionID = "subscription_id"
	PathParamTeamID         = "team_id"
	PathParamRecordType     = "record_type"
	PathParamRecordID       = "record_id"

	// Websocket events
	WSEventConnect    = "connect"
	WSEventDisconnect = "disconnect"
)

var (
	ValidSubscriptionRecordTypes = map[string]bool{
		SubscriptionRecordTypeIncident:      true,
		SubscriptionRecordTypeProblem:       true,
		SubscriptionRecordTypeChangeRequest: true,
	}

	ValidSubscriptionEvents = map[string]bool{
		SubscriptionEventPriority:        true,
		SubscriptionEventState:           true,
		SubscriptionEventCommented:       true,
		SubscriptionEventAssignedTo:      true,
		SubscriptionEventAssignmentGroup: true,
	}

	FormattedEventNames = map[string]string{
		SubscriptionEventPriority:        "Priority changed",
		SubscriptionEventState:           "State changed",
		SubscriptionEventCommented:       "New comment",
		SubscriptionEventAssignedTo:      "Assigned to changed",
		SubscriptionEventAssignmentGroup: "Assignment group changed",
	}

	FormattedRecordTypes = map[string]string{
		SubscriptionRecordTypeProblem:       "Problem",
		SubscriptionRecordTypeIncident:      "Incident",
		SubscriptionRecordTypeChangeRequest: "Change Request",
	}

	ErrUpdateSetNotUploaded error = errors.New("update set not uploaded")
)

type ServiceNowOAuthToken string
