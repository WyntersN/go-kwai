/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2022-11-08 13:17:31
 * @LastEditTime: 2022-11-22 11:16:26
 * @FilePath: \GoKwai\addons\open_kwai\core\result\order.go
 */
package result

type Order_CursorList struct {
	Result    int                  `json:"result"`
	Msg       string               `json:"msg"`
	ErrorMsg  string               `json:"error_msg"`
	Code      string               `json:"code"`
	Data      Order_CursorListData `json:"data"`
	RequestID string               `json:"requestId"`
	SubMsg    string               `json:"sub_msg"`
	SubCode   string               `json:"sub_code"`
}

type Order_CursorListData struct {
	Cursor    string `json:"cursor"`
	OrderList []struct {
		OrderBaseInfo struct {
			DiscountFee                   int    `json:"discountFee"`
			BuyerNick                     string `json:"buyerNick"`
			PayTime                       int64  `json:"payTime"`
			Remark                        string `json:"remark"`
			Oid                           int64  `json:"oid"`
			SellerOpenID                  string `json:"sellerOpenId"`
			ExpressFee                    int    `json:"expressFee"`
			BuyerImage                    string `json:"buyerImage"`
			PayType                       int    `json:"payType"`
			ValidPromiseShipmentTimeStamp int64  `json:"validPromiseShipmentTimeStamp"`
			SellerNick                    string `json:"sellerNick"`
			RecvTime                      int64  `json:"recvTime"`
			BuyerOpenID                   string `json:"buyerOpenId"`
			CpsType                       int    `json:"cpsType"`
			PromiseTimeStampOfDelivery    int    `json:"promiseTimeStampOfDelivery"`
			RefundTime                    int    `json:"refundTime"`
			RiskCode                      int    `json:"riskCode"`
			UpdateTime                    int64  `json:"updateTime"`
			TheDayOfDeliverGoodsTime      int    `json:"theDayOfDeliverGoodsTime"`
			CommentStatus                 int    `json:"commentStatus"`
			SendTime                      int64  `json:"sendTime"`
			PreSale                       int    `json:"preSale"`
			CoType                        int    `json:"coType"`
			CreateTime                    int64  `json:"createTime"`
			TotalFee                      int    `json:"totalFee"`
			SellerDelayPromiseTimeStamp   int64  `json:"sellerDelayPromiseTimeStamp"`
			PayChannel                    string `json:"payChannel"`
			ActivityType                  int    `json:"activityType"`
			Status                        int    `json:"status"`
		} `json:"orderBaseInfo"`
		OrderRefundList []interface{} `json:"orderRefundList"`
		OrderAddress    struct {
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
		} `json:"orderAddress"`
		OrderLogisticsInfo []struct {
			ExpressNo   string `json:"expressNo"`
			ExpressCode int    `json:"expressCode"`
		} `json:"orderLogisticsInfo"`
		OrderItemInfo struct {
			ItemPicURL    string `json:"itemPicUrl"`
			ItemType      int    `json:"itemType"`
			DiscountFee   int    `json:"discountFee"`
			OriginalPrice int    `json:"originalPrice"`
			ItemTitle     string `json:"itemTitle"`
			OrderItemID   int64  `json:"orderItemId"`
			Num           int    `json:"num"`
			ItemExtra     struct {
				CategoryInfo struct {
					ItemCid      int    `json:"itemCid"`
					CategoryName string `json:"categoryName"`
				} `json:"categoryInfo"`
			} `json:"itemExtra"`
			WarehouseCode string `json:"warehouseCode"`
			ItemID        int64  `json:"itemId"`
			RelItemID     int64  `json:"relItemId"`
			RelSkuID      int    `json:"relSkuId"`
			Price         int    `json:"price"`
			ItemLinkURL   string `json:"itemLinkUrl"`
			SkuNick       string `json:"skuNick"`
			SkuDesc       string `json:"skuDesc"`
			GoodsCode     string `json:"goodsCode"`
			SkuID         int64  `json:"skuId"`
		} `json:"orderItemInfo"`
		OrderNote struct {
			SellerNote []interface{} `json:"sellerNote"`
		} `json:"orderNote"`
		OrderCpsInfo struct {
			DistributorName string `json:"distributorName"`
			DistributorID   int    `json:"distributorId"`
		} `json:"orderCpsInfo"`
	} `json:"orderList"`
	PageSize  int   `json:"pageSize"`
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
}

type Order_Detail struct {
	Result    int              `json:"result"`
	Msg       string           `json:"msg"`
	ErrorMsg  string           `json:"error_msg"`
	Code      string           `json:"code"`
	Data      Order_DetailData `json:"data"`
	RequestID string           `json:"requestId"`
	SubMsg    string           `json:"sub_msg"`
	SubCode   string           `json:"sub_code"`
}

type Order_DetailData struct {
	AddressUpdateAuditInfo struct {
		TimeoutOfAudit int `json:"timeoutOfAudit"`
		AuditTime      int `json:"auditTime"`
		AuditStatus    int `json:"auditStatus"`
	} `json:"addressUpdateAuditInfo"`
	OrderBaseInfo struct {
		DiscountFee                   int    `json:"discountFee"`
		BuyerNick                     string `json:"buyerNick"`
		PayTime                       int64  `json:"payTime"`
		Remark                        string `json:"remark"`
		Oid                           int64  `json:"oid"`
		SellerOpenID                  string `json:"sellerOpenId"`
		ExpressFee                    int    `json:"expressFee"`
		BuyerImage                    string `json:"buyerImage"`
		ShopNewBuyer                  bool   `json:"shopNewBuyer"`
		PayType                       int    `json:"payType"`
		ValidPromiseShipmentTimeStamp int64  `json:"validPromiseShipmentTimeStamp"`
		SellerNick                    string `json:"sellerNick"`
		RecvTime                      int    `json:"recvTime"`
		BuyerOpenID                   string `json:"buyerOpenId"`
		CpsType                       int    `json:"cpsType"`
		PromiseTimeStampOfDelivery    int    `json:"promiseTimeStampOfDelivery"`
		RefundTime                    int    `json:"refundTime"`
		CarrierType                   int    `json:"carrierType"`
		OrderTaxInfo                  struct {
			SellerBearAmount int `json:"sellerBearAmount"`
			TaxAmount        int `json:"taxAmount"`
		} `json:"orderTaxInfo"`
		PlatformNewBuyer            bool   `json:"platformNewBuyer"`
		RiskCode                    int    `json:"riskCode"`
		UpdateTime                  int64  `json:"updateTime"`
		TheDayOfDeliverGoodsTime    int    `json:"theDayOfDeliverGoodsTime"`
		CommentStatus               int    `json:"commentStatus"`
		SendTime                    int64  `json:"sendTime"`
		PreSale                     int    `json:"preSale"`
		CoType                      int    `json:"coType"`
		CreateTime                  int64  `json:"createTime"`
		TotalFee                    int    `json:"totalFee"`
		SellerDelayPromiseTimeStamp int64  `json:"sellerDelayPromiseTimeStamp"`
		PayChannel                  string `json:"payChannel"`
		ActivityType                int    `json:"activityType"`
		Status                      int    `json:"status"`
	} `json:"orderBaseInfo"`
	OrderRefundList []interface{} `json:"orderRefundList"`
	OrderAddress    struct {
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
	} `json:"orderAddress"`
	OrderLogisticsInfo []struct {
		LogisticsID int64  `json:"logisticsId"`
		ExpressNo   string `json:"expressNo"`
		ExpressCode int    `json:"expressCode"`
	} `json:"orderLogisticsInfo"`
	OrderItemInfo struct {
		ItemPicURL    string `json:"itemPicUrl"`
		ItemType      int    `json:"itemType"`
		DiscountFee   int    `json:"discountFee"`
		OriginalPrice int    `json:"originalPrice"`
		ItemTitle     string `json:"itemTitle"`
		OrderItemID   int64  `json:"orderItemId"`
		Num           int    `json:"num"`
		ItemExtra     struct {
			CategoryInfo struct {
				ItemCid      int    `json:"itemCid"`
				CategoryName string `json:"categoryName"`
			} `json:"categoryInfo"`
		} `json:"itemExtra"`
		ServiceInfo struct {
			ServiceRule         string `json:"serviceRule"`
			Freight             bool   `json:"freight"`
			FreightProviderType int    `json:"freightProviderType"`
			FirstOrderGuarantee bool   `json:"firstOrderGuarantee"`
			CompensateFake      struct {
				Times int    `json:"times"`
				Desc  string `json:"desc"`
				Link  string `json:"link"`
			} `json:"compensateFake"`
			InstantDelivery int  `json:"instantDelivery"`
			InstantRefund   bool `json:"instantRefund"`
			ServiceRuleInfo struct {
				RefundRule               string `json:"refundRule"`
				DepositRule              string `json:"depositRule"`
				DeliverGoodsIntervalTime int    `json:"deliverGoodsIntervalTime"`
				TheDayOfDeliverGoodsTime int    `json:"theDayOfDeliverGoodsTime"`
				SaleFlag                 bool   `json:"saleFlag"`
				PromiseDeliveryTime      int    `json:"promiseDeliveryTime"`
				ImmediatelyOnOfflineFlag int    `json:"immediatelyOnOfflineFlag"`
				DeliveryMethod           string `json:"deliveryMethod"`
				SupportVerification      int    `json:"supportVerification"`
				CertMerchantCode         string `json:"certMerchantCode"`
				CertExpireType           int    `json:"certExpireType"`
				CertStartTime            int    `json:"certStartTime"`
				CertEndTime              int    `json:"certEndTime"`
				CertExpDays              int    `json:"certExpDays"`
			} `json:"serviceRuleInfo"`
			FreightProviderTypeDesc    string `json:"freightProviderTypeDesc"`
			ShowInsuranceRefuseTag     bool   `json:"showInsuranceRefuseTag"`
			InsuranceRefuseReason      string `json:"insuranceRefuseReason"`
			FreightProviderExplainDesc string `json:"freightProviderExplainDesc"`
		} `json:"serviceInfo"`
		GoodStoreCode int    `json:"goodStoreCode"`
		WarehouseCode string `json:"warehouseCode"`
		ItemID        int64  `json:"itemId"`
		RelItemID     int64  `json:"relItemId"`
		RelSkuID      int    `json:"relSkuId"`
		Price         int    `json:"price"`
		ItemLinkURL   string `json:"itemLinkUrl"`
		SkuNick       string `json:"skuNick"`
		SkuDesc       string `json:"skuDesc"`
		GoodsCode     string `json:"goodsCode"`
		SkuID         int64  `json:"skuId"`
	} `json:"orderItemInfo"`
	OrderNote struct {
		SellerNote []interface{} `json:"sellerNote"`
	} `json:"orderNote"`
	OrderDeliveryInfo struct {
		EnableAppendPackage bool `json:"enableAppendPackage"`
		MaxPackageNum       int  `json:"maxPackageNum"`
	} `json:"orderDeliveryInfo"`
	OrderFlag struct {
		Flag []string `json:"flag"`
	} `json:"orderFlag"`
	OrderCpsInfo struct {
		DistributorName string `json:"distributorName"`
		DistributorID   int    `json:"distributorId"`
	} `json:"orderCpsInfo"`
}
