package autorestapi

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func GenAutoApi(r *gin.Engine, prefix string, model Model) {
	GenAutoGET(r, prefix, model)
	GenAutoPOST(r, prefix, model)
	GenAutoPUT(r, prefix, model)
	GenAutoPATCH(r, prefix, model)
	GenAutoDELETE(r, prefix, model)
}

func GenAutoGET(r *gin.Engine, prefix string, model Model) {
	api, err := GenMethodApi(GET, model)
	if err != nil {
		panic(err)
	}
	group := genGroup(r, prefix, model)
	group.GET("/", api)
}

func GenAutoPOST(r *gin.Engine, prefix string, model Model) {
	api, err := GenMethodApi(POST, model)
	if err != nil {
		panic(err)
	}
	group := genGroup(r, prefix, model)
	group.POST("/", api)
}

func GenAutoPUT(r *gin.Engine, prefix string, model Model) {
	api, err := GenMethodApi(PUT, model)
	if err != nil {
		panic(err)
	}
	group := genGroup(r, prefix, model)
	group.PUT("/", api)
}

func GenAutoPATCH(r *gin.Engine, prefix string, model Model) {
	api, err := GenMethodApi(PATCH, model)
	if err != nil {
		panic(err)
	}
	group := genGroup(r, prefix, model)
	group.PATCH("/", api)
}

func GenAutoDELETE(r *gin.Engine, prefix string, model Model) {
	api, err := GenMethodApi(DELETE, model)
	if err != nil {
		panic(err)
	}
	group := genGroup(r, prefix, model)
	group.DELETE("/", api)
}

func GenMethodApi(method int, model Model) (func(*gin.Context), error) {
	return GenApi(method, model)
}

func GenApi(method int, model Model) (api func(*gin.Context), err error) {
	if method < METHOD_MIN || method > METHOD_MAX {
		err = fmt.Errorf("method have to in (%d, %d)", METHOD_MIN, METHOD_MAX)
		return
	}
	switch method {
	case GET:
		api = func(c *gin.Context) {
			var vPageIndex, vPageSize, sorted string
			query := make(map[string]string)
			get_query := c.Request.URL.Query()
			for key, values := range get_query {
				if len(values) < 1 {
					continue
				}
				switch key {
				case "page_index":
					vPageIndex = values[0]
				case "page_size":
					vPageSize = values[0]
				case "sorted":
					sorted = values[0]
				default:
					query[key] = values[0]
				}
			}
			vPageIndex = c.Query("page_index")
			vPageSize = c.Query("page_size")
			iPageIndex, _ := strconv.Atoi(vPageIndex)
			iPageSize, _ := strconv.Atoi(vPageSize)
			ok := checkPageSizeAndPageIndex(iPageSize, iPageIndex)
			if !ok {
				errMsg := "page_index 和 page_size 必须大于 0"
				RespErr(c, errMsg)
				return
			}
			result, count, err := FindPage(model, query, iPageSize, iPageIndex, sorted)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			RespListData(c, result, count)
		}
	case POST:
		api = func(c *gin.Context) {
			m, err := model.New()
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			c.BindJSON(m)
			err = m.Check()
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			err = Save(m)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			RespData(c, nil)
		}
	case PUT:
		api = func(c *gin.Context) {
			m, err := model.New()
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			c.BindJSON(m)
			err = m.Check()
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			_id := c.PostForm("_id")
			id := bson.ObjectId(_id)
			query := bson.M{"_id": id}
			om, err := FindOne(model, query)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			update := bson.M{"$set": om}
			if om != nil {
				err = UpdateOne(model, query, update)
			}
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			RespData(c, nil)
			return
		}
	case PATCH:
		api = func(c *gin.Context) {
			data := make(map[string]interface{})
			err := c.BindJSON(data)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			count, err := Count(model, data["query"])
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			if count <= 0 {
				RespData(c, nil)
				return
			}
			update := map[string]interface{}{
				"$set": data["update"],
			}
			err = UpdateOne(model, data["query"], update)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			RespData(c, nil)
			return
		}
	case DELETE:
		api = func(c *gin.Context) {
			query := bson.M{}
			c.BindJSON(query)
			err := DeleteAll(model, query)
			if err != nil {
				errMsg := fmt.Sprintf("%v", err)
				RespErr(c, errMsg)
				return
			}
			RespData(c, nil)
			return
		}
	}
	return
}

func genGroup(r *gin.Engine, prefix string, model Model) *gin.RouterGroup {
	modelName := getStructName(model)
	url := genUrl(prefix, modelName)
	group := r.Group(url)
	return group
}

func genUrl(prefix, name string) string {
	if prefix == "" {
		return name
	}
	return fmt.Sprintf("%s/%s", prefix, name)
}
