package services

import (
	"database/sql"
	"investify/api/types"
	"investify/api/types/errors"
	db "investify/db/sqlc"
	"investify/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BusinessService interface {
	CreateBusinessService(ctx *gin.Context, req types.CreateBusinessRequest) (types.CreateBusinessResponse, error)
	GetBusinessService(ctx *gin.Context) (types.GetBusinessResponse, error)
	GetBusinessServiceByOwner(ctx *gin.Context) (types.GetBusinessFeedResponse, error)
	GetInvestorFeedService(ctx *gin.Context) (types.GetInvestorFeedResponse, error)
}

type BusinessServiceImpl struct {
	store db.Store
}

func NewBusinessService(store db.Store) BusinessService {
	return &BusinessServiceImpl{store: store}
}

func (b *BusinessServiceImpl) CreateBusinessService(ctx *gin.Context, req types.CreateBusinessRequest) (types.CreateBusinessResponse, error) {
	//initiate the databse Trasaction
	//cnnect the authentication middleware  with this where role to create business should be the owner
	//Extract the user_id form the acess token
	//check the existance of that userId
	//get owner object with the user id
	//inser the adress the into the dtabse
	//extarct the adress id and insert it in the Business table databse
	//transaction done
	var respObject types.CreateBusinessResponse

	err := b.store.ExecTx(ctx, func(tx *db.Queries) error {

		user, err := util.CurrentUser(ctx, b.store)
		if err != nil {
			return errors.ErrUserNotFound
		}

		owner, err := tx.GetOwnerByUserId(ctx, user.UserID)

		if err != nil {
			return errors.ErrCreateOwner
		}

		address, err := tx.CreateAddress(ctx, db.CreateAddressParams{
			AddressStreet:  req.AdressDetails.AddressStreet,
			AddressCity:    req.AdressDetails.AddressCity,
			AddressState:   req.AdressDetails.AddressState,
			AddressCountry: req.AdressDetails.AddressCountry,
			AddressZipcode: req.AdressDetails.AddressZipcode,
		})
		if err != nil {
			return errors.ErrCreateAddress
		}
		respObject.AddressInfo = address
		business, err := tx.CreateBusiness(ctx, db.CreateBusinessParams{
			BusinessOwnerID:        owner.OwnerID,
			BusinessOwnerFirstname: req.BusinessDetails.BusinessOwnerFirstname,
			BusinessOwnerLastname:  req.BusinessDetails.BusinessOwnerLastname,
			BusinessEmail:          req.BusinessDetails.BusinessEmail,
			BusinessName:           req.BusinessDetails.BusinessName,
			BusinessContact:        req.BusinessDetails.BusinessContact,
			BusinessAddressID:      address.AddressID,
			BusinessRatings:        req.BusinessDetails.BusinessRatings,
			BusinessMinamount:      req.BusinessDetails.BusinessMinamount,
		})

		if err != nil {
			return errors.ErrCreateBusiness
		}
		respObject.BusinessInfo = business

		return nil //commit transaction
	})

	if err != nil {
		return types.CreateBusinessResponse{}, err
	}

	return respObject, nil
}
func (b *BusinessServiceImpl) GetBusinessService(ctx *gin.Context) (types.GetBusinessResponse, error) {
	idstr := ctx.Param("id")
	// Convert the string ID to int64
	log.Println(idstr)
	var respObject types.GetBusinessResponse
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return types.GetBusinessResponse{}, errors.ErrInvalidID
	}
	business, err := b.store.GetBusinessById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetBusinessResponse{}, errors.ErrGetBusiness
		}
		return types.GetBusinessResponse{}, errors.ErrGetBusiness
	}
	address, err := b.store.GetAddressById(ctx, business.BusinessAddressID)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetBusinessResponse{}, errors.ErrGetBusiness
		}
		return types.GetBusinessResponse{}, errors.ErrGetBusiness
	}
	// b.store.get
	respObject.BusinessInfo = business
	respObject.AddressInfo = address
	return respObject, nil
}

// asingle owner can have mulltiple business
func (b *BusinessServiceImpl) GetBusinessServiceByOwner(ctx *gin.Context) (types.GetBusinessFeedResponse, error) {
	user, err := util.CurrentUser(ctx, b.store)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrUserNotFound
	}
	owner, err := b.store.GetOwnerByUserId(ctx, user.UserID)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrGetBusinessByOwner
	}

	var respObject types.GetBusinessFeedResponse
	business, err := b.store.GetBusinessByOwnerId(ctx, owner.OwnerID)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrGetBusinessByOwner
	}
	for _, elem := range business {
		respObject.BusinessInfo = append(respObject.BusinessInfo, elem)
	}
	return respObject, nil
}
func (b *BusinessServiceImpl) GetInvestorFeedService(ctx *gin.Context) (types.GetInvestorFeedResponse, error) {

	var respObject types.GetInvestorFeedResponse
	investors, err := b.store.GetInvestorFeed(ctx)
	if err != nil {
		return types.GetInvestorFeedResponse{}, errors.ErrGetInvestorFeed
	}
	//filter what to send in the feed
	for _, elem := range investors {
		respObject.InvestorInfo = append(respObject.InvestorInfo, elem)
	}
	return respObject, nil
}
