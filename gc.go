package gocardless

import (
	"bytes"
	"errors"
	"fmt"
)

type Service struct {
	BankAuthorisations      BankAuthorisationService
	BankDetailsLookups      BankDetailsLookupService
	BillingRequests         BillingRequestService
	BillingRequestFlows     BillingRequestFlowService
	BillingRequestTemplates BillingRequestTemplateService
	Blocks                  BlockService
	Creditors               CreditorService
	CreditorBankAccounts    CreditorBankAccountService
	CurrencyExchangeRates   CurrencyExchangeRateService
	Customers               CustomerService
	CustomerBankAccounts    CustomerBankAccountService
	CustomerNotifications   CustomerNotificationService
	Events                  EventService
	InstalmentSchedules     InstalmentScheduleService
	Institutions            InstitutionService
	Mandates                MandateService
	MandateImports          MandateImportService
	MandateImportEntries    MandateImportEntryService
	MandatePdfs             MandatePdfService
	PayerAuthorisations     PayerAuthorisationService
	Payments                PaymentService
	Payouts                 PayoutService
	PayoutItems             PayoutItemService
	RedirectFlows           RedirectFlowService
	Refunds                 RefundService
	ScenarioSimulators      ScenarioSimulatorService
	Subscriptions           SubscriptionService
	TaxRates                TaxRateService
	Webhooks                WebhookService
}

func init() {
	initUserAgent()
}

func New(config Config) (*Service, error) {
	if config == nil {
		return nil, errors.New("invalid configuration")
	}

	s := &Service{
		BankAuthorisations: &BankAuthorisationServiceImpl{
			config: config,
		}, BankDetailsLookups: &BankDetailsLookupServiceImpl{
			config: config,
		}, BillingRequests: &BillingRequestServiceImpl{
			config: config,
		}, BillingRequestFlows: &BillingRequestFlowServiceImpl{
			config: config,
		}, BillingRequestTemplates: &BillingRequestTemplateServiceImpl{
			config: config,
		}, Blocks: &BlockServiceImpl{
			config: config,
		}, Creditors: &CreditorServiceImpl{
			config: config,
		}, CreditorBankAccounts: &CreditorBankAccountServiceImpl{
			config: config,
		}, CurrencyExchangeRates: &CurrencyExchangeRateServiceImpl{
			config: config,
		}, Customers: &CustomerServiceImpl{
			config: config,
		}, CustomerBankAccounts: &CustomerBankAccountServiceImpl{
			config: config,
		}, CustomerNotifications: &CustomerNotificationServiceImpl{
			config: config,
		}, Events: &EventServiceImpl{
			config: config,
		}, InstalmentSchedules: &InstalmentScheduleServiceImpl{
			config: config,
		}, Institutions: &InstitutionServiceImpl{
			config: config,
		}, Mandates: &MandateServiceImpl{
			config: config,
		}, MandateImports: &MandateImportServiceImpl{
			config: config,
		}, MandateImportEntries: &MandateImportEntryServiceImpl{
			config: config,
		}, MandatePdfs: &MandatePdfServiceImpl{
			config: config,
		}, PayerAuthorisations: &PayerAuthorisationServiceImpl{
			config: config,
		}, Payments: &PaymentServiceImpl{
			config: config,
		}, Payouts: &PayoutServiceImpl{
			config: config,
		}, PayoutItems: &PayoutItemServiceImpl{
			config: config,
		}, RedirectFlows: &RedirectFlowServiceImpl{
			config: config,
		}, Refunds: &RefundServiceImpl{
			config: config,
		}, ScenarioSimulators: &ScenarioSimulatorServiceImpl{
			config: config,
		}, Subscriptions: &SubscriptionServiceImpl{
			config: config,
		}, TaxRates: &TaxRateServiceImpl{
			config: config,
		}, Webhooks: &WebhookServiceImpl{
			config: config,
		},
	}

	return s, nil
}

type APIError struct {
	Message          string            `json:"message"`
	DocumentationUrl string            `json:"documentation_url"`
	Type             string            `json:"type"`
	RequestID        string            `json:"request_id"`
	Errors           []ValidationError `json:"errors"`
	Code             int               `json:"code"`
}

func (err *APIError) Error() string {
	if len(err.Errors) == 0 {
		return err.Message
	}
	var msg bytes.Buffer
	fmt.Fprintf(&msg, "%s:", err.Message)
	for _, err := range err.Errors {
		fmt.Fprintf(&msg, "\n * %s: %s", err.Field, err.Message)
	}
	return msg.String()
}

type ValidationError struct {
	Message        string `json:"message"`
	Field          string `json:"field"`
	RequestPointer string `json:"request_pointer"`
}
