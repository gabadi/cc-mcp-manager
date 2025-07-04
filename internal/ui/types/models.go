package types

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// AppState represents the current application state
type AppState int

const (
	MainNavigation AppState = iota
	SearchMode
	SearchActiveNavigation // Combined search + navigation
	ModalActive
)

// ToggleOperationState represents the current toggle operation state
type ToggleOperationState int

const (
	ToggleIdle ToggleOperationState = iota
	ToggleLoading
	ToggleSuccess
	ToggleError
	ToggleRetrying
)

// LoadingType represents the type of loading operation
type LoadingType int

const (
	LoadingStartup LoadingType = iota
	LoadingRefresh
)

// SpinnerState represents the current spinner animation frame
type SpinnerState int

const (
	SpinnerFrame1 SpinnerState = iota
	SpinnerFrame2
	SpinnerFrame3
	SpinnerFrame4
)

// LoadingOverlay represents the loading overlay state
type LoadingOverlay struct {
	Active      bool
	Message     string
	Spinner     SpinnerState
	Cancellable bool
	Type        LoadingType
}

// GetSpinnerChar returns the character for the current spinner state
func (s SpinnerState) GetSpinnerChar() string {
	switch s {
	case SpinnerFrame1:
		return "◐"
	case SpinnerFrame2:
		return "◓"
	case SpinnerFrame3:
		return "◑"
	case SpinnerFrame4:
		return "◒"
	default:
		return "◐"
	}
}

// Model represents the main application model
type Model struct {
	// Window dimensions
	Width  int
	Height int

	// Application state
	State AppState

	// Navigation
	ActiveColumn          int
	SelectedItem          int
	FilteredSelectedIndex int // Track selection position in filtered results

	// Search
	SearchQuery       string
	SearchActive      bool
	SearchInputActive bool // Toggle text input vs navigation
	SearchResults     []string

	// Layout
	Columns     []Column
	ColumnCount int

	// MCP inventory (placeholder for future stories)
	MCPItems []MCPItem

	// Modal state
	ActiveModal ModalType

	// Form state for add MCP workflow
	FormData   FormData
	FormErrors map[string]string

	// Success message state
	SuccessMessage string
	SuccessTimer   int // Timer for auto-hiding success message

	// Edit mode state (Epic 1 Story 4)
	EditMode    bool   // True when editing an existing MCP
	EditMCPName string // Name of the MCP being edited

	// Claude integration state (Epic 2 Story 1)
	ClaudeAvailable bool
	ClaudeStatus    ClaudeStatus
	LastClaudeSync  time.Time
	ClaudeSyncError string

	// Toggle operation state (Epic 2 Story 2)
	ToggleState     ToggleOperationState
	ToggleMCPName   string
	ToggleError     string
	ToggleRetrying  bool
	LastToggleSync  time.Time
	ToggleStartTime time.Time

	// Loading overlay state (Epic 2 Story 6)
	LoadingOverlay *LoadingOverlay
}

// ModalType represents the type of modal being displayed
type ModalType int

const (
	NoModal ModalType = iota
	AddModal
	AddMCPTypeSelection
	AddCommandForm
	AddSSEForm
	AddJSONForm
	EditModal
	DeleteModal
)

// FormData represents the current form data during MCP addition
type FormData struct {
	Name        string
	Command     string
	Args        string
	URL         string
	JSONConfig  string
	Environment string // UI input as string, converted to map[string]string on save
	ActiveField int    // Track which field is currently focused for Tab navigation
}

// MCPItem represents an MCP in the inventory
type MCPItem struct {
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Active      bool              `json:"active"`
	Command     string            `json:"command"`
	Args        []string          `json:"args,omitempty"` // Changed from string to []string for MCP standard compliance
	URL         string            `json:"url,omitempty"`
	JSONConfig  string            `json:"json_config,omitempty"`
	Environment map[string]string `json:"env,omitempty"` // New field for environment variables
}

// Column represents a UI column
type Column struct {
	Title string
	Items []string
	Width int
}

// ClaudeStatus represents the status of Claude CLI integration
type ClaudeStatus struct {
	Available    bool      `json:"available"`
	Version      string    `json:"version,omitempty"`
	ActiveMCPs   []string  `json:"active_mcps,omitempty"`
	LastCheck    time.Time `json:"last_check"`
	Error        string    `json:"error,omitempty"`
	InstallGuide string    `json:"install_guide,omitempty"`
}

// TimerTickMsg represents a timer tick message for countdown functionality
type TimerTickMsg struct {
	ID string // Unique identifier for the timer
}

// LoadingProgressMsg represents a loading progress message
type LoadingProgressMsg struct {
	Type    LoadingType
	Message string
	Done    bool
}

// LoadingSpinnerMsg represents a spinner animation tick message
type LoadingSpinnerMsg struct {
	Type LoadingType
}

// LoadingStepMsg represents a loading step progression message
type LoadingStepMsg struct {
	Type LoadingType
	Step int
}

// getDefaultMCPs returns the default MCP items for fallback
func getDefaultMCPs() []MCPItem {
	return []MCPItem{
		{Name: "context7", Type: "SSE", Active: true, Command: "npx @context7/mcp-server"},
		{Name: "github-mcp", Type: "CMD", Active: true, Command: "github-mcp"},
		{Name: "ht-mcp", Type: "CMD", Active: true, Command: "ht-mcp"},
		{Name: "filesystem", Type: "CMD", Active: false, Command: "filesystem-mcp"},
		{Name: "docker-mcp", Type: "CMD", Active: false, Command: "docker-mcp"},
		{Name: "redis-mcp", Type: "CMD", Active: false, Command: "redis-mcp"},
		{Name: "jira-mcp", Type: "CMD", Active: false, Command: "jira-mcp"},
		{Name: "aws-mcp", Type: "JSON", Active: false, Command: "aws-mcp"},
		{Name: "k8s-mcp", Type: "CMD", Active: false, Command: "k8s-mcp"},
		{Name: "confluence", Type: "SSE", Active: false, Command: "confluence-mcp"},
		{Name: "mongodb", Type: "CMD", Active: false, Command: "mongodb-mcp"},
		{Name: "terraform", Type: "CMD", Active: false, Command: "terraform-mcp"},
		{Name: "gitlab-mcp", Type: "CMD", Active: false, Command: "gitlab-mcp"},
		{Name: "linear-mcp", Type: "CMD", Active: false, Command: "linear-mcp"},
		{Name: "postgres", Type: "CMD", Active: false, Command: "postgres-mcp"},
		{Name: "elastic", Type: "JSON", Active: false, Command: "elastic-mcp"},
		{Name: "bitbucket", Type: "CMD", Active: false, Command: "bitbucket-mcp"},
		{Name: "asana-mcp", Type: "CMD", Active: false, Command: "asana-mcp"},
		{Name: "notion-mcp", Type: "CMD", Active: false, Command: "notion-mcp"},
		{Name: "anthropic", Type: "HTTP", Active: false, Command: "anthropic-mcp"},
		{Name: "sourcegraph", Type: "CMD", Active: false, Command: "sourcegraph-mcp"},
		{Name: "todoist", Type: "CMD", Active: false, Command: "todoist-mcp"},
		{Name: "slack-mcp", Type: "CMD", Active: false, Command: "slack-mcp"},
		{Name: "openai-mcp", Type: "HTTP", Active: false, Command: "openai-mcp"},
		{Name: "codeberg", Type: "CMD", Active: false, Command: "codeberg-mcp"},
		{Name: "calendar", Type: "CMD", Active: false, Command: "calendar-mcp"},
		{Name: "discord", Type: "CMD", Active: false, Command: "discord-mcp"},
		{Name: "gemini-mcp", Type: "HTTP", Active: false, Command: "gemini-mcp"},
		{Name: "gitness", Type: "CMD", Active: false, Command: "gitness-mcp"},
		{Name: "email-mcp", Type: "CMD", Active: false, Command: "email-mcp"},
		{Name: "teams-mcp", Type: "CMD", Active: false, Command: "teams-mcp"},
		{Name: "claude-mcp", Type: "HTTP", Active: false, Command: "claude-mcp"},
		{Name: "fossil-mcp", Type: "CMD", Active: false, Command: "fossil-mcp"},
		{Name: "browser", Type: "CMD", Active: false, Command: "browser-mcp"},
		{Name: "zoom-mcp", Type: "CMD", Active: false, Command: "zoom-mcp"},
	}
}

// NewModel creates a new application model
func NewModel() Model {
	return Model{
		State:             MainNavigation,
		ActiveColumn:      0,
		SelectedItem:      0,
		SearchQuery:       "",
		SearchActive:      false,
		SearchInputActive: false,
		Columns:           make([]Column, 1),
		ColumnCount:       1,
		MCPItems:          getDefaultMCPs(), // This will be replaced by storage loading
		FormErrors:        make(map[string]string),
	}
}

// NewModelWithMCPs creates a new application model with provided MCP items
func NewModelWithMCPs(mcpItems []MCPItem) Model {
	model := NewModel()
	model.MCPItems = mcpItems
	model.FormErrors = make(map[string]string)
	return model
}

// StartLoadingOverlay starts a loading overlay with the given type
func (m *Model) StartLoadingOverlay(loadingType LoadingType) {
	m.LoadingOverlay = &LoadingOverlay{
		Active:      true,
		Message:     getInitialLoadingMessage(loadingType),
		Spinner:     SpinnerFrame1,
		Cancellable: true,
		Type:        loadingType,
	}
}

// UpdateLoadingMessage updates the loading overlay message
func (m *Model) UpdateLoadingMessage(message string) {
	if m.LoadingOverlay != nil && m.LoadingOverlay.Active {
		m.LoadingOverlay.Message = message
	}
}

// StopLoadingOverlay stops the loading overlay
func (m *Model) StopLoadingOverlay() {
	if m.LoadingOverlay != nil {
		m.LoadingOverlay.Active = false
		m.LoadingOverlay = nil
	}
}

// AdvanceSpinner advances the spinner to the next frame
func (m *Model) AdvanceSpinner() {
	if m.LoadingOverlay != nil && m.LoadingOverlay.Active {
		m.LoadingOverlay.Spinner = (m.LoadingOverlay.Spinner + 1) % 4
	}
}

// IsLoadingOverlayActive returns true if loading overlay is active
func (m *Model) IsLoadingOverlayActive() bool {
	return m.LoadingOverlay != nil && m.LoadingOverlay.Active
}

// getInitialLoadingMessage returns the initial message for the loading type
func getInitialLoadingMessage(loadingType LoadingType) string {
	switch loadingType {
	case LoadingStartup:
		return "Initializing MCP Manager..."
	case LoadingRefresh:
		return "Refreshing MCP status..."
	default:
		return "Loading..."
	}
}

// Init initializes the application
func (m Model) Init() tea.Cmd {
	return nil
}
