package usecase

import (
	"app/config"
	"app/domain"
	"app/infrastructure/persistence"
)

type Shop struct {
	Id uint `json:"id"`
	Name string `json:"name"`
}

//func CreateNewShop (c echo.Context) error {
//	param := new(Shop)
//	if err := c.Bind(param); err != nil {
//		log.Println("bad request")
//		bytes, err := json.Marshal(param)
//		if err != nil {
//			fmt.Println("JSON marshal error: ", err)
//			return nil
//		}
//		return c.String(http.StatusBadRequest, string(bytes))
//	}
//	insertRecord := &domain.Shop{ShopId: param.ShopId, Name:param.Name, Price: param.Price}
//	log.Println(insertRecord)
//	conn, err := config.ConnectMySql()
//	if err != nil {
//		return nil
//	}
//	defer conn.Close()
//	repo := persistence.ShopRepositoryWithRDB(conn)
//	bytes, err := json.Marshal(repo.Create(insertRecord))
//	if err != nil {
//		fmt.Println("JSON marshal error: ", err)
//		return nil
//	}
//	return c.String(http.StatusCreated, string(bytes))
//}

/**
 * ショップ情報を取得する
 */
func FetchShopInfo(shopId int) domain.Shop {
	conn, err := config.ConnectMySql()
	if err != nil {
		return domain.Shop{}
	}
	defer conn.Close()
	repo := persistence.ShopRepositoryWithRDB(conn)
	res := repo.GetByShopId(shopId)
	return res
}

