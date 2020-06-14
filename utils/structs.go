package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"sync"
)

type GetItemRequest struct {
	XMLName       xml.Name `xml:"GetItemRequest"`
	Text          string   `xml:",chardata"`
	Xmlns         string   `xml:"xmlns,attr"`
	ErrorLanguage string   `xml:"ErrorLanguage"`
	WarningLevel  string   `xml:"WarningLevel"`
	DetailLevel   string   `xml:"DetailLevel"`
	ItemID        string   `xml:"ItemID"`
}
type GetItemResponse struct {
	XMLName   xml.Name `xml:"GetItemResponse"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Timestamp string   `xml:"Timestamp"`
	Ack       string   `xml:"Ack"`
	Version   string   `xml:"Version"`
	Build     string   `xml:"Build"`
	Item      struct {
		Text            string `xml:",chardata"`
		AutoPay         string `xml:"AutoPay"`
		BuyerProtection string `xml:"BuyerProtection"`
		BuyItNowPrice   struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"BuyItNowPrice"`
		Country        string `xml:"Country"`
		Currency       string `xml:"Currency"`
		Description    string `xml:"Description"`
		GiftIcon       string `xml:"GiftIcon"`
		HitCounter     string `xml:"HitCounter"`
		ItemID         string `xml:"ItemID"`
		ListingDetails struct {
			Text                   string `xml:",chardata"`
			Adult                  string `xml:"Adult"`
			BindingAuction         string `xml:"BindingAuction"`
			CheckoutEnabled        string `xml:"CheckoutEnabled"`
			ConvertedBuyItNowPrice struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"ConvertedBuyItNowPrice"`
			ConvertedStartPrice struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"ConvertedStartPrice"`
			ConvertedReservePrice struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"ConvertedReservePrice"`
			HasReservePrice             string `xml:"HasReservePrice"`
			StartTime                   string `xml:"StartTime"`
			EndTime                     string `xml:"EndTime"`
			ViewItemURL                 string `xml:"ViewItemURL"`
			HasUnansweredQuestions      string `xml:"HasUnansweredQuestions"`
			HasPublicMessages           string `xml:"HasPublicMessages"`
			ViewItemURLForNaturalSearch string `xml:"ViewItemURLForNaturalSearch"`
		} `xml:"ListingDetails"`
		ListingDesigner struct {
			Text     string `xml:",chardata"`
			LayoutID string `xml:"LayoutID"`
			ThemeID  string `xml:"ThemeID"`
		} `xml:"ListingDesigner"`
		ListingDuration    string `xml:"ListingDuration"`
		ListingType        string `xml:"ListingType"`
		Location           string `xml:"Location"`
		PaymentMethods     string `xml:"PaymentMethods"`
		PayPalEmailAddress string `xml:"PayPalEmailAddress"`
		PrimaryCategory    struct {
			Text         string `xml:",chardata"`
			CategoryID   string `xml:"CategoryID"`
			CategoryName string `xml:"CategoryName"`
		} `xml:"PrimaryCategory"`
		PrivateListing        string `xml:"PrivateListing"`
		ProductListingDetails struct {
			Text     string `xml:",chardata"`
			UPC      string `xml:"UPC"`
			BrandMPN struct {
				Text  string `xml:",chardata"`
				Brand string `xml:"Brand"`
				MPN   string `xml:"MPN"`
			} `xml:"BrandMPN"`
			IncludeeBayProductDetails string `xml:"IncludeeBayProductDetails"`
		} `xml:"ProductListingDetails"`
		Quantity     string `xml:"Quantity"`
		ReservePrice struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"ReservePrice"`
		ReviseStatus struct {
			Text        string `xml:",chardata"`
			ItemRevised string `xml:"ItemRevised"`
		} `xml:"ReviseStatus"`
		Seller struct {
			Text                    string `xml:",chardata"`
			AboutMePage             string `xml:"AboutMePage"`
			Email                   string `xml:"Email"`
			FeedbackScore           string `xml:"FeedbackScore"`
			PositiveFeedbackPercent string `xml:"PositiveFeedbackPercent"`
			FeedbackPrivate         string `xml:"FeedbackPrivate"`
			IDVerified              string `xml:"IDVerified"`
			EBayGoodStanding        string `xml:"eBayGoodStanding"`
			NewUser                 string `xml:"NewUser"`
			RegistrationDate        string `xml:"RegistrationDate"`
			Site                    string `xml:"Site"`
			Status                  string `xml:"Status"`
			UserID                  string `xml:"UserID"`
			UserIDChanged           string `xml:"UserIDChanged"`
			UserIDLastChanged       string `xml:"UserIDLastChanged"`
			VATStatus               string `xml:"VATStatus"`
			SellerInfo              struct {
				Text                  string `xml:",chardata"`
				AllowPaymentEdit      string `xml:"AllowPaymentEdit"`
				CheckoutEnabled       string `xml:"CheckoutEnabled"`
				CIPBankAccountStored  string `xml:"CIPBankAccountStored"`
				GoodStanding          string `xml:"GoodStanding"`
				LiveAuctionAuthorized string `xml:"LiveAuctionAuthorized"`
				MerchandizingPref     string `xml:"MerchandizingPref"`
				QualifiesForB2BVAT    string `xml:"QualifiesForB2BVAT"`
				StoreOwner            string `xml:"StoreOwner"`
				StoreURL              string `xml:"StoreURL"`
				SafePaymentExempt     string `xml:"SafePaymentExempt"`
				TopRatedSeller        string `xml:"TopRatedSeller"`
			} `xml:"SellerInfo"`
			MotorsDealer string `xml:"MotorsDealer"`
		} `xml:"Seller"`
		SellingStatus struct {
			Text         string `xml:",chardata"`
			BidCount     string `xml:"BidCount"`
			BidIncrement struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"BidIncrement"`
			ConvertedCurrentPrice struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"ConvertedCurrentPrice"`
			CurrentPrice struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"CurrentPrice"`
			LeadCount    string `xml:"LeadCount"`
			MinimumToBid struct {
				Text       string `xml:",chardata"`
				CurrencyID string `xml:"currencyID,attr"`
			} `xml:"MinimumToBid"`
			QuantitySold                string `xml:"QuantitySold"`
			ReserveMet                  string `xml:"ReserveMet"`
			SecondChanceEligible        string `xml:"SecondChanceEligible"`
			ListingStatus               string `xml:"ListingStatus"`
			QuantitySoldByPickupInStore string `xml:"QuantitySoldByPickupInStore"`
		} `xml:"SellingStatus"`
		ShippingDetails struct {
			Text                   string `xml:",chardata"`
			ApplyShippingDiscount  string `xml:"ApplyShippingDiscount"`
			GlobalShipping         string `xml:"GlobalShipping"`
			CalculatedShippingRate struct {
				Text        string `xml:",chardata"`
				WeightMajor struct {
					Text              string `xml:",chardata"`
					MeasurementSystem string `xml:"measurementSystem,attr"`
					Unit              string `xml:"unit,attr"`
				} `xml:"WeightMajor"`
				WeightMinor struct {
					Text              string `xml:",chardata"`
					MeasurementSystem string `xml:"measurementSystem,attr"`
					Unit              string `xml:"unit,attr"`
				} `xml:"WeightMinor"`
			} `xml:"CalculatedShippingRate"`
			SalesTax struct {
				Text                  string `xml:",chardata"`
				SalesTaxPercent       string `xml:"SalesTaxPercent"`
				ShippingIncludedInTax string `xml:"ShippingIncludedInTax"`
			} `xml:"SalesTax"`
			ShippingServiceOptions struct {
				Text                string `xml:",chardata"`
				ShippingService     string `xml:"ShippingService"`
				ShippingServiceCost struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"ShippingServiceCost"`
				ShippingServiceAdditionalCost struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"ShippingServiceAdditionalCost"`
				ShippingServicePriority string `xml:"ShippingServicePriority"`
				ExpeditedService        string `xml:"ExpeditedService"`
				ShippingTimeMin         string `xml:"ShippingTimeMin"`
				ShippingTimeMax         string `xml:"ShippingTimeMax"`
			} `xml:"ShippingServiceOptions"`
			ShippingType                           string   `xml:"ShippingType"`
			ThirdPartyCheckout                     string   `xml:"ThirdPartyCheckout"`
			ShippingDiscountProfileID              string   `xml:"ShippingDiscountProfileID"`
			InternationalShippingDiscountProfileID string   `xml:"InternationalShippingDiscountProfileID"`
			ExcludeShipToLocation                  []string `xml:"ExcludeShipToLocation"`
			SellerExcludeShipToLocationsPreference string   `xml:"SellerExcludeShipToLocationsPreference"`
		} `xml:"ShippingDetails"`
		ShipToLocations []string `xml:"ShipToLocations"`
		Site            string   `xml:"Site"`
		StartPrice      struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"StartPrice"`
		Storefront struct {
			Text             string `xml:",chardata"`
			StoreCategoryID  string `xml:"StoreCategoryID"`
			StoreCategory2ID string `xml:"StoreCategory2ID"`
			StoreURL         string `xml:"StoreURL"`
		} `xml:"Storefront"`
		TimeLeft                    string `xml:"TimeLeft"`
		Title                       string `xml:"Title"`
		HitCount                    string `xml:"HitCount"`
		LocationDefaulted           string `xml:"LocationDefaulted"`
		GetItFast                   string `xml:"GetItFast"`
		BuyerResponsibleForShipping string `xml:"BuyerResponsibleForShipping"`
		SKU                         string `xml:"SKU"`
		PostalCode                  string `xml:"PostalCode"`
		PictureDetails              struct {
			Text               string   `xml:",chardata"`
			GalleryType        string   `xml:"GalleryType"`
			GalleryURL         string   `xml:"GalleryURL"`
			PhotoDisplay       string   `xml:"PhotoDisplay"`
			PictureURL         []string `xml:"PictureURL"`
			PictureSource      string   `xml:"PictureSource"`
			ExternalPictureURL string   `xml:"ExternalPictureURL"`
		} `xml:"PictureDetails"`
		DispatchTimeMax     string `xml:"DispatchTimeMax"`
		ProxyItem           string `xml:"ProxyItem"`
		BuyerGuaranteePrice struct {
			Text       string `xml:",chardata"`
			CurrencyID string `xml:"currencyID,attr"`
		} `xml:"BuyerGuaranteePrice"`
		BuyerRequirementDetails struct {
			Text                      string `xml:",chardata"`
			ShipToRegistrationCountry string `xml:"ShipToRegistrationCountry"`
			MaximumItemRequirements   struct {
				Text                 string `xml:",chardata"`
				MaximumItemCount     string `xml:"MaximumItemCount"`
				MinimumFeedbackScore string `xml:"MinimumFeedbackScore"`
			} `xml:"MaximumItemRequirements"`
			MaximumUnpaidItemStrikesInfo struct {
				Text   string `xml:",chardata"`
				Count  string `xml:"Count"`
				Period string `xml:"Period"`
			} `xml:"MaximumUnpaidItemStrikesInfo"`
		} `xml:"BuyerRequirementDetails"`
		IntangibleItem string `xml:"IntangibleItem"`
		ReturnPolicy   struct {
			Text                               string `xml:",chardata"`
			RefundOption                       string `xml:"RefundOption"`
			Refund                             string `xml:"Refund"`
			ReturnsWithinOption                string `xml:"ReturnsWithinOption"`
			ReturnsWithin                      string `xml:"ReturnsWithin"`
			ReturnsAcceptedOption              string `xml:"ReturnsAcceptedOption"`
			ReturnsAccepted                    string `xml:"ReturnsAccepted"`
			ShippingCostPaidByOption           string `xml:"ShippingCostPaidByOption"`
			ShippingCostPaidBy                 string `xml:"ShippingCostPaidBy"`
			InternationalReturnsAcceptedOption string `xml:"InternationalReturnsAcceptedOption"`
		} `xml:"ReturnPolicy"`
		ConditionID                   string `xml:"ConditionID"`
		ConditionDisplayName          string `xml:"ConditionDisplayName"`
		PostCheckoutExperienceEnabled string `xml:"PostCheckoutExperienceEnabled"`
		SellerProfiles                struct {
			Text                  string `xml:",chardata"`
			SellerShippingProfile struct {
				Text                string `xml:",chardata"`
				ShippingProfileID   string `xml:"ShippingProfileID"`
				ShippingProfileName string `xml:"ShippingProfileName"`
			} `xml:"SellerShippingProfile"`
			SellerReturnProfile struct {
				Text              string `xml:",chardata"`
				ReturnProfileID   string `xml:"ReturnProfileID"`
				ReturnProfileName string `xml:"ReturnProfileName"`
			} `xml:"SellerReturnProfile"`
			SellerPaymentProfile struct {
				Text               string `xml:",chardata"`
				PaymentProfileID   string `xml:"PaymentProfileID"`
				PaymentProfileName string `xml:"PaymentProfileName"`
			} `xml:"SellerPaymentProfile"`
		} `xml:"SellerProfiles"`
		ShippingServiceCostOverrideList struct {
			Text                        string `xml:",chardata"`
			ShippingServiceCostOverride struct {
				Text                    string `xml:",chardata"`
				ShippingServicePriority string `xml:"ShippingServicePriority"`
				ShippingServiceType     string `xml:"ShippingServiceType"`
				ShippingServiceCost     struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"ShippingServiceCost"`
				ShippingServiceAdditionalCost struct {
					Text       string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"ShippingServiceAdditionalCost"`
			} `xml:"ShippingServiceCostOverride"`
		} `xml:"ShippingServiceCostOverrideList"`
		ShippingPackageDetails struct {
			Text              string `xml:",chardata"`
			ShippingIrregular string `xml:"ShippingIrregular"`
			ShippingPackage   string `xml:"ShippingPackage"`
			WeightMajor       struct {
				Text              string `xml:",chardata"`
				MeasurementSystem string `xml:"measurementSystem,attr"`
				Unit              string `xml:"unit,attr"`
			} `xml:"WeightMajor"`
			WeightMinor struct {
				Text              string `xml:",chardata"`
				MeasurementSystem string `xml:"measurementSystem,attr"`
				Unit              string `xml:"unit,attr"`
			} `xml:"WeightMinor"`
		} `xml:"ShippingPackageDetails"`
		HideFromSearch      string `xml:"HideFromSearch"`
		OutOfStockControl   string `xml:"OutOfStockControl"`
		EBayPlus            string `xml:"eBayPlus"`
		EBayPlusEligible    string `xml:"eBayPlusEligible"`
		IsSecureDescription string `xml:"IsSecureDescription"`
	} `xml:"Item"`
}

type GetStoreRequest struct {
	XMLName               xml.Name `xml:"GetStoreRequest"`
	Text                  string   `xml:",chardata"`
	Xmlns                 string   `xml:"xmlns,attr"`
	ErrorLanguage         string   `xml:"ErrorLanguage"`
	WarningLevel          string   `xml:"WarningLevel"`
	CategoryStructureOnly string   `xml:"CategoryStructureOnly"`
}
type GetStoreResponse struct {
	XMLName   xml.Name `xml:"GetStoreResponse"`
	Text      string   `xml:",chardata"`
	Xmlns     string   `xml:"xmlns,attr"`
	Timestamp string   `xml:"Timestamp"`
	Ack       string   `xml:"Ack"`
	Version   string   `xml:"Version"`
	Build     string   `xml:"Build"`
	Store     struct {
		Text              string `xml:",chardata"`
		Name              string `xml:"Name"`
		SubscriptionLevel string `xml:"SubscriptionLevel"`
		Description       string `xml:"Description"`
		CustomCategories  struct {
			Text           string `xml:",chardata"`
			CustomCategory []struct {
				Text          string `xml:",chardata"`
				CategoryID    int `xml:"CategoryID"`
				Name          string `xml:"Name"`
				Order         string `xml:"Order"`
				ChildCategory []struct {
					Text          string `xml:",chardata"`
					CategoryID    string `xml:"CategoryID"`
					Name          string `xml:"Name"`
					Order         string `xml:"Order"`
					ChildCategory []struct {
						Text       string `xml:",chardata"`
						CategoryID string `xml:"CategoryID"`
						Name       string `xml:"Name"`
						Order      string `xml:"Order"`
					} `xml:"ChildCategory"`
				} `xml:"ChildCategory"`
			} `xml:"CustomCategory"`
		} `xml:"CustomCategories"`
	} `xml:"Store"`
}


func (itemResp *GetItemResponse) ToFile(wg *sync.WaitGroup, fileDirectory string) error {
	defer wg.Done()
	file, err := xml.MarshalIndent(itemResp, "", " ")
	ioutil.WriteFile(fmt.Sprintf("%s/%s.xml", fileDirectory, itemResp.Item.ItemID), file, 0644)
	return err
}

type GetItemGroupResponse struct {
	CommonDescriptions []GIGRCommonDescriptions `json:"commonDescriptions"`
	Items              []GIGRItems              `json:"items"`
	Warnings           []GIGRWarnings           `json:"warnings"`
}
type GIGRCommonDescriptions struct {
	Description string   `json:"description"`
	ItemIds     []string `json:"itemIds"`
}
type GIGRAdditionalImages struct {
	Height   int    `json:"height"`
	ImageURL string `json:"imageUrl"`
	Width    int    `json:"width"`
}
type GIGRConstraint struct {
	ExpirationDate string `json:"expirationDate"`
}
type GIGRCouponDiscountAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}
type GIGRAvailableCoupons struct {
	Constraint     GIGRConstraint           `json:"constraint"`
	DiscountAmount GIGRCouponDiscountAmount `json:"discountAmount"`
	DiscountType   string                   `json:"discountType"`
	Message        string                   `json:"message"`
	RedemptionCode string                   `json:"redemptionCode"`
	TermsWebURL    string                   `json:"termsWebUrl"`
}
type GIGRCurrentBidPrice struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGREstimatedAvailabilities struct {
	AvailabilityThreshold       int      `json:"availabilityThreshold"`
	AvailabilityThresholdType   string   `json:"availabilityThresholdType"`
	DeliveryOptions             []string `json:"deliveryOptions"`
	EstimatedAvailabilityStatus string   `json:"estimatedAvailabilityStatus"`
	EstimatedAvailableQuantity  int      `json:"estimatedAvailableQuantity"`
	EstimatedSoldQuantity       int      `json:"estimatedSoldQuantity"`
}
type GIGRImage struct {
	Height   int    `json:"height"`
	ImageURL string `json:"imageUrl"`
	Width    int    `json:"width"`
}
type GIGRItemLocation struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	Country         string `json:"country"`
	County          string `json:"county"`
	PostalCode      string `json:"postalCode"`
	StateOrProvince string `json:"stateOrProvince"`
}
type GIGRLocalizedAspect struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
type GIGRDiscountAmount struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGROriginalPrice struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRMarketingPrice struct {
	DiscountAmount     GIGRDiscountAmount `json:"discountAmount"`
	DiscountPercentage string             `json:"discountPercentage"`
	OriginalPrice      GIGROriginalPrice  `json:"originalPrice"`
	PriceTreatment     string             `json:"priceTreatment"`
}
type GIGRMinimumPriceToBid struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRPrice struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRItemGroupAdditionalImages struct {
	Height   int    `json:"height"`
	ImageURL string `json:"imageUrl"`
	Width    int    `json:"width"`
}
type GIGRItemGroupImage struct {
	Height   int    `json:"height"`
	ImageURL string `json:"imageUrl"`
	Width    int    `json:"width"`
}
type GIGRPrimaryItemGroup struct {
	ItemGroupAdditionalImages []GIGRItemGroupAdditionalImages `json:"itemGroupAdditionalImages"`
	ItemGroupHref             string                          `json:"itemGroupHref"`
	ItemGroupID               string                          `json:"itemGroupId"`
	ItemGroupImage            GIGRItemGroupImage              `json:"itemGroupImage"`
	ItemGroupTitle            string                          `json:"itemGroupTitle"`
	ItemGroupType             string                          `json:"itemGroupType"`
}
type GIGRRatingHistograms struct {
	Count  int    `json:"count"`
	Rating string `json:"rating"`
}
type GIGRPrimaryProductReviewRating struct {
	AverageRating    string                 `json:"averageRating"`
	RatingHistograms []GIGRRatingHistograms `json:"ratingHistograms"`
	ReviewCount      int                    `json:"reviewCount"`
}
type GIGRProductIdentity struct {
	IdentifierType  string `json:"identifierType"`
	IdentifierValue string `json:"identifierValue"`
}
type GIGRAdditionalProductIdentities struct {
	ProductIdentity []GIGRProductIdentity `json:"productIdentity"`
}
type GIGRAspects struct {
	LocalizedName   string   `json:"localizedName"`
	LocalizedValues []string `json:"localizedValues"`
}
type GIGRAspectGroups struct {
	Aspects            []GIGRAspects `json:"aspects"`
	LocalizedGroupName string        `json:"localizedGroupName"`
}
type GIGRProduct struct {
	AdditionalImages            []GIGRAdditionalImages            `json:"additionalImages"`
	AdditionalProductIdentities []GIGRAdditionalProductIdentities `json:"additionalProductIdentities"`
	AspectGroups                []GIGRAspectGroups                `json:"aspectGroups"`
	Brand                       string                            `json:"brand"`
	Description                 string                            `json:"description"`
	Gtins                       []string                          `json:"gtins"`
	Image                       GIGRImage                         `json:"image"`
	Mpns                        []string                          `json:"mpns"`
	Title                       string                            `json:"title"`
}
type GIGRReturnPeriod struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}
type GIGRReturnTerms struct {
	ExtendedHolidayReturnsOffered bool             `json:"extendedHolidayReturnsOffered"`
	RefundMethod                  string           `json:"refundMethod"`
	RestockingFeePercentage       string           `json:"restockingFeePercentage"`
	ReturnInstructions            string           `json:"returnInstructions"`
	ReturnMethod                  string           `json:"returnMethod"`
	ReturnPeriod                  GIGRReturnPeriod `json:"returnPeriod"`
	ReturnsAccepted               bool             `json:"returnsAccepted"`
	ReturnShippingCostPayer       string           `json:"returnShippingCostPayer"`
}
type GIGRSellerProvidedLegalAddress struct {
	AddressLine1    string `json:"addressLine1"`
	AddressLine2    string `json:"addressLine2"`
	City            string `json:"city"`
	Country         string `json:"country"`
	CountryName     string `json:"countryName"`
	County          string `json:"county"`
	PostalCode      string `json:"postalCode"`
	StateOrProvince string `json:"stateOrProvince"`
}
type GIGRVatDetails struct {
	IssuingCountry string `json:"issuingCountry"`
	VatID          string `json:"vatId"`
}
type GIGRSellerLegalInfo struct {
	Email                      string                         `json:"email"`
	Fax                        string                         `json:"fax"`
	Imprint                    string                         `json:"imprint"`
	LegalContactFirstName      string                         `json:"legalContactFirstName"`
	LegalContactLastName       string                         `json:"legalContactLastName"`
	Name                       string                         `json:"name"`
	Phone                      string                         `json:"phone"`
	RegistrationNumber         string                         `json:"registrationNumber"`
	SellerProvidedLegalAddress GIGRSellerProvidedLegalAddress `json:"sellerProvidedLegalAddress"`
	TermsOfService             string                         `json:"termsOfService"`
	VatDetails                 []GIGRVatDetails               `json:"vatDetails"`
}
type GIGRSeller struct {
	FeedbackPercentage string              `json:"feedbackPercentage"`
	FeedbackScore      int                 `json:"feedbackScore"`
	SellerAccountType  string              `json:"sellerAccountType"`
	SellerLegalInfo    GIGRSellerLegalInfo `json:"sellerLegalInfo"`
	Username           string              `json:"username"`
}
type GIGRAdditionalShippingCostPerUnit struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRImportCharges struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRShippingCost struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRShipToLocationUsedForEstimate struct {
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
}
type GIGRShippingOptions struct {
	AdditionalShippingCostPerUnit GIGRAdditionalShippingCostPerUnit `json:"additionalShippingCostPerUnit"`
	CutOffDateUsedForEstimate     string                            `json:"cutOffDateUsedForEstimate"`
	FulfilledThrough              string                            `json:"fulfilledThrough"`
	GuaranteedDelivery            bool                              `json:"guaranteedDelivery"`
	ImportCharges                 GIGRImportCharges                 `json:"importCharges"`
	MaxEstimatedDeliveryDate      string                            `json:"maxEstimatedDeliveryDate"`
	MinEstimatedDeliveryDate      string                            `json:"minEstimatedDeliveryDate"`
	QuantityUsedForEstimate       int                               `json:"quantityUsedForEstimate"`
	ShippingCarrierCode           string                            `json:"shippingCarrierCode"`
	ShippingCost                  GIGRShippingCost                  `json:"shippingCost"`
	ShippingCostType              string                            `json:"shippingCostType"`
	ShippingServiceCode           string                            `json:"shippingServiceCode"`
	ShipToLocationUsedForEstimate GIGRShipToLocationUsedForEstimate `json:"shipToLocationUsedForEstimate"`
	TrademarkSymbol               string                            `json:"trademarkSymbol"`
	Type                          string                            `json:"type"`
}
type GIGRRegionExcluded struct {
	RegionName string `json:"regionName"`
	RegionType string `json:"regionType"`
}
type GIGRRegionIncluded struct {
	RegionName string `json:"regionName"`
	RegionType string `json:"regionType"`
}
type GIGRShipToLocations struct {
	RegionExcluded []GIGRRegionExcluded `json:"regionExcluded"`
	RegionIncluded []GIGRRegionIncluded `json:"regionIncluded"`
}
type GIGRRegion struct {
	RegionName string `json:"regionName"`
	RegionType string `json:"regionType"`
}
type GIGRTaxJurisdiction struct {
	Region            GIGRRegion `json:"region"`
	TaxJurisdictionID string     `json:"taxJurisdictionId"`
}
type GIGRTaxes struct {
	EbayCollectAndRemitTax   bool                `json:"ebayCollectAndRemitTax"`
	IncludedInPrice          bool                `json:"includedInPrice"`
	ShippingAndHandlingTaxed bool                `json:"shippingAndHandlingTaxed"`
	TaxJurisdiction          GIGRTaxJurisdiction `json:"taxJurisdiction"`
	TaxPercentage            string              `json:"taxPercentage"`
	TaxType                  string              `json:"taxType"`
}
type GIGRUnitPrice struct {
	ConvertedFromCurrency string `json:"convertedFromCurrency"`
	ConvertedFromValue    string `json:"convertedFromValue"`
	Currency              string `json:"currency"`
	Value                 string `json:"value"`
}
type GIGRParameters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type GIGRWarnings struct {
	Category     string           `json:"category"`
	Domain       string           `json:"domain"`
	ErrorID      int              `json:"errorId"`
	InputRefIds  []string         `json:"inputRefIds"`
	LongMessage  string           `json:"longMessage"`
	Message      string           `json:"message"`
	OutputRefIds []string         `json:"outputRefIds"`
	Parameters   []GIGRParameters `json:"parameters"`
	Subdomain    string           `json:"subdomain"`
}
type GIGRItems struct {
	AdditionalImages           []GIGRAdditionalImages         `json:"additionalImages"`
	AdultOnly                  bool                           `json:"adultOnly"`
	AgeGroup                   string                         `json:"ageGroup"`
	AvailableCoupons           []GIGRAvailableCoupons         `json:"availableCoupons"`
	BidCount                   int                            `json:"bidCount"`
	Brand                      string                         `json:"brand"`
	BuyingOptions              []string                       `json:"buyingOptions"`
	CategoryID                 string                         `json:"categoryId"`
	CategoryPath               string                         `json:"categoryPath"`
	Color                      string                         `json:"color"`
	Condition                  string                         `json:"condition"`
	ConditionID                string                         `json:"conditionId"`
	CurrentBidPrice            GIGRCurrentBidPrice            `json:"currentBidPrice"`
	Description                string                         `json:"description"`
	EligibleForInlineCheckout  bool                           `json:"eligibleForInlineCheckout"`
	EnabledForGuestCheckout    bool                           `json:"enabledForGuestCheckout"`
	EnergyEfficiencyClass      string                         `json:"energyEfficiencyClass"`
	Epid                       string                         `json:"epid"`
	EstimatedAvailabilities    []GIGREstimatedAvailabilities  `json:"estimatedAvailabilities"`
	Gender                     string                         `json:"gender"`
	Gtin                       string                         `json:"gtin"`
	Image                      GIGRImage                      `json:"image"`
	InferredEpid               string                         `json:"inferredEpid"`
	ItemAffiliateWebURL        string                         `json:"itemAffiliateWebUrl"`
	ItemEndDate                string                         `json:"itemEndDate"`
	ItemID                     string                         `json:"itemId"`
	ItemLocation               GIGRItemLocation               `json:"itemLocation"`
	ItemWebURL                 string                         `json:"itemWebUrl"`
	LegacyItemID               string                         `json:"legacyItemId"`
	LocalizedAspects           []GIGRLocalizedAspect          `json:"localizedAspects"`
	LotSize                    int                            `json:"lotSize"`
	MarketingPrice             GIGRMarketingPrice             `json:"marketingPrice"`
	Material                   string                         `json:"material"`
	MinimumPriceToBid          GIGRMinimumPriceToBid          `json:"minimumPriceToBid"`
	Mpn                        string                         `json:"mpn"`
	Pattern                    string                         `json:"pattern"`
	Price                      GIGRPrice                      `json:"price"`
	PriceDisplayCondition      string                         `json:"priceDisplayCondition"`
	PrimaryItemGroup           GIGRPrimaryItemGroup           `json:"primaryItemGroup"`
	PrimaryProductReviewRating GIGRPrimaryProductReviewRating `json:"primaryProductReviewRating"`
	Product                    GIGRProduct                    `json:"product"`
	ProductFicheWebURL         string                         `json:"productFicheWebUrl"`
	QualifiedPrograms          []string                       `json:"qualifiedPrograms"`
	QuantityLimitPerBuyer      int                            `json:"quantityLimitPerBuyer"`
	ReservePriceMet            bool                           `json:"reservePriceMet"`
	ReturnTerms                GIGRReturnTerms                `json:"returnTerms"`
	Seller                     GIGRSeller                     `json:"seller"`
	SellerItemRevision         string                         `json:"sellerItemRevision"`
	ShippingOptions            []GIGRShippingOptions          `json:"shippingOptions"`
	ShipToLocations            GIGRShipToLocations            `json:"shipToLocations"`
	ShortDescription           string                         `json:"shortDescription"`
	Size                       string                         `json:"size"`
	SizeSystem                 string                         `json:"sizeSystem"`
	SizeType                   string                         `json:"sizeType"`
	Subtitle                   string                         `json:"subtitle"`
	Taxes                      []GIGRTaxes                    `json:"taxes"`
	Title                      string                         `json:"title"`
	TopRatedBuyingExperience   bool                           `json:"topRatedBuyingExperience"`
	UniqueBidderCount          int                            `json:"uniqueBidderCount"`
	UnitPrice                  GIGRUnitPrice                  `json:"unitPrice"`
	UnitPricingMeasure         string                         `json:"unitPricingMeasure"`
	Warnings                   []GIGRWarnings                 `json:"warnings"`
}
