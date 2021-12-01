package api

//1. 场馆列表
//type ChamberListReq struct {
//	SchoolId    dto.ID `form:"schoolId" binding:"required"`
//	ChamberId   dto.ID `json:"chamberId" form:"chamberId"` // 场馆id
//	Capacity    int    `json:"capacity" form:"capacity"`   // 场馆容量大于这个值
//	ChamberName string `form:"chamberName"`                // 场馆名称
//	Status      int8   `form:"status"`                     // 场馆状态
//	PageNum     int    `form:"pageNum"`
//	PageSize    int    `form:"pageSize"`
//}
//type ChamberListResp struct {
//	Total    int       `json:"total"`
//	PageNum  int       `json:"pageNum"`
//	PageSize int       `json:"pageSize"`
//	List     []Chamber `json:"list"`
//}
//type Chamber struct {
//	ChamberId    dto.ID `json:"chamberId" form:"chamberId"`       // 场馆id
//	ChamberName  string `json:"chamberName" form:"chamberName"`   // 场馆名称
//	Capacity     int    `json:"capacity" form:"capacity"`         // 场馆容量
//	Status       int8   `json:"status" form:"status"`             // 场馆状态
//	StatusName   string `json:"statusName" form:"statusName"`     // 场馆状态名称
//	Remark       string `json:"remark" form:"remark"`             // 场馆描述
//	BuildingName string `json:"buildingName" form:"buildingName"` // 建筑名称
//	Floor        string `json:"floor" form:"floor"`               // 楼层
//	HouseNumber  string `json:"houseNumber" form:"houseNumber"`   // 门牌号
//}
//
//// 2. 删除场馆
//type DelChamberReq struct {
//	ChamberId dto.ID `json:"chamberId"` // 场馆id
//}
//type DelChamberResp struct {
//}
//
//// 3. 批量导入场馆
//type BatchAddChamberReq struct {
//}
//type BatchAddChamberResp struct {
//	Result     bool             `json:"result"`     // true: 场馆，false: 失败
//	SuccessCnt int              `json:"successCnt"` // 成功数量
//	FailCnt    int              `json:"failCnt"`    // 失败数量
//	FailList   []AddFailChamber `json:"failList"`   // 失败列表
//}
//type AddFailChamber struct {
//	SerialNumber int    `json:"serialNumber"` // 序号
//	RowNum       int    `json:"rowNum"`       // 第几行出错
//	Reason       string `json:"reason"`       // 出错原因
//	ChamberList  Chamber
//}
//
//// 4. 新增场馆信息展示
//type AddChamberInfoReq struct {
//}
//type AddChamberInfoResp struct {
//	ChamberStatus []ChamberStatus `json:"chamberStatus"`
//}
//
//type ChamberStatus struct {
//	Status     int8   `json:"status"`     // 场馆状态
//	StatusName string `json:"statusName"` // 场馆状态名称
//}
//
//// 5. 新增场馆
//type AddChamberReq struct {
//	Chamber Chamber
//}
//type AddChamberResp struct {
//}
//
//// 6. 编辑场馆信息展示
//type EditChamberInfoReq struct {
//	ChamberId dto.ID `form:"chamberId"` // 场馆id
//}
//type EditChamberInfoResp struct {
//	Chamber Chamber
//	ChamberStatus []ChamberStatus `json:"chamberStatus"`
//}
//
//// 7. 编辑场馆
//type EditChamberReq struct {
//	Chamber Chamber
//
//}
//type EditChamberResp struct {
//}
