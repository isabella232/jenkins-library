// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/gcs"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/SAP/jenkins-library/pkg/validation"
	"github.com/bmatcuk/doublestar"
	"github.com/spf13/cobra"
)

type checkmarxExecuteScanOptions struct {
	Assignees                     []string `json:"assignees,omitempty"`
	AvoidDuplicateProjectScans    bool     `json:"avoidDuplicateProjectScans,omitempty"`
	FilterPattern                 string   `json:"filterPattern,omitempty"`
	FullScanCycle                 string   `json:"fullScanCycle,omitempty"`
	FullScansScheduled            bool     `json:"fullScansScheduled,omitempty"`
	GeneratePdfReport             bool     `json:"generatePdfReport,omitempty"`
	GithubAPIURL                  string   `json:"githubApiUrl,omitempty"`
	GithubToken                   string   `json:"githubToken,omitempty"`
	Incremental                   bool     `json:"incremental,omitempty"`
	MaxRetries                    int      `json:"maxRetries,omitempty"`
	Owner                         string   `json:"owner,omitempty"`
	Password                      string   `json:"password,omitempty"`
	Preset                        string   `json:"preset,omitempty"`
	ProjectName                   string   `json:"projectName,omitempty"`
	PullRequestName               string   `json:"pullRequestName,omitempty"`
	Repository                    string   `json:"repository,omitempty"`
	ServerURL                     string   `json:"serverUrl,omitempty"`
	SourceEncoding                string   `json:"sourceEncoding,omitempty"`
	TeamID                        string   `json:"teamId,omitempty"`
	TeamName                      string   `json:"teamName,omitempty"`
	Username                      string   `json:"username,omitempty"`
	VerifyOnly                    bool     `json:"verifyOnly,omitempty"`
	VulnerabilityThresholdEnabled bool     `json:"vulnerabilityThresholdEnabled,omitempty"`
	VulnerabilityThresholdHigh    int      `json:"vulnerabilityThresholdHigh,omitempty"`
	VulnerabilityThresholdLow     int      `json:"vulnerabilityThresholdLow,omitempty"`
	VulnerabilityThresholdMedium  int      `json:"vulnerabilityThresholdMedium,omitempty"`
	VulnerabilityThresholdResult  string   `json:"vulnerabilityThresholdResult,omitempty" validate:"possible-values=FAILURE"`
	VulnerabilityThresholdUnit    string   `json:"vulnerabilityThresholdUnit,omitempty"`
	IsOptimizedAndScheduled       bool     `json:"isOptimizedAndScheduled,omitempty"`
}

type checkmarxExecuteScanInflux struct {
	step_data struct {
		fields struct {
			checkmarx bool
		}
		tags struct {
		}
	}
	checkmarx_data struct {
		fields struct {
			high_issues                          int
			high_not_false_postive               int
			high_not_exploitable                 int
			high_confirmed                       int
			high_urgent                          int
			high_proposed_not_exploitable        int
			high_to_verify                       int
			medium_issues                        int
			medium_not_false_postive             int
			medium_not_exploitable               int
			medium_confirmed                     int
			medium_urgent                        int
			medium_proposed_not_exploitable      int
			medium_to_verify                     int
			low_issues                           int
			low_not_false_postive                int
			low_not_exploitable                  int
			low_confirmed                        int
			low_urgent                           int
			low_proposed_not_exploitable         int
			low_to_verify                        int
			information_issues                   int
			information_not_false_postive        int
			information_not_exploitable          int
			information_confirmed                int
			information_urgent                   int
			information_proposed_not_exploitable int
			information_to_verify                int
			lines_of_code_scanned                int
			files_scanned                        int
			initiator_name                       string
			owner                                string
			scan_id                              string
			project_id                           string
			projectName                          string
			team                                 string
			team_full_path_on_report_date        string
			scan_start                           string
			scan_time                            string
			checkmarx_version                    string
			scan_type                            string
			preset                               string
			deep_link                            string
			report_creation_time                 string
		}
		tags struct {
		}
	}
}

func (i *checkmarxExecuteScanInflux) persist(path, resourceName string) {
	measurementContent := []struct {
		measurement string
		valType     string
		name        string
		value       interface{}
	}{
		{valType: config.InfluxField, measurement: "step_data", name: "checkmarx", value: i.step_data.fields.checkmarx},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_issues", value: i.checkmarx_data.fields.high_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_not_false_postive", value: i.checkmarx_data.fields.high_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_not_exploitable", value: i.checkmarx_data.fields.high_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_confirmed", value: i.checkmarx_data.fields.high_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_urgent", value: i.checkmarx_data.fields.high_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_proposed_not_exploitable", value: i.checkmarx_data.fields.high_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "high_to_verify", value: i.checkmarx_data.fields.high_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_issues", value: i.checkmarx_data.fields.medium_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_not_false_postive", value: i.checkmarx_data.fields.medium_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_not_exploitable", value: i.checkmarx_data.fields.medium_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_confirmed", value: i.checkmarx_data.fields.medium_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_urgent", value: i.checkmarx_data.fields.medium_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_proposed_not_exploitable", value: i.checkmarx_data.fields.medium_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "medium_to_verify", value: i.checkmarx_data.fields.medium_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_issues", value: i.checkmarx_data.fields.low_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_not_false_postive", value: i.checkmarx_data.fields.low_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_not_exploitable", value: i.checkmarx_data.fields.low_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_confirmed", value: i.checkmarx_data.fields.low_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_urgent", value: i.checkmarx_data.fields.low_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_proposed_not_exploitable", value: i.checkmarx_data.fields.low_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "low_to_verify", value: i.checkmarx_data.fields.low_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_issues", value: i.checkmarx_data.fields.information_issues},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_not_false_postive", value: i.checkmarx_data.fields.information_not_false_postive},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_not_exploitable", value: i.checkmarx_data.fields.information_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_confirmed", value: i.checkmarx_data.fields.information_confirmed},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_urgent", value: i.checkmarx_data.fields.information_urgent},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_proposed_not_exploitable", value: i.checkmarx_data.fields.information_proposed_not_exploitable},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "information_to_verify", value: i.checkmarx_data.fields.information_to_verify},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "lines_of_code_scanned", value: i.checkmarx_data.fields.lines_of_code_scanned},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "files_scanned", value: i.checkmarx_data.fields.files_scanned},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "initiator_name", value: i.checkmarx_data.fields.initiator_name},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "owner", value: i.checkmarx_data.fields.owner},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_id", value: i.checkmarx_data.fields.scan_id},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "project_id", value: i.checkmarx_data.fields.project_id},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "projectName", value: i.checkmarx_data.fields.projectName},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "team", value: i.checkmarx_data.fields.team},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "team_full_path_on_report_date", value: i.checkmarx_data.fields.team_full_path_on_report_date},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_start", value: i.checkmarx_data.fields.scan_start},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_time", value: i.checkmarx_data.fields.scan_time},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "checkmarx_version", value: i.checkmarx_data.fields.checkmarx_version},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "scan_type", value: i.checkmarx_data.fields.scan_type},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "preset", value: i.checkmarx_data.fields.preset},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "deep_link", value: i.checkmarx_data.fields.deep_link},
		{valType: config.InfluxField, measurement: "checkmarx_data", name: "report_creation_time", value: i.checkmarx_data.fields.report_creation_time},
	}

	errCount := 0
	for _, metric := range measurementContent {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(metric.measurement, fmt.Sprintf("%vs", metric.valType), metric.name), metric.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting influx environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Error("failed to persist Influx environment")
	}
}

type checkmarxExecuteScanReports struct {
}

func (p *checkmarxExecuteScanReports) persist(stepConfig checkmarxExecuteScanOptions, gcpJsonKeyFilePath string, gcsBucketId string, gcsFolderPath string, gcsSubFolder string) {
	if gcsBucketId == "" {
		log.Entry().Info("persisting reports to GCS is disabled, because gcsBucketId is empty")
		return
	}
	log.Entry().Info("Uploading reports to Google Cloud Storage...")
	content := []gcs.ReportOutputParam{
		{FilePattern: "**/piper_checkmarx_report.html", ParamRef: "", StepResultType: "checkmarx"},
		{FilePattern: "**/CxSASTResults_*.xml", ParamRef: "", StepResultType: "checkmarx"},
		{FilePattern: "**/ScanReport.*", ParamRef: "", StepResultType: "checkmarx"},
		{FilePattern: "**/toolrun_checkmarx_*.json", ParamRef: "", StepResultType: "checkmarx"},
	}
	envVars := []gcs.EnvVar{
		{Name: "GOOGLE_APPLICATION_CREDENTIALS", Value: gcpJsonKeyFilePath, Modified: false},
	}
	gcsClient, err := gcs.NewClient(gcs.WithEnvVars(envVars))
	if err != nil {
		log.Entry().Errorf("creation of GCS client failed: %v", err)
		return
	}
	defer gcsClient.Close()
	structVal := reflect.ValueOf(&stepConfig).Elem()
	inputParameters := map[string]string{}
	for i := 0; i < structVal.NumField(); i++ {
		field := structVal.Type().Field(i)
		if field.Type.String() == "string" {
			paramName := strings.Split(field.Tag.Get("json"), ",")
			paramValue, _ := structVal.Field(i).Interface().(string)
			inputParameters[paramName[0]] = paramValue
		}
	}
	if err := gcs.PersistReportsToGCS(gcsClient, content, inputParameters, gcsFolderPath, gcsBucketId, gcsSubFolder, doublestar.Glob, os.Stat); err != nil {
		log.Entry().Errorf("failed to persist reports: %v", err)
	}
}

// CheckmarxExecuteScanCommand Checkmarx is the recommended tool for security scans of JavaScript, iOS, Swift and Ruby code.
func CheckmarxExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "checkmarxExecuteScan"

	metadata := checkmarxExecuteScanMetadata()
	var stepConfig checkmarxExecuteScanOptions
	var startTime time.Time
	var influx checkmarxExecuteScanInflux
	var reports checkmarxExecuteScanReports
	var logCollector *log.CollectorHook
	var splunkClient *splunk.Splunk
	telemetryClient := &telemetry.Telemetry{}

	var createCheckmarxExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Checkmarx is the recommended tool for security scans of JavaScript, iOS, Swift and Ruby code.",
		Long: `Checkmarx is a Static Application Security Testing (SAST) tool to analyze i.e. Java- or TypeScript, Swift, Golang, Ruby code,
and many other programming languages for security flaws based on a set of provided rules/queries that can be customized and extended.

This step by default enforces a specific audit baseline for findings and therefore ensures that:

* No 'To Verify' High and Medium issues exist in your project
* Total number of High and Medium 'Confirmed' or 'Urgent' issues is zero
* 10% of all Low issues are 'Confirmed' or 'Not Exploitable'

You can adapt above thresholds specifically using the provided configuration parameters and i.e. check for ` + "`" + `absolute` + "`" + `
thresholds instead of ` + "`" + `percentage` + "`" + ` whereas we strongly recommend you to stay with the defaults provided.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			GeneralConfig.GitHubAccessTokens = ResolveAccessTokens(GeneralConfig.GitHubTokens)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.GithubToken)
			log.RegisterSecret(stepConfig.Password)
			log.RegisterSecret(stepConfig.Username)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient = &splunk.Splunk{}
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			validation, err := validation.New(validation.WithJSONNamesForStructFields(), validation.WithPredefinedErrorMessages())
			if err != nil {
				return err
			}
			if err = validation.ValidateStruct(stepConfig); err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			stepTelemetryData := telemetry.CustomData{}
			stepTelemetryData.ErrorCode = "1"
			handler := func() {
				influx.persist(GeneralConfig.EnvRootPath, "influx")
				reports.persist(stepConfig, GeneralConfig.GCPJsonKeyFilePath, GeneralConfig.GCSBucketId, GeneralConfig.GCSFolderPath, GeneralConfig.GCSSubFolder)
				config.RemoveVaultSecretFiles()
				stepTelemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				stepTelemetryData.ErrorCategory = log.GetErrorCategory().String()
				stepTelemetryData.PiperCommitHash = GitCommit
				telemetryClient.SetData(&stepTelemetryData)
				telemetryClient.Send()
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunkClient.Send(telemetryClient.GetData(), logCollector)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetryClient.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient.Initialize(GeneralConfig.CorrelationID,
					GeneralConfig.HookConfig.SplunkConfig.Dsn,
					GeneralConfig.HookConfig.SplunkConfig.Token,
					GeneralConfig.HookConfig.SplunkConfig.Index,
					GeneralConfig.HookConfig.SplunkConfig.SendLogs)
			}
			checkmarxExecuteScan(stepConfig, &stepTelemetryData, &influx)
			stepTelemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addCheckmarxExecuteScanFlags(createCheckmarxExecuteScanCmd, &stepConfig)
	return createCheckmarxExecuteScanCmd
}

func addCheckmarxExecuteScanFlags(cmd *cobra.Command, stepConfig *checkmarxExecuteScanOptions) {
	cmd.Flags().StringSliceVar(&stepConfig.Assignees, "assignees", []string{``}, "Defines the assignees for the Github Issue created/updated with the results of the scan as a list of login names.")
	cmd.Flags().BoolVar(&stepConfig.AvoidDuplicateProjectScans, "avoidDuplicateProjectScans", true, "Whether duplicate scans of the same project state shall be avoided or not")
	cmd.Flags().StringVar(&stepConfig.FilterPattern, "filterPattern", `!**/node_modules/**, !**/.xmake/**, !**/*_test.go, !**/vendor/**/*.go, **/*.html, **/*.xml, **/*.go, **/*.py, **/*.js, **/*.scala, **/*.ts`, "The filter pattern used to zip the files relevant for scanning, patterns can be negated by setting an exclamation mark in front i.e. `!test/*.js` would avoid adding any javascript files located in the test directory")
	cmd.Flags().StringVar(&stepConfig.FullScanCycle, "fullScanCycle", `5`, "Indicates how often a full scan should happen between the incremental scans when activated")
	cmd.Flags().BoolVar(&stepConfig.FullScansScheduled, "fullScansScheduled", true, "Whether full scans are to be scheduled or not. Should be used in relation with `incremental` and `fullScanCycle`")
	cmd.Flags().BoolVar(&stepConfig.GeneratePdfReport, "generatePdfReport", true, "Whether to generate a PDF report of the analysis results or not")
	cmd.Flags().StringVar(&stepConfig.GithubAPIURL, "githubApiUrl", `https://api.github.com`, "Set the GitHub API URL.")
	cmd.Flags().StringVar(&stepConfig.GithubToken, "githubToken", os.Getenv("PIPER_githubToken"), "GitHub personal access token as per https://help.github.com/en/github/authenticating-to-github/creating-a-personal-access-token-for-the-command-line")
	cmd.Flags().BoolVar(&stepConfig.Incremental, "incremental", true, "Whether incremental scans are to be applied which optimizes the scan time but might reduce detection capabilities. Therefore full scans are still required from time to time and should be scheduled via `fullScansScheduled` and `fullScanCycle`")
	cmd.Flags().IntVar(&stepConfig.MaxRetries, "maxRetries", 3, "Maximum number of HTTP request retries upon intermittend connetion interrupts")
	cmd.Flags().StringVar(&stepConfig.Owner, "owner", os.Getenv("PIPER_owner"), "Set the GitHub organization.")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "The password to authenticate")
	cmd.Flags().StringVar(&stepConfig.Preset, "preset", os.Getenv("PIPER_preset"), "The preset to use for scanning, if not set explicitly the step will attempt to look up the project's setting based on the availability of `checkmarxCredentialsId`")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", os.Getenv("PIPER_projectName"), "The name of the Checkmarx project to scan into")
	cmd.Flags().StringVar(&stepConfig.PullRequestName, "pullRequestName", os.Getenv("PIPER_pullRequestName"), "Used to supply the name for the newly created PR project branch when being used in pull request scenarios")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Set the GitHub repository.")
	cmd.Flags().StringVar(&stepConfig.ServerURL, "serverUrl", os.Getenv("PIPER_serverUrl"), "The URL pointing to the root of the Checkmarx server to be used")
	cmd.Flags().StringVar(&stepConfig.SourceEncoding, "sourceEncoding", `1`, "The source encoding to be used, if not set explicitly the project's default will be used")
	cmd.Flags().StringVar(&stepConfig.TeamID, "teamId", os.Getenv("PIPER_teamId"), "The group ID related to your team which can be obtained via the Pipeline Syntax plugin as described in the `Details` section")
	cmd.Flags().StringVar(&stepConfig.TeamName, "teamName", os.Getenv("PIPER_teamName"), "The full name of the team to assign newly created projects to which is preferred to teamId")
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "The username to authenticate")
	cmd.Flags().BoolVar(&stepConfig.VerifyOnly, "verifyOnly", false, "Whether the step shall only apply verification checks or whether it does a full scan and check cycle")
	cmd.Flags().BoolVar(&stepConfig.VulnerabilityThresholdEnabled, "vulnerabilityThresholdEnabled", true, "Whether the thresholds are enabled or not. If enabled the build will be set to `vulnerabilityThresholdResult` in case a specific threshold value is exceeded")
	cmd.Flags().IntVar(&stepConfig.VulnerabilityThresholdHigh, "vulnerabilityThresholdHigh", 100, "The specific threshold for high severity findings")
	cmd.Flags().IntVar(&stepConfig.VulnerabilityThresholdLow, "vulnerabilityThresholdLow", 10, "The specific threshold for low severity findings")
	cmd.Flags().IntVar(&stepConfig.VulnerabilityThresholdMedium, "vulnerabilityThresholdMedium", 100, "The specific threshold for medium severity findings")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityThresholdResult, "vulnerabilityThresholdResult", `FAILURE`, "The result of the build in case thresholds are enabled and exceeded")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityThresholdUnit, "vulnerabilityThresholdUnit", `percentage`, "The unit for the threshold to apply.")
	cmd.Flags().BoolVar(&stepConfig.IsOptimizedAndScheduled, "isOptimizedAndScheduled", false, "Whether the pipeline runs in optimized mode and the current execution is a scheduled one")

	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("projectName")
	cmd.MarkFlagRequired("serverUrl")
	cmd.MarkFlagRequired("username")
}

// retrieve step metadata
func checkmarxExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "checkmarxExecuteScan",
			Aliases:     []config.Alias{},
			Description: "Checkmarx is the recommended tool for security scans of JavaScript, iOS, Swift and Ruby code.",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "checkmarxCredentialsId", Description: "Jenkins 'Username with password' credentials ID containing username and password to communicate with the Checkmarx backend.", Type: "jenkins"},
				},
				Resources: []config.StepResources{
					{Name: "checkmarx", Type: "stash"},
				},
				Parameters: []config.StepParameters{
					{
						Name:        "assignees",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     []string{``},
					},
					{
						Name:        "avoidDuplicateProjectScans",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "filterPattern",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `!**/node_modules/**, !**/.xmake/**, !**/*_test.go, !**/vendor/**/*.go, **/*.html, **/*.xml, **/*.go, **/*.py, **/*.js, **/*.scala, **/*.ts`,
					},
					{
						Name:        "fullScanCycle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `5`,
					},
					{
						Name:        "fullScansScheduled",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "generatePdfReport",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "githubApiUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `https://api.github.com`,
					},
					{
						Name: "githubToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "githubTokenCredentialsId",
								Type: "secret",
							},

							{
								Name:    "githubVaultSecretName",
								Type:    "vaultSecret",
								Default: "github",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "access_token"}},
						Default:   os.Getenv("PIPER_githubToken"),
					},
					{
						Name:        "incremental",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "maxRetries",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     3,
					},
					{
						Name: "owner",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/owner",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubOrg"}},
						Default:   os.Getenv("PIPER_owner"),
					},
					{
						Name: "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "checkmarxCredentialsId",
								Param: "password",
								Type:  "secret",
							},

							{
								Name:    "checkmarxVaultSecretName",
								Type:    "vaultSecret",
								Default: "checkmarx",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_password"),
					},
					{
						Name:        "preset",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_preset"),
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "checkmarxProject"}, {Name: "checkMarxProjectName", Deprecated: true}},
						Default:     os.Getenv("PIPER_projectName"),
					},
					{
						Name:        "pullRequestName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_pullRequestName"),
					},
					{
						Name: "repository",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "github/repository",
							},
						},
						Scope:     []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{{Name: "githubRepo"}},
						Default:   os.Getenv("PIPER_repository"),
					},
					{
						Name:        "serverUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{{Name: "checkmarxServerUrl"}},
						Default:     os.Getenv("PIPER_serverUrl"),
					},
					{
						Name:        "sourceEncoding",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `1`,
					},
					{
						Name:        "teamId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "checkmarxGroupId"}, {Name: "groupId", Deprecated: true}},
						Default:     os.Getenv("PIPER_teamId"),
					},
					{
						Name:        "teamName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_teamName"),
					},
					{
						Name: "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "checkmarxCredentialsId",
								Param: "username",
								Type:  "secret",
							},

							{
								Name:    "checkmarxVaultSecretName",
								Type:    "vaultSecret",
								Default: "checkmarx",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_username"),
					},
					{
						Name:        "verifyOnly",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     false,
					},
					{
						Name:        "vulnerabilityThresholdEnabled",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     true,
					},
					{
						Name:        "vulnerabilityThresholdHigh",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     100,
					},
					{
						Name:        "vulnerabilityThresholdLow",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     10,
					},
					{
						Name:        "vulnerabilityThresholdMedium",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "int",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     100,
					},
					{
						Name:        "vulnerabilityThresholdResult",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `FAILURE`,
					},
					{
						Name:        "vulnerabilityThresholdUnit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `percentage`,
					},
					{
						Name: "isOptimizedAndScheduled",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "custom/isOptimizedAndScheduled",
							},
						},
						Scope:     []string{"PARAMETERS"},
						Type:      "bool",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   false,
					},
				},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "influx",
						Type: "influx",
						Parameters: []map[string]interface{}{
							{"name": "step_data", "fields": []map[string]string{{"name": "checkmarx"}}},
							{"name": "checkmarx_data", "fields": []map[string]string{{"name": "high_issues"}, {"name": "high_not_false_postive"}, {"name": "high_not_exploitable"}, {"name": "high_confirmed"}, {"name": "high_urgent"}, {"name": "high_proposed_not_exploitable"}, {"name": "high_to_verify"}, {"name": "medium_issues"}, {"name": "medium_not_false_postive"}, {"name": "medium_not_exploitable"}, {"name": "medium_confirmed"}, {"name": "medium_urgent"}, {"name": "medium_proposed_not_exploitable"}, {"name": "medium_to_verify"}, {"name": "low_issues"}, {"name": "low_not_false_postive"}, {"name": "low_not_exploitable"}, {"name": "low_confirmed"}, {"name": "low_urgent"}, {"name": "low_proposed_not_exploitable"}, {"name": "low_to_verify"}, {"name": "information_issues"}, {"name": "information_not_false_postive"}, {"name": "information_not_exploitable"}, {"name": "information_confirmed"}, {"name": "information_urgent"}, {"name": "information_proposed_not_exploitable"}, {"name": "information_to_verify"}, {"name": "lines_of_code_scanned"}, {"name": "files_scanned"}, {"name": "initiator_name"}, {"name": "owner"}, {"name": "scan_id"}, {"name": "project_id"}, {"name": "projectName"}, {"name": "team"}, {"name": "team_full_path_on_report_date"}, {"name": "scan_start"}, {"name": "scan_time"}, {"name": "checkmarx_version"}, {"name": "scan_type"}, {"name": "preset"}, {"name": "deep_link"}, {"name": "report_creation_time"}}},
						},
					},
					{
						Name: "reports",
						Type: "reports",
						Parameters: []map[string]interface{}{
							{"filePattern": "**/piper_checkmarx_report.html", "type": "checkmarx"},
							{"filePattern": "**/CxSASTResults_*.xml", "type": "checkmarx"},
							{"filePattern": "**/ScanReport.*", "type": "checkmarx"},
							{"filePattern": "**/toolrun_checkmarx_*.json", "type": "checkmarx"},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
