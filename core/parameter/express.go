/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-12-07 13:15:29
 * @LastEditTime: 2022-12-13 11:13:06
 * @FilePath: \GoKwai\addons\open_kwai\core\parameter\express.go
 */
package parameter

type Express_GetEbillOrderRequest struct {
	GetEbillOrderRequest []Express_Ebill_Get `json:"getEbillOrderRequest"`
}

type Express_Ebill_Get struct {
	MerchantCode         string                       `json:"merchantCode"`
	TotalPackageQuantity int                          `json:"totalPackageQuantity"`
	PackageCode          string                       `json:"packageCode"`
	TemplateUrl          string                       `json:"templateUrl"`
	NetSiteCode          string                       `json:"netSiteCode"`
	NetSiteName          string                       `json:"netSiteName"`
	ItemList             []Express_Ebill_Get_ItemList `json:"itemList"`
	ExtData              string                       `json:"extData"`
	ReceiverContract     struct {
		Name   string `json:"name"`
		Mobile string `json:"mobile"`
	} `json:"receiverContract"`
	SenderContract struct {
		Name   string `json:"name"`
		Mobile string `json:"mobile"`
	} `json:"senderContract"`
	ExpressCompanyCode string `json:"expressCompanyCode"`
	OrderChannel       string `json:"orderChannel"`
	MerchantName       string `json:"merchantName"`
	TradeOrderCode     string `json:"tradeOrderCode"`
	SenderAddress      struct {
		ProvinceName  string `json:"provinceName"`
		CityName      string `json:"cityName"`
		DistrictName  string `json:"districtName"`
		DetailAddress string `json:"detailAddress"`
	} `json:"senderAddress"`
	ReceiverAddress struct {
		ProvinceName  string `json:"provinceName"`
		CityName      string `json:"cityName"`
		DistrictName  string `json:"districtName"`
		DetailAddress string `json:"detailAddress"`
	} `json:"receiverAddress"`
	SettleAccount *string `json:"settleAccount"`
	RequestID     string  `json:"requestId"`
}
type Express_Ebill_Get_ItemList struct {
	ItemTitle    string `json:"itemTitle"`
	ItemQuantity uint   `json:"itemQuantity"`
}

type Express_Ebill_WayBillData struct {
	EncryptedData string `json:"encryptedData"`
	Signature     string `json:"signature"`
	Key           string `json:"key"`
	AddData       struct {
		SenderInfo struct {
			Address struct {
				CityName      string `json:"cityName"`
				DetailAddress string `json:"detailAddress"`
				DistrictName  string `json:"districtName"`
				ProvinceName  string `json:"provinceName"`
			} `json:"address"`
			Contact struct {
				Mobile string `json:"mobile"`
				Name   string `json:"name"`
			} `json:"contact"`
		} `json:"senderInfo"`
	} `json:"addData"`
	TemplateURL string      `json:"templateURL"`
	Params      interface{} `json:"params"`
}
type Express_Ebill_WayBillInfo struct {
	TemplateURL string `json:"templateURL"`
	CustomData  struct {
		SInfo string `json:"sInfo"`
		/*
			ContactName        string      `json:"contactName"`
			ZipCode            interface{} `json:"zipCode"`
			Province           string      `json:"province"`
			City               string      `json:"city"`
			Country            string      `json:"country"`
			Addr               string      `json:"addr"`
			Address            string      `json:"address"`
			Phone              interface{} `json:"phone"`
			MobilePhone        string      `json:"mobilePhone"`
			MAndP              string      `json:"mAndP"`
			SellerCompany      interface{} `json:"sellerCompany"`
			ShortName          interface{} `json:"shortName"`
			Cus1               interface{} `json:"cus1"`
			Cus2               interface{} `json:"cus2"`
			Cus3               interface{} `json:"cus3"`
			ReceiverName       string      `json:"receiverName"`
			BuyerNick          string      `json:"buyerNick"`
			SellerNick         string      `json:"sellerNick"`
			ReceiverZip        interface{} `json:"receiverZip"`
			ReceiverState      string      `json:"receiverState"`
			ReceiverCity       string      `json:"receiverCity"`
			ReceiverDistrict   string      `json:"receiverDistrict"`
			CBillMark          string      `json:"cBillMark"`
			ReceiverAddress    string      `json:"receiverAddress"`
			FullAddress        string      `json:"fullAddress"`
			ReceiverPhone      interface{} `json:"receiverPhone"`
			ReceiverMobile     string      `json:"receiverMobile"`
			ReceiverMAndP      string      `json:"receiverMAndP"`
			PrintAbleNumStr    string      `json:"printAbleNumStr"`
			TotalWeightStr     string      `json:"totalWeightStr"`
			TotalVolumeStr     string      `json:"totalVolumeStr"`
			Title              string      `json:"title"`
			InvoiceName        interface{} `json:"invoiceName"`
			ExpressNo          string      `json:"expressNo"`
			AllMoneySum        int         `json:"allMoneySum"`
			AllMoneyCapitalSum string      `json:"allMoneyCapitalSum"`
			SimpleProv         string      `json:"simpleProv"`
			AllOrderID         string      `json:"allOrderId"`
			TradeID            string      `json:"tradeId"`
			FullSellerMemo     interface{} `json:"fullSellerMemo"`
			FullBuyerMessage   string      `json:"fullBuyerMessage"`
			PayTimeStr         string      `json:"payTimeStr"`
			Gift               interface{} `json:"gift"`
			TotalPaymentStr    string      `json:"totalPaymentStr"`
			ReceiverCompany    interface{} `json:"receiverCompany"`
			WaybillCode        string      `json:"waybillCode"`
			PackageCode        string      `json:"packageCode"`
			EnoBC              string      `json:"enoBC"`
			Billsite           interface{} `json:"billsite"`
			SortSite           interface{} `json:"sortSite"`
			SortCode           interface{} `json:"sortCode"`
			Sbc                interface{} `json:"sbc"`
			PkgName            interface{} `json:"pkgName"`
			FourSegmentCode    interface{} `json:"fourSegmentCode"`
			Pcc                interface{} `json:"pcc"`
			BillMark           interface{} `json:"billMark"`
			Custid             interface{} `json:"custid"`
			Transport          interface{} `json:"transport"`
			PayMethod          interface{} `json:"payMethod"`
			ExpressType        interface{} `json:"expressType"`
			ParentWaybillCode  interface{} `json:"parentWaybillCode"`
			MainNo             interface{} `json:"mainNo"`
			SourceSortCode     interface{} `json:"sourceSortCode"`
			TargetSortCode     interface{} `json:"targetSortCode"`
			Road               interface{} `json:"road"`
			SiteName           interface{} `json:"siteName"`
			SiteType           interface{} `json:"siteType"`
			PromiseTimeType    interface{} `json:"promiseTimeType"`
			CodMoney           interface{} `json:"codMoney"`
			Watermark          interface{} `json:"watermark"`
			PkgNo              interface{} `json:"pkgNo"`
			FromBranchName     interface{} `json:"fromBranchName"`
			ToBranchName       interface{} `json:"toBranchName"`
			OrderSign          interface{} `json:"orderSign"`
			Cargo              interface{} `json:"cargo"`
			CoverCode          interface{} `json:"coverCode"`
			SfSegmentCode      interface{} `json:"sfSegmentCode"`
			FwSegmentCode      interface{} `json:"fwSegmentCode"`
			YtoYzd             interface{} `json:"ytoYzd"`
			ZtoYzd             interface{} `json:"ztoYzd"`
			YundaYzd           interface{} `json:"yundaYzd"`
			CodingMapping      interface{} `json:"codingMapping"`
			CodingMappingOut   interface{} `json:"codingMappingOut"`
			FreightPayMoney    interface{} `json:"freightPayMoney"`
			PrintAbleItems     []struct {
				ID                 interface{} `json:"id"`
				CreateTime         interface{} `json:"createTime"`
				LastUpdateTime     interface{} `json:"lastUpdateTime"`
				Price              string      `json:"price"`
				TotalFee           string      `json:"totalFee"`
				DiscountFee        string      `json:"discountFee"`
				AdjustFee          interface{} `json:"adjustFee"`
				Payment            string      `json:"payment"`
				Num                int         `json:"num"`
				NumIid             string      `json:"numIid"`
				SkuID              string      `json:"skuId"`
				OuterIid           interface{} `json:"outerIid"`
				OuterSkuID         string      `json:"outerSkuId"`
				Title              string      `json:"title"`
				Status             interface{} `json:"status"`
				RefundStatus       interface{} `json:"refundStatus"`
				SkuPropertiesName  string      `json:"skuPropertiesName"`
				StoreCode          interface{} `json:"storeCode"`
				PicPath            string      `json:"picPath"`
				Tid                string      `json:"tid"`
				IsDaixiao          interface{} `json:"isDaixiao"`
				LogisticsCompany   interface{} `json:"logisticsCompany"`
				InvoiceNo          interface{} `json:"invoiceNo"`
				OrderAttr          string      `json:"orderAttr"`
				PromiseCollectTime interface{} `json:"promiseCollectTime"`
				ShortName          interface{} `json:"shortName"`
				ItemWeight         interface{} `json:"itemWeight"`
				ItemVolume         interface{} `json:"itemVolume"`
				OrderID            string      `json:"orderId"`
				OrderShortName     interface{} `json:"orderShortName"`
				MarketName         interface{} `json:"marketName"`
				GateName           interface{} `json:"gateName"`
				CanSendGoods       interface{} `json:"canSendGoods"`
				CantSendReason     interface{} `json:"cantSendReason"`
				Gift               bool        `json:"gift"`
				Cost               interface{} `json:"cost"`
				FilteredSkuPN      []string    `json:"filteredSkuPN"`
				SkuIdx             string      `json:"skuIdx"`
				ShowName           string      `json:"showName"`
			} `json:"printAbleItems"`
			CurrentTotalTradePage string `json:"currentTotalTradePage"`
			PrintTimes            string `json:"printTimes"`
			PrintDate             string `json:"printDate"`
			PrintTime             string `json:"printTime"`
		*/
	} `json:"customData"`
}

type Express_Ebill_ReceiveInfo struct {
	Address              string `json:"address"`
	Consignee            string `json:"consignee"`
	DistrictCode         int    `json:"districtCode"`
	City                 string `json:"city"`
	CityCode             int    `json:"cityCode"`
	ProvinceCode         int    `json:"provinceCode"`
	Mobile               string `json:"mobile"`
	EncryptedMobile      string `json:"encryptedMobile"`
	EncryptedConsignee   string `json:"encryptedConsignee"`
	DesensitiseConsignee string `json:"desensitiseConsignee"`
	EncryptedAddress     string `json:"encryptedAddress"`
	Province             string `json:"province"`
	District             string `json:"district"`
	DesensitiseMobile    string `json:"desensitiseMobile"`
	DesensitiseAddress   string `json:"desensitiseAddress"`
}
