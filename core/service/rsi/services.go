/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2013-12-03 23:20
 * description :
 * history :
 */

package rsi

import (
	"github.com/jsix/gof"
	"go2o/core/dao"
	"go2o/core/infrastructure/domain"
	"go2o/core/query"
	"go2o/core/repository"
)

var (
	PromService *promotionService
	// 基础服务
	FoundationService *foundationService
	// 会员服务
	MemberService *memberService
	// 商户服务
	MerchantService *merchantService
	// 商店服务
	ShopService *shopService
	// 产品服务
	ProductService *productService
	// 商品服务
	ItemService *itemService
	// 购物服务
	ShoppingService *shoppingService
	// 售后服务
	AfterSalesService *afterSalesService
	// 支付服务
	PaymentService *paymentService
	// 消息服务
	MssService *mssService
	// 快递服务
	ExpressService *expressService
	// 配送服务
	ShipmentService *shipmentService
	// 内容服务
	ContentService *contentService
	// 广告服务
	AdService *adService

	// 个人金融服务
	PersonFinanceService *personFinanceService
	// 门户数据服务
	PortalService *portalService

	CommonDao *dao.CommonDao
)

// 处理错误
func handleError(err error) error {
	return domain.HandleError(err, "service")
	//if err != nil && gof.CurrentApp.Debug() {
	//	gof.CurrentApp.Log().Println("[ Go2o][ Rep][ Error] -", err.Error())
	//}
	//return err
}

func Init(ctx gof.App) {
	Context := ctx
	db := Context.Db()
	orm := db.GetOrm()
	sto := Context.Storage()

	/** Repository **/
	proMRepo := repository.NewProModelRepo(db, orm)
	valueRepo := repository.NewValueRepo(db, sto)
	userRepo := repository.NewUserRepo(db)
	notifyRepo := repository.NewNotifyRepo(db)
	mssRepo := repository.NewMssRepo(db, notifyRepo, valueRepo)
	expressRepo := repository.NewExpressRepo(db, valueRepo)
	shipRepo := repository.NewShipmentRepo(db, expressRepo)
	memberRepo := repository.NewMemberRepo(sto, db, mssRepo, valueRepo)
	productRepo := repository.NewProductRepo(db, proMRepo, valueRepo)
	itemWsRepo := repository.NewItemWholesaleRepo(db)
	itemRepo := repository.NewGoodsItemRepo(db, productRepo,
		proMRepo, itemWsRepo, expressRepo, valueRepo)
	tagSaleRepo := repository.NewTagSaleRepo(db, valueRepo)
	promRepo := repository.NewPromotionRepo(db, itemRepo, memberRepo)
	catRepo := repository.NewCategoryRepo(db, valueRepo, sto)
	//afterSalesRepo := repository.NewAfterSalesRepo(db)

	shopRepo := repository.NewShopRepo(db, sto)
	wholesaleRepo := repository.NewWholesaleRepo(db)
	mchRepo := repository.NewMerchantRepo(db, sto, wholesaleRepo, shopRepo, userRepo, memberRepo, mssRepo, valueRepo)
	cartRepo := repository.NewCartRepo(db, memberRepo, mchRepo, itemRepo)
	personFinanceRepo := repository.NewPersonFinanceRepository(db, memberRepo)
	deliveryRepo := repository.NewDeliverRepo(db)
	contentRepo := repository.NewContentRepo(db)
	adRepo := repository.NewAdvertisementRepo(db, sto)
	orderRepo := repository.NewOrderRepo(sto, db, mchRepo, nil, productRepo, cartRepo, itemRepo,
		promRepo, memberRepo, deliveryRepo, expressRepo, shipRepo, valueRepo)
	paymentRepo := repository.NewPaymentRepo(sto, db, memberRepo, orderRepo, valueRepo)
	asRepo := repository.NewAfterSalesRepo(db, orderRepo, memberRepo, paymentRepo)

	orderRepo.SetPaymentRepo(paymentRepo)

	/* 初始化数据 */
	memberRepo.GetManager().GetAllBuyerGroups()

	/** Query **/
	memberQue := query.NewMemberQuery(db)
	mchQuery := query.NewMerchantQuery(ctx)
	contentQue := query.NewContentQuery(db)
	goodsQuery := query.NewItemQuery(db)
	shopQuery := query.NewShopQuery(ctx)
	orderQuery := query.NewOrderQuery(db)
	afterSalesQuery := query.NewAfterSalesQuery(db)

	/** Service **/
	ProductService = NewProService(proMRepo, catRepo, productRepo)
	FoundationService = NewFoundationService(valueRepo)
	PromService = NewPromotionService(promRepo)
	ShoppingService = NewShoppingService(orderRepo, cartRepo,
		productRepo, itemRepo, mchRepo, orderQuery)
	AfterSalesService = NewAfterSalesService(asRepo, afterSalesQuery, orderRepo)
	MerchantService = NewMerchantService(mchRepo, memberRepo, mchQuery, orderQuery)
	ShopService = NewShopService(shopRepo, mchRepo, shopQuery)
	MemberService = NewMemberService(MerchantService, memberRepo, memberQue, orderQuery, valueRepo)
	ItemService = NewSaleService(catRepo, itemRepo, goodsQuery, tagSaleRepo, proMRepo, mchRepo, valueRepo)
	PaymentService = NewPaymentService(paymentRepo, orderRepo)
	MssService = NewMssService(mssRepo)
	ExpressService = NewExpressService(expressRepo)
	ShipmentService = NewShipmentService(shipRepo, deliveryRepo)
	ContentService = NewContentService(contentRepo, contentQue)
	AdService = NewAdvertisementService(adRepo, sto)
	PersonFinanceService = NewPersonFinanceService(personFinanceRepo, memberRepo)

	//m := memberRepo.GetMember(1)
	//d := m.ProfileManager().GetDeliverAddress()[0]
	//v := d.GetValue()
	//v.Province = 440000
	//v.City = 440600
	//v.District = 440605
	//d.SetValue(&v)
	//d.Save()

	CommonDao = dao.NewCommDao(orm, sto, adRepo, catRepo)
	PortalService = NewPortalService(CommonDao)
}
