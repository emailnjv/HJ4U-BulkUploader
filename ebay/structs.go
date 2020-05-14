package ebay

import (
	"encoding/xml"
)

type GetItemRequest struct {
	XMLName       xml.Name `xml:"GetItemRequest"`
	Text          string   `xml:",chardata"`
	Xmlns         string   `xml:"xmlns,attr"`
	ErrorLanguage string   `xml:"ErrorLanguage"`
	WarningLevel  string   `xml:"WarningLevel"`
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
